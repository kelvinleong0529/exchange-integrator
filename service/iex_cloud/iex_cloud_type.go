package iex_cloud

type (
	IEXCloudManager struct {
		apiToken string
	}

	IEXCloudRequest struct {
		endpoint string
		method   string
	}

	CEOCompensationReq struct {
		Symbol string `validate:"required"`
	}

	CEOCompensationRes struct {
		Symbol              string  `json:"symbol"`
		Name                string  `json:"name"`
		CompanyName         string  `json:"companyName"`
		Location            string  `json:"location"`
		Salary              float64 `json:"salary"`
		Bonus               float64 `json:"bonus"`
		StockAwards         float64 `json:"stockAwards"`
		OptionAwards        float64 `json:"optionAwards"`
		NonEquityIncentives float64 `json:"nonEquityIncentives"`
		PensionAndDeferred  float64 `json:"pensionAndDeferred"`
		OtherCompensation   float64 `json:"otherCompensation"`
		Total               float64 `json:"total"`
		Year                string  `json:"year"`
	}
)
