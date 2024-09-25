package model

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type User struct {
	ID          int         `json:"id"`
	Username    string      `json:"username"`
	Name        string      `json:"name"`
	Password    string      `json:"password"`
	UserType    int         `json:"user_type"`
	ContactInfo ContactInfo `json:"contact_info"`
}
