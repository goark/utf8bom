# utf8bom -- Strip leading UTF-8 BOM (Byte Order Mark)

[![check vulns](https://github.com/goark/utf8bom/workflows/vulns/badge.svg)](https://github.com/goark/utf8bom/actions)
[![lint status](https://github.com/goark/utf8bom/workflows/lint/badge.svg)](https://github.com/goark/utf8bom/actions)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/goark/utf8bom/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/goark/utf8bom.svg)](https://github.com/goark/utf8bom/releases/latest)

This package is forked from [github.com/spkg/bom](https://github.com/spkg/bom) package.

## Usage

### Import

```go
import "github.com/goark/utf8bom"
```

### Strip leading UTF-8 BOM

```go
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
```

[utf8bom]: https://github.com/goark/utf8bom "Strip leading UTF-8 BOM"
