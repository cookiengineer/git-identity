package profiles

import "encoding/json"
import "os"
import "fmt"

func Save(profile Profile) bool {

	var result bool = false

	fmt.Println(Folder)
	fmt.Println(profile.User.Name)

	if Folder != "" && profile.User.Name != "" {

		buffer, err1 := json.MarshalIndent(profile, "", "\t")

		if err1 == nil {

			err2 := os.WriteFile(Folder+"/"+profile.User.Name+".json", buffer, 0666)

			if err2 == nil {
				result = true
			} else {
				fmt.Println("Warning: Could not save config to \""+Folder+"/"+profile.User.Name+".json\"!")
			}

		} else {
			fmt.Println("Warning: Could not save config to \""+Folder+"/"+profile.User.Name+".json\"!")
		}

	}

	return result

}
