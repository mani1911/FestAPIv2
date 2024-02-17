package dto

type TShirtsUpdateRequest struct {
	Size           string `json:"size"`
	Code           string `json:"code"`
	ScreenshotLink string `json:"screenshotLink"`
	RecaptchaCode  string `json:"recaptchaCode"`
}

type TShirtsUpdateDTO struct {
	UserID         uint   `json:"userID"`
	Size           string `json:"size"`
	Code           string `json:"code"`
	ScreenshotLink string `json:"screenshotLink"`
	RecaptchaCode  string `json:"recaptchaCode"`
}
