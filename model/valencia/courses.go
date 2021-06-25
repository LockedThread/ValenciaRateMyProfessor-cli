package valencia

import (
	"fmt"
)

var Campuses = map[string]int{
	"WC":  1544,
	"EC":  1544,
	"DTC": 17862,
	"LNC": 16828,
	"PC":  17341,
	"WP":  13626,
	"OC":  13651,
}

type Course struct {
	CRN       string
	Subject   string
	Course    int
	Title     string
	CampusID  string
	Credits   float64
	Honors    bool
	Professor FullName
}

func (c Course) String() string {
	return fmt.Sprintf("CRN: %s, Subject: %s, Course: %d, Title: %s, CampusID: %s, Credits: %f, Honors: %t, Professor: %s", c.CRN, c.Subject, c.Course, c.Title, c.CampusID, c.Credits, c.Honors, c.Professor)
}

func (c Course) CampusName() {

}
