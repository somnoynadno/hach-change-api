package auxiliary

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username" gorm:"unique_index;not null;"`
}
