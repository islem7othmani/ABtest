package dgraph

import (
	"fmt"
	"strings"
)

func Euid(uid string) string {
	return uid[1 : len(uid)-1]
}

// relation
// 0x
// _:new
func Tuid(uid string) string {
	if strings.HasPrefix(uid, "0x") {
		return fmt.Sprintf(`<%s>`, uid)
	} else if strings.HasPrefix(uid, "_:") {
		return uid
	}
	return fmt.Sprintf(`_:new_%s`, uid)
}

func Slag(predicate string) string {
	return fmt.Sprintf(`dt_%s`, predicate)
}
