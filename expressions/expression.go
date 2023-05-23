package expression

import "strings"

type Expression interface {
	print(builder strings.Builder)
}
