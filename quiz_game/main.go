package main

import (
  "fmt"
  "io/ioutil"
	"strings"
  "container/list"
)

func Check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  dat, err := ioutil.ReadFile("problems.csv")
  Check(err)
  lines := strings.Split(string(dat), "\n")

  quiz := list.New()
  for _, line := range lines {
    if line != "" {
      split := strings.Split(string(line), ",")
      quiz.PushBack([2]string{split[0], split[1]})
    }
  }

	// Iterate through list and print its contents.
	for e := quiz.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
