package actions

import "git-identity/profiles"
import "fmt"
import "os"
import "strings"

func ShowKey(alias string) bool {

	var result bool = false

	_, ok := profiles.Profiles[alias]

	if ok == true {

		key := profiles.Folder+"/"+alias+".key.pub"

		buffer, err := os.ReadFile(key)

		if err == nil {

			fmt.Println(strings.TrimSpace(string(buffer)))
			result = true

		} else {

			fmt.Println("Warning: Identity \"" + alias + "\" has no SSH key pair!?")
			fmt.Println("Warning: Please execute ssh-keygen -t ed25519 -f \"" + key + "\" manually.")
			result = false

		}

	}

	return result

}
