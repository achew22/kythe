#!/bin/sh -xe

# e.g.
#
# args=-std=c++11 /path/to/serve_cmdline.sh input1.cc input2.cc

here=$(cd `dirname $0` && pwd)

$here/extract_cmdline.sh /tmp/test "$@"
$here/index.sh /tmp/test
$here/serve.sh /tmp/test
