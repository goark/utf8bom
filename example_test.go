package utf8bom_test

import (
	"fmt"
	"io"
	"os"

	"github.com/goark/utf8bom"
)

func ExampleReader() {
	rc, err := func(path string) (io.ReadCloser, error) {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		return utf8bom.Strip(file), nil
	}("testdata/sample.txt")
	if err != nil {
		return
	}
	defer rc.Close()

	b, err := io.ReadAll(rc)
	if err != nil {
		return
	}
	fmt.Println(string(b))
	// Output:
	// hello
}
