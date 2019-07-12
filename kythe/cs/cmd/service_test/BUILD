load("//tools:build_rules/shims.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "service_test",
    srcs = [
        "service_test.go",
    ],
    deps = [
        "//kythe/cs/service",
        "@org_golang_x_net//html:go_default_library",
    ],
)
