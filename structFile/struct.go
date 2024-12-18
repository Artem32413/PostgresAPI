package structFile

type Car struct {
	ID      string `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Mileage int    `json:"mileage"`
	Owners  int    `json:"owners"`
}

type Furniture struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Length       int    `json:"length"`
}

type Flower struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	ArrivalDate string  `json:"arrivaldate"`
}
