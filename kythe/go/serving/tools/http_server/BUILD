load("//tools:build_rules/shims.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "http_server",
    srcs = [
        "csearch.go",
        "http_server.go",
    ],
    deps = [
        "//kythe/go/services/filetree",
        "//kythe/go/services/graph",
        "//kythe/go/services/graphstore",
        "//kythe/go/services/graphstore/proxy",
        "//kythe/go/services/xrefs",
        "//kythe/go/serving/filetree",
        "//kythe/go/serving/graph",
        "//kythe/go/serving/xrefs",
        "//kythe/go/storage/leveldb",
        "//kythe/go/storage/table",
        "//kythe/go/util/flagutil",
        "//kythe/proto:filetree_go_proto",
        "//kythe/proto:xref_go_proto",
        "@com_github_google_codesearch//index:go_default_library",
        "@com_github_google_codesearch//regexp:go_default_library",
        "@org_golang_x_net//http2:go_default_library",
    ],
)
