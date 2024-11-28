package main

import (
	"fmt"
	"syscall"
)

func main() {

    fmt.Println("Starting program...")

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