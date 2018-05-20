package main

import (
    "log"
    "os"
    "fmt"
)

func main() {
    file, err := os.Create("result.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()


    os.Setenv("url1", "www.bbc.com")
    os.Setenv("url2", "www.bing.com")
    os.Setenv("url3", "www.sky.com")

    fmt.Printf("%#v\n", os.Getenv("url1"))
    fmt.Printf("%#v\n", os.Getenv("url2"))
    fmt.Printf("%#v\n", os.Getenv("url3"))

}
