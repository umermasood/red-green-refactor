package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("A", 3)
	want := "AAA"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("Go! ", 3))
	// Output:
	// Go! Go! Go!
}
