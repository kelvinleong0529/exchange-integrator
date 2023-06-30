package interactive_broker_third_party

type (
	ExchangeType string
)

const (
	ExchangeTypeAmericanStockExchange ExchangeType = `AMEX`
	ExchangeTypeBulletinBoard         ExchangeType = `ARCAEDGE`
	ExchangeTypeNewYorkStockExchange  ExchangeType = `ARCAEDGE`
	ExchangeTypeNASDAQ                ExchangeType = `ISLAND`
	ExchangeTypePinkSheets            ExchangeType = `ARCAEDGE`
	ExchangeTypeOTCQB                 ExchangeType = `ARCAEDGE`
	ExchangeTypeOTCQX                 ExchangeType = `ARCAEDGE`
	ExchangeTypeTorontoExchange       ExchangeType = `TSE`
	ExchangeTypeVentureExchange       ExchangeType = `VENTURE`
	ExchangeTypeNEX                   ExchangeType = `AEQNEO`
)
