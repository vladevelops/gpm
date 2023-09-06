package internal

import (
	"fmt"
	"os/exec"
	"time"
)

func InstallPackage(pkgNameFullPath string) {
	fmt.Println("install requested")
	goGet := exec.Command("go", "get", pkgNameFullPath)

	ch := make(chan error)
	printProgress := true
	go func() {
		ch <- goGet.Run()
		printProgress = false
	}()
	go func() {

		counter := 0
		fmt.Print("\033[s") // save the cursor position
		for {

			fmt.Print("\033[u\033[K")
			if !printProgress {
				break
			}
			if counter == 50 {
				counter = 0
			}
			renderString := ""
			for i := 0; i < counter; i++ {
				renderString += "="
			}
			counter++
			fmt.Print("|" + renderString + ">")
			time.Sleep(15 * time.Microsecond)

		}
	}()
	<-ch
	fmt.Print("\033[u\033[K")
}
