matter
======

TOML Frontmatter or Endmatter unMarshaller.

```go
func Unmarshal(b []byte, v interface{}) (content []byte, err error)
```
```go
import	"github.com/gotamer/matter"

// Matter unmarshals matter to v, and returns a reader of contents.
func Matter(r io.Reader, v interface{}) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		s, err := matter.Unmarshal(b, v)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		pw.Write(s)
		pw.Close()
	}()

	return pr
}

```



Either input will work fine:
```go

var inputFront = `

+++
name = "frontmatter"
tags = ['input', 'front']
+++

Matter first.


`


var inputEnd = `

Content first.

+++
name = "endmatter"
tags = ['input', 'end']
+++

`

```