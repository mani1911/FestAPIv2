package dto

type TShirtsUpdateRequest struct {
	Size string `json:"size"`
}

type TShirtsUpdateDTO struct {
	UserID uint   `json:"userID"`
	Size   string `json:"size"`
}
