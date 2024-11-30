package main

import (
	"fmt"
    "time"
    "math/rand"
	"syscall"
)

var data = []string{
    "The only thing we have to fear is fear itself",
    "Ask not what your country can do for you; ask what you can do for your country",
    "Tis better to have loved and lost than never to have loved at all",
}


func main() {
    fmt.Println("Starting program...")

    rand.Seed(time.Now().UnixNano())
    randIdx := rand.Intn(len(data))

    text := data[randIdx]

    fmt.Println(text)
    
    for {
        key, err := readKeyStroke()
        if err != nil {
            fmt.Println("Error reading keystroke:", err)
            return
        }
        fmt.Printf("%c", key)
    }
}

func readKeyStroke() (byte, error) {
    var buf [1]byte
    _, err := syscall.Read(syscall.Stdin, buf[:])
    if err != nil {
        return 0, err
    }
    return buf[0], nil

}