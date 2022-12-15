package service

import "github.com/asaskevich/govalidator"

type CreateUserCommand struct {
	Name   string `json:"name" valid:"required,maxstringlength(50)"`
	Age    int    `json:"age"  valid:"required"`
	TeamID string `json:"team_id"  valid:"required,maxstringlength(50)"`
}

type CreateTeamCommand struct {
	Name  string `json:"name"  valid:"required,maxstringlength(100)"`
	Type  string `json:"'type'" valid:"required,maxstringlength(100)"`
	HubID string `json:"hub_id"  valid:"required,maxstringlength(50)"`
}

type CreateHubCommand struct {
	Name     string `json:"name"  valid:"required,maxstringlength(100)"`
	Location string `json:"location"  valid:"required,maxstringlength(100)"`
}

func (command CreateUserCommand) Validation() error {
	_, err := govalidator.ValidateStruct(command)
	return err
}

func (command CreateTeamCommand) Validation() error {
	_, err := govalidator.ValidateStruct(&command)
	return err
}

func (command CreateHubCommand) Validation() error {
	_, err := govalidator.ValidateStruct(command)
	return err
}

type SearchCommand struct {
	Location string
	Type     string
}

type HubsContract struct {
	Name     string          `json:"name"`
	Location string          `json:"location"`
	Teams    []TeamsContract `json:"teams"`
}

type TeamsContract struct {
	Name  string          `json:"name"`
	Type  string          `json:"'type'"`
	Users []UsersContract `json:"users"`
}

type UsersContract struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type SearchCommandResponse struct {
	Data struct {
		Hubs []HubsContract `json:"hubs"`
	} `json:"data"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name" validate:"required,max=100"`
	Age    int    `json:"age"  validate:"required"`
	TeamID string `json:"team_id"  validate:"required,max=50" `
}

type CreateTeamResponse struct {
	TeamID string `json:"team_id"`
	Name   string `json:"name"  validate:"required,max=100"`
	Type   string `json:"'type'" validate:"required,max=100"`
	HubID  string `json:"hub_id"  validate:"required,max=50"`
}

type CreateHubResponse struct {
	HubId    string `json:"hub_id"`
	Name     string `json:"name"  validate:"required,max=100"`
	Location string `json:"location"  validate:"required,max=100"`
}
