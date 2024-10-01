package request

type HealthCenter struct {
	Name      string  `form:"name" valid:"required~field nama tidak ditemukan"`
	Longitude float64 `form:"longitude" valid:"required~field longitude tidak ditemukan"`
	Latitude  float64 `form:"latitude" valid:"required~field latitude tidak ditemukan"`
}
