package models

import (
	"categoryAPI/globals"
	"fmt"
	"strconv"
)

func GetCategories(cmd string, values ...any) []Category {
	results := make([]Category, 0)

	if cmd == "" {
		globals.Db.Model(new(Category)).Find(&results)
	} else {
		globals.Db.Model(new(Category)).Where(cmd, values...).Find(&results)
	}

	return results
}

func Exists(cmd string, values ...any) bool {
	result := new(Category)
	result.Name = ""

	globals.Db.Model(new(Category)).Where(cmd, values...).First(result)

	return result.Name != ""
}

func IsSubCategory(id int) bool {
	id_s := strconv.Itoa(id)

	return Exists(`sub_categories like ? or
										   sub_categories like ? or
										   sub_categories = ?`,
		"%,"+id_s+",%",
		"%,"+id_s+"]",
		"["+id_s+"]")
}

func ToSQLList[T any, S []T](s S) string {
	result := "("

	for _, element := range s {
		result += fmt.Sprintf("%v", element) + ","
	}

	return result[:len(result)-1] + ")"
}
