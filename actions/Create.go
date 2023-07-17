package actions

import "git-identity/profiles"
import "git-identity/prompt"
import "fmt"
import "os/exec"

func Create(alias string) bool {

	var result bool = false

	_, ok := profiles.Profiles[alias]

	if ok == false {

		name := prompt.String("What is your full username?")
		email := prompt.String("What is your email address?")

		key := profiles.Folder+"/"+alias+".key"
		cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", key, "-q", "-N", "")
		_, err := cmd.Output()

		if err == nil {

			var profile profiles.Profile

			profile.Core.SSHCommand="ssh -i \""+key+"\" -F /dev/null"
			profile.User.Name = name
			profile.User.Email = email
			profile.User.UseConfigOnly = true

			result = profiles.Save(alias, profile)

		} else {

			fmt.Println("Warning: Could not generate SSH key pair!")
			fmt.Println("Warning: Please execute ssh-keygen -t ed25519 -f \"" + key + "\" manually.")
			result = false

		}

	} else {

		fmt.Println("Warning: Identity \"" + alias + "\" already exists!")
		result = false

	}

	return result

}
