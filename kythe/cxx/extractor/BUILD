package(default_visibility = ["//kythe:default_visibility"])

cc_library(
    name = "index_pack",
    srcs = ["index_pack.cc"],
    hdrs = ["index_pack.h"],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deprecation = "Please use :kzip_{reader,writer}",
    deps = [
        "//external:libuuid",
        "//external:zlib",
        "//kythe/cxx/common:lib",
        "//kythe/proto:analysis_cc_proto",
        "//third_party/llvm",
        "@boringssl//:crypto",
        "@com_google_absl//absl/memory",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_library(
    name = "index_pack_testlib",
    testonly = 1,
    srcs = ["index_pack_test.cc"],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":index_pack",
        "//external:zlib",
        "//third_party:gtest",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_github_tencent_rapidjson//:rapidjson",
        "@com_google_absl//absl/memory",
        "@com_google_protobuf//:protobuf",
    ],
)

cc_test(
    name = "index_pack_test",
    size = "small",
    deps = [":index_pack_testlib"],
)

cc_library(
    name = "cxx_details",
    srcs = [
        "cxx_details.cc",
    ],
    hdrs = [
        "cxx_details.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    visibility = [
        "//kythe:default_visibility",
        "//third_party/llvm/src:__pkg__",
    ],
    deps = [
        "//kythe/proto:cxx_cc_proto",
        "//third_party/llvm",
        "@com_github_google_glog//:glog",
    ],
)

cc_library(
    name = "path_utils",
    srcs = ["path_utils.cc"],
    hdrs = ["path_utils.h"],
    deps = [
        "//kythe/cxx/common:lib",
        "//third_party/llvm",
    ],
)

cc_library(
    name = "path_utils_testlib",
    testonly = 1,
    srcs = [
        "path_utils_test.cc",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":path_utils",
        "//third_party:gtest",
        "//third_party/llvm",
    ],
)

cc_test(
    name = "path_utils_test",
    size = "small",
    deps = [
        ":path_utils_testlib",
    ],
)

cc_library(
    name = "supported_language",
    srcs = ["language.cc"],
    hdrs = ["language.h"],
    deps = ["//third_party/llvm"],
)

cc_library(
    name = "command_line_utils",
    srcs = [
        "CommandLineUtils.cc",
    ],
    hdrs = [
        "CommandLineUtils.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        "//third_party/llvm",
        "@com_github_google_glog//:glog",
    ],
)

cc_library(
    name = "commandline_testlib",
    testonly = 1,
    srcs = [
        "CommandLineUtilsTest.cc",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":command_line_utils",
        "//third_party:gtest",
        "//third_party/llvm",
    ],
)

cc_test(
    name = "commandline_test",
    size = "small",
    deps = [
        ":commandline_testlib",
    ],
)

cc_library(
    name = "lib",
    srcs = [
        "cxx_extractor.cc",
    ],
    hdrs = [
        "cxx_extractor.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":command_line_utils",
        ":index_pack",
        ":path_utils",
        ":supported_language",
        "//external:zlib",
        "//kythe/cxx/common:index_writer",
        "//kythe/cxx/common:json_proto",
        "//kythe/cxx/common:kzip_writer",
        "//kythe/cxx/common:lib",
        "//kythe/cxx/indexer/cxx:clang_utils",
        "//kythe/proto:analysis_cc_proto",
        "//kythe/proto:buildinfo_cc_proto",
        "//kythe/proto:cxx_cc_proto",
        "//third_party/llvm",
        "//third_party/llvm/src:clang_builtin_headers",
        "//third_party/llvm/src:preprocessor_util_lib",
        "@boringssl//:crypto",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_absl//absl/memory",
        "@com_google_absl//absl/strings",
        "@com_google_protobuf//:protobuf",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_library(
    name = "cmdlib",
    srcs = [
        "cxx_extractor_main.cc",
    ],
    hdrs = [
        "cxx_extractor.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lib",
        ":supported_language",
        "//external:zlib",
        "//kythe/proto:analysis_cc_proto",
        "//third_party/llvm",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_library(
    name = "bazelcmdlib",
    srcs = [
        "cxx_extractor_bazel_main.cc",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lib",
        ":supported_language",
        "//external:zlib",
        "//third_party/bazel:extra_actions_base_cc_proto",
        "//third_party/llvm",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_library(
    name = "objcbazelcmdlib",
    srcs = [
        "objc_extractor_bazel_main.cc",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lib",
        ":objc_bazel_support_library",
        ":supported_language",
        "//external:zlib",
        "//third_party/bazel:extra_actions_base_cc_proto",
        "//third_party/llvm",
        "@com_github_gflags_gflags//:gflags",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_library(
    name = "objc_bazel_support_library",
    srcs = ["objc_bazel_support.cc"],
    hdrs = ["objc_bazel_support.h"],
    deps = [
        "//third_party/bazel:extra_actions_base_cc_proto",
        "//third_party/llvm",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_test(
    name = "objc_bazel_support_library_test",
    size = "small",
    srcs = ["objc_bazel_support_test.cc"],
    deps = [
        ":objc_bazel_support_library",
        "//third_party:gtest",
        "@com_github_google_glog//:glog",
    ],
)

action_listener(
    name = "extract_kindex",
    extra_actions = [":extra_action"],
    mnemonics = ["CppCompile"],
    visibility = ["//visibility:public"],
)

extra_action(
    name = "extra_action",
    cmd = "$(location :cxx_extractor_bazel) \
        $(EXTRA_ACTION_FILE) \
        $(output $(ACTION_ID).c++.kindex) \
        $(location //external:vnames_config)",
    data = ["//external:vnames_config"],
    out_templates = ["$(ACTION_ID).c++.kindex"],
    tools = [":cxx_extractor_bazel"],
)

action_listener(
    name = "extract_kindex_objc",
    extra_actions = [":extra_action_objc"],
    mnemonics = ["ObjcCompile"],
    visibility = ["//visibility:public"],
)

extra_action(
    name = "extra_action_objc",
    cmd = "$(location :objc_extractor_bazel) \
        $(EXTRA_ACTION_FILE) \
        $(output $(ACTION_ID).objc.kindex) \
        $(location //external:vnames_config) \
        $(location //third_party/bazel:get_devdir) \
        $(location //third_party/bazel:get_sdkroot)",
    data = [
        "//external:vnames_config",
        "//third_party/bazel:get_devdir",
        "//third_party/bazel:get_sdkroot",
    ],
    out_templates = ["$(ACTION_ID).objc.kindex"],
    tools = [":objc_extractor_bazel"],
)

action_listener(
    name = "extract_kzip",
    extra_actions = [":extra_action_kzip"],
    mnemonics = ["CppCompile"],
    visibility = ["//visibility:public"],
)

extra_action(
    name = "extra_action_kzip",
    cmd = "$(location :cxx_extractor_bazel) \
        $(EXTRA_ACTION_FILE) \
        $(output $(ACTION_ID).c++.kzip) \
        $(location //external:vnames_config)",
    data = ["//external:vnames_config"],
    out_templates = ["$(ACTION_ID).c++.kzip"],
    tools = [":cxx_extractor_bazel"],
)

action_listener(
    name = "extract_kindex_objc_kzip",
    extra_actions = [":extra_action_objc_kzip"],
    mnemonics = ["ObjcCompile"],
    visibility = ["//visibility:public"],
)

extra_action(
    name = "extra_action_objc_kzip",
    cmd = "$(location :objc_extractor_bazel) \
        $(EXTRA_ACTION_FILE) \
        $(output $(ACTION_ID).objc.kzip) \
        $(location //external:vnames_config) \
        $(location //third_party/bazel:get_devdir) \
        $(location //third_party/bazel:get_sdkroot)",
    data = [
        "//external:vnames_config",
        "//third_party/bazel:get_devdir",
        "//third_party/bazel:get_sdkroot",
    ],
    out_templates = ["$(ACTION_ID).objc.kzip"],
    tools = [":objc_extractor_bazel"],
)

cc_binary(
    name = "cxx_extractor",
    deps = [
        ":cmdlib",
    ],
)

cc_binary(
    name = "cxx_extractor_bazel",
    deps = [
        ":bazelcmdlib",
    ],
)

cc_binary(
    name = "objc_extractor_bazel",
    deps = [
        ":objcbazelcmdlib",
    ],
)

cc_library(
    name = "testlib",
    testonly = 1,
    srcs = [
        "cxx_extractor_test.cc",
    ],
    hdrs = [
        "cxx_extractor.h",
    ],
    copts = [
        "-Wno-non-virtual-dtor",
        "-Wno-unused-variable",
        "-Wno-implicit-fallthrough",
    ],
    deps = [
        ":lib",
        "//external:zlib",
        "//kythe/proto:analysis_cc_proto",
        "//third_party:gtest",
        "//third_party/llvm",
        "@com_github_google_glog//:glog",
        "@com_google_protobuf//:protobuf",
        "@com_googlesource_code_re2//:re2",
    ],
)

cc_test(
    name = "test",
    size = "small",
    deps = [
        ":testlib",
    ],
)
