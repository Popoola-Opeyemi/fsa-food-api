package helpers

import (
	"fmt"
	"fsa-food-api/model"
)

func ProcessData(data model.EstablishmentsResponse) (
	[]model.Ratings, int, error) {

	total := data.Meta.TotalCount
	ratingCounts := make(map[string]int)
	for _, establishment := range data.Establishments {
		ratingCounts[establishment.RatingValue]++
	}

	var result []model.Ratings
	for rating, count := range ratingCounts {
		result = append(result, model.Ratings{Rating: rating, Count: count})
	}

	fmt.Println("result", result)
	fmt.Println("total", total)

	return result, total, nil
}

func GetPercentages(ratings []model.Ratings, count int) (
	[]model.RatingPercentage, error) {

	percentages := []model.RatingPercentage{}

	for _, itm := range ratings {
		percentage := float64((itm.Count)) / float64(count) * 100
		percentages = append(percentages, model.RatingPercentage{
			Name:  itm.Rating,
			Value: percentage,
		})
	}

	return percentages, nil

}
