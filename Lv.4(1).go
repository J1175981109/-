package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fp, err:= os.Create("proverb.txt")

	if err != nil {
		fmt.Println("文件创建失败。")
		return
	}

	bytes := []byte("don't communicate by sharing memory share memory by communicating")

	bw, err := fp.Write(bytes)

	log.Printf("Wrote %d bytes.\n", bw)

	fp.Close()
}