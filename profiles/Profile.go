package profiles

type Profile struct {
	Core struct {
		SSHCommand string `json:"sshCommand"`
	} `json:"core"`
	User struct {
		Name          string `json:"name"`
		Email         string `json:"email"`
		UseConfigOnly bool   `json:"useConfigOnly"`
	} `json:"user"`
}
