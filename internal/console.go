package internal

import (
	"bufio"
	"errors"
	"fmt"
	"gpm/requests"
	"os"
	"strconv"
	"strings"
)

func ConsoleOutput(pkgName string, description string) (bool, error) {
	fmt.Println("=> The following package was found: " + pkgName)
	fmt.Println(description)
	fmt.Print("Insatall [ Y/n ] ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	switch input {
	case "":
		return true, nil
	case "y":
		return true, nil
	case "Y":
		return true, nil
	case "n":
		return false, nil
	case "N":
		return false, nil
	default:
		return false, nil
	}
}
func ConsoleOutputMultipleChoise(pkgNames []requests.PackageData) (int, error) {
	fmt.Println()
	fmt.Println("=> The following packages were found: ")
	fmt.Println()
	for i, finded_package := range pkgNames {
		fmt.Printf("[%v] Name:  %v \n", i+1, finded_package.PkgName)
		fmt.Printf("Description: %v\n\n", finded_package.PkgDescription)

	}
	fmt.Print("=> Provide install number : ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	input_int, err := strconv.Atoi(input)

	if err != nil {

		return -1, errors.New("not an integer")
	}
	for i := 0; i < len(pkgNames); i++ {
		if input_int == i {
			return input_int, nil
		}
	}
	return -1, errors.New("provide a valid index")
}
