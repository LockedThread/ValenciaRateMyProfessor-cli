package children

import (
	"cli/database"
	"cli/model/rate_my_professor"
	"cli/model/valencia"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var FileSearch = &cobra.Command{
	Use:   "filesearch [file]",
	Short: "Finds the professor with the best rate my professor score",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		if len(args) == 0 {
			fmt.Println("Please provide the name of the professor you want to search for.")
			return
		}

		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}

		courseArray := ParseTable(string(data))

		courseInstructorMap := make(map[string]valencia.Course)

		for _, course := range courseArray {
			courseInstructorMap[course.Professor.FormattedString()] = course
		}

		var professors []rate_my_professor.Professor

		for _, course := range courseInstructorMap {
			professorFullName := course.Professor

			p, err := database.FindProfessor(valencia.Campuses[course.CampusID], professorFullName.FirstAndLast())
			if err != nil {
				log.Fatalf("Fuck: %s\n", err)
			}
			fmt.Println(p)
			if p != nil {
				professor := *p
				_, err = strconv.ParseFloat(professor.OverallRating, 64)
				if err == nil {
					professors = append(professors, professor)
				} else {
					log.Printf("Fuck: %s\n", err)
				}
			}
		}
		fmt.Println("shit")
		sort.SliceStable(professors, func(i, j int) bool {
			float1, _ := strconv.ParseFloat(professors[i].OverallRating, 64)
			float2, _ := strconv.ParseFloat(professors[j].OverallRating, 64)
			return float1 > float2
		})
		fmt.Printf("Found Professors:\n")
		for _, professor := range professors {
			fmt.Printf("%s %s | %s %d\n", professor.FirstName, professor.LastName, professor.OverallRating, professor.RatingsCount)
		}
	},
}

func ParseTable(data string) []valencia.Course {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	courses := make([]valencia.Course, 1)

	table := doc.Find("table")

	table.Find("tr").Each(func(i int, rowSelection *goquery.Selection) {
		if i < 2 {
			return
		}

		tableData := rowSelection.Find("td")

		courses = append(courses, GetCourse(tableData))
	})
	return courses
}

func GetCourse(tableData *goquery.Selection) (course valencia.Course) {
	cont := true
	tableData.Each(func(i int, selection *goquery.Selection) {
		ParseTableRow(&course, &cont, i, selection)
	})

	if len(course.Title) == 0 {
		return
	}

	return course
}

func ParseTableRow(course *valencia.Course, cont *bool, i int, selection *goquery.Selection) () {
	text := selection.Text()

	if i == 0 {
		if len(text) == 2 {
			*cont = false
			return
		} else {
			*cont = true
		}
	}
	if *cont == false {
		return
	}

	switch i {
	case 1:
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, " ", "", -1)

		atoi, err := strconv.ParseInt(text, 0, 64)
		if err != nil {
			log.Fatal(err)
		}

		course.Course = int(atoi)
		break
	case 2:
		course.Subject = text
		break
	case 3:
		course.CRN = text
		break
	case 4:
		break
	case 5:
		course.CampusID = text
		break
	case 6:
		atoi, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal(err)
		}
		course.Credits = atoi
		break
	case 7:
		course.Title = text
		break
	case 16:
		course.Professor = valencia.GetFullNameFromString(text)
		break
	}
}
