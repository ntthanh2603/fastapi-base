package entity

// Account defines the account entity
type Account struct {
	AccountId     int64  `gorm:"column:account_id;primaryKey;autoIncrement"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Email         string `gorm:"column:email"`
	Status        int    `gorm:"column:status"`
	Language      string `gorm:"column:language"`
	LastLoginTime int64  `gorm:"column:last_login_time"`
	CreateTime    int64  `gorm:"column:create_time"`
	UpdateTime    int64  `gorm:"column:update_time"`
	IsDeleted     int    `gorm:"column:is_deleted"`
}

// GORM override table name
func (Account) TableName() string {
	return "drunk_user"
}
