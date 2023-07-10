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

	// portfolio
	PortfolioAccountsRes struct {
		ID                     string                 `json:"id"`
		AccountID              string                 `json:"accountId"`
		AccountVan             string                 `json:"accountVan"`
		AccountTitle           string                 `json:"accountTitle"`
		AccountAlias           string                 `json:"accountAlias"`
		AccountOpenTime        uint64                 `json:"accountStatus"`
		TradingType            string                 `json:"tradingType"`
		FinancialAdvisorClient bool                   `json:"faclient"`
		CovestorAccount        bool                   `json:"covestor"`
		Parent                 PortfolioAccountParent `json:"parent"`
		Desc                   string                 `json:"desc"`
		DisplayName            AccountDisplayName     `json:"displayName"`
		Currency               `json:"currency"`
		AccountType            `json:"type"`
		AccountStatus          `json:"clearingStatus"`
	}

	PortfolioAccountParent struct {
		MMC         []string `json:"mmc"`
		AccountID   string   `json:"accountId"`
		IsMParent   bool     `json:"isMParent"`
		IsMChild    bool     `json:"isMChild"`
		IsMultiplex bool     `json:"isMultiplex"`
	}

	AccountInformationReq struct {
		AccountID string `validate:"required"`
	}

	SwitchAccountReq struct {
		AccountID string
	}

	SwitchAccountRes struct {
		Set       bool   `json:"set"`
		AccountID string `json:"acctId"`
	}

	// order
	PlaceOrdersReq struct {
		AccountID                  string `validate:"required"`
		ContractID                 string
		ContractIDExchange         string
		SecurityType               string
		CustomerOrderID            string
		ParentID                   string
		IsSingleGroup              bool
		OutsideRegularTradingHours bool
		Price                      float64
		AuxPrice                   float64
		Ticker                     string
		TrailingAmount             float64
		Referrer                   string
		Quantity                   float64
		CashQuantity               float64
		ForexQuantity              float64
		IsCurrencyConversion       bool
		Strategy                   string
		StrategyParameters         interface{}
		ListingExchange
		OrderType
		OrderSide
		OrderAllocationMethod
		TimeInForce
	}

	PlaceOrdersRes struct {
		ID      string   `json:"id"`
		Message []string `json:"message"`
	}

	OrderStatusReq struct {
		OrderID string `validate:"required"`
	}

	OrderStatusRes struct {
		SubType                      string          `json:"sub_type"`
		RequestID                    string          `json:"request_id"`
		OrderID                      uint64          `json:"order_id"`
		ContractIDExchange           string          `json:"conidex"`
		Symbol                       string          `json:"symbol"`
		Side                         OrderStatusSide `json:"status"`
		ContractDescription1         string          `json:"contract_description_1"`
		AccountID                    string          `json:"account"`
		OptionAccount                string          `json:"option_acct"`
		CompanyName                  string          `json:"company_name"`
		Size                         float64         `json:"size,string"`
		TotalSize                    float64         `json:"total_size,string"`
		Currency                     string          `json:"currency"`
		LimitPrice                   float64         `json:"limit_price,string"`
		StopPrice                    float64         `json:"stop_price,string"`
		CumulativeFill               float64         `json:"cum_fill,string"`
		OrderStatusDescription       string          `json:"order_status_description"`
		ForegroundColor              string          `json:"fg_color"`
		BackgroundColor              string          `json:"bg_color"`
		OrderNotEditable             bool            `json:"order_not_editable"`
		EditableFields               string          `json:"editable_fields"`
		CannotCancelOrder            bool            `json:"cannot_cancel_order"`
		OutsideRegularTradingHours   bool            `json:"outside_rth"`
		DeactivateOrder              bool            `json:"deactivate_order"`
		UsePriceManagementAlgo       bool            `json:"use_price_mgmt_algo"`
		SecurityType                 string          `json:"sec_type"`
		AvailableChartPeriods        string          `json:"available_chart_periods"`
		OrderDescription             string          `json:"order_description"`
		OrderDescriptionWithContract string          `json:"order_description_with_contract"`
		AlertActive                  uint64          `json:"alert_active"`
		ChildOrderType               OrderType       `json:"child_order_type"`
		SizeAndFills                 string          `json:"size_and_fills"`
		ExitStrategyDisplayPrice     float64         `json:"exit_strategy_display_price,string"`
		ExitStrategyChartDescription string          `json:"exit_strategy_chart_description"`
		ExitStrategyToolAvailability string          `json:"exit_strategy_tool_availability"`
		AllowedDuplicateOpposite     bool            `json:"allowed_duplicate_opposite"`
		OrderTimeUnix                uint64          `json:"order_time"`
		OneCancelsAllGroupID         string          `json:"oca_group_id"`
		ListingExchange              `json:"listing_exchange"`
		OrderStatus                  `json:"order_status"`
		TimeInForce                  `json:"tif"`
	}

	ModifyOrderReq struct {
		AccountID                  string `validate:"required"`
		OrderID                    string `validate:"required"`
		ContractID                 string
		OutsideRegularTradingHours bool
		Price                      float64
		AuxPrice                   float64
		Ticker                     string
		Quantity                   float64
		Deactivate                 bool
		OrderType
		OrderSide
		ListingExchange
		TimeInForce
	}

	ModifyOrderRes struct {
		OrderID      string `json:"order_id"`
		LocalIrderID string `json:"local_order_id"`
		OrderStatus  `json:"order_status"`
	}

	CancelOrderReq struct {
		AccountID string `validate:"required"`
		OrderID   string `validate:"required"`
	}

	CancelOrderRes struct {
		OrderID    string `json:"order_id"`
		Message    string `json:"msg"`
		ContractID uint64 `json:"conid"`
		Account    string `json:"account"`
	}

	// alert
	CreateOrModifyAlertReq struct {
		AccountID                   string `validate:"required"`
		OrderID                     uint64
		AlertName                   string
		AlertMessage                string
		AlertRepeatable             uint8
		Email                       string
		SendMessage                 uint8
		ExpireTime                  string
		OutsideRegularTradingHours  uint8
		TraderWorkstationOrdersOnly uint8
		ShowPopup                   uint8
		ToolID                      uint8
		PlayAudio                   string
		Conditions                  []struct {
			AlertType
			ContractIDExchange string
			Operator           string `validate:"oneof=>= <="`
			TriggerMethod      string
			Value              string `validate:"required"`
			AlertLogicBind
			TimeZone string
		}
		TimeInForce
	}

	CreateOrModifyAlertRes struct {
		RequestID      uint64 `json:"request_id"`
		OrderID        uint64 `json:"order_id"`
		Success        bool   `json:"success"`
		Text           string `json:"text"`
		OrderStatus    string `json:"order_status"`
		WarningMessage string `json:"warning_message"`
	}
)
