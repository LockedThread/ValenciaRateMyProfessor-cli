package children

import (
	"cli/database"
	"cli/scrape"
	"fmt"
	"github.com/spf13/cobra"
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
		professors := scrape.ScrapeProfessors(atoi)
		for _, professor := range professors.Professors {
			fmt.Println(*professor)
		}

		fmt.Println(professors.Total)

		_ = database.InsertScrapeData(atoi, professors)
	},
}
