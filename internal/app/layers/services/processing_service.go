package services

import (
	"log"
	"math"
	"sort"

	"github.com/iki-rumondor/go-tbc/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/models"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/response"
)

type ProcessingService struct {
	Repo interfaces.ProcessingInterface
}

func NewProcessingService(repo interfaces.ProcessingInterface) *ProcessingService {
	return &ProcessingService{
		Repo: repo,
	}
}

func initCentroid(k int, data []int64) []float64 {
	centroid := make([]float64, k)
	for i := 0; i < k; i++ {
		centroid[i] = float64(data[i])
	}
	return centroid
}

func indexOf(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func kMeansClustering(data []int64) []int {
	k := 3
	maxIter := 20

	centroids := initCentroid(k, data)
	prevCentroids := make([]float64, k)
	clusters := make([]int, len(data))

	for iter := 0; iter < maxIter; iter++ {

		// Assign each point to the nearest centroid
		for i, point := range data {
			minDist := math.MaxFloat64
			for j := 0; j < k; j++ {
				// Hitung jarak Euclidean (abs dalam kasus 1D)
				dist := math.Abs(float64(point) - centroids[j])
				if dist < minDist {
					minDist = dist
					clusters[i] = j
				}
			}
		}

		// Simpan centroids sebelumnya untuk cek konvergensi
		copy(prevCentroids, centroids)

		// Perbarui centroid dengan menghitung rata-rata setiap cluster
		count := make([]int, k)
		sum := make([]float64, k)

		for i, point := range data {
			clusterIdx := clusters[i]
			sum[clusterIdx] += float64(point)
			count[clusterIdx]++
		}

		for j := 0; j < k; j++ {
			if count[j] > 0 {
				centroids[j] = sum[j] / float64(count[j])
			}
		}

		// Cek apakah centroid sudah konvergen (tidak berubah)
		converged := true
		for j := 0; j < k; j++ {
			if math.Abs(centroids[j]-prevCentroids[j]) > 1e-6 {
				converged = false
				break
			}
		}

		// Jika sudah konvergen, berhenti
		if converged {
			break
		}
	}

	// Urutkan centroid beserta cluster assignment
	sortedCentroids := make([]float64, k)
	copy(sortedCentroids, centroids)

	// Simpan urutan index centroid setelah diurutkan
	sortedIdx := make([]int, k)
	for i := 0; i < k; i++ {
		sortedIdx[i] = i
	}

	// Urutkan centroid beserta indexnya
	sort.Slice(sortedIdx, func(i, j int) bool {
		return centroids[sortedIdx[i]] < centroids[sortedIdx[j]]
	})

	// Update assignment cluster berdasarkan urutan centroid yang baru
	newClusters := make([]int, len(data))
	for i, cluster := range clusters {
		newClusters[i] = indexOf(sortedIdx, cluster)
	}

	return newClusters
}

func (s *ProcessingService) KmeansClustering(year string) error {
	cases, err := s.Repo.GetCasesByYear(year)
	if err != nil {
		return response.SERVICE_INTERR
	}

	if len(*cases) < 3 {
		return response.BADREQ_ERR("Jumlah data kurang dari 3")
	}

	var childCount []int64
	var adultCount []int64
	var maleCount []int64
	var femaleCount []int64
	var totalCount []int64

	for _, item := range *cases {
		childCount = append(childCount, item.ChildCount)
		adultCount = append(adultCount, item.AdultCount)
		maleCount = append(maleCount, item.MaleCount)
		femaleCount = append(femaleCount, item.FemaleCount)
		totalCount = append(totalCount, item.FemaleCount+item.MaleCount)
	}

	childClusters := kMeansClustering(childCount)
	adultClusters := kMeansClustering(adultCount)
	maleClusters := kMeansClustering(maleCount)
	femaleClusters := kMeansClustering(femaleCount)
	totalClusters := kMeansClustering(totalCount)

	var resultModel []models.Result

	for iter, item := range *cases {
		resultModel = append(resultModel, models.Result{
			Cluster: byte(childClusters[iter]),
			Type:    "child",
			CaseID:  item.ID,
		})

		resultModel = append(resultModel, models.Result{
			Cluster: byte(adultClusters[iter]),
			Type:    "adult",
			CaseID:  item.ID,
		})

		resultModel = append(resultModel, models.Result{
			Cluster: byte(maleClusters[iter]),
			Type:    "male",
			CaseID:  item.ID,
		})

		resultModel = append(resultModel, models.Result{
			Cluster: byte(femaleClusters[iter]),
			Type:    "female",
			CaseID:  item.ID,
		})

		resultModel = append(resultModel, models.Result{
			Cluster: byte(totalClusters[iter]),
			Type:    "total",
			CaseID:  item.ID,
		})
	}

	if err := s.Repo.GenerateResult(year, &resultModel); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil

}



// func calculateDistance(centroid []float64, data []int64) map[byte][]int64 {
// 	var clusters = make(map[byte][]int64)
// 	for _, item := range data {
// 		// Hitung  jarak antara data dan centroid
// 		var result []int64
// 		for _, c := range centroid {
// 			distance := math.Abs(float64(item) - c)
// 			result = append(result, int64(distance))
// 		}

// 		// Cek angka terkecil dan indexnya
// 		minValue := result[0]
// 		var minIndex byte = 0
// 		for i := 1; i < len(result); i++ {
// 			if result[i] < minValue {
// 				minValue = result[i]
// 				minIndex = byte(i)
// 			}
// 		}

// 		// Simpan data pada clusters sesuai dengan posisi index
// 		clusters[minIndex] = append(clusters[minIndex], item)
// 	}

// 	return clusters
// }

// func mean(data []int64) float64 {

// }

// func kMeansClustering(data []int64) {
// 	k := 3
// 	centroid := initCentroid(k, data)
// 	clusters := calculateDistance(centroid, data)
// 	log.Println(clusters)
// }