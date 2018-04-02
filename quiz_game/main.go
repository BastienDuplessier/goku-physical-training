package main

import (
  "fmt"
  "io/ioutil"
	"strings"
  "container/list"
  "os"
)

func Check(e error) {
  if e != nil {
    panic(e)
  }
}

func DetermineFilename()string {
  args := os.Args[1:]

  for _, arg := range args {
    if strings.HasPrefix(arg, "--quizfile") {
      return strings.Split(arg, "=")[1]
    }
  }
  return "problems.csv"
}

func main() {
  dat, err := ioutil.ReadFile(DetermineFilename())
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
