package models

type Token struct {
	UserID         int    `json:"userID"`
	InviteToken    string `json:"inviteToken"`
	ExpirationDate int    `json:"expirationDate"`
}

type User struct {
	UserID    int    `json:"userID"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	UserType  int    `json:"userType"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
