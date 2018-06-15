package utils

import (
	"bufio"
	"strings"
)

// FilterSearch takes helm search output and only returns 1 or all the package lines
func FilterSearch(searchResults string, matchSelector string, first bool) []string {
	var resultsList []string
	scanner := bufio.NewScanner(strings.NewReader(searchResults))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, matchSelector) {
			resultsList = append(resultsList, line)
			if first {
				return resultsList
			}
		}
	}

	return resultsList
}

// GetLineInfo take one line from helm search and extract all the fields
func GetLineInfo(line string) (map[string]interface{}, error) {
	var info map[string]interface{}
	info = make(map[string]interface{})
	infoSlice := strings.Fields(line)

	info["name"] = infoSlice[0]
	info["chart_version"] = infoSlice[1]
	info["app_version"] = infoSlice[2]
	info["description"] = ""

	for i := 3; i < len(infoSlice); i++ {
		info["description"] = info["description"].(string) + " " + infoSlice[i]
	}

	return info, nil
}
