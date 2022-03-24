matter
======

TOML Frontmatter or Endmatter unMarshaller.

```go
func Unmarshal(b []byte, v interface{}) (content []byte, err error)
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