#!/bin/sh -e

here=$(cd `dirname $0` && pwd)

(cd $here && bazel build -c opt \
  //kythe/cxx/indexer/cxx:indexer \
  //kythe/cs/cmd/index \
)
bin="$here/../../../bazel-bin"

rm -rf /tmp/test
mkdir -p /tmp/test/tus

$here/extract_cmdline.sh /tmp/test "$@"
$bin/kythe/cxx/indexer/cxx/indexer /tmp/test/kindex/*.kindex | \
  $bin/kythe/cs/cmd/index/index dumptu /tmp/test/out
$bin/kythe/cxx/indexer/cxx/indexer /tmp/test/kindex/*.kindex | \
  $bin/kythe/cs/cmd/index/index tu /tmp/test/out > /tmp/test/tus/test.tu
$bin/kythe/cs/cmd/index/index dumpcorpus /tmp/test/out /tmp/test/tus
