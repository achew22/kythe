load("@io_bazel_rules_go//go:def.bzl", "go_binary")

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
        "@go_codesearch//:index",
    ],
)
