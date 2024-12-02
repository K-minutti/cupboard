package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
    "strings"
)

var data = []string{
    "The only thing we have to fear is fear itself",
    "Ask not what your country can do for you; ask what you can do for your country",
    "Tis better to have loved and lost than never to have loved at all",
}

var Green = "\033[32m"

func main() {
    fmt.Println("Starting program...")

    rand.Seed(time.Now().UnixNano())
    randIdx := rand.Intn(len(data))
    text := data[randIdx]


    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

    go func() {
        <- sigChan
        fmt.Println("\nRecevied interrupt signal...")
        os.Exit(0)
    }()


    // matched := make([]bool, len(text))

    // printText := func(){
    //     for i, char := range text {
    //         if matched[i] {
    //             fmt.Printf("\033[32m%c\033[0m", char)
    //         } else {
    //             fmt.Printf("%c", char)
    //         }
    //     }
    //     fmt.Println()
    // }
    // printText()
    
    fmt.Println(text)

    for {
        key, err := readKeyStroke()
        if err != nil {
            fmt.Println("Error reading keystroke:", err)
            return
        }
        
        // fmt.Printf("%c", key)

        if strings.Contains(text, string(key)) {
            fmt.Printf("%s%c%s", Green, key, Green)
        } else {
            fmt.Printf("%c", key)
        }

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