package main

import "io/ioutil"
import "fmt"

func main() {
    files,_ := ioutil.ReadDir("/Users/softeam/Downloads")
    fmt.Println(len(files))
}
