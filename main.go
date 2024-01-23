package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"

	"github.com/RoaringBitmap/roaring/roaring64"
)

func main() {
	var items []uint64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		item, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			panic(err)
		}

		items = append(items, item)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	bitmap := roaring64.NewBitmap()
	for _, item := range items {
		bitmap.Add(item)
	}

	var bitmapBuf bytes.Buffer
	if _, err := bitmap.WriteTo(&bitmapBuf); err != nil {
		panic(err)
	}

	readBitmap := roaring64.NewBitmap()
	if _, err := readBitmap.ReadFrom(&bitmapBuf); err != nil {
		panic(err)
	}
}
