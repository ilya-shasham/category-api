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
	subcategories := []int{}

	if err := json.Unmarshal([]byte(ctx.Param("subcategories")), &subcategories); err != nil {
		helpers.Throw(err, ctx, helpers.ClientFault)
		return
	}

	var category *models.Category = new(models.Category)

	{
		matches := models.GetCategories("id = ?", id)

		if len(matches) == 0 {
			helpers.Throw(errors.New("id not found"), ctx, helpers.ClientFault)
			return
		}

		*category = matches[0]
	}

	if name != "" {
		category.Name = name
	}

	if description != "" {
		category.Description = description
	}

	if len(subcategories) != 0 {
		if subcategories[0] != -1 {
			category.SubCategories = helpers.UnsafeMarshal(subcategories)
		}
	} else {
		category.SubCategories = helpers.UnsafeMarshal(subcategories)
	}

	globals.Db.Save(category)
	helpers.Success(nil, ctx)
}
