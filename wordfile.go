package randomizer

import (
	"bufio"
	"bytes"
	"math/rand"
	"os"
	"text/scanner"
	"time"
	"unicode"
)

type WordFile struct {
	size       int
	words      []string
	exclusions map[string]struct{}
}

func (this *WordFile) Open(path string) []string {

	this.words = make([]string, 1000)
	if f, err := os.Open(path); err != nil {
		panic(err)
	} else {

		s := bufio.NewScanner(f)
		s.Split(bufio.ScanWords)
		var wordCount int
		for wordCount = 0; s.Scan() != false; wordCount++ {
			//fmt.Printf("%s\n", s.Text())
			text := s.Text()
			if len(text) > 3 {
				this.words = append(this.words, s.Text())
			}
		}
	}
	this.size = len(this.words)
	return this.words
}

func (this *WordFile) unpunctuate(input string) string {

	b := bytes.NewBuffer([]byte(input))
	s := new(scanner.Scanner).Init(b)
	s.Mode = scanner.ScanChars

	chars := make([]rune, 0)
	for _, c := range input {
		if unicode.IsLetter(c) {
			chars = append(chars, c)
		}
	}

	return string(chars)
}

func (this *WordFile) Exclude(exclude ...string) *WordFile {
	for _, x := range exclude {
		this.exclusions[x] = struct{}{}
	}
	return this
}

func (this *WordFile) Count(words []string) map[string]int {
	return make(map[string]int)
}

func (this *WordFile) Total(words []string) int {
	return this.size
}

func (this *WordFile) Random() string {
	if this.size == 0 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return this.words[r.Intn(this.size)]
}
