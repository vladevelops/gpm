package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getStringFromHttp(queryParam string, method string) (string, error) {
	url := "https://pkg.go.dev/search?q=" + queryParam + "&m="
	// Method Not Allowed
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
func getRawHTML(queryParam string) string {
	rawHtmlStringPost, err := getStringFromHttp(queryParam, "POST")
	if strings.TrimSpace(rawHtmlStringPost) == "Method Not Allowed" {
		rawHtmlStringPost, err = getStringFromHttp(queryParam, "GET")
	}
	if err != nil {
		fmt.Println(err)
	}

	return rawHtmlStringPost
}

type PackageData struct {
	PkgName        string
	PkgDescription string
}

func ParseHttpString(queryParam string, maxResultLength int) []PackageData {

	rawHtmlString := getRawHTML(queryParam)

	getLinesFromString := strings.Split(rawHtmlString, "\n")

	var packageOptions []PackageData
	var resultPackage, description string
	for i := range getLinesFromString {
		// search result

		searchResultFinded := strings.Contains(getLinesFromString[i], "search result")
		extra_hrferFinded := strings.Contains(getLinesFromString[i], "href")
		if searchResultFinded && extra_hrferFinded {
			var cleanSpaces string = strings.TrimSpace(getLinesFromString[i])
			resultPackage = strings.Split(cleanSpaces, "\"")[1][1:]
		}

		searchResultDescription := strings.Contains(getLinesFromString[i], "SearchSnippet-synopsis")

		if searchResultDescription {
			description = strings.TrimSpace(getLinesFromString[i+1])
		}

		if description != "" && resultPackage != "" {
			d := PackageData{resultPackage, description}
			packageOptions = append(packageOptions, d)
			description = ""
			resultPackage = ""
		}
		if len(packageOptions) == maxResultLength {
			break
		}

	}
	// Parse the string
	return packageOptions
}
