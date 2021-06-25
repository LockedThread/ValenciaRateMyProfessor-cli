package children

import (
	"cli/database"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var FindCmd = &cobra.Command{
	Use:   "find [schoolId] [Name]",
	Short: "Finds the professor with the name",
	Run: func(cmd *cobra.Command, args []string) {
		atoi, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Unable to parse %s as int", args[0])
			return
		}
		var name string
		for i := 1; i < len(args); i++ {
			name += args[i]
		}
		if len(name) == 0 {
			fmt.Println("Please provide the name of the professor you want to search for.")
			return
		}
		professor := database.FindProfessor(atoi, name)
		fmt.Println(professor)
	},
}
