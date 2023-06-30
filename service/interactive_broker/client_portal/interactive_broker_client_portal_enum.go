package interactive_broker_client_portal

type (
	LoginType uint
	Period    string
	KlineBar  string
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

	// endpoints
	endpointKeepaliveSession     = `/tickle`
	endpointTerminateSession     = `/logout`
	endpointValidateSSO          = `/sso/validate`
	endpointAuthenticationStatus = `/iserver/auth/status`
	endpointReauthenticate       = `iserver/reauthenticate`

	endpointMarketDataHistory = `hmds/history`
)
