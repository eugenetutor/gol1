package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fileName := "document.pdf"
	extension := filepath.Ext(fileName)
	if extension == ".jpg" || extension == ".png" || extension == ".gif"{
		fmt.Println("It's picture")
	}else if extension == ".mp4" || extension == ".mov" || extension == ".avi"{
		fmt.Println("It's video")
	}else if extension == ".mp3" || extension == ".wav" || extension == ".flac"{
		fmt.Println("It's music")
	}else if extension == ".doc" || extension == ".pdf" || extension == ".txt"{
		fmt.Println("It's document")
	}else {
		fmt.Println("Unknown")
	}
}
