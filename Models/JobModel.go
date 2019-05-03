package Models

type JobModel struct {
	Jid          int
	Cpyid        int
	Title        string
	Describe     string
	Salary       string
	Location     string
	OtherDetails string
	ReleaseDate  string
	StartDate    string
	Deadline     string

	JobURL       string
}
