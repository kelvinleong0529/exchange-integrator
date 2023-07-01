package interactive_broker_third_party

type (
	ListingExchange string
)

const (
	ListingExchangeAmericanStockExchange ListingExchange = `AMEX`
	ListingExchangeBulletinBoard         ListingExchange = `ARCAEDGE`
	ListingExchangeNewYorkStockExchange  ListingExchange = `ARCAEDGE`
	ListingExchangeNASDAQ                ListingExchange = `ISLAND`
	ListingExchangePinkSheets            ListingExchange = `ARCAEDGE`
	ListingExchangeOTCQB                 ListingExchange = `ARCAEDGE`
	ListingExchangeOTCQX                 ListingExchange = `ARCAEDGE`
	ListingExchangeTorontoExchange       ListingExchange = `TSE`
	ListingExchangeVentureExchange       ListingExchange = `VENTURE`
	ListingExchangeNEX                   ListingExchange = `AEQNEO`
)
