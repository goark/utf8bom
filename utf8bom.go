package utf8bom

import (
	"bufio"
	"bytes"
	"io"
)

// Reader is decorator of io.Reader for stripping leading UTF-8 BOM.
type Reader struct {
	*bufio.Reader
	closer func() error
}

// Strip strips leading UTF-8 BOM and returns Reader object.
func Strip(r io.Reader) *Reader {
	closer := func() error { return nil }
	if c, ok := r.(io.Closer); ok {
		closer = c.Close
	}
	br := &Reader{Reader: bufio.NewReader(r), closer: closer}
	b, err := br.Peek(3)
	if err != nil {
		return br
	}
	if bytes.Equal(b, []byte{0xef, 0xbb, 0xbf}) { // compare BOM
		_, _ = br.Discard(3)
	}
	return br
}

// Close closes this object.
func (r *Reader) Close() error {
	return r.closer()
}

/* MIT License
 *
 * Copyright 2022 Spiegel (forked from github.com/spkg/bom package)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
