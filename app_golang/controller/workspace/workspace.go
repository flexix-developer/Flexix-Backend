package workspace

type workspacebody struct {
	Fname string `json:"fname" validate:"require"`
	Lname string `json:"lname" validate:"require"`
	Email string `json:"email" validate:"require"`
	Pass  string `json:"pass" validate:"require"`
}