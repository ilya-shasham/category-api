package handlers

import (
	"categoryAPI/globals"
	"categoryAPI/helpers"
	"categoryAPI/models"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func DeleteCategory(ctx *gin.Context) {
	ids := []int{}
	err := json.Unmarshal([]byte(ctx.Param("ids")), &ids)

	if err != nil {
		helpers.Throw(err, ctx, helpers.ClientFault)
		return
	}

	globals.Db.Exec("delete from categories where id in " + models.ToSQLList(ids))

	helpers.Success(nil, ctx)
}
