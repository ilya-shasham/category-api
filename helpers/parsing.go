package helpers

import (
	"categoryAPI/models"

	"github.com/gin-gonic/gin"
)

func ParseParams[T any](ctx *gin.Context, names []string, parser func(string) (T, error), results ...*T) error {
	for i, name := range names {
		parsed, err := parser(ctx.Param(name))

		if err != nil {
			Throw(err, ctx, ClientFault)
			return err
		}

		*results[i] = parsed
	}

	return nil
}

type preppedCategory struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	SubCategories []int  `json:"subcategories"`
}

func PrepCategories(categories ...models.Category) []preppedCategory {
	result := make([]preppedCategory, len(categories))

	for i := 0; i < len(categories); i++ {
		result[i] = preppedCategory{
			ID:            int(categories[i].ID),
			Name:          categories[i].Name,
			Description:   categories[i].Description,
			SubCategories: UnsafeUnmarshal[[]int](categories[i].SubCategories),
		}
	}

	return result
}
