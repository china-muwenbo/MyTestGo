package gotest

import (
	"testing"
	"fmt"
	"log"
)

func TestBase(t *testing.T) {
	  fmt.Println("This is TestBase func")
	  Base1()
	  base()
}

func BenchmarkHello(b *testing.B) {
	fmt.Println("b.N=",b.N)

	for i := 0; i < b.N; i++ {
		fmt.Println("test")
	}

}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
	log.Println("hello")
}
func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) { fmt.Println("B=1") })
	t.Run("A=2", func(t *testing.T) { fmt.Println("b=2") })
	t.Run("B=1", func(t *testing.T) { fmt.Println("B=1") })
	// <tear-down code>
}
