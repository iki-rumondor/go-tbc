package response

type HealthCenter struct {
	Uuid      string `json:"uuid" `
	Name      string `json:"name" `
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
