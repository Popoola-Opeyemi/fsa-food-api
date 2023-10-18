package client

import (
	"encoding/json"
	"fsa-food-api/model"
	"strconv"

	"io/ioutil"
	"net/http"
	"net/url"
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

func GetEstablishment(id int, pageNumber int) (model.EstablishmentsResponse, error) {
	var response model.EstablishmentsResponse

	baseURL := "https://api.ratings.food.gov.uk/Establishments"

	u, _ := url.Parse(baseURL)
	q := u.Query()
	q.Add("localAuthorityId", strconv.Itoa(id))

	if pageNumber > 1 {
		q.Add("pageNumber", strconv.Itoa(pageNumber))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return response, err
	}

	req.Header.Add("x-api-version", "2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetLocalAuthority(id int) (model.EstablishmentsResponse, error) {

	data := model.EstablishmentsResponse{}
	var e []model.Establishments
	pageNumber := 1

	for {

		resp, err := GetEstablishment(id, pageNumber)
		if err != nil {
			return model.EstablishmentsResponse{}, err
		}
		e = append(e, resp.Establishments...)
		pageNumber++

		if resp.Meta.PageNumber == resp.Meta.TotalPages {
			data.Meta = resp.Meta
			data.Meta.SchemeType = resp.Establishments[0].SchemeType
			break
		}
	}
	data.Establishments = e
	return data, nil

}
