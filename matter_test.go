package matter_test

import (
	"testing"
	"fmt"
	
	"matter"
)

type front struct {
	Name string 
	Tags []string 
}

var inputEnd = `

Rest of the content.

+++
name = "endmatter"
tags = ['input', 'end']
+++

`

var inputFront = `

+++
name = "frontmatter"
tags = ['input', 'front']
+++

Rest of the content.


`

func TestFront(t *testing.T){
	ty := new(front)
	res, err := matter.Unmarshal([]byte(inputFront), ty)
	if err != nil {
		t.Error( err)
	}
	fmt.Printf("\nres: %s\n", res)
	fmt.Printf("ty: %s\n", ty)
}

func TestEnd(t *testing.T){
	ty := new(front)
	res, err := matter.Unmarshal([]byte(inputEnd), ty)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("\nres: %s\n", res)
	fmt.Printf("ty: %s\n", ty)
}
