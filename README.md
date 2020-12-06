# AdventOfCode2020

My personal repo containing solutions to the [Advent of Code 2020](https://adventofcode.com/2020) (AoC) programming event.
This year I chose to use this event to deepen my knowledge of and experience with the [Go](https://golang.org/) programming language.
Also because each challenge provides a short example with an expected result, I generally follow a 
[TDD](https://en.wikipedia.org/wiki/Test-driven_development) workflow where I write the test first and then write code to make it pass.
Generally my code is setup as follows:
* **dayDD(a|b).go** - main program file where DD corresponds to the day 01-25
* **dayDD(a|b)_test.go** - unit test
* **data.txt** - actual data provided by AoC
* **test.txt** - test/example data provided by AoC

The head of my programs are strucured as follows:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
  var result string
  
  // implementation here
	
  return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
```

Because AoC always provides text (which I put in a *.txt file) as input and
always accepts a string as the solution, by standard implementation
function always has the the function signature `doit(fileName string) string` 
where it accepts a file name as input and provides a string as output.  In
this way `main()` is always the same.

Likewise my standard testing file is:

```go
package main

import (
	"testing"
)

func Test(t *testing.T) {
	ans := doit("test.txt")
	exp := "expected answer"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}
```

Where from problem to problem only the the `exp := "expected answer"` line has to change.

