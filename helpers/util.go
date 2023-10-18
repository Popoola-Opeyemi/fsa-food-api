package helpers

import (
	"fsa-food-api/model"
	"sync"
)

var HISRating = []string{"Pass and Eat Safe", "Pass", "Exempt", "Improvement Required"}
var HRSRating = []string{"1", "2", "3", "4", "5", "6", "Exempt"}

func ProcessData(data model.EstablishmentsResponse) ([]model.Ratings, int, error) {
	total := data.Meta.TotalCount
	ratingCounts := make(map[string]int)
	ratingCountsMutex := sync.RWMutex{}

	var wg sync.WaitGroup

	for _, establishment := range data.Establishments {
		wg.Add(1)
		go func(establishment model.Establishments) {
			defer wg.Done()

			ratingCountsMutex.Lock()
			ratingCounts[establishment.RatingValue]++
			ratingCountsMutex.Unlock()
		}(establishment)
	}

	if data.Meta.SchemeType == "FHIS" {
		ratingCounts = filterMap(ratingCounts, HISRating)
	}

	if data.Meta.SchemeType == "FHRS" {
		ratingCounts = filterMap(ratingCounts, HRSRating)
	}

	wg.Wait()

	result := make([]model.Ratings, 0)

	ratingCountsMutex.RLock()
	for rating, count := range ratingCounts {
		result = append(result, model.Ratings{Rating: rating, Count: count})
	}
	ratingCountsMutex.RUnlock()

	return result, total, nil
}

func GetPercentages(ratings []model.Ratings, count int) (
	[]model.RatingPercentage, error) {

	percentages := make([]model.RatingPercentage, len(ratings))

	for idx, itm := range ratings {
		percentage := float64((itm.Count)) / float64(count) * 100
		percentages[idx] = model.RatingPercentage{
			Name:  itm.Rating,
			Value: percentage,
		}
	}

	return percentages, nil

}

func filterMap(inputMap map[string]int, allowedKeys []string) map[string]int {
	resultMap := make(map[string]int)
	for _, key := range allowedKeys {
		if val, ok := inputMap[key]; ok {
			resultMap[key] = val
		}
	}
	return resultMap
}
