package category

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddCategoryRequest struct {
	CategoryID int64 `json:"category_id"`
}