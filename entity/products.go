package entity

// Product content
type Product struct {
	Id          int    `json:"id"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}


type Products struct {
	Product []*Product  `json:"product"`
}
