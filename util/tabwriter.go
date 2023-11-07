package util

import (
	"bytes"
	"text/tabwriter"
)

func TabWriter(lines ...string) string {
	buf := &bytes.Buffer{}
	wri := tabwriter.NewWriter(buf, 0, 0, 2, ' ', 0)
	for i := 0; i < len(lines); i++ {
		_, _ = wri.Write([]byte(lines[i] + "\n"))
	}
	_ = wri.Flush()
	return buf.String()
}
