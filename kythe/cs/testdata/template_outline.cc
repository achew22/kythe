// navigate: kythe/cs/testdata/template_outline.cc

// click 0/3: f
// url: http://localhost/kythe/cs/testdata/template_outline.cc#19

// click 2/3: f
// url: http://localhost/kythe/cs/testdata/template_outline.cc#19

// click 1/3: f
// check: <h1>Declarations</h1>
// check: void <b>f</b>(int i);
// check: <h1>References</h1>
// check: this-&gt;<b>f</b>(42);

template <typename T> struct S {
  void f(int i);
};

template <typename T> void S<T>::f(int i) {
  this->f(42);
}

template struct S<int>;
