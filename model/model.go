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
		Rating     string  `json:"rating"`
		Percentage float64 `json:"percentage"`
	}
	EstablishmentsResponse struct {
		Establishments []struct {
			FHRSID                   int    `json:"FHRSID"`
			LocalAuthorityBusinessID string `json:"LocalAuthorityBusinessID"`
			RatingValue              string `json:"RatingValue"`
			RatingKey                string `json:"RatingKey"`
		} `json:"establishments"`
	}
)
