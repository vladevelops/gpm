package internal

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// modified version of this example - https://medium.com/@j.d.livni/create-a-load-bar-in-go-f158837ff4c4
//

func printProgressBar(prefix string, length int) {
	total := 100
	fill := "-"
	for i := 0; i < total; i++ {
		iteration := i + 1
		filledLength := int(length * iteration / total)
		end := "*"
		if iteration == total {
			end = "-"
		}
		bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
		fmt.Printf("\r%s [%s]", prefix, bar)

		time.Sleep(15 * time.Millisecond)
	}
}

var wt sync.WaitGroup

func InstallPackage(pkgNameFullPath string) {
	fmt.Println("Installing....")
	goGet := exec.Command("go", "get", pkgNameFullPath)
	wt.Add(1)
	go func() {
		result := goGet.Run()

		if result == nil {
			wt.Done()
			fmt.Println()
		} else {
			wt.Done()
			log.Fatal("go get command failed")
		}
	}()
	go func() {
		for {
			printProgressBar("Progress", 25)
		}
	}()

	wt.Wait()
}
