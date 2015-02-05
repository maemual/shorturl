package shorturl

import (
	"testing"
)

func TestNew(t *testing.T) {
	short := New("githhub.com/maemual/shorturl", 0)
	if short != "ZRNdM1" {
		t.Error("Wrong Result")
	}
}

func BenchmarkNew(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		New("https://github.com/maemual/shorturl/blob/master/shorturl.go", 0)
	}
}
