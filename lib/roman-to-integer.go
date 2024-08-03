package lib

import "fmt"

func Hello(s string) string {
	text := fmt.Sprintf("Hello my name is %s", s)

	return text
}
func RomanToInt(s string) int {
	roman_dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	letter_to_int := []int{}
	for _, v := range s {
		converted := roman_dict[string(v)]
		letter_to_int = append(letter_to_int, converted)
	}

	skip_count := 0

	for i := 0; i < len(s); i++ {
		if skip_count > 0 {
			skip_count--
			continue
		}

		if i == len(letter_to_int)-1 {
			sum += letter_to_int[i]
			break
		}

		if letter_to_int[i] > letter_to_int[i+1] || letter_to_int[i] == letter_to_int[i+1] {
			sum += letter_to_int[i]
		} else if letter_to_int[i] < letter_to_int[i+1] {
			sum += letter_to_int[i+1] - letter_to_int[i]
			skip_count = 1
		}
	}
	return sum

}
