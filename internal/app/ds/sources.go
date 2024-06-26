package ds

import "strings"

type Source int

const (
	NoSource Source = iota
	Codeforces
	Codewars
	ProjectEuler
	Timus
	InformaticsMskRu
	Leetcode
	Hackerrank
)

var (
	SourcesMap = map[string]Source{
		"codeforces":       Codeforces,
		"codewars":         Codewars,
		"projecteuler":     ProjectEuler,
		"timus":            Timus,
		"informaticsmskru": InformaticsMskRu,
		"leetcode":         Leetcode,
		"hackerrank":       Hackerrank,
	}
)

func ParseSourceString(str string) (Source, bool) {
	c, ok := SourcesMap[strings.ToLower(str)]
	return c, ok
}
