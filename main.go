package main

import "git-identity/actions"
import "os"
import "strings"
import "fmt"

func showUsage() {

	fmt.Println("Usage: git-identity [Action] [Alias] [...]")
	fmt.Println("")
	fmt.Println("| Action           | Description                                         |")
	fmt.Println("|:-----------------|:----------------------------------------------------|")
	fmt.Println("| create <alias>   | creates a new alias                                 |")
	fmt.Println("| import <alias>   | imports global git user config as a new alias       |")
	fmt.Println("| clone <alias>    | clones a private repository with SSH key of alias   |")
	fmt.Println("| config <alias>   | modifies current git repository to always use alias |")
	fmt.Println("| show             | prints out available aliases                        |")
	fmt.Println("| show-key <alias> | prints out the public SSH key for given alias       |")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("")
	fmt.Println("# 1.A) create a new alias")
	fmt.Println("git identity create us3rn4me;")
	fmt.Println("")
	fmt.Println("# 1.B) import a new alias from existing global ssh/git config")
	fmt.Println("git identity import us3rn4me;")
	fmt.Println("")
	fmt.Println("# 2. copy/paste the ssh public key to github/gitlab")
	fmt.Println("git identity show-key us3rn4m3;")
	fmt.Println("")
	fmt.Println("# 3. clone a private repository with the given alias")
	fmt.Println("git identity clone us3rn4m3 git@github.com:private/repo.git ./repo;")
	fmt.Println("")
	fmt.Println("# 4. modify repository config use the given alias")
	fmt.Println("cd /path/to/existing/repo;")
	fmt.Println("git identity config us3rn4m3;")

}

func main() {

	var alias string
	var action string
	var arguments []string

	if len(os.Args) == 2 {

		action = strings.TrimSpace(os.Args[1])

	} else if len(os.Args) == 3 {

		action = strings.TrimSpace(os.Args[1])
		alias = strings.TrimSpace(os.Args[2])

	} else if len(os.Args) > 3 {

		action = strings.TrimSpace(os.Args[1])
		alias = strings.TrimSpace(os.Args[2])
		arguments = os.Args[3:]

	}

	if action == "create" {

		if alias != "" {

			result := actions.Create(alias)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else if action == "import" {

		if alias != "" {

			result := actions.Import(alias)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else if action == "clone" {

		if alias != "" {

			result := actions.Clone(alias, arguments)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else if action == "config" {

		if alias != "" {

			result := actions.Config(alias)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else if action == "show" {

		result := actions.Show()

		if result == true {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	} else if action == "showkey" || action == "show-key" {

		if alias != "" {

			result := actions.ShowKey(alias)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else {

			showUsage()
			os.Exit(1)

		}

	} else {

		showUsage()
		os.Exit(1)

	}

}
