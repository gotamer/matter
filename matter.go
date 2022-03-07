// Package matter provides TOML frontmatter or endmatter unmarshalling.
package matter

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// Unmarshal parses TOML frontmatter or endmatter and returns the content.
// When no front or endmatter delimiters are present the original content is returned.
func Unmarshal(b []byte, v interface{}) (content []byte, err error) {
	// Delimiter
	var delim = []byte("+++")

	parts := bytes.SplitN(b, delim, 3)
	if len(parts) == 1 {
		content = parts[0]
	} else if len(parts[0]) > len(parts[2]) {
		content = parts[0]
	} else {
		content = parts[2]
	}
	_, err = toml.Decode(string(parts[1]), v)
	return
}

// func Marshal(v interface{}) (content []byte, err error) {}
