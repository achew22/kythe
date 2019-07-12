load("//tools:build_rules/shims.bzl", "go_binary")

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
        "//kythe/proto:storage_go_proto",
        "@com_github_google_codesearch//index:go_default_library",
    ],
)
