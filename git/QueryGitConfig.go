package git

import "os/exec"
import "strings"

func QueryGitConfig(git_config string, query string) string {

	var result string

	cmd := exec.Command("git", "config", "--file", git_config, query)
	buffer, err := cmd.Output()

	if err == nil {
		result = strings.TrimSpace(string(buffer))
	}

	return result

}
