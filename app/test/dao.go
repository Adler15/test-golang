package test

import "stduentProject/db"

var test Test
var tests []Test

func createTest(test Test) int64 {
	db := db.GetDB()
	create := db.Create(&test)
	return create.RowsAffected
}

func updateTest(test map[string]interface{}) int64 {
	db := db.GetDB()
	id := test["id"]
	t := Test{ID: uint(id.(float64))}
	delete(test, "id")
	updates := db.Model(&t).Updates(test)
	return updates.RowsAffected
}

func deleteTest(id string) int64 {
	db := db.GetDB()
	del := db.Delete(&Test{}, id)
	return del.RowsAffected
}

func findAllTest() []Test {
	db := db.GetDB()
	db.Find(&tests)
	return tests
}

func findOneTest(id string) Test {
	db := db.GetDB()
	db.First(&test, id)
	return test
}

func findByWhere(where map[string]interface{}) []Test {
	db := db.GetDB()
	db.Find(&tests, where)
	return tests
}
