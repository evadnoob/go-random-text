[//]: -*- mode: gfm -*-


a utility for generating random data, sort of.

Example:

```
	wordfile := &WordFile{exclusions: make(map[string]struct{})}
	wordfile.Exclude("the")
	wordfile.Open("war_and_peace_leo_tolstoy.txt")

	word := wordfile.Random()

	fmt.Printf("%s\n", word)


```
