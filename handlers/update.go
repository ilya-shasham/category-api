package handlers

import (
	"categoryAPI/globals"
	"categoryAPI/helpers"
	"categoryAPI/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(ctx *gin.Context) {
	id := 0

	if err := helpers.ParseParams[int](ctx, []string{"id"}, strconv.Atoi, &id); err != nil {
		helpers.Throw(err, ctx, helpers.ClientFault)
		return
	}

	name, description := strings.ToLower(ctx.Param("name")), ctx.Param("description")

	if name == "" {
		helpers.Throw(errors.New("name cannot be empty"), ctx, helpers.ClientFault)
		return
	}

	subcategories := []int{}

	if err := json.Unmarshal([]byte(ctx.Param("subcategories")), &subcategories); err != nil {
		helpers.Throw(err, ctx, helpers.ClientFault)
		return
	}

	var category *models.Category

	{
		matches := models.GetCategories("id = ?", id)

		if len(matches) == 0 {
			helpers.Throw(errors.New("id not found"), ctx, helpers.ClientFault)
			return
		}

		*category = matches[0]
	}

	category.Name, category.Description, category.SubCategories = name, description, helpers.UnsafeMarshal(subcategories)

	globals.Db.Save(category)
	helpers.Success(nil, ctx)
}
