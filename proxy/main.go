package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/proxy/youtube"
)

func main() {
	y := youtube.NewYoutubeCached()

	fmt.Println("should retrieve from source")
	fmt.Println(y.ListVideos())

	fmt.Println("\nshould retrieve from cache")
	fmt.Println(y.ListVideos())
}
