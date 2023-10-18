package helpers

import (
	"fmt"
	"fsa-food-api/model"
	"sort"
	"strconv"
	"sync"
)

var HISRating = map[string]struct{}{
	"Pass and Eat Safe":    {},
	"Pass":                 {},
	"Improvement Required": {},
}

var HRSRating = map[string]struct{}{
	"1":      {},
	"2":      {},
	"3":      {},
	"4":      {},
	"5":      {},
	"6":      {},
	"Exempt": {},
}

func ProcessData(data model.EstablishmentsResponse) ([]model.Ratings, int, error) {
	ratingCounts := countRatings(data.Establishments)
	ratingCounts = applyFilter(data.Meta.SchemeType, ratingCounts)

	// Calculate the total for the filtered data
	filteredTotal := calculateTotal(ratingCounts)

	result := convertToSortedRatings(ratingCounts, filteredTotal)
	return result, filteredTotal, nil
}

func countRatings(establishments []model.Establishments) map[string]int {
	ratingCounts := make(map[string]int)
	ratingCountsMutex := sync.RWMutex{}
	var wg sync.WaitGroup

	for _, establishment := range establishments {
		wg.Add(1)
		go func(establishment model.Establishments) {
			defer wg.Done()
			ratingCountsMutex.Lock()
			ratingCounts[establishment.RatingValue]++
			ratingCountsMutex.Unlock()
		}(establishment)
	}

	wg.Wait()
	return ratingCounts
}

func applyFilter(schemeType string, ratingCounts map[string]int) map[string]int {
	switch schemeType {
	case "FHIS":
		return filterMap(ratingCounts, HISRating)
	case "FHRS":
		return filterMap(ratingCounts, HRSRating)
	default:
		return ratingCounts
	}
}

func calculateTotal(ratingCounts map[string]int) int {
	total := 0
	for _, count := range ratingCounts {
		total += count
	}
	return total
}

func convertToSortedRatings(ratingCounts map[string]int, total int) []model.Ratings {
	ratingsSlice := make([]model.Ratings, 0)

	for rating, count := range ratingCounts {
		ratingsSlice = append(ratingsSlice, model.Ratings{Rating: rating, Count: count})
	}

	sort.Slice(ratingsSlice, func(i, j int) bool {
		// Sort in descending order (from top to bottom)
		return ratingsSlice[i].Count > ratingsSlice[j].Count
	})

	return ratingsSlice
}

func GetPercentages(ratings []model.Ratings, count int) []model.RatingPercentage {
	percentages := make([]model.RatingPercentage, len(ratings))

	for idx, itm := range ratings {
		percentage := float64(itm.Count) / float64(count) * 100
		val, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", percentage), 64)
		percentages[idx] = model.RatingPercentage{
			Name:  itm.Rating,
			Value: val,
		}
	}

	return percentages
}

func filterMap(inputMap map[string]int, allowedKeys map[string]struct{}) map[string]int {
	resultMap := make(map[string]int)
	for key := range inputMap {
		if _, ok := allowedKeys[key]; ok {
			resultMap[key] = inputMap[key]
		}
	}
	return resultMap
}
