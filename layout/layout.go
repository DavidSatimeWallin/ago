package layout

import (
	"fmt"
	"strings"

	"github.com/dvwallin/ago/html"
	"github.com/dvwallin/ago/types"
)

// GenerateHeader gives back the parsed header
func GenerateHeader(ph types.Placeholders) {
	output := parse(html.Header, ph)
	// If github_account != empty == fill GITHUB_LINK with nav
	// <nav>
	//     <a href="https://github.com/[[GITHUB_ACCOUNT]]">My Github</a>
	// </nav>
	fmt.Println(output)

}

func parse(input string, ph types.Placeholders) string {
	for _, v := range ph {
		input = strings.ReplaceAll(input, v.Tag, v.Value)
	}
	return input
}
