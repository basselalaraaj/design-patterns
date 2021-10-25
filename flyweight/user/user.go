package user

import "strings"

type user struct {
	names    []int
	allNames *[]string
}

func newUser(allNames *[]string, names []int) *user {
	return &user{
		allNames: allNames,
		names:    names,
	}
}

func (u *user) GetFullName() string {
	var fullName []string
	for _, v := range u.names {
		fullName = append(fullName, (*u.allNames)[v])
	}
	return strings.Join(fullName, " ")
}
