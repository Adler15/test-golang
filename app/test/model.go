package test

type Test struct {
	ID uint `gorm:"column:id" json:"id"`
	Test1  string `gorm:"column:test1" json:"test1"`
	Test2 string `gorm:"column:test2" json:"test2"`
	Test3 string `gorm:"column:test3" json:"test3"`
	Test4 string `gorm:"column:test4" json:"test4"`
	Test5 string `gorm:"column:test5" json:"test5"`
}

func (test Test) TableName() string {
	return "cfl_test"
}
