package lib

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseInput(input string) (string, string) {
	// har ikke laget regex stringen selv

	// regex string som skal gjøre at input blir splittet på space, med mindre det er quotations rundt.
	re := regexp.MustCompile(`^(\w+)\s+"([^"]+)"$|^(\w+)\s+(.+)$|^(\w+)$`)
	matches := re.FindStringSubmatch(input)

	if matches == nil {
		return "", ""
	}

	// alt under er en del av noe jeg ikke har laget. har ikke nok peiling på parsing til å skjønne dette
	command := ""
	for _, match := range matches[1:5] {
		if match != "" {
			command = match
			break
		}
	}

	arg := ""
	if matches[2] != "" {
		arg = matches[2]
	} else if matches[4] != "" {
		arg = matches[4]
	}

	return command, arg
}


func FormatResponse(response OMDbResponse) string {
	if response.Response != "True" {
		return "No results found."
	}

	var formatted strings.Builder
	fmt.Fprintf(&formatted, "Total Results: %s\n", response.TotalResults)
	fmt.Fprintln(&formatted, "Search Results:")

	for _, movie := range response.Search {
		fmt.Fprintf(&formatted, "\nTitle: %s\nYear: %s\nIMDB ID: %s\nType: %s\nPoster: %s\n",
			movie.Title, movie.Year, movie.ImdbID, movie.Type, movie.Poster)
	}

	return formatted.String()
}