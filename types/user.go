package types

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"unique"`
}
