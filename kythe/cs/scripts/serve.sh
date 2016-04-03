#!/bin/sh -xe

here=$(cd `dirname $0` && pwd)

(cd $here && bazel build -c opt \
  //kythe/cs/cmd/serve \
)
bin="$here/../../../bazel-bin"

$bin/kythe/cs/cmd/serve/serve -index_dir "$1/out"
