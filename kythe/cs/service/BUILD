package(default_visibility = ["//visibility:public"])

load("//tools:build_rules/shims.bzl", "go_library")

licenses(["notice"])

go_library(
    name = "service",
    srcs = ["service.go"],
    deps = [
        "//kythe/cs/cshtml",
        "//kythe/cs/defindex",
        "//kythe/cs/refpack",
        "@com_github_google_codesearch//index:go_default_library",
        "@com_github_google_codesearch//regexp:go_default_library",
    ],
)
