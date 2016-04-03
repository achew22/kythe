// navigate: kythe/cs/testdata/function.cc
// check: <span class="unref">function2</span>

// click 1/2: function1
// url: http://localhost/kythe/cs/testdata/function.cc#12

// click 0/2: function1
// check: <h1>References</h1>
// click 0/1: 15   function1();
// url: http://localhost/kythe/cs/testdata/function.cc#15

void function1() {}

void function2() {
  function1();
}
