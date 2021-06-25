package valencia

import (
	"fmt"
	"strings"
)

type FullName struct {
	firstName     string
	middleInitial string
	lastName      string
}

func (f FullName) String() string {
	return "{" + f.FormattedString() + "}"
}

func (f FullName) FirstAndLast() string {
	return fmt.Sprintf("%s %s", f.firstName, f.lastName)}

func (f FullName) FormattedString() string {
	if len(f.middleInitial) > 0 {
		return fmt.Sprintf("%s %s %s", f.firstName, f.middleInitial, f.lastName)
	}
	return f.FirstAndLast()
}

func TrimSpaces(s string) string {
	var fullString string
	for i := range s {
		cursor := s[i]
		if cursor == ' ' {
			if s[i+1] == ' ' {
				continue
			}
		}
		fullString += string(cursor)
	}
	return fullString
}

func GetFullNameFromString(name string) FullName {
	name = strings.Replace(name, " (P)", "", 1)
	name = strings.TrimSpace(name)
	name = TrimSpaces(name)

	fullNameStruct := FullName{}

	split := strings.Split(name, " ")

	if len(split) == 1 {
		fullNameStruct.firstName = split[0]
		return fullNameStruct
	}

	splitLength := len(split)

	fullNameStruct.firstName = split[0]
	fullNameStruct.lastName = split[splitLength-1]

	for i := 1; i < splitLength-1; i++ {
		if i == splitLength-2 {
			fullNameStruct.middleInitial += split[i]

		} else {
			fullNameStruct.middleInitial += split[i] + " "
		}
	}

	if fullNameStruct.FormattedString() != name {
		fmt.Printf("name=%s\n", name)
		fmt.Printf("FormattedString=%s\n", fullNameStruct.FormattedString())
	}

	return fullNameStruct
}
