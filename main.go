package main

import (
    "fmt"
    "flag"
)

func main() {
    var input string
    flag.StringVar(&input, "i", "", "input name")
    flag.Parse()
    
    output := fmt.Sprintf("Hello %s :)", input)
    fmt.Println(output)
}
