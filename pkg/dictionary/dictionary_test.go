package dictionary_test

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/sgielen/vierkantle/pkg/dictionary"
)

type wordReader struct {
	Words []string
}

func (r *wordReader) ReadWord() dictionary.WordReadResult {
	if len(r.Words) == 0 {
		return dictionary.WordReadResult{}
	}
	var word string
	word, r.Words = r.Words[0], r.Words[1:]
	return dictionary.WordReadResult{Word: word}
}

func TestDictionary(t *testing.T) {
	reader := wordReader{Words: []string{"hello", "hell", "hellos", "potato"}}
	dict := dictionary.NewPrefixDictionary()
	dict.Read(&reader, dictionary.NormalWord, false)

	if diff := deep.Equal(dict.HasWord("a"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("ah"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("ahe"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("l"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}

	if diff := deep.Equal(dict.HasWord("h"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("he"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hel"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hell"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NormalWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hello"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NormalWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hellos"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NormalWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}

	if diff := deep.Equal(dict.HasWord("p"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potat"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potato"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NormalWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potatos"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
}

func TestDictionaryWordTypes(t *testing.T) {
	dict := dictionary.NewPrefixDictionary()
	reader := wordReader{Words: []string{"hello", "hell", "hellos", "potato"}}
	dict.Read(&reader, dictionary.BonusWord, false)
	reader = wordReader{Words: []string{"hello", "pot", "foo"}}
	dict.Read(&reader, dictionary.NormalWord, false)
	reader = wordReader{Words: []string{"potato", "foo", "mumble"}}
	dict.Read(&reader, dictionary.SensitiveWord, true /* upgrade only */)
	if diff := deep.Equal(dict.HasWord("hello"), &dictionary.HasWordsWithPrefixResult{
		// normal goes over bonus
		HasThisWord:        dictionary.NormalWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("hell"), &dictionary.HasWordsWithPrefixResult{
		// only a bonus word
		HasThisWord:        dictionary.BonusWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("potato"), &dictionary.HasWordsWithPrefixResult{
		// sensitive goes over bonus
		HasThisWord:        dictionary.SensitiveWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("foo"), &dictionary.HasWordsWithPrefixResult{
		// sensitive goes over normal
		HasThisWord:        dictionary.SensitiveWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("mum"), &dictionary.HasWordsWithPrefixResult{
		// upgrade only
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("mumble"), &dictionary.HasWordsWithPrefixResult{
		// upgrade only
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
}

type mockReader struct {
	r       *rand.Rand
	letters []rune
}

func (m *mockReader) ReadWord() dictionary.WordReadResult {
	length := m.r.Intn(12)
	res := strings.Builder{}
	for ; length > 0; length-- {
		res.WriteRune(m.letters[m.r.Intn(len(m.letters))])
	}
	return dictionary.WordReadResult{
		Word: res.String(),
	}
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
				dict := dictionary.NewPrefixDictionary()
				dict.Read(reader, dictionary.NormalWord, false)
				dict.HasWord("dictionaries")
			}
		})
	}
}

func TestDictionaryForceType(t *testing.T) {
	fh, err := os.CreateTemp(os.TempDir(), "vierkantletest")
	if err != nil {
		t.Fatal(err)
	}
	tempfile := fh.Name()

	if err := dictionary.AddForceTypeToFile(tempfile, "testno", dictionary.NoWord); err != nil {
		t.Fatal(err)
	}
	if err := dictionary.AddForceTypeToFile(tempfile, "testnorm", dictionary.NormalWord); err != nil {
		t.Fatal(err)
	}
	if err := dictionary.AddForceTypeToFile(tempfile, "testbon", dictionary.BonusWord); err != nil {
		t.Fatal(err)
	}
	if err := dictionary.AddForceTypeToFile(tempfile, "testsens", dictionary.SensitiveWord); err != nil {
		t.Fatal(err)
	}

	if err := dictionary.AddForceTypeToFile(tempfile, "test2bon", dictionary.NormalWord); err != nil {
		t.Fatal(err)
	}
	if err := dictionary.AddForceTypeToFile(tempfile, "test2bon", dictionary.BonusWord); err != nil {
		t.Fatal(err)
	}

	dict := dictionary.NewPrefixDictionary()
	if err := dict.ReadForceTypeFromFile(tempfile); err != nil {
		t.Fatal(err)
	}

	if diff := deep.Equal(dict.HasWord("testno"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NoWord,
		HasWordsWithPrefix: true,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("testnorm"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.NormalWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("testbon"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.BonusWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("testsens"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.SensitiveWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(dict.HasWord("test2bon"), &dictionary.HasWordsWithPrefixResult{
		HasThisWord:        dictionary.BonusWord,
		HasWordsWithPrefix: false,
	}); diff != nil {
		t.Fatal(diff)
	}
}
