package interactive_broker_client_portal

type (
	InteractiveBrokerClientPortalManager struct{}

	InteractiveBrokerClientPortalRequest struct {
		endpoint string
		method   string
	}

	// session
	TerminateSessionRes struct {
		Confirmed bool `json:"confirmed"`
	}

	ValidateSSORes struct {
		LoginType `json:"LOGIN_TYPE"`
		Username  string `json:"USER_NAME"`
		UserID    uint64 `json:"USER_ID"`
		Expire    uint64 `json:"EXPIRE"` // time in milliseconds until session expires
		Result    bool   `json:"RESULT"` // true is session was validated, false if not
		AuthTime  uint64 `json:"AUTH_TIME"`
	}

	AuthenticationStatusRes struct {
		Authenticated bool     `json:"authenticated"`
		Connected     bool     `json:"connected"`
		Competing     bool     `json:"competing"`
		Fail          string   `json:"fail"`
		Message       string   `json:"message"`
		Prompts       []string `json:"prompts"`
	}

	// market data
	MarketDataHistoryReq struct {
		ContractID                 string `validate:"required"`
		OutsideRegularTradingHours bool
		Period                     `validate:"required"`
		KlineBar
	}

	MarketDataHistoryRes struct {
		Open            float64                 `json:"open"`
		StartTime       string                  `json:"startTime"`
		StartTimeUnix   uint64                  `json:"startTimeVal"`
		EndTime         string                  `json:"endTime"`
		EndTimeUnix     uint64                  `json:"endTimeVal"`
		Points          uint64                  `json:"points"`
		Data            []MarketDataHistoryData `json:"data"`
		MarketDataDelay uint64                  `json:"mktDataDelay"`
	}

	MarketDataHistoryData struct {
		TimeUnix uint64  `json:"t"`
		Open     float64 `json:"o"`
		Close    float64 `json:"c"`
		High     float64 `json:"h"`
		Low      float64 `json:"l"`
		Volume   float64 `json:"v"`
	}
)
