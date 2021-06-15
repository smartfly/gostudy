package doc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"
)

func TestTokenOp(t *testing.T) {
	const jsonStream = `
		{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
	`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}

func TestCompact(t *testing.T) {
	const jsonStream = `
		{
			"Message": "Hello", 
			"Array": [1, 2, 3], 
			"Null": null, 
			"Number": 1.234
		}
	`
	var buf bytes.Buffer
	_ = json.Compact(&buf, []byte(jsonStream))
	fmt.Println(buf.String())
}
