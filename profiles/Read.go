package profiles

import "encoding/json"
import "os"

func Read(alias string) bool {

	var result bool = false

	if Folder != "" && alias != "" {

		buffer, err1 := os.ReadFile(Folder+"/"+alias+".json")

		if err1 == nil {

			var profile Profile

			err2 := json.Unmarshal(buffer, &profile)


			if err2 == nil {
				Profiles[alias] = profile
				result = true
			}

		}

	}

	return result

}
