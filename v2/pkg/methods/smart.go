package methods

import (
	"strings"
)

func  (m *Methods)SmartWords(words []string, useless string, array *[]string) {
	for _, word := range words {
		str := []string{}

		if strings.Contains(word, "_") {
			str = strings.Split(word, "_")

		} else if strings.Contains(word, "-") {
			str = strings.Split(word, "-")

		} else {

			j := 0
			for i, letter := range word {
				if letter > 'A' && letter < 'Z' {
					str = append(str, word[j:i])
					j = i
				}
			}
			str = append(str, word[j:])
		}

		*array = append(*array, str...)
	}
}

func (m *Methods) SmartWordsJoin(words []string, values string, array *[]string) {

	vals := strings.SplitN(values, ":", 2)
	applyCase := strings.ToLower(vals[0])
	joinWith := vals[1]

	fn1 := func(s string) string {
		return s
	}
	fn2 := fn1

	caseMap := map[string]func(string) string{
		"l": strings.ToLower,
		"u": strings.ToUpper,
		"t": strings.Title,
		"c": strings.Title,
	}
	if _, exists := caseMap[applyCase]; exists {
		fn2 = caseMap[applyCase]
		fn1 = caseMap[applyCase]
	}
	if applyCase == "c" {
		fn1 = strings.ToLower
	}

	for _, word := range words {
		str := []string{}

		if strings.Contains(word, "_") {
			str = strings.Split(word, "_")

		} else if strings.Contains(word, "-") {
			str = strings.Split(word, "-")

		} else {

			j := 0
			for i, letter := range word {
				if letter > 'A' && letter < 'Z' {
					str = append(str, word[j:i])
					j = i
				}
			}
			str = append(str, word[j:])
		}

		str[0] = fn1(str[0])
		for i, word := range str[1:] {
			str[i+1] = fn2(word)
		}
		*array = append(*array, strings.Join(str, joinWith))
	}
}
