package main

import (
	"fmt"
	"io"
)

func updateChangelog(w io.Writer, version, shortDesc, longDesc string) error {
	_, err := fmt.Fprintf(w, "## %s\n- %s\n", version, shortDesc)
	if err != nil {
		return err
	}
	if longDesc != "" {
		_, err = fmt.Fprintf(w, "  %s\n", longDesc)
	}
	return err
}
