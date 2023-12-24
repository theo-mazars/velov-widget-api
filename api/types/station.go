package types

type BikeStation struct {
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Position struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"position"`
	Status      string `json:"status"`
	LastUpdate  string `json:"lastUpdate"`
	Connected   bool   `json:"connected"`
	TotalStands struct {
		Availabilities struct {
			Bikes           int `json:"bikes"`
			Stands          int `json:"stands"`
			MechanicalBikes int `json:"mechanicalBikes"`
			ElectricalBikes int `json:"electricalBikes"`
		}
		Capacity int `json:"capacity"`
	} `json:"totalStands"`
}

type StationStruct struct {
	Number    int     `json:"number"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
