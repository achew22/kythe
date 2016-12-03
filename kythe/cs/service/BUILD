package(default_visibility = ["//visibility:public"])

load("//tools:build_rules/go.bzl", "go_package_library")

licenses(["notice"])

go_package_library(
    name = "service",
    srcs = ["service.go"],
    deps = [
        "//kythe/cs/cshtml",
        "//kythe/cs/defindex",
        "//kythe/cs/refpack",
        "@go_codesearch//:index",
        "@go_codesearch//:regexp",
    ],
)
