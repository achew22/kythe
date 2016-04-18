// navigate: kythe/cs/testdata/for_range.cc

// click 1/2: s
// url: http://localhost/kythe/cs/testdata/for_range.cc#16

// click 0/1: :
// check: <h1>Definitions</h1>
// check: char *<b>begin</b>();
// check: char *<b>end</b>();

struct S {
  char *begin();
  char *end();
};

void f(S s) {
  for (char c : s) {
    c++;
  }
}
