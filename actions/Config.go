package actions

import "git-identity/profiles"
import "fmt"
import "os/exec"
import "strings"

func Config(alias string) bool {

	var result bool = false

	profile, ok := profiles.Profiles[alias]

	if ok == true {

		cmd1 := exec.Command("git", "rev-parse", "--absolute-git-dir")
		buffer1, err1 := cmd1.Output()

		if err1 == nil {

			git_folder := strings.TrimSpace(string(buffer1))

			if git_folder != "" {
				result = PatchGitFolder(git_folder, profile)
			}

		} else {

			fmt.Println("Warning: Not a git repository!")
			result = false

		}

	}

	return result

}
