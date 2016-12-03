load("@io_bazel_rules_go//go:def.bzl", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "service_test",
    srcs = [
        "service_test.go",
    ],
    deps = [
        "//kythe/cs/service",
	"@go_x_net//:html",
    ]
)
