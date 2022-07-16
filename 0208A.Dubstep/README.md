# 208A. Dubstep

## Problem

Vasya works as a DJ in the best Berland nightclub, and he often uses dubstep music in his performance. Recently, he has decided to take a couple of old songs and make dubstep remixes from them.

Let's assume that a song consists of some number of words. To make the dubstep remix of this song, Vasya inserts a certain number of words "WUB" before the first word of the song (the number may be zero), after the last word (the number may be zero), and between words (at least one between any pair of neighbouring words), and then the boy glues together all the words, including "WUB", in one string and plays the song at the club.

For example, a song with words "I AM X" can transform into a dubstep remix as "WUBWUBIWUBAMWUBWUBX" and cannot transform into "WUBWUBIAMWUBX".

Recently, Petya has heard Vasya's new dubstep track, but since he isn't into modern music, he decided to find out what was the initial song that Vasya remixed. Help Petya restore the original song.

[Link to CodeForces](https://codeforces.com/problemset/problem/208/A)

## Input

The input consists of a single non-empty string, consisting only of uppercase English letters, the string's length doesn't exceed 200 characters. It is
guaranteed that before Vasya remixed the song, no word contained substring "WUB" in it; Vasya didn't change the word order. It is also guaranteed that initially
the song had at least one word.

## Output

Print the words of the initial song that Vasya used to make a dubsteb remix. Separate the words with a space.

## Examples

### Input

```
WUBWUBABCWUB
```

### Output

```
ABC
```

### Input

```
WUBWEWUBAREWUBWUBTHEWUBCHAMPIONSWUBMYWUBFRIENDWUB
```

### Output

```
WE ARE THE CHAMPIONS MY FRIEND
```

## Note

In the first sample: "WUBWUBABCWUB" = "WUB" + "WUB" + "ABC" + "WUB". That means that the song originally consisted of a single word "ABC", and all words "WUB" were added by Vasya.

In the second sample Vasya added a single word "WUB" between all neighbouring words, in the beginning and in the end, except for words "ARE" and "THE" — between them Vasya added two "WUB".

## Solutions

### Constraints

  - Time limit per test: 2 seconds
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

### Go 1.17.5

|  Problem  |    Lang   |  Verdict | Time  | Memory |
|:---------:|:---------:|:--------:|:-----:|:------:|
| 208A - 18 |    Go     | Accepted | 60 ms | 76 KB  |

[Link to source code](solution.go)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func Fprintf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Fscanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var song string
	Fscanf("%s\n", &song)

	isNewWord := true
	isSongStarted := false

	i := 0
	for i < len(song) {
		if i < len(song)-2 && song[i] == 'W' && song[i+1] == 'U' && song[i+2] == 'B' {
			i += 3
			isNewWord = true
		} else {
			if isNewWord {
				if isSongStarted {
					Fprintf(" ")
				} else {
					isSongStarted = true
				}

				isNewWord = false
			}

			Fprintf("%c", song[i])
			i++
		}
	}
	Fprintf("\n")
}

func main() {
	Solution()
	writer.Flush()
}
```
