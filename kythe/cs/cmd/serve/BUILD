load("//tools:build_rules/shims.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "serve",
    srcs = [
        "serve.go",
    ],
    deps = [
        "//kythe/cs/service",
    ],
)
