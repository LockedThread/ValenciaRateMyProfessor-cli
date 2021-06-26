package rate_my_professor

import (
	"fmt"
	"strconv"
)

type Professors struct {
	Professors []*Professor `json:"professors"`
	Total      int          `json:"searchResultsTotal"`
	Remaining  int          `json:"remaining"`
	Type       string       `json:"type"`
}

type Professor struct {
	Department      string `json:"tDept" bson:"department"`
	Tsid            string `json:"tSid"` // TODO: Figure out what this is
	InstitutionName string `json:"institution_name" bson:"institutionName"`
	FirstName       string `json:"tFname" bson:"firstName"`
	MiddleName      string `json:"tMiddlename" bson:"middleName"`
	LastName        string `json:"tLname" bson:"lastName"`
	TeacherID       int    `json:"tid" bson:"teacherId"`
	RatingsCount    int    `json:"tNumRatings" bson:"ratingCount"`
	RatingClass     string `json:"rating_class" bson:"ratingClass"`
	ContentType     string `json:"contentType"`
	CategoryType    string `json:"categoryType"`
	OverallRating   string `json:"overall_rating" bson:"overallRating"`
}

func (p Professor) CalculateRating() float64 {
	float, err := strconv.ParseFloat(p.OverallRating, 64)
	if err != nil {
		return 0.0
	}
	return float * float64(p.RatingsCount)
}

func (p Professor) String() string {
	return fmt.Sprintf("Department: %s, Tsid: %s, InstitutionName: %s, FirstName: %s, MiddleName: %s, LastName: %s, TeacherID: %d, RatingsCount: %d, RatingClass: %s, ContentType: %s, CategoryType: %s, OverallRating: %f",
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
