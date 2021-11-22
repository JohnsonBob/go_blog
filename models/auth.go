package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(userName string, password string) bool {
	auth := Auth{}
	db.Select("id").Where(Auth{Username: userName, Password: password}).First(&auth)
	return auth.ID > 0
}
