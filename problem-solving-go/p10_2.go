package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}
	mem := make([]bool, len(s)+1)
	mem[0] = true
	for i := 1; i < len(mem); i++ {
		for j := i - 1; j >= 0; j-- {
			if mem[j] && dict[s[j:i]] {
				mem[i] = true
				break
			}
		}
	}
	return mem[len(mem)-1]
}

func main() {
	s := "IAMHERE"
	dict := []string{"I", "A", "AM", "HE", "HERE"}

	sentences := wordBreak(s, dict)

	fmt.Println(sentences)

}
