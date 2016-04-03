// navigate: kythe/cs/testdata/field.cc

// click 1/2: field
// url: http://localhost/kythe/cs/testdata/field.cc#12

// click 0/2: field
// check: <h1>References</h1>
// click 0/1: 16   c->field = 1;
// url: http://localhost/kythe/cs/testdata/field.cc#16

class cls {
  int field;
};

void function(cls *c) {
  c->field = 1;
}
