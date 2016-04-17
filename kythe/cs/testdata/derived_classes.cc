// navigate: kythe/cs/testdata/derived_classes.cc

// click 0/3: B
// check: <h1>Derived Classes</h1>
// check: struct <b>D1</b> : B {};
// check: struct <b>D2</b> : B {};

struct B {};

struct D1 : B {};
struct D2 : B {};
