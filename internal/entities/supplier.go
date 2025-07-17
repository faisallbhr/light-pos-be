package entities

type Supplier struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}
