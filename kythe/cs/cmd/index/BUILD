load("/tools/build_rules/go", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "index",
    srcs = [
        "index.go",
    ],
    deps = [
        "//kythe/cs/cshtml",
        "//kythe/cs/defindex",
        "//kythe/cs/refpack",
        "//kythe/go/platform/delimited",
        "//kythe/proto:storage_proto_go",
        "//third_party/go/src/github.com/google/codesearch/index",
    ],
)
