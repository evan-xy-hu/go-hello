package hello

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	name := "Mike"
	want := fmt.Sprintf("Hello %s, welcome!\n", name)
	got := SayHello(name)
	if got != want {
		t.Fatalf("Not match: want: %s, but got: %s", want, got)
	}

}
