#!/bin/sh -xe

here=$(cd `dirname $0` && pwd)

(cd $here && bazel build -c opt \
  //kythe/cxx/extractor:cxx_extractor \
)
bin="$here/../../../bazel-bin"

export KYTHE_OUTPUT_DIRECTORY="$1/kindex"
rm -rf "$KYTHE_OUTPUT_DIRECTORY"
mkdir -p "$KYTHE_OUTPUT_DIRECTORY"
shift

for i in "$@"; do
  KYTHE_CORPUS=test KYTHE_ROOT_DIRECTORY=. $bin/kythe/cxx/extractor/cxx_extractor --with_executable /usr/bin/clang $args "$i"
done
