package actions

import "git-identity/git"
import "git-identity/profiles"
import "fmt"
import "os/exec"
import "sort"
import "strings"

func Show() bool {

	var active string
	var aliases []string

	for alias, _ := range profiles.Profiles {
		aliases = append(aliases, alias)
	}

	sort.Strings(aliases)

	cmd1 := exec.Command("git", "rev-parse", "--absolute-git-dir")
	buffer1, err1 := cmd1.Output()

	if err1 == nil {

		git_folder := strings.TrimSpace(string(buffer1))

		if git_folder != "" {

			git_config := git_folder + "/config"
			name := git.QueryGitConfig(git_config, "user.name")
			email := git.QueryGitConfig(git_config, "user.email")

			for alias, profile := range profiles.Profiles {

				if profile.User.Name == name && profile.User.Email == email {
					active = alias
					break
				}

			}

		}

	}

	for a := 0; a < len(aliases); a++ {

		alias := aliases[a]

		if alias == active {
			fmt.Println("* "+alias)
		} else {
			fmt.Println("  "+alias)
		}

	}

	return true

}
