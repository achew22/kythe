// navigate: kythe/cs/testdata/multiple_definitions.cc

// click 2/3: bar
// check: <title>Multiple Definitions</title>

// click 0/2: bar
// url: http://localhost/kythe/cs/testdata/multiple_definitions.cc#14
// back

// click 1/2: bar
// url: http://localhost/kythe/cs/testdata/multiple_definitions.cc#18

struct c1 {
  void bar() {}
};

struct c2 {
  void bar() {}
};

struct c3 {
  void bar() {}
};

template <typename T> void foo() {
  T().bar();
}

void function() {
  foo<c1>();
  foo<c2>();
}
