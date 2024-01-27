package handlers

import (
	"categoryAPI/globals"
	"categoryAPI/helpers"
	"categoryAPI/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddCategory(ctx *gin.Context) {
	req := new(postRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		helpers.Throw(err, ctx, helpers.ClientFault)
		return
	}

	// Obviously not a good idea to make it possible
	// to have `tech` and `Tech` as seperate categories.
	req.Name = strings.ToLower(req.Name)

	if err := req.validate(); err != nil {
		helpers.Throw(err, ctx, helpers.ClientFault)
		return
	}

	globals.Db.Save(&models.Category{
		Name:          req.Name,
		Description:   req.Description,
		SubCategories: helpers.UnsafeMarshal(req.SubCategories),
	})

	helpers.Success(nil, ctx)
}
