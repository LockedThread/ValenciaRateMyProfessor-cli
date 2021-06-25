package rate_my_professor

import "fmt"

type Professors struct {
	Professors []*Professor `json:"professors"`
	Total      int          `json:"searchResultsTotal"`
	Remaining  int          `json:"remaining"`
	Type       string       `json:"type"`
}

type Professor struct {
	Department      string `json:"tDept"`
	Tsid            string `json:"tSid"` // TODO: Figure out what this is
	InstitutionName string `json:"institution_name"`
	FirstName       string `json:"tFname"`
	MiddleName      string `json:"tMiddlename"`
	LastName        string `json:"tLname"`
	TeacherID       int    `json:"tid"`
	RatingsCount    int    `json:"tNumRatings"`
	RatingClass     string `json:"rating_class"`
	ContentType     string `json:"contentType"`
	CategoryType    string `json:"categoryType"`
	OverallRating   string `json:"overall_rating"`
}

func (p Professor) String() string {
	return fmt.Sprintf("Department: %s, Tsid: %s, InstitutionName: %s, FirstName: %s, MiddleName: %s, LastName: %s, TeacherID: %d, RatingsCount: %d, RatingClass: %s, ContentType: %s, CategoryType: %s, OverallRating: %s",
		p.Department,
		p.Tsid,
		p.InstitutionName,
		p.FirstName,
		p.MiddleName,
		p.LastName,
		p.TeacherID,
		p.RatingsCount,
		p.RatingClass,
		p.ContentType,
		p.CategoryType,
		p.OverallRating,
	)
}
