package Models

type NewsModel struct {
	NewsID      int
	CpyID       int
	CpyName     string
	NewsTitle   string
	NewsContent string
	ReleaseDate string

	//Generate
	NewsURL     string

}