package main

import (
	"github.com/sy9777m/learnGo/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	word := "hello"
	errAdd := dictionary.Add(word, "first")
	errUpdate := dictionary.Update(word, "Second")
	search, _ := dictionary.Search(word)

}
