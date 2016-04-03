load("/tools/build_rules/go", "go_binary")

package(default_visibility = ["//kythe:default_visibility"])

go_binary(
    name = "serve",
    srcs = [
        "serve.go",
    ],
    deps = [
        "//kythe/cs/service",
    ]
)
