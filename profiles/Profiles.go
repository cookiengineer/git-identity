package profiles

import "os"
import "strings"

var Folder string
var Profiles map[string]Profile

func init() {

	Profiles = make(map[string]Profile, 0)

	var folder string

	xdg_config := os.Getenv("XDG_CONFIG_HOME")
	home := os.Getenv("HOME")
	user := os.Getenv("USER")

	if xdg_config != "" {
		folder = xdg_config + "/git-identity"
	} else if home != "" {
		folder = home + "/.config/git-identity"
	} else if user != "" {
		folder = "/home/" + user + "/.config/git-identity"
	}


	if folder != "" {

		Folder = folder

		stat, err1 := os.Stat(folder)

		if err1 == nil && stat.IsDir() {

			entries, err2 := os.ReadDir(folder)

			if err2 == nil {

				for e := 0; e < len(entries); e++ {

					var file = entries[e].Name()

					if strings.HasSuffix(file, ".json") {
						Read(file[0:len(file)-5])
					}

				}

			}

		} else {
			os.MkdirAll(folder, 0750)
		}

	}

}
