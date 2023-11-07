package util

import (
	"fmt"
)

func Scan(hint string) (s string) {
	fmt.Print(hint)
	_, _ = fmt.Scan(&s)
	return
}
