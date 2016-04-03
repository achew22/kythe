// navigate: kythe/cs/testdata/class.cc

// click 1/2: cls
// url: http://localhost/kythe/cs/testdata/class.cc#11

// click 0/2: cls
// check: <h1>References</h1>
// click 0/1: 14   cls c;
// url: http://localhost/kythe/cs/testdata/class.cc#14

class cls {};

void function() {
  cls c;
}
