package handlers

import (
	"categoryAPI/models"
	"errors"
	"fmt"
)

type postRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	SubCategories []int  `json:"subcategories"`
}

func (r *postRequest) validate() error {
	if r.Name == "" {
		return errors.New("name cannot be empty")
	}

	if len(r.SubCategories) > 255 {
		return errors.New("too many subcategories. max is 255")
	}

	if models.Exists("name", r.Name) {
		return fmt.Errorf("%s already registered as a category name", r.Name)
	}

	for _, id := range r.SubCategories {
		if !models.Exists("id", id) {
			return fmt.Errorf("%d does not exist", id)
		}

		if models.IsSubCategory(id) {
			return fmt.Errorf("%d already a subcategory of something else", id)
		}
	}

	return nil
}
