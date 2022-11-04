package utf8bom_test

import (
	"fmt"
	"io"
	"os"

	"github.com/goark/utf8bom"
)

func ExampleReader() {
	file, err := os.Open("testdata/sample.txt")
	if err != nil {
		return
	}
	r := utf8bom.Strip(file)
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		return
	}
	fmt.Println(string(b))
	// Output:
	// hello
}
