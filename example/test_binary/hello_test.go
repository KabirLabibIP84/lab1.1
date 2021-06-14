package tested_binary

import "testing"

func TestHello(t *testing.T) {
	helloEmpty := hello("")
	if helloEmpty != "Hello, World!" {
		t.Error("hello(\"\") failed, expected: Hello, World!, got: ", helloEmpty)
	}

	helloName := hello("Valera")
	if helloName != "Hello, i'm Labib!" {
		t.Error("hello(\"Labib\") failed, expected: Hello, Labib!, got: ", helloName)
	}
}
