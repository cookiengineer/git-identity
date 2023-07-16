package actions

import "git-identity/profiles"
import "os"
import "os/exec"
import "strings"
import "fmt"

func Clone(alias string, arguments []string) bool {

	var result bool = false

	profile, ok := profiles.Profiles[alias]

	if ok == true {

		err1 := os.Setenv("GIT_SSH_COMMAND", profile.Core.SSHCommand)

		if err1 == nil {

			cmd_arguments := []string{"clone"}
			cmd_arguments = append(cmd_arguments, arguments...)

			var stdout_buf strings.Builder
			var stderr_buf strings.Builder

			cmd := exec.Command("git", cmd_arguments...)
			cmd.Stdout = &stdout_buf
			cmd.Stderr = &stderr_buf

			err2 := cmd.Run()

			if err2 == nil {

				buffer := strings.TrimSpace(stderr_buf.String())
				lines := strings.Split(buffer, "\n")

				if len(lines) > 0 {

					if strings.HasPrefix(lines[0], "Cloning into '") && strings.HasSuffix(lines[0], "'...") {

						repo_folder := lines[0][14:len(lines[0])-4]

						var git_folder string

						if strings.HasPrefix(repo_folder, "./") {

							pwd, err := os.Getwd()

							if err == nil {
								git_folder = pwd+"/"+repo_folder[2:]+"/.git"
							} else {
								git_folder = repo_folder+"/.git"
							}

						} else if strings.HasPrefix(repo_folder, "/") {
							git_folder = repo_folder+"/.git"
						}

						if git_folder != "" {

							result = PatchGitFolder(git_folder, profile)

						} else {

							fmt.Println("Warning: Please execute git identity config \"" + alias + "\" manually inside the repository folder.")
							result = true

						}

					}

				}

			}

		}

	}

	return result

}
