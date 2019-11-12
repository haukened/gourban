# gourban [![Build Status](https://travis-ci.com/haukened/gourban.svg?branch=master)](https://travis-ci.com/haukened/gourban) [![GoDoc](https://godoc.org/github.com/haukened/gourban?status.svg)](https://godoc.org/github.com/haukened/gourban)   
GoLang Package for querying UrbanDictionary with minimal dependencies.

gourban provides a golang interface to the urbandictionary API.  the Query(string) function returns a slice of all definitions from urbandictionary.com, and accepts single words or phrases like ["yeet", "yolo", "ok boomer"].

## Simple use:
```
import (
    "fmt"
    
    "github.com/haukened/gourban"
)

func main() {
    ...
    definitions := gourban.Query("yeet")
    for def := range definitions {
        fmt.Printf("%s: %s\n", def.Word, def.Description)
    }
}
```
