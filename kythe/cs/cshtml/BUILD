package(default_visibility = ["//visibility:public"])

load("//tools:build_rules/go.bzl", "go_package_library", "go_test")

licenses(["notice"])

go_package_library(
    name = "cshtml",
    srcs = ["cshtml.go"],
)

go_test(
    name = "cshtml_test",
    srcs = ["cshtml_test.go"],
    library = "cshtml",
)
