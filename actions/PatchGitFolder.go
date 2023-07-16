package actions

import "git-identity/profiles"
import "os/exec"

func PatchGitFolder(git_folder string, profile profiles.Profile) bool {

	var result bool = false

	if git_folder != "" {

		var result_core_sshcommand bool = false
		var result_user_name bool = false
		var result_user_email bool = false

		git_config := git_folder+"/config"

		if profile.User.Name != "" {

			cmd := exec.Command("git", "config", "--file", git_config, "user.name", profile.User.Name)
			_, err := cmd.Output()

			if err == nil {
				result_user_name = true
			}

		} else {
			result_user_name = true
		}

		if profile.User.Email != "" {

			cmd := exec.Command("git", "config", "--file", git_config, "user.email", profile.User.Email)
			_, err := cmd.Output()

			if err == nil {
				result_user_email = true
			}

		} else {
			result_user_email = true
		}

		if profile.Core.SSHCommand != "" {

			cmd := exec.Command("git", "config", "--file", git_config, "core.sshCommand", profile.Core.SSHCommand)
			_, err := cmd.Output()

			if err == nil {
				result_core_sshcommand = true
			}

		} else {
			result_core_sshcommand = true
		}

		if result_user_name == true && result_user_email == true && result_core_sshcommand == true {
			result = true
		}

	}

	return result

}
