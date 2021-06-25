package scrape

import (
	"cli/model/rate_my_professor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ScrapeProfessors(schoolId int) rate_my_professor.Professors {
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
		professorArray := make([]*rate_my_professor.Professor, len(professorsUnmarshalled.Professors))
		copy(professorArray, professors.Professors)
		professors.Professors = professorArray
	} else {
		for _, professor := range professorsUnmarshalled.Professors {
			if professor != nil {
				professors.Professors = append(professors.Professors, professor)
			}
		}
	}
	if professorsUnmarshalled.Remaining <= 20 {
		return professors
	}

	return startScraping(professors, schoolId, page+1)
}