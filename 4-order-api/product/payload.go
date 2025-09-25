package product

type ProductCreateRequest struct {
	Name        string   `json:"name"  validate:"required"`
	Description string   `json:"description"  validate:"required"`
	Images      []string `json:"images"  validate:"required"`
}
type ProductUpdateRequest struct {
	Name        string   `json:"name"  validate:"required"`
	Description string   `json:"description" `
	Images      []string `json:"images" `
}
