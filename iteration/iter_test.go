package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	exceptedt := "aaaaaaaaaa"
	if repeated != exceptedt {
		t.Errorf("excepted  %q  but got %q", exceptedt, repeated)
	}
}

func ExampleRepeat() {
	repeated1 := Repeat("b")
	fmt.Println(repeated1)
	//Output : aaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
