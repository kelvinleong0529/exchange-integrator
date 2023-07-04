package interactive_broker_client_portal

import (
	"exchange-integrator/utils/network"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// When connected to an IServer Brokerage Session, this endpoint will allow you to submit orders.
// CP WEB API supports various advanced orderTypes
//
// 1) Bracket - Attach additional opposite-side order(s) by using a single cOID sent with the parent and set the
// same value for parentId in each child order(s).
//
// 2) Cash Quantity - Send orders using monetary value by specifying cashQty instead of quantity, e.g. cashQty: 200.
// The endpoint /iserver/contract/rules returns list of valid orderTypes in cqtTypes.
//
// 3) Currency Conversion - Convert cash from one currency to another by including isCcyConv = true. To specify the
// cash quantity use fxQTY instead of quantity, e.g. fxQTY: 100.
//
// 4) Fractional - Contracts that support fractional shares can be traded by specifying quantity as a float,
// e.g. quantity: 0.001. The endpoint /iserver/contract/rules returns a list of valid orderTypes in fraqTypes.
//
// 5) IB Algos - Attached user-defined settings to your trades by using any of IBKR's Algo Orders. Use the endpoint
// /iserver/contract/{conid}/algos to identify the available strategies for a contract.
//
// 6) One-Cancels-All (OCA) - Group multiple unrelated orders by passing order request info in an array and including
// isSingleGroup = true for each order. All orders will be assigned the same oca_group_id.
func (i *InteractiveBrokerClientPortalManager) PlaceOrders(queryParam *PlaceOrdersReq) (*PlaceOrdersRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf(endpointPlaceOrders, queryParam.AccountID),
	}
	result := new(PlaceOrdersRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

func (i *InteractiveBrokerClientPortalRequest) OrderStatus(queryParam *OrderStatusReq) (*OrderStatusRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf(endpointOrderStatus, queryParam.OrderID),
	}
	result := new(OrderStatusRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// Modifies an open order. Must call /iserver/accounts endpoint prior to modifying an order.
// Use /iservers/account/orders endpoint to review open-order(s).
func (i *InteractiveBrokerClientPortalManager) ModifyOrder(queryParam *ModifyOrderReq) ([]*ModifyOrderRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf(endpointUpdateOrder, queryParam.AccountID, queryParam.OrderID),
	}
	result := make([]*ModifyOrderRes, 0)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}

// Cancels an open order. Must call /iserver/accounts endpoint prior to cancelling an order.
// Use /iservers/account/orders endpoint to review open-order(s) and get latest order status.
func (i *InteractiveBrokerClientPortalManager) CancelOrder(queryParam *CancelOrderReq) (*CancelOrderRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	request := &InteractiveBrokerClientPortalRequest{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf(endpointUpdateOrder, queryParam.AccountID, queryParam.OrderID),
	}
	result := new(CancelOrderRes)
	return result, network.SubmitRequestAndUnmarshalResponse(request, result)
}
