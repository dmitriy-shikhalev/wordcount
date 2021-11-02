package main

import (
    "errors"
//    "flag"
    "fmt"
    "os"
    "strings"
//    "bufio"
)

func main() {
    words, err := readInput()
    if err != nil {
        fail(err)
    }
    
    fmt.Println(len(words))
}

// match returns true if src matches pattern,
// false otherwise.
func match(pattern string, src string) bool {
    return strings.Contains(src, pattern)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() ([]string, error) {
    //reader := bufio.NewReader(os.Stdin)
    //fmt.Print("Enter text: ")
    //text, _ := reader.ReadString('\n')
    //var words = strings.Fields(text)
    if len(os.Args) > 1 {
      s := os.Args[1]
      words := strings.Fields(s)
      return words, nil
    } else {
      return nil, errors.New("Uncorrect input")
    }
}

// fail prints the error and exits.
func fail(err error) {
    fmt.Errorf("%v", err)
}
