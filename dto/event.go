package dto

type EventRegistrationDTO struct {
	EventID uint
	UserID  uint
}

type EventRegistrationRequest struct {
	EventID uint `json:"event_id"`
}

type AbstractDetailsRequest struct {
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
