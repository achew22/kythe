load("//tools:build_rules/shims.bzl", "go_test", "go_library")

package(default_visibility = ["//kythe:default_visibility"])

go_library(
    name = "columnar",
    srcs = ["columnar_encoding.go"],
    deps = [
        "//kythe/go/util/keys",
        "//kythe/go/util/kytheuri",
        "//kythe/proto:schema_go_proto",
        "//kythe/proto:storage_go_proto",
        "//kythe/proto:xref_serving_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

go_test(
    name = "columnar_test",
    size = "small",
    srcs = ["columnar_test.go"],
    library = ":columnar",
    visibility = ["//visibility:private"],
    deps = [
        "//kythe/go/util/compare",
        "//kythe/proto:common_go_proto",
        "//kythe/proto:serving_go_proto",
    ],
)
