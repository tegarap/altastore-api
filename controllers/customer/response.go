package customer

type Response struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
	Gender   string `json:"gender" sql:"type:ENUM('male', 'female')"`
}
