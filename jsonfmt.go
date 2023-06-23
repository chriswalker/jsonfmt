// There are many `jsonfmt`s written in Go. This is mine.
//
// Assumes input from stdin (piped in), and outputs to stdout. This
// is primarily for use in editors to format JSON, in the vein of
// `gofmt`.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var in []byte
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "jsonfmt: unable to read from stdin: %v\n", err)
		os.Exit(1)
	}

	if len(in) > 0 {
		var buf bytes.Buffer
		if err := json.Indent(&buf, in, "", "  "); err != nil {
			fmt.Fprintf(os.Stderr, "jsonfmt: unable to indent: %v\n", err)
			os.Exit(1)
		}
		j := buf.Bytes()
		j = append(j, '\n')
		os.Stdout.Write(j)
	}
}
