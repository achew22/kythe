load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def maybe(repo_rule, name, **kwargs):
    """Defines a repository if it does not already exist.
    """
    if name not in native.existing_rules():
        repo_rule(name = name, **kwargs)

def kythe_rule_repositories():
    """Defines external repositories for Kythe Bazel rules.

    These repositories must be loaded before calling external.bzl%kythe_dependencies.
    """

    skylib_version = "0.8.0"
    http_archive(
        name = "bazel_skylib",
        type = "tar.gz",
        url = "https://github.com/bazelbuild/bazel-skylib/releases/download/{}/bazel-skylib.{}.tar.gz".format(skylib_version, skylib_version),
        sha256 = "2ef429f5d7ce7111263289644d233707dba35e39696377ebab8b0bc701f7818e",
    )

    git_repository(
        name = "io_bazel_rules_go",
        remote = "https://github.com/bazelbuild/rules_go",
        commit = "e2795c08da49e5806ebb52810f68782cfb8f3a6e",
    )

    maybe(
        http_archive,
        name = "bazel_gazelle",
        sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
        urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    )

    maybe(
        http_archive,
        name = "build_bazel_rules_nodejs",
        sha256 = "e04a82a72146bfbca2d0575947daa60fda1878c8d3a3afe868a8ec39a6b968bb",
        urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/0.31.1/rules_nodejs-0.31.1.tar.gz"],
    )
