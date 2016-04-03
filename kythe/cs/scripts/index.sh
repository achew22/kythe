#!/bin/sh -xe

here=$(cd `dirname $0` && pwd)

(cd $here && bazel build -c opt \
  //kythe/cxx/extractor:cxx_extractor \
  //kythe/cxx/indexer/cxx:indexer \
  //kythe/cs/cmd/index \
  //kythe/cs/cmd/serve \
)
bin="$here/../../../bazel-bin"

rm -rf $1/out $1/tus
mkdir -p $1/out $1/tus

set +x
for i in $1/kindex/*.kindex ; do
  echo "$bin/kythe/cxx/indexer/cxx/indexer $i | $bin/llvmcs/cmd/index/index tu $1/out > $1/tus/`basename $i`.tu"
done | parallel --gnu -v
set -x

$bin/kythe/cs/cmd/index/index corpus $1/out $1/tus $repos
