#! /bin/bash

set -ex

here="$(cd `dirname $0` && pwd)"
bin="${here}/bazel-bin/external/io_kythe"
#corpus=$(mktemp -d /tmp/kythe.XXXXXXX)
corpus=/tmp/kythe.static
query=(
  "//kythe/go/languageserver/..."
)

(cd $here && bazelisk build \
  @io_kythe//kythe/go/indexer/cmd/go_indexer \
  @io_kythe//kythe/cs/cmd/index \
  @io_kythe//kythe/cs/cmd/serve \
)

bazelisk build \
  --keep_going \
  --experimental_action_listener=//kythe/extractors:extract_kzip_go \
  "${query[@]}"

mkdir "${corpus}/tus" || true
mkdir "${corpus}/out" || true

for i in $( find bazel-out/k8-fastbuild/extra_actions/kythe/extractors/extract_kzip_go_extra_action | grep \\.kzip\$ ) ; do
  echo "${bin}/kythe/go/indexer/cmd/go_indexer/go_indexer $i | ${bin}/kythe/cs/cmd/index/index tu ${corpus}/out > ${corpus}/tus/$(basename $i).tu"
done | parallel --gnu -v

"${bin}/kythe/cs/cmd/index/index" corpus "${corpus}/out" "${corpus}/tus" # "="

echo "To see your cross referencing server:"
echo "${bin}/kythe/cs/cmd/serve/serve -index_dir ${corpus}/out"
echo "and then open http://localhost:8080"
