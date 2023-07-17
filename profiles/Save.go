package profiles

import "encoding/json"
import "os"
import "fmt"

func Save(alias string, profile Profile) bool {

	var result bool = false

	if Folder != "" && alias != "" {

		buffer, err1 := json.MarshalIndent(profile, "", "\t")

		if err1 == nil {

			err2 := os.WriteFile(Folder+"/"+alias+".json", buffer, 0666)

			if err2 == nil {
				result = true
			} else {
				fmt.Println("Warning: Could not save config to \""+Folder+"/"+alias+".json\"!")
			}

		} else {
			fmt.Println("Warning: Could not save config to \""+Folder+"/"+alias+".json\"!")
		}

	}

	return result

}
