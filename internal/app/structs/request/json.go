package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type HealthCenter struct {
	Name      string  `json:"name" valid:"required~field nama tidak ditemukan"`
	Longitude float64 `json:"longitude" valid:"required~field longitude tidak ditemukan"`
	Latitude  float64 `json:"latitude" valid:"required~field latitude tidak ditemukan"`
}
