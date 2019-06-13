package Models

type StdUserModel struct {
	StdID         int
	Username      string
	Password      string
	FirstName     string
	LastName      string
	MailAddress   string
	CollegeName   string
	Degree        string
	Department    string
	Major         string
	GraduateDate  string
	LastLoginDate string

	Hobbies string
	Skills  string

	//Generate
	StdURL    string
	ApplyDate string
}
