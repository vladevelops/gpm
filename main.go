package main

import (
	"fmt"
	"gpm/internal"
	"gpm/requests"
	"log"
	"os"
	"strconv"
)

const (
	INSTALL         = "i"
	INSTALL_SEARCH  = "is"
	SEARCH_CAPACITY = "c"
)

func main() {
	var COMMANDS_ARRAY = []string{INSTALL, INSTALL_SEARCH, SEARCH_CAPACITY}

	args := os.Args[1:]
	recivedArgument, err := internal.ArgumentsFormatter(args, COMMANDS_ARRAY)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, commandStruct := range recivedArgument {
		switch commandStruct.Commad {
		case INSTALL:
			pkgNames := requests.ParseHttpString(commandStruct.Value, 1)
			result := pkgNames[0]
			pkgName, description := result.PkgName, result.PkgDescription
			isConfirmed, err := internal.ConsoleOutput(pkgName, description)
			if err != nil {
				log.Fatal("Aborted")
			}
			if isConfirmed {
				internal.InstallPackage(pkgName)
				fmt.Println("Usage: import \"" + pkgName + "\"")

			} else {
				log.Fatal("Aborted")
			}
		case INSTALL_SEARCH:
			maxResultLength := 4
			for _, subCommand := range recivedArgument {
				switch subCommand.Commad {
				case SEARCH_CAPACITY:
					input_int, err := strconv.Atoi(subCommand.Value)

					if err != nil {

						fmt.Println("Warning: mx command must be a number, system defaulted to 4")
					} else {
						if input_int <= 0 {
							fmt.Println("Warning: mx command must be a number greater then 0, system defaulted to 4")
						} else {
							maxResultLength = input_int
						}
					}

				}
			}
			pkgNames := requests.ParseHttpString(commandStruct.Value, maxResultLength)
			indexOfPackage, err := internal.ConsoleOutputMultipleChoise(pkgNames)
			if err != nil {
				log.Fatal(err)
			}
			internal.InstallPackage(pkgNames[indexOfPackage].PkgName)

			fmt.Println("Usage: import \"" + pkgNames[indexOfPackage].PkgName + "\"")
		}

	}
}
