package(default_visibility = ["@//visibility:public"])

load("@//third_party:go/build.bzl", "external_go_package")

licenses(["notice"])

exports_files(["LICENSE"])

external_go_package(
    name = "index",
    base_pkg = "github.com/google/codesearch",
    deps = [
        ":sparse",
    ],
    exclude_srcs = [
        "mmap_bsd.go",
        "mmap_windows.go",
    ],
)

external_go_package(
    name = "regexp",
    base_pkg = "github.com/google/codesearch",
    deps = [
        ":sparse",
    ],
)

external_go_package(
    name = "sparse",
    base_pkg = "github.com/google/codesearch",
)
