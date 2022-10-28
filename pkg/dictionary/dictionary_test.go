package dictionary_test

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/sgielen/vierkantle/pkg/dictionary"
)

type wordReader struct {
	Words []string
}

func (r *wordReader) ReadWord() string {
	if len(r.Words) == 0 {
		return ""
	}
	var word string
	word, r.Words = r.Words[0], r.Words[1:]
	return word
}

func TestDictionary(t *testing.T) {
	reader := wordReader{Words: []string{"hello", "hell", "hellos", "potato"}}
	dict := dictionary.NewPrefixDictionary(&reader, 4, 8)

	if diff := deep.Equal(dict.HasWord("a"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("ah"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("ahe"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("l"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}

	if diff := deep.Equal(dict.HasWord("h"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("he"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hel"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hell"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        true,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hello"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        true,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hellos"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        true,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}

	if diff := deep.Equal(dict.HasWord("p"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potat"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potato"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        true,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potatos"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        false,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
}
