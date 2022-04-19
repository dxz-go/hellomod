package hellomod

import "testing"

func TestHello(t *testing.T) {
    want := "V2: Hello World!"
    if Hello() != want {
        t.Errorf("Hello() != %s", want)
    }
}
