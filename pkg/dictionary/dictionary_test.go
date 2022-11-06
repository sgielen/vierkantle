package dictionary_test

import (
	"fmt"
	"math/rand"
	"strings"
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

type mockReader struct {
	r       *rand.Rand
	letters []rune
}

func (m *mockReader) ReadWord() string {
	length := m.r.Intn(12)
	res := strings.Builder{}
	for ; length > 0; length-- {
		res.WriteRune(m.letters[m.r.Intn(len(m.letters))])
	}
	return res.String()
}

var seeds = []int64{0, 1000, 2000, 5000}

func BenchmarkDictionary(b *testing.B) {
	for _, seed := range seeds {
		reader := &mockReader{
			r: rand.New(rand.NewSource(seed)),
			letters: []rune{
				'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
				'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			},
		}
		b.Run(fmt.Sprintf("seed_%d", seed), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				dictionary := dictionary.NewPrefixDictionary(reader, 4, 12)
				dictionary.HasWord("dictionaries")
			}
		})
	}
}
