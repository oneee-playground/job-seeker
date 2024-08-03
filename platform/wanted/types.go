package wanted

type (
	listResponse struct {
		Data  []listData `json:"data"`
		Links struct {
			Prev *string `json:"prev"`
			Next *string `json:"next"`
		} `json:"links"`
	}

	listData struct {
		ID int `json:"id"`
	}
)

type (
	detailResponse struct {
		Job detailJob `json:"job"`
	}

	detailJob struct {
		ID      int     `json:"id"`
		Status  string  `json:"status"`
		DueTime *string `json:"due_time"`
		Detail  struct {
			Position        string `json:"position"`
			Intro           string `json:"intro"`
			MainTasks       string `json:"main_tasks"`
			Requirements    string `json:"requirements"`
			PreferredPoints string `json:"preferred_points"`
			Benefits        string `json:"benefits"`
			HireRounds      string `json:"hire_rounds"`
		} `json:"detail"`
		Company struct {
			Name string `json:"name"`
		} `json:"company"`
		Address struct {
			Country      string `json:"country"`
			CountryCode  string `json:"country_code"`
			Location     string `json:"location"`
			District     string `json:"district"`
			FullLocation string `json:"full_location"`
		} `json:"address"`
		AnnualTo   int `json:"annual_to"`
		AnnualFrom int `json:"annual_from"`
	}
)
