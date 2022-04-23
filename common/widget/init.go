package widget

import "strings"

func SnakeCase(s []string) string {
	var r [] string
	for _, v := range s {
		r = append(r, strings.Title(v))
	}
	return strings.Join(r, "")
}
