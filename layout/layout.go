package layout

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dvwallin/ago/agotypes"
	"github.com/dvwallin/ago/config"
	"github.com/dvwallin/ago/tmpl"
)

// GenerateHeader gives back the parsed header
func GenerateHeader() string {
	cfg := config.GetCfg()
	output := parse(tmpl.Header, cfg)
	return output
}

// GenerateFooter gives back the parsed header
func GenerateFooter() string {
	cfg := config.GetCfg()
	output := parse(tmpl.Footer, cfg)
	return output
}

func parse(input string, cfg agotypes.Config) string {
	v := reflect.ValueOf(cfg)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		input = strings.ReplaceAll(
			input,
			fmt.Sprintf(
				"[[%s]]",
				strings.ToUpper(typeOfS.Field(i).Name),
			),
			v.Field(i).Interface().(string),
		)
	}
	return input
}
