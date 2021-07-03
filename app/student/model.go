package student

type StudentTest struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Age     int	`json:"age"`
	Address string `json:"address"`
}

//TableName 为StudentTest绑定表名
func (student StudentTest) TableName() string {
	return "student_test"
}

