package users

type User struct {
	Uid          string `gorm:"column:uid;primaryKey"`
	Email        string `gorm:"column:email;unique"`
	PasswordHash string `gorm:"column:password_hash"`
}
