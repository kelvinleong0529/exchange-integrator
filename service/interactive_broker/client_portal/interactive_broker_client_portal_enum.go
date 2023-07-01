package interactive_broker_client_portal

type (
	LoginType uint
	Period    string
	KlineBar  string

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

	endpointPlaceOrders = `/iserver/account/%s/orders`
	endpointOrderStatus = `/iserver/account/order/status/%s`
)
