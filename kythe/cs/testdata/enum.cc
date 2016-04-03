// navigate: kythe/cs/testdata/enum.cc

// click 1/2: enumerator
// url: http://localhost/kythe/cs/testdata/enum.cc#12

// click 0/2: enumerator
// check: <h1>References</h1>
// click 0/1: 16   return enumerator;
// url: http://localhost/kythe/cs/testdata/enum.cc#16

enum e {
  enumerator,
};

e function() {
  return enumerator;
}
