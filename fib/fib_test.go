package fib

import "testing"

/*
func TestFib(t *testing.T) {

	//
	input := 0
	got := Fib(input)
	want := 0
	if got != want {
		t.Errorf("Fib(%d): want %d, got %d", input, want, got)
	}

	//
	input = 1
    got = Fib(input)
    want = 1
    if got != want {
    t.Errorf("Fib(%d): want %d, got %d", input, want, got)
    }

  }
*/

func TestFib(t *testing.T) {
	type test struct {
		input int
		want int
	  }

	  var tests = []test {
		{0,0},
		{1,1},
		{2,1},
		{3,2},
	  }

	  for _, tc := range tests {
		  got := Fib(tc.input)
		  if got != tc.want {
			  t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
		  }
	  }
}