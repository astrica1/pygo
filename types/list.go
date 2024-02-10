package pygo

import "fmt"

type List []any

func (l List) String() string {
	if len(l) == 0 {
		return "[]"
	}
	result := "["
	for _, i := range l {
		result += fmt.Sprintf("%v, ", i)
	}
	return result[:len(result)-2] + "]"
}

func (l *List) Append(values ...any) {
	(*l) = append((*l), values...)
}

func (l *List) Clear() {
	(*l) = []any{}
}
