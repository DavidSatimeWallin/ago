package layout

import (
	"fmt"
	"strings"

	"github.com/dvwallin/ago/html"
	"github.com/dvwallin/ago/types"
)

// GenerateHeader gives back the parsed header
func GenerateHeader(ph types.Placeholders) {
	fmt.Println(parse(html.Header, ph))
}

func parse(input string, ph types.Placeholders) string {
	for _, v := range ph {
		input = strings.ReplaceAll(input, v.Tag, v.Value)
	}
	return input
}
