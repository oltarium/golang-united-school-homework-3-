package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int
	for key, _ := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		result = append(result, input[key])
	}
	return
}
