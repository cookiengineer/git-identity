package actions

import "git-identity/profiles"
import "fmt"
import "sort"

func Show() bool {

	var aliases []string

	for alias, _ := range profiles.Profiles {
		aliases = append(aliases, alias)
	}

	sort.Strings(aliases)

	for a := 0; a < len(aliases); a++ {
		fmt.Println(aliases[a])
	}

	return true

}
