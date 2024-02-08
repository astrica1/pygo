package pygo

import "fmt"

func Type(v any) string {
	return fmt.Sprintf("%T", v)
}
