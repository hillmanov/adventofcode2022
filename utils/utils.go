package utils

import (
	"bufio"
	"embed"
	"strconv"
)

type Number interface {
	int | int64 | float64
}

func ReadLines(f embed.FS, filename string) ([]string, error) {
	file, err := f.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return s, nil
}

func ReadNumbers[T Number](f embed.FS, filename string) ([]T, error) {
	file, err := f.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	i := []T{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		i = append(i, T(n))
	}

	return i, nil
}

func ReadContents(f embed.FS, filename string) (string, error) {
	contents, err := f.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func ReplaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func ParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MinOf[T Number](numbers []T) T {
	var min = numbers[0]
	for _, value := range numbers {
		if min > value {
			min = value
		}
	}
	return min
}

func MaxOf[T Number](numbers []T) T {
	var max = numbers[0]
	for _, value := range numbers {
		if max < value {
			max = value
		}
	}
	return max
}

func MinMax[T Number](numbers []T) (T, T) {
	var max = numbers[0]
	var min = numbers[0]
	for _, value := range numbers {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func SumOf[T Number](numbers []T) T {
	sum := T(0)
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Abs[T Number](n T) T {
	if n < 0 {
		return n * -1
	}
	return n
}

func UniqueOf[T comparable](collection []T) []T {
	keys := make(map[T]bool)
	unique := []T{}
	for _, entry := range collection {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique = append(unique, entry)
		}
	}
	return unique
}

func CopyOf[T any](collection []T) []T {
	copf := make([]T, len(collection))
	copy(copf, collection)
	return copf
}

func IndexOf[T comparable](haystack []T, needle T) int {
	for index, value := range haystack {
		if value == needle {
			return index
		}
	}
	return -1
}

func Intersection[T comparable](a, b []T) []T {
	var result []T
	for _, x := range a {
		if IndexOf(b, x) != -1 {
			result = append(result, x)
		}
	}
	return UniqueOf(result)
}

func Reverse[T any](collection []T) []T {
	for i := 0; i < len(collection)/2; i++ {
		j := len(collection) - i - 1
		collection[i], collection[j] = collection[j], collection[i]
	}
	return collection
}

func Pop[T any](collection []T) (T, []T) {
	return collection[0], collection[1:]
}

func Shift[T any](collection []T, value T) []T {
	collection = append([]T{value}, collection...)
	return collection
}
