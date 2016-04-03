// navigate: kythe/cs/testdata/variable.cc

// click 1/2: variable
// url: http://localhost/kythe/cs/testdata/variable.cc#11

// click 0/2: variable
// check: <h1>References</h1>
// click 0/1: 14   variable = 1;
// url: http://localhost/kythe/cs/testdata/variable.cc#14

int variable;

void function() {
  variable = 1;
}
