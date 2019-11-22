package tpl

import "regexp"

var patterns = map[string]string{
	`where[\r]*[\n]*[\s]*order by`:                   "order by",
	`where[\r]*[\n]*[\s]*group by`:                   "group by",
	`where[\r]*[\n]*[\s]*limit`:                      "limit",
	`where[\r]*[\n]*[\s]*or$`:                        "",
	`where[\r]*[\n]*[\s]*and$`:                       "",
	`where[\r]*[\n]*[\s]*or[\r]*[\n]*[\s]*order by`:  "order by",
	`where[\r]*[\n]*[\s]*or[\r]*[\n]*[\s]*group by`:  "group by",
	`where[\r]*[\n]*[\s]*or[\r]*[\n]*[\s]*limit`:     "limit",
	`where[\r]*[\n]*[\s]*and[\r]*[\n]*[\s]*order by`: "order by",
	`where[\r]*[\n]*[\s]*and[\r]*[\n]*[\s]*group by`: "group by",
	`where[\r]*[\n]*[\s]*and[\r]*[\n]*[\s]*limit`:    "limit",
	`where[\r]*[\n]*[\s]*or[\n]*[\s]*`:               "where ",
	`where[\r]*[\n]*[\s]*and[\n]*[\s]*`:              "where ",
	`where[\r]*[\n]*[\s]*$`:                          "",
}

func replaceSpecialCharacter(s string) string {
	var result = s
	for k, v := range patterns {
		brackets, _ := regexp.Compile(k)
		result = brackets.ReplaceAllStringFunc(result, func(s string) string {
			return v
		})
	}
	return result

}
