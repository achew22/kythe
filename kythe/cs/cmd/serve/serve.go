/*
 * Copyright 2016 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// HTTP server that uses the kythe.io/kythe/cs/service package to serve a corpus
// index.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"kythe.io/kythe/cs/service"
)

var (
	indexDir        = flag.String("index_dir", "", "Path to index to serve")
	indexStorageDir = flag.String("index_storage_dir", "", "Path to storage when managing indexes")
	indexerLogDir   = flag.String("indexer_log_dir", "", "Path to directory for indexer log storage")
	indexerCommand  = flag.String("indexer_command", "", "Command to run indexer")
	indexerInterval = flag.Int("indexer_interval", 86400, "Interval between indexer runs, in seconds")
	indexerMod      = flag.Int("indexer_mod", 0, "Indexer runs when (time.Now().Unix() % indexer_interval) = this")
)

func httpServer(s *service.Service) {
	var srv http.Server
	srv.Addr = ":8080"
	srv.Handler = s
	srv.ListenAndServe()
}

func runStaticIndex(path string) {
	var s service.Service
	go httpServer(&s)

	ix, err := service.LoadIndex(path)
	if err != nil {
		log.Fatalf("error loading index %s: %s\n", path, err)
	}
	s.SetIndex(ix)
	select {}
}

func runManagedIndex(storageDir, logDir, cmd string, interval, mod int) {
	var s service.Service
	go httpServer(&s)

	// find the latest index
	files, err := ioutil.ReadDir(storageDir)
	var ix *service.Index
	if len(files) != 0 {
		indexpath := storageDir + "/" + files[len(files)-1].Name()
		ix, err = service.LoadIndex(indexpath)
		if err != nil {
			log.Printf("error loading latest index %s: %s\n", indexpath, err)
		} else {
			files = files[:len(files)-1]
		}
	}

	// clean up any remaining indices
	for _, f := range files {
		indexpath := storageDir + "/" + f.Name()
		log.Printf("removing old index %s\n", indexpath)
		err := os.RemoveAll(indexpath)
		if err != nil {
			log.Printf("error removing old index: %s\n", indexpath, err)
		}
	}

	if ix != nil {
		s.SetIndex(ix)
	}

	lastindexpath := ""

	for {
		timeUntilNextIndex := 0
		if ix != nil {
			modNow := time.Now().Unix() - int64(mod)
			nextIndexTime := (modNow / int64(interval) * int64(interval)) + int64(interval)
			timeUntilNextIndex = int(nextIndexTime - modNow)
		}

		if lastindexpath != "" {
			timeUntilCleanup := 60
			if timeUntilNextIndex < timeUntilCleanup {
				timeUntilNextIndex = 0
			} else {
				timeUntilNextIndex -= timeUntilCleanup
			}

			time.Sleep(time.Duration(timeUntilCleanup) * time.Second)
			log.Printf("removing previous index %s\n", lastindexpath)
			err := os.RemoveAll(lastindexpath)
			if err != nil {
				log.Printf("error removing previous index: %s\n", lastindexpath, err)
			}
			lastindexpath = ""
		}

		time.Sleep(time.Duration(timeUntilNextIndex) * time.Second)

		now := time.Now().Unix()
		path := fmt.Sprintf("%s/%020d", storageDir, now)
		logpath := fmt.Sprintf("%s/%020d.log", logDir, now)

		logf, err := os.Create(logpath)
		if err != nil {
			log.Printf("error opening log file: %s\n", err)
		}

		fullcmd := cmd + " " + path
		log.Printf("running indexer command %#v\n", fullcmd)

		cmdobj := exec.Command("sh", "-c", fullcmd)
		if logf != nil {
			cmdobj.Stdout = logf
			cmdobj.Stderr = logf
		}
		err = cmdobj.Run()
		if err != nil {
			log.Printf("error running indexer command: %s\n", err)
			continue
		}

		if logf != nil {
			logf.Close()
		}

		log.Printf("loading new index from %s\n", path)
		newix, err := service.LoadIndex(path)
		if err != nil {
			log.Printf("error loading new index: %s\n", path, err)
		} else {
			if ix != nil {
				lastindexpath = ix.Dir()
			}
			ix = newix
			s.SetIndex(newix)
			log.Printf("serving new index\n")
		}
		runtime.GC()
	}
}

func main() {
	flag.Parse()
	if *indexDir != "" {
		runStaticIndex(*indexDir)
	} else if *indexStorageDir != "" && *indexerLogDir != "" && *indexerCommand != "" {
		runManagedIndex(*indexStorageDir, *indexerLogDir, *indexerCommand, *indexerInterval, *indexerMod)
	} else {
		log.Fatal("need -index_dir or -index_storage_dir -indexer_log_dir -indexer_command")
	}
}
