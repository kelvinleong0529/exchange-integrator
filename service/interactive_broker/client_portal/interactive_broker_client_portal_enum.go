package interactive_broker_client_portal

type (
	LoginType uint
	Period    string
	KlineBar  string
	Currency  string

	// portfolio
	AccountDisplayName string
	AccountType        string
	AccountStatus      string

	// order
	ListingExchange       string
	OrderType             string
	OrderSide             string
	OrderAllocationMethod string
	OrderStatus           string
	OrderStatusSide       string
	TimeInForce           string
)

const (
	LoginTypeLive LoginType = iota + 1
	LoginTypePaper

	PeriodMinutes Period = `min`
	PeriodHours   Period = `h`
	PeriodDays    Period = `d`
	PeriodWeeks   Period = `w`
	PeriodMonths  Period = `m`
	PeriodYears   Period = `y`

	KlineBarMinutes KlineBar = `min`
	KlineBarHours   KlineBar = `h`
	KlineBarDays    KlineBar = `d`
	KlineBarWeeks   KlineBar = `w`
	KlineBarMonths  KlineBar = `m`

	CurrencyAUD Currency = `AUD`
	CurrencyGBP Currency = `GBP`
	CurrencyCAD Currency = `CAD`
	CurrencyCNH Currency = `CNH`
	CurrencyCZK Currency = `CZK`
	CurrencyDKK Currency = `DKK`
	CurrencyEUR Currency = `EUR`
	CurrencyHKD Currency = `HKD`
	CurrencyHUF Currency = `HUF`
	CurrencyINR Currency = `INR`
	CurrencyILS Currency = `ILS`
	CurrencyJPY Currency = `JPY`
	CurrencyMXN Currency = `MXN`
	CurrencyNOK Currency = `NOK`
	CurrencyNZD Currency = `NZD`
	CurrencyPLN Currency = `PLN`
	CurrencyRUB Currency = `RUB`
	CurrencySGD Currency = `SGD`
	CurrencySEK Currency = `SEK`
	CurrencyCHF Currency = `CHF`
	CurrencyUSD Currency = `USD`

	// portfolio
	AccountDisplayNameAccountTitle AccountDisplayName = `accountTitle`
	AccountDisplayNameAccountVan   AccountDisplayName = `accountVan`
	AccountDisplayNameAccountID    AccountDisplayName = `accountId`

	AccountTypeIndividual   AccountType = `INDIVIDUAL`
	AccountTypeJoint        AccountType = `JOINT`
	AccountTypeOrganization AccountType = `ORG`
	AccountTypeTrust        AccountType = `TRUST`
	AccountTypeDemo         AccountType = `DEMO`

	AccountStatusOpen      AccountStatus = `O`
	AccountStatusPendingP  AccountStatus = `P`
	AccountStatusPendingN  AccountStatus = `N`
	AccountStatusAbandoned AccountStatus = `A`
	AccountStatusRejected  AccountStatus = `R`
	AccountStatusClosed    AccountStatus = `C`

	// order
	ListingExchangeSMART    ListingExchange = `SMART`
	ListingExchangeAMEX     ListingExchange = `AMEX`
	ListingExchangeNYSE     ListingExchange = `NYSE`
	ListingExchangeCBOE     ListingExchange = `CBOE`
	ListingExchangeISE      ListingExchange = `ISE`
	ListingExchangeCHX      ListingExchange = `CHX`
	ListingExchangeARCA     ListingExchange = `ARCA`
	ListingExchangeISLAND   ListingExchange = `ISLAND`
	ListingExchangeDRCTEDGE ListingExchange = `DRCTEDGE`
	ListingExchangeBEX      ListingExchange = `BEX`
	ListingExchangeBATS     ListingExchange = `BATS`
	ListingExchangeEDGEA    ListingExchange = `EDGEA`
	ListingExchangeCSFBALGO ListingExchange = `CSFBALGO`
	ListingExchangeJEFFALGO ListingExchange = `JEFFALGO`
	ListingExchangeBYX      ListingExchange = `BYX`
	ListingExchangeIEX      ListingExchange = `IEX`
	ListingExchangeFOXRIVER ListingExchange = `FOXRIVER`
	ListingExchangeTPLUS1   ListingExchange = `TPLUS1`
	ListingExchangeNYSENAT  ListingExchange = `NYSENAT`
	ListingExchangePSX      ListingExchange = `PSX`

	OrderTypeLimit      OrderType = `LMT`
	OrderTypeMarket     OrderType = `MKT`
	OrderTypeStop       OrderType = `STP`
	OrderTypeStopLimit  OrderType = `STOP_LIMIT`
	OrderTypeMidPrice   OrderType = `MIDPRICE`
	OrderTypeTrail      OrderType = `TRAIL`
	OrderTypeTrailLimit OrderType = `TRAILLMT`

	OrderSideSell OrderSide = `SELL`
	OrderSideBuy  OrderSide = `BUY`

	OrderAllocationMethodNetLiquidity    OrderAllocationMethod = `NetLiquidity`
	OrderAllocationMethodAvailableEquity OrderAllocationMethod = `AvailableEquity`
	OrderAllocationMethodEqualQuantity   OrderAllocationMethod = `EqualQuantity`
	OrderAllocationMethodPercentChange   OrderAllocationMethod = `PctChange`

	OrderStatusPendingSubmit OrderStatus = `PendingSubmit`
	OrderStatusPendingCancel OrderStatus = `PendingCancel`
	OrderStatusSubmitted     OrderStatus = `Submitted`
	OrderStatusFilled        OrderStatus = `Filled`

	OrderStatusSideBuy     OrderStatusSide = `B`
	OrderStatusSideBSell   OrderStatusSide = `S`
	OrderStatusSideExpired OrderStatusSide = `X`

	TimeInForceGoodTillCancel     TimeInForce = `GTC`
	TimeInForceOpenPriceGuarantee TimeInForce = `OPG`
	TimeInForceDay                TimeInForce = `DAY`
	TimeInForceImmediateOrCancel  TimeInForce = `IOC`

	// endpoints
	endpointKeepaliveSession     = `/tickle`
	endpointTerminateSession     = `/logout`
	endpointValidateSSO          = `/sso/validate`
	endpointAuthenticationStatus = `/iserver/auth/status`
	endpointReauthenticate       = `iserver/reauthenticate`

	endpointMarketDataHistory = `hmds/history`

	endpointPortfolioAccounts  = `portfolio/accounts`
	endpointAccountInformation = `portfolio/%s/meta`
	endpointSwitchAccount      = `iserver/account`

	endpointPlaceOrders = `/iserver/account/%s/orders`
	endpointOrderStatus = `/iserver/account/order/status/%s`
	endpointUpdateOrder = `/iserver/account/%s/order/%s`
)
