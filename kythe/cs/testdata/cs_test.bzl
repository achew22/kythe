def cs_test(name, src):
  native.sh_test(
      name = name,
      srcs = ["cs_test.sh"],
      data = [src] + [
          "//kythe/cs/cmd/index",
          "//kythe/cs/cmd/service_test",
          "//kythe/cxx/indexer/cxx:indexer",
      ],
      args = ["$(location %s)" % src]
  )
