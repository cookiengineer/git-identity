package actions

import "git-identity/git"
import "git-identity/profiles"
import "git-identity/prompt"
import "fmt"
import "io/ioutil"
import "os"

var key_names []string = []string{
	"id_dsa",
	"id_ecdsa",
	"id_ecdsa_sk",
	"id_ed25519",
	"id_ed25519_sk",
	"id_rsa",
}

func Import(alias string) bool {

	var result bool = false

	_, ok := profiles.Profiles[alias]

	if ok == false {

		var ssh_folder string
		var git_config string

		var result_ssh bool = false
		var result_git bool = false

		home := os.Getenv("HOME")
		user := os.Getenv("USER")

		if home != "" {
			git_config = home + "/.gitconfig"
			ssh_folder = home + "/.ssh"
		} else if user != "" {
			git_config = "/home/" + user + "/.gitconfig"
			ssh_folder = "/home/" + user + "/.ssh"
		}

		if ssh_folder != "" {

			for k := 0; k < len(key_names); k++ {

				in_key := ssh_folder + "/" + key_names[k]
				out_key := profiles.Folder+"/"+alias+".key"

				stat, err1 := os.Stat(in_key)

				if err1 == nil && stat.IsDir() == false {

					private_key_buffer, err2 := ioutil.ReadFile(in_key)
					public_key_buffer, err3 := ioutil.ReadFile(in_key+".pub")

					if err2 == nil && err3 == nil {

						err4 := ioutil.WriteFile(out_key, private_key_buffer, 0600)
						err5 := ioutil.WriteFile(out_key+".pub", public_key_buffer, 0644)

						if err4 == nil && err5 == nil {
							result_ssh = true
						}

					} else {

						fmt.Println("Warning: Could not import SSH key pair!")
						result = false

					}

					break

				}

			}

		}

		if git_config != "" {

			name := git.QueryGitConfig(git_config, "user.name")
			email := git.QueryGitConfig(git_config, "user.email")
			key := profiles.Folder+"/"+alias+".key"

			if name == "" {
				name = prompt.String("What is your full username?")
			}

			if email == "" {
				email = prompt.String("What is your email address?")
			}

			if name != "" && email != "" {

				var profile profiles.Profile

				profile.Core.SSHCommand="ssh -i \""+key+"\" -F /dev/null"
				profile.User.Name = name
				profile.User.Email = email
				profile.User.UseConfigOnly = true

				result_git = profiles.Save(profile)

			}

		}

		if result_ssh == true && result_git == true {
			result = true
		}

	} else {

		fmt.Println("Warning: Identity \"" + alias + "\" already exists!")
		result = false

	}

	return result

}
