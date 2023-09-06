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
	INSTALL                 = "i"
	INSTALL_DEEP            = "id"
	INSTALL_DEEP_MAX_LENGTH = "mx"
)

func main() {
	var COMMANDS_ARRAY = []string{INSTALL, INSTALL_DEEP, INSTALL_DEEP_MAX_LENGTH}

	args := os.Args[1:]
	recivedArgument, err := internal.ArgumentsFormatter(args, COMMANDS_ARRAY)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, commandStruct := range recivedArgument {
		switch commandStruct.Commad {
		case INSTALL:
			pkgName, description := requests.ParseHttpString(commandStruct.Value)
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
		case INSTALL_DEEP:
			maxResultLength := 4
			for _, subCommand := range recivedArgument {
				switch subCommand.Commad {
				case INSTALL_DEEP_MAX_LENGTH:
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
			pkgNames := requests.ParseHttpStringDeep(commandStruct.Value, maxResultLength)
			indexOfPackage, err := internal.ConsoleOutputMultipleChoise(pkgNames)
			if err != nil {
				log.Fatal(err)
			}
			internal.InstallPackage(pkgNames[indexOfPackage].PkgName)

			fmt.Println("Usage: import \"" + pkgNames[indexOfPackage].PkgName + "\"")
		}

	}
}
