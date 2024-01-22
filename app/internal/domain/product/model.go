package product

type Product struct {
	ID          string
	Name        string
	Description string
}

type CreateProductDTO struct {
	Name        string
	Description string
}

type UpdateProductDTO struct {
	ID          string
	Name        string
	Description string
	ImageURL    string
}

func NewCategory(dto CreateProductDTO) Product {
	return Product{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

func UpdateCategory(dto UpdateProductDTO) Product {
	return Product{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
	}
}
