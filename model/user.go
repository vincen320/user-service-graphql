package model

type (
	User struct {
		ID      int64   `json:"id"`
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address string  `json:"address"`
		Salary  float64 `json:"salary"`
		Hobbies []Hobby `json:"hobbies"`
	}

	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
