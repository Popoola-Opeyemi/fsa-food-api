package model

type (
	Authority struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	AuthorityRating struct {
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	}
	FSAAuthorities struct {
		Authorities []FSAAuthority `json:"authorities"`
	}

	FSAAuthority struct {
		ID   int    `json:"LocalAuthorityId"`
		Name string `json:"Name"`
	}

	Ratings struct {
		Rating string `json:"rating"`
		Count  int    `json:"count"`
	}

	RatingPercentage struct {
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	}
	Establishments struct {
		LocalAuthorityBusinessID string `json:"LocalAuthorityBusinessID"`
		RatingValue              string `json:"RatingValue"`
		RatingKey                string `json:"RatingKey"`
	}
	Meta struct {
		ItemCount  int `json:"itemCount"`
		TotalPages int `json:"totalPages"`
		PageNumber int `json:"pageNumber"`
		TotalCount int `json:"totalCount"`
	}
	EstablishmentsResponse struct {
		Meta           Meta             `json:"meta"`
		Establishments []Establishments `json:"establishments"`
	}
)
