package hello_test

import (
	"fmt"

	"github.com/raizyr/wercker_go_test"
)

func ExampleHelloWorld() {
	fmt.Println(hello.Hello())
	// Output: hello world
}
