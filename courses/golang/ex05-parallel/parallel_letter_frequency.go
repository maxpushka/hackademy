package letter

import "sync"

func Frequency(text string) map[rune]uint {
	frequencies := make(map[rune]uint)
	for _, ch := range text {
		frequencies[ch]++
	}
	return frequencies
}

func mergeMaps(maps ...map[rune]uint) map[rune]uint {
	res := make(map[rune]uint)

	for _, m := range maps {
		for key, value := range m {
			res[key] += value
		}
	}

	return res
}

func ConcurrentFrequency(texts []string) map[rune]uint {
	var wg sync.WaitGroup
	res := make([]map[rune]uint, len(texts))
	for index, text := range texts {
		wg.Add(1)
		go func(index int, text string) {
			res[index] = Frequency(text)
			wg.Done()
		}(index, text)
	}
	wg.Wait()
	return mergeMaps(res...)
}
