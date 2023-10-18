package client

import (
	"encoding/json"
	"fmt"
	"fsa-food-api/model"

	"io/ioutil"
	"net/http"
)

func GetAuthorities() (model.FSAAuthorities, error) {
	var fsaAuthorities model.FSAAuthorities

	req, err := http.NewRequest(http.MethodGet,
		"http://api.ratings.food.gov.uk/Authorities", nil)
	if err != nil {
		return fsaAuthorities, err
	}

	req.Header.Set("x-api-version", "2")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fsaAuthorities, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fsaAuthorities, err
	}

	err = json.Unmarshal(body, &fsaAuthorities)
	if err != nil {
		return fsaAuthorities, err
	}

	return fsaAuthorities, nil
}

func GetLocalAuthoritymodel(id string) (model.EstablishmentsResponse, error) {
	var response model.EstablishmentsResponse

	url := fmt.Sprintf(`https://api.ratings.food.gov.uk/Establishments?
	localAuthorityId=%s`, id)

	req, err := http.NewRequest(http.MethodGet,
		url, nil)
	if err != nil {
		return response, err
	}

	req.Header.Set("x-api-version", "2")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
