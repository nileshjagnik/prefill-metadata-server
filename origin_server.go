package main

import (
	"fmt"
	"github.com/amahi/go-metadata"
	"time"
)

const size int = 1000000

func main() {
	Lib, err := metadata.Init(size, "metadata.db")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		duration, _ := time.ParseDuration("300ms")
		Lib.Prefill("./example_media_folder/movie", "movie", duration, false)
		Lib.Prefill("./example_media_folder/tv", "tv", duration, false)
		data, err := Lib.GetMetadata("Argo", "movie")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(data)
		}
	}
}
