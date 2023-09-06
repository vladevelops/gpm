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
	fmt.Println("Finded the following package: " + pkgName)
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
	fmt.Println("Finded the following packages: ")
	for i, finded_package := range pkgNames {
		fmt.Printf("[%v] Name:  %v \n", i, finded_package.PkgName)
		fmt.Printf("Description: %v\n\n", finded_package.PkgDescription)

	}
	fmt.Print("Whitch index to install [ x ] ")

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
	return -1, errors.New("not a valid index")
}
