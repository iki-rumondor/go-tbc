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

type Case struct {
	Uuid         string        `json:"uuid" `
	Year         string        `json:"year"`
	ChildCount   int64         `json:"child_count"`
	AdultCount   int64         `json:"adult_count"`
	MaleCount    int64         `json:"male_count"`
	FemaleCount  int64         `json:"female_count"`
	Total        int64         `json:"total"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
	HealthCenter *HealthCenter `json:"health_center"`
}
