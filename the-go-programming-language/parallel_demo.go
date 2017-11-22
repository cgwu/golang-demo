package main

import (
	"fmt"
	"time"
	"log"
	"sync"
	"os"
)

func processImgThumbnail(img string) {
	//time.Sleep(time.Millisecond * 50)
	time.Sleep(time.Second * 5)
	//fmt.Println("processed img:", img)
	log.Println("processed img:", img)
}

func main() {
	//var imgFiles []string
	imgFiles := make([]string, 0, 100)
	for x := 0; x < 100; x++ {
		imgFiles = append(imgFiles, fmt.Sprintf("img_%d.jpg", x))
	}
	log.Println("img files number:", len(imgFiles))

	ch := make(chan struct{})
	for _, f := range imgFiles {
		//fmt.Println(i, v)
		go func(f string) {
			processImgThumbnail(f)
			ch <- struct{}{}
		}(f)
	}
	// Wait for goroutines to complete
	for range imgFiles {
		//log.Println(i, f)
		<-ch
	}
	log.Println("main done")
}


/*
// 非常有启发!
// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
*/