// navigate: ?q=symbol

// check: <a href="/kythe/cs/testdata/search.cc#9"><span class="lines">    9</span> int <b>symbol</b>;</a>
// check: <a href="/kythe/cs/testdata/search.cc#14"><span class="lines">   14</span> long <b>symbol</b>;</a>
// check: <a href="/kythe/cs/testdata/search.cc#20"><span class="lines">   20</span> char <b>symbol</b>;</a>
// check: <a href="/kythe/cs/testdata/search.cc#10"><span class="lines">   10</span> int <b>symbolic</b>;</a>
// check: <a href="/kythe/cs/testdata/search.cc#1"><span class="lines">    1</span> // navigate: ?q=symbol<b></b></a>

int symbol;
int symbolic;

namespace foo {

long symbol;

}

struct bar {

char symbol;

};

void f(bar *b) {
  return symbol + symbolic + foo::symbol + b->symbol;
}
