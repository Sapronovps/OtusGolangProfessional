package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func Top10(input string) []string {
	words := findWords(input)
	frequencyWords := make(map[string]int)

	for i := range words {
		frequencyWords[words[i]]++
	}

	return sortWordsByKeyAndValue(frequencyWords)
}

func Top10NotSensetive(input string) []string {
	words := findWords(input)
	frequencyWords := make(map[string]int)

	for i := range words {
		if words[i] == "-" {
			continue
		}
		frequencyWords[strings.ToLower(RemovePunctuationMark(words[i]))]++
	}

	return sortWordsByKeyAndValue(frequencyWords)
}

// Удаление знаков пунктуации по краям слова.
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

// Сортировка слов по ключам и значениям.
func sortWordsByKeyAndValue(frequencyWords map[string]int) []string {
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

// Поиск слов в заданной строке.
func findWords(input string) []string {
	re := regexp.MustCompile(`[^\s]+`)
	return re.FindAllString(input, -1)
}
