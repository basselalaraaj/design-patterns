package user

import (
	"strings"
)

type game struct {
	allNames []string
}

func NewGame() *game {
	return &game{}
}

func (g *game) CreateUser(name string) *user {
	var userNames []int
	names := strings.Split(name, " ")
	for _, v := range names {
		var found bool
		for k, p := range g.allNames {
			if v == p {
				userNames = append(userNames, k)
				found = true
			}
		}

		if !found {
			g.allNames = append(g.allNames, v)
			userNames = append(userNames, len(g.allNames)-1)
		}
	}

	return newUser(&g.allNames, userNames)
}

func (g *game) CountNames() int {
	return len(g.allNames)
}
