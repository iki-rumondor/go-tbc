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

type Case struct {
	Year             string `json:"year" valid:"required~field tahun tidak ditemukan"`
	ChildCount       int64  `json:"child_count" valid:"required~field jumlah anak-anak tidak ditemukan"`
	AdultCount       int64  `json:"adult_count" valid:"required~field jumlah dewasa tidak ditemukan"`
	MaleCount        int64  `json:"male_count" valid:"required~field jumlah laki-laki tidak ditemukan"`
	FemaleCount      int64  `json:"female_count" valid:"required~field jumlah perempuan tidak ditemukan"`
	HealthCenterUuid string `json:"health_center_uuid" valid:"required~field uuid puskesmas tidak ditemukan"`
}
