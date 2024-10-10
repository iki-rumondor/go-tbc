package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type Case struct {
	Year             string `json:"year" valid:"required~field tahun tidak ditemukan"`
	ChildCount       int64  `json:"child_count" valid:"numeric~field jumlah anak-anak tidak ditemukan"`
	AdultCount       int64  `json:"adult_count" valid:"numeric~field jumlah dewasa tidak ditemukan"`
	MaleCount        int64  `json:"male_count" valid:"numeric~field jumlah laki-laki tidak ditemukan"`
	FemaleCount      int64  `json:"female_count" valid:"numeric~field jumlah perempuan tidak ditemukan"`
	HealthCenterUuid string `json:"health_center_uuid" valid:"required~field uuid puskesmas tidak ditemukan"`
}

type Clustering struct {
	Year string `json:"year" valid:"required~field tahun tidak ditemukan"`
}
