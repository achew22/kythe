package(default_visibility = ["//visibility:public"])

load("//tools:build_rules/go.bzl", "go_package_library", "go_test")

licenses(["notice"])

go_package_library(
    name = "refpack",
    srcs = ["refpack.go"],
)

go_test(
    name = "refpack_test",
    srcs = ["refpack_test.go"],
    library = "refpack",
)
