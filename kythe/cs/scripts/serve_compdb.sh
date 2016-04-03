#!/bin/sh -xe

# e.g.
#
# cd /path/to/project
# mkdir gen
# cd gen
# cmake -G Ninja ..
# /path/to/serve_compdb.sh .. < compile_commands.json

here=$(cd `dirname $0` && pwd)

$here/extract_compdb.sh /tmp/test "$1"
$here/index.sh /tmp/test
$here/serve.sh /tmp/test
