package interactive_broker_third_party

type (
	InteractiveBrokerThirdPartyManager struct {
	}

	InteractiveBrokerThirdPartyRequest struct {
		manager    *InteractiveBrokerThirdPartyManager
		method     string
		endpoint   string
		queryParam string
	}

	// OAuth
	ObtainRequestTokenReq struct {
		ConsumerKey     string
		SignatureMethod string
		Signature       string
		Timestamp       string
		Nonce           string
		Callback        string
	}

	ObtainRequestTokenRes struct {
		Token string `json:"oauth_token"`
	}

	ObtainAccessTokenReq struct {
		ConsumerKey     string
		Token           string
		SignatureMethod string
		Signature       string
		Timestamp       string
		Nonce           string
		Verifier        string
	}

	ObtainAccessTokenRes struct {
		Token       string `json:"oauth_token"`
		TokenSecret string `json:"oauth_token_secret"`
	}

	ObtainLiveSessionTokenReq struct {
		ConsumerKey            string
		SignatureMethod        string
		Signature              string
		Timestamp              string
		Nonce                  string
		DiffieHellmanChallenge string
	}

	ObtainLiveSessionTokenRes struct {
		DiffieHellmanResponse     string `json:"diffie_hellman_response"`
		LiveSessionTokenSignature string `json:"live_session_token_signature"`
	}

	// financial instruments
	FinancialInstrumentsReq struct {
		Type       string
		Symbol     string
		Currency   string
		ContractID uint64
		ListingExchange
	}

	FinancialInstrumentsRes struct {
		Results []FinancialInstrumentsData
	}

	FinancialInstrumentsData struct {
		ContractID      uint64 `json:"ContractId"`
		Ticker          string `json:"Ticker"`
		SecurityType    string `json:"SecurityType"`
		CompanyName     string `json:"CompanyName"`
		Currency        string `json:"Currency"`
		ListingExchange `json:"Exchange"`
	}
)
