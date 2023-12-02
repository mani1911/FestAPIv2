package dto

type EventRegistrationDTO struct {
	EventID     uint
	UserID      uint
	TeamMembers []string
	TeamName    string
}

type EventRegistrationRequest struct {
	EventID     uint     `json:"event_id"`
	TeamMembers []string `json:"team_members"`
	TeamName    string   `json:"team_name"`
}

type AbstractDetailsRequest struct {
	EventID uint `param:"event_id"`
}

type EventStatusRequest struct {
	EventID uint `param:"event_id"`
}

type AbstractDetailsResponse struct {
	ForwardEmail    string `json:"forward_email"`
	MaxParticipants uint   `json:"max_participants"`
}

type GetEventDetailsResponse struct {
	EventID   uint   `json:"event_id"`
	EventName string `json:"event_name"`
}

type EventStatusResponse struct {
	IsRegistered bool     `json:"is_registered"`
	IsTeam       bool     `json:"is_team"`
	TeamID       uint     `json:"team_id"`
	TeamMembers  []string `json:"team_members"`
}
