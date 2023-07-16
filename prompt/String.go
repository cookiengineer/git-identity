package prompt

import "bufio"
import "fmt"
import "os"
import "strings"

func String(question string) string {

	var answer string

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println(question+" ")

		tmp1, err := reader.ReadString('\n')

		if err == nil && tmp1 != "" {

			tmp2 := strings.TrimSpace(tmp1)

			if tmp2 != "" {
				answer = tmp2
				break
			}

		}

	}

	return answer

}

