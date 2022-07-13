# 71A. Way Too Long Words

## Problem

Sometimes some words like "localization" or "internationalization" are so long that writing them many times in one text is quite tiresome.

Let's consider a word too long, if its length is strictly more than 10 characters. All too long words should be replaced with a special abbreviation.

This abbreviation is made like this: we write down the first and the last letter of a word and between them we write the number of letters between the first and the last letters. That number is in decimal system and doesn't contain any leading zeroes.

Thus, "localization" will be spelt as "l10n", and "internationalization» will be spelt as "i18n".

You are suggested to automatize the process of changing the words with abbreviations. At that all too long words should be replaced by the abbreviation and the words that are not too long should not undergo any changes.

[Link to CodeForces](https://codeforces.com/problemset/problem/71/A)

## Input

The first line contains an integer $n$ ($1 \leq n \leq 100$). Each of the following n lines contains one word. All the words consist of lowercase Latin letters and possess the lengths of from 1 to 100 characters.

## Output

Print $n$ lines. The $i$-th line should contain the result of replacing of the $i$-th word from the input data.


## Examples

### Input

```
4
word
localization
internationalization
pneumonoultramicroscopicsilicovolcanoconiosis
```

### Output

```
word
l10n
i18n
p43s
```

## Constraints

  - Time limit per test: 1 second
  - Memory limit per test: 256 megabytes
  - Input: standard input
  - Output: standard output

## Solutions

### GNU C++17 7.3.0

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|  71A-14 | GNU C++17 | Accepted | 15 ms |  4 KB  |

[Link to source code](solution.cpp)

```c++
#include <iostream>


void solution() {
    int case_count = 0;
    std::cin >> case_count;

    std::string word;
    for (int case_index = 0; case_index < case_count; ++case_index) {
        std::cin >> word;
        if (word.size() <= 10) {
            std::cout << word;
        } else {
            std::cout << *word.begin() << word.length() - 2 <<*word.rbegin();
        }

        std::cout << '\n';
    }
}


void setup() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(NULL);
}


int main() {
    setup();
    solution();
}
```

### C# Mono 6.8

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|  71A-14 |  Mono C#  | Accepted | 31 ms |  0 KB  |

[Link to source code](solution.cs)

```c#
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


namespace Codeforces
{
    class Program
    {
        static void Solution()
        {
            int caseCount = Convert.ToInt32(System.Console.ReadLine());
            for (int caseIndex = 0; caseIndex < caseCount; ++caseIndex)
            {
                String word = System.Console.ReadLine();
                if (word.Length <= 10)
                {
                    System.Console.WriteLine(word);
                }
                else
                {
                    System.Console.WriteLine("{0}{1}{2}", word[0], word.Length - 2, word[word.Length - 1]);
                }
            }
        }

        static void Main(string[] args)
        {
            Solution();
        }
    }
}
```

### Go 1.17.5

| Problem |    Lang   |  Verdict | Time  | Memory |
|:-------:|:---------:|:--------:|:-----:|:------:|
|  71A-14 |    Go     | Accepted | 15 ms | 60 KB  |

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

func Printf(format string, x ...interface{}) {
	fmt.Fprintf(writer, format, x...)
}

func Scanf(format string, x ...interface{}) {
	fmt.Fscanf(reader, format, x...)
}

func Solution() {
	var caseCount int
	Scanf("%d\n", &caseCount)

	var word string
	for caseIndex := 0; caseIndex < caseCount; caseIndex++ {
		Scanf("%s\n", &word)
		if len(word) <= 10 {
			Printf("%s\n", word)
		} else {
			Printf("%c%d%c\n", word[0], len(word)-2, word[len(word)-1])
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
