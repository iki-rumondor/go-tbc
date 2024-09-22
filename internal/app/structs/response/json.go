package response

type User struct {
	Uuid     string `json:"uuid" `
	Name     string `json:"name" `
	Username string `json:"username"`
	RoleName string `json:"role_name"`
}

type HealthCenter struct {
	Uuid      string  `json:"uuid" `
	Name      string  `json:"name" `
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}
