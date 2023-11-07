package util

import (
	"bufio"
	"fmt"
	"os"

	"github.com/starudream/go-lib/core/v2/utils/signalutil"
)

func Scan(hint string) (s string) {
	done := make(chan struct{})

	go func() {
		in := bufio.NewScanner(os.Stdin)
		for fmt.Print(hint); in.Scan(); fmt.Print(hint) {
			if s = in.Text(); s != "" {
				close(done)
				break
			}
		}
	}()

	select {
	case <-signalutil.Done():
	case <-done:
	}

	return
}
