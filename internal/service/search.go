package service

import (
	"context"

	"dev/internal/persistence"
)

func Search(ctx context.Context, search SearchCommand) (result SearchCommandResponse, err error) {
	hubResult, err := persistence.Hub().SearchHub(ctx, search.Location, search.Type)

	if err != nil {
		return result, ErrDatabase
	}

	result = SearchCommandResponse{
		Data: struct {
			Hubs []HubsContract `json:"hubs"`
		}{Hubs: []HubsContract{}},
	}

	for _, hubs := range hubResult {
		contractHub := HubsContract{
			Name:     hubs.Name,
			Location: hubs.Location,
			Teams:    []TeamsContract{},
		}

		for _, team := range hubs.Teams {
			contractTeam := TeamsContract{
				Name:  team.Name,
				Type:  team.Type,
				Users: []UsersContract{},
			}
			for _, user := range team.Users {
				contractTeam.Users = append(contractTeam.Users, UsersContract{
					Name: user.Name,
					Age:  user.Age,
				})
			}
			contractHub.Teams = append(contractHub.Teams, contractTeam)
		}
		result.Data.Hubs = append(result.Data.Hubs, contractHub)
	}

	return
}
