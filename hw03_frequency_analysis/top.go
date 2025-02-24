package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"unicode"
)

func Top10(input string, isSensetive bool) []string {
	re := regexp.MustCompile(`[^\s]+`)
	words := re.FindAllString(input, -1)
	frequencyWords := make(map[string]int)

	for i := 0; i < len(words); i++ {
		if isSensetive {
			frequencyWords[words[i]]++
		} else {
			if words[i] == "-" {
				continue
			}
			frequencyWords[FirstLetterToLower(RemovePunctuationMark(words[i]))]++
		}
	}

	// Шаг 1: Получаем ключи карты
	keys := make([]string, 0, len(frequencyWords))
	for k := range frequencyWords {
		keys = append(keys, k)
	}

	// Шаг 2: Сортируем ключи на основе значений
	sort.Slice(keys, func(i, j int) bool {
		if frequencyWords[keys[i]] == frequencyWords[keys[j]] {
			return keys[i] < keys[j] // Сортировка по ключам, если значения равны
		}
		return frequencyWords[keys[i]] > frequencyWords[keys[j]] // Сортировка по возрастанию значений
	})

	top10 := make([]string, 0, 10)
	for _, k := range keys {
		if (len(top10)) == 10 {
			break
		}
		top10 = append(top10, k)
	}

	return top10
}

func FirstLetterToLower(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func RemovePunctuationMark(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	if !unicode.IsLetter(r[0]) {
		r = r[1:]
	}
	rLen := len(r)
	if !unicode.IsLetter(r[rLen-1]) {
		r = r[:rLen-1]
	}

	return string(r)
}
