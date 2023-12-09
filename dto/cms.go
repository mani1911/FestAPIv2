package dto

type AddEventRequest struct {
	EventID         uint   `json:"event_id"`
	EventName       string `json:"event_name"`
	IsTeam          bool   `json:"is_team"`
	MaxTeamSize     uint   `json:"max_team_size"`
	MaxParticipants uint   `json:"max_participants"`
	ForwardEmail    string `json:"forward_email"`
}
