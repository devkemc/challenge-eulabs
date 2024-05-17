package entity

type Product struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
	Stock       int32   `json:"stock"`
	Active      bool    `json:"active"`
}
