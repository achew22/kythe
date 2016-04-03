#!/bin/sh

infile="$1"
outdir="$TEST_TMPDIR/out"
tudir="$TEST_TMPDIR/tus"

rm -rf "$tudir" "$outdir"
mkdir -p "$tudir" "$outdir"

kythe/cxx/indexer/cxx/indexer -i "$infile" | \
  kythe/cs/cmd/index/index tu "$outdir" > "$tudir/test.tu"
kythe/cs/cmd/index/index corpus "$outdir" "$tudir"
kythe/cs/cmd/service_test/service_test "$outdir" "$infile"
