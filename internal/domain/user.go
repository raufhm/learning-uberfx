package domain

type User struct {
	UID       string `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
