package sortnumASC

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func SortB(number string) (string, error) {

	if _, err := strconv.ParseFloat(number, 64); err != nil {
		fmt.Printf("%q is not a number.\n", number)
		return "", errors.New("error, please input number only")
	}

	newNumber := []rune(number)
	sort.Slice(newNumber, func(i int, j int) bool { return newNumber[i] < newNumber[j] })
	result := string(newNumber)
	return result, nil
}
