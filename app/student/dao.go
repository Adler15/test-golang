package student

import "stduentProject/db"

var st StudentTest
var sts []StudentTest

func findAll() []StudentTest {
	db := db.GetDB()
	db.Find(&sts)
	return sts
}

func findOne(id string) StudentTest {
	db := db.GetDB()
	st := StudentTest{}
	db.First(&st, id)
	return st
}

func findByName(name string) []StudentTest {
	db := db.GetDB()
	db.Where("name=?", name).Find(&sts)
	return sts
}

func create(st StudentTest) int64 {
	db := db.GetDB()
	result := db.Create(&st)
	return result.RowsAffected
}

func update(student map[string]interface{}) int64 {
	db := db.GetDB()
	id := student["id"]
	stu := StudentTest{ID: uint(id.(float64))}
	delete(student,"id")
	result := db.Model(&stu).Updates(student)
	return result.RowsAffected
}

func deleteSt(id string) int64 {
	db := db.GetDB()
	result := db.Delete(&StudentTest{}, id)
	return result.RowsAffected

}
