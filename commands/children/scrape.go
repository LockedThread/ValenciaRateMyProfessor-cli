package children

import (
	"cli/model/rate_my_professor"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strconv"
)

var ScrapeCmd = &cobra.Command{
	Use:   "scrape [school-id]",
	Short: "Scrapes the school for all of the data on the professors",
	Run: func(cmd *cobra.Command, args []string) {
		atoi, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Unable to parse %s as int", args[0])
			return
		}
		professors := scrapeProfessors(atoi)
		for _, professor := range professors.Professors {
			fmt.Println(professor)
		}

		fmt.Println(professors.Total)


	},
}

func scrapeProfessors(schoolId int) rate_my_professor.Professors {
	professors := rate_my_professor.Professors{}

	return startScraping(professors, schoolId, 1)
}

func startScraping(professors rate_my_professor.Professors, schoolId int, page int) rate_my_professor.Professors {
	//goland:noinspection GoPrintFunctions
	url := fmt.Sprintf("https://www.ratemyprofessors.com/filter/professor/?&page=%d&filter=teacherlastname_sort_s+asc&query=**&queryoption=TEACHER&queryBy=schoolId&sid=%d", page, schoolId)

	response, err := http.Get(url)
	if err != nil {
		_ = fmt.Errorf("unable to make request: %s", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		_ = fmt.Errorf("unable to read request: %s", err)
	}
	err = response.Body.Close()
	if err != nil {
		_ = fmt.Errorf("unable close response body: %s", err)
	}
	var professorsUnmarshalled rate_my_professor.Professors

	err = json.Unmarshal(data, &professorsUnmarshalled)
	if err != nil {
		_ = fmt.Errorf("unable unmarshal professor data: %s", err)
	}

	if page == 1 {
		professors = professorsUnmarshalled
		professorArray := make([]rate_my_professor.Professor, professorsUnmarshalled.Total)
		copy(professorArray, professors.Professors)
		professors.Professors = professorArray
	} else {
		professors.Professors = append(professors.Professors, professorsUnmarshalled.Professors...)
		professors.Remaining = professorsUnmarshalled.Remaining
		if professors.Remaining <= 20 {
			return professors
		}
	}

	//return professors
	return startScraping(professors, schoolId, page+1)
}
