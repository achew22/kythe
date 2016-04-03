#!/bin/sh -xe

here=$(cd `dirname $0` && pwd)

(cd $here && bazel build -c opt \
  //kythe/cxx/extractor:cxx_extractor \
  //third_party/jq:jq \
)
bin="$here/../../../bazel-bin"

export KYTHE_ROOT_DIRECTORY="$(cd $2 && pwd)"

rm -rf "$1/kindex"
mkdir -p "$1/kindex"
export KYTHE_OUTPUT_DIRECTORY="$(cd $1 && pwd)/kindex"

JQ=$bin/third_party/jq/jq KYTHE_EXTRACTOR=$bin/kythe/cxx/extractor/cxx_extractor KYTHE_CORPUS=test \
  $here/../../../kythe/extractors/cmake/extract_compilation_database.sh -
