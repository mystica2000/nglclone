package interfaces

type User struct {
	Id    string `json:"sub"`
	Image string `json:"picture"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type QuestionResponse struct {
	Response string `json:"response"`
	Done     int    `json:"done"`
	Id       int64  `json:"id"`
}

type UserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active int    `json:"active"`
	Image  string `json:"image"`
}

type ToggleURLResponse struct {
	Active int `json:"active"`
}

type UserForm struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	PublicID string `json:"public"`
}

type Clients struct {
	ID      string
	Channel chan []QuestionResponse
}
