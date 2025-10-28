package hello

import "testing"

func TestSayHelloTo(t *testing.T) {
	words := SayHelloTo("jack")

	expected := "hello jack"

	if words != expected {
		t.Fatalf("expect: %s, got: %s", expected, words)
	}
}
