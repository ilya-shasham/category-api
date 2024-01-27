package handlers

import (
	"categoryAPI/helpers"
	"categoryAPI/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCategoryById(ctx *gin.Context) {
	var id int

	if err := helpers.ParseParams(ctx, []string{"id"}, strconv.Atoi, &id); err != nil {
		return
	}

	result := models.GetCategories("id = ?", id)

	if len(result) == 0 {
		helpers.Success(nil, ctx)
		return
	}

	helpers.Success(helpers.PrepCategories(result[0])[0], ctx)
}

func GetCategoriesContaining(ctx *gin.Context) {
	contains := strings.ToLower(ctx.Param("contains"))
	results := models.GetCategories("name line ?", "%"+contains+"%")

	helpers.Success(helpers.PrepCategories(results...), ctx)
}

func GetCategoriesRanged(ctx *gin.Context) {
	var from, to int

	if err := helpers.ParseParams(ctx, []string{"from", "to"}, strconv.Atoi, &from, &to); err != nil {
		return
	}

	results := models.GetCategories("id between ? and ?", from, to)

	helpers.Success(helpers.PrepCategories(results...), ctx)
}

func GetAllCategories(ctx *gin.Context) {
	results := models.GetCategories("")

	helpers.Success(helpers.PrepCategories(results...), ctx)
}
