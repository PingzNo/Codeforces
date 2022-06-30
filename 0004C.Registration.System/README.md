# 4C. Registration system

## Problem

A new e-mail service "Berlandesk" is going to be opened in Berland in the near future. The site administration wants to launch their project as soon as possible, that's why they ask you to help. You're suggested to implement the prototype of site registration system. The system should work on the following principle.

Each time a new user wants to register, he sends to the system a request with his name. If such a name does not exist in the system database, it is inserted into the database, and the user gets the response OK, confirming the successful registration. If the name already exists in the system database, the system makes up a new user name, sends it to the user as a prompt and also inserts the prompt into the database. The new name is formed by the following rule. Numbers, starting with 1, are appended one after another to name (name1, name2, ...), among these numbers the least i is found so that namei does not yet exist in the database.

[Link to CodeForces](https://codeforces.com/problemset/problem/4/C)

## Input

The first line contains number $n$ ($1 \leq n \leq 105$). The following $n$ lines contain the requests to the system. Each request is a non-empty line, and consists of not more than 32 characters, which are all lowercase Latin letters.

## Output

Print $n$ lines, which are system responses to the requests: OK in case of successful registration, or a prompt with a new name, if the requested name is already taken.

## Examples

### Input

```
4
abacaba
acaba
abacaba
acab
```

### Output

```
OK
OK
abacaba1
OK
```

### Input

```
6
first
first
second
second
third
third
```

### Output

```
OK
first1
OK
second1
OK
third1
```

## Solutions

### Constraints

  - Time limit per test: 5 second
  - Memory limit per test: 64 megabytes
  - Input: standard input
  - Output: standard output

### GNU C++17 7.3.0

| Problem |    Lang   |  Verdict |  Time  | Memory  |
|:-------:|:---------:|:--------:|:------:|:-------:|
| 4C - 11 |    Go     | Accepted | 248 ms | 3512 KB |

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
	var nameCount int
	Fscanf("%d\n", &nameCount)

	names := map[string]int{}

	var name string
	for nameIndex := 0; nameIndex < nameCount; nameIndex++ {
		Fscanf("%s\n", &name)
		if count, nameFound := names[name]; !nameFound {
			names[name] = 1
			Fprintf("OK\n")
		} else {
			names[name]++
			Fprintf("%s%d\n", name, count)
		}
	}
}

func main() {
	Solution()
	writer.Flush()
}
```
