// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: clients/grpc/proto/orders.proto

package investapiconnect

import (
	__ "./"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// OrdersStreamServiceName is the fully-qualified name of the OrdersStreamService service.
	OrdersStreamServiceName = "tinkoff.public.invest.api.contract.v1.OrdersStreamService"
	// OrdersServiceName is the fully-qualified name of the OrdersService service.
	OrdersServiceName = "tinkoff.public.invest.api.contract.v1.OrdersService"
)

// OrdersStreamServiceClient is a client for the
// tinkoff.public.invest.api.contract.v1.OrdersStreamService service.
type OrdersStreamServiceClient interface {
	// Stream сделок пользователя
	TradesStream(context.Context, *connect_go.Request[__.TradesStreamRequest]) (*connect_go.ServerStreamForClient[__.TradesStreamResponse], error)
}

// NewOrdersStreamServiceClient constructs a client for the
// tinkoff.public.invest.api.contract.v1.OrdersStreamService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrdersStreamServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) OrdersStreamServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ordersStreamServiceClient{
		tradesStream: connect_go.NewClient[__.TradesStreamRequest, __.TradesStreamResponse](
			httpClient,
			baseURL+"/tinkoff.public.invest.api.contract.v1.OrdersStreamService/TradesStream",
			opts...,
		),
	}
}

// ordersStreamServiceClient implements OrdersStreamServiceClient.
type ordersStreamServiceClient struct {
	tradesStream *connect_go.Client[__.TradesStreamRequest, __.TradesStreamResponse]
}

// TradesStream calls tinkoff.public.invest.api.contract.v1.OrdersStreamService.TradesStream.
func (c *ordersStreamServiceClient) TradesStream(ctx context.Context, req *connect_go.Request[__.TradesStreamRequest]) (*connect_go.ServerStreamForClient[__.TradesStreamResponse], error) {
	return c.tradesStream.CallServerStream(ctx, req)
}

// OrdersStreamServiceHandler is an implementation of the
// tinkoff.public.invest.api.contract.v1.OrdersStreamService service.
type OrdersStreamServiceHandler interface {
	// Stream сделок пользователя
	TradesStream(context.Context, *connect_go.Request[__.TradesStreamRequest], *connect_go.ServerStream[__.TradesStreamResponse]) error
}

// NewOrdersStreamServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrdersStreamServiceHandler(svc OrdersStreamServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/tinkoff.public.invest.api.contract.v1.OrdersStreamService/TradesStream", connect_go.NewServerStreamHandler(
		"/tinkoff.public.invest.api.contract.v1.OrdersStreamService/TradesStream",
		svc.TradesStream,
		opts...,
	))
	return "/tinkoff.public.invest.api.contract.v1.OrdersStreamService/", mux
}

// UnimplementedOrdersStreamServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrdersStreamServiceHandler struct{}

func (UnimplementedOrdersStreamServiceHandler) TradesStream(context.Context, *connect_go.Request[__.TradesStreamRequest], *connect_go.ServerStream[__.TradesStreamResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tinkoff.public.invest.api.contract.v1.OrdersStreamService.TradesStream is not implemented"))
}

// OrdersServiceClient is a client for the tinkoff.public.invest.api.contract.v1.OrdersService
// service.
type OrdersServiceClient interface {
	// Метод выставления заявки.
	PostOrder(context.Context, *connect_go.Request[__.PostOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error)
	// Метод отмены биржевой заявки.
	CancelOrder(context.Context, *connect_go.Request[__.CancelOrderRequest]) (*connect_go.Response[__.CancelOrderResponse], error)
	// Метод получения статуса торгового поручения.
	GetOrderState(context.Context, *connect_go.Request[__.GetOrderStateRequest]) (*connect_go.Response[__.OrderState], error)
	// Метод получения списка активных заявок по счёту.
	GetOrders(context.Context, *connect_go.Request[__.GetOrdersRequest]) (*connect_go.Response[__.GetOrdersResponse], error)
	// Метод изменения выставленной заявки.
	ReplaceOrder(context.Context, *connect_go.Request[__.ReplaceOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error)
}

// NewOrdersServiceClient constructs a client for the
// tinkoff.public.invest.api.contract.v1.OrdersService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrdersServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) OrdersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ordersServiceClient{
		postOrder: connect_go.NewClient[__.PostOrderRequest, __.PostOrderResponse](
			httpClient,
			baseURL+"/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder",
			opts...,
		),
		cancelOrder: connect_go.NewClient[__.CancelOrderRequest, __.CancelOrderResponse](
			httpClient,
			baseURL+"/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder",
			opts...,
		),
		getOrderState: connect_go.NewClient[__.GetOrderStateRequest, __.OrderState](
			httpClient,
			baseURL+"/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState",
			opts...,
		),
		getOrders: connect_go.NewClient[__.GetOrdersRequest, __.GetOrdersResponse](
			httpClient,
			baseURL+"/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders",
			opts...,
		),
		replaceOrder: connect_go.NewClient[__.ReplaceOrderRequest, __.PostOrderResponse](
			httpClient,
			baseURL+"/tinkoff.public.invest.api.contract.v1.OrdersService/ReplaceOrder",
			opts...,
		),
	}
}

// ordersServiceClient implements OrdersServiceClient.
type ordersServiceClient struct {
	postOrder     *connect_go.Client[__.PostOrderRequest, __.PostOrderResponse]
	cancelOrder   *connect_go.Client[__.CancelOrderRequest, __.CancelOrderResponse]
	getOrderState *connect_go.Client[__.GetOrderStateRequest, __.OrderState]
	getOrders     *connect_go.Client[__.GetOrdersRequest, __.GetOrdersResponse]
	replaceOrder  *connect_go.Client[__.ReplaceOrderRequest, __.PostOrderResponse]
}

// PostOrder calls tinkoff.public.invest.api.contract.v1.OrdersService.PostOrder.
func (c *ordersServiceClient) PostOrder(ctx context.Context, req *connect_go.Request[__.PostOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error) {
	return c.postOrder.CallUnary(ctx, req)
}

// CancelOrder calls tinkoff.public.invest.api.contract.v1.OrdersService.CancelOrder.
func (c *ordersServiceClient) CancelOrder(ctx context.Context, req *connect_go.Request[__.CancelOrderRequest]) (*connect_go.Response[__.CancelOrderResponse], error) {
	return c.cancelOrder.CallUnary(ctx, req)
}

// GetOrderState calls tinkoff.public.invest.api.contract.v1.OrdersService.GetOrderState.
func (c *ordersServiceClient) GetOrderState(ctx context.Context, req *connect_go.Request[__.GetOrderStateRequest]) (*connect_go.Response[__.OrderState], error) {
	return c.getOrderState.CallUnary(ctx, req)
}

// GetOrders calls tinkoff.public.invest.api.contract.v1.OrdersService.GetOrders.
func (c *ordersServiceClient) GetOrders(ctx context.Context, req *connect_go.Request[__.GetOrdersRequest]) (*connect_go.Response[__.GetOrdersResponse], error) {
	return c.getOrders.CallUnary(ctx, req)
}

// ReplaceOrder calls tinkoff.public.invest.api.contract.v1.OrdersService.ReplaceOrder.
func (c *ordersServiceClient) ReplaceOrder(ctx context.Context, req *connect_go.Request[__.ReplaceOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error) {
	return c.replaceOrder.CallUnary(ctx, req)
}

// OrdersServiceHandler is an implementation of the
// tinkoff.public.invest.api.contract.v1.OrdersService service.
type OrdersServiceHandler interface {
	// Метод выставления заявки.
	PostOrder(context.Context, *connect_go.Request[__.PostOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error)
	// Метод отмены биржевой заявки.
	CancelOrder(context.Context, *connect_go.Request[__.CancelOrderRequest]) (*connect_go.Response[__.CancelOrderResponse], error)
	// Метод получения статуса торгового поручения.
	GetOrderState(context.Context, *connect_go.Request[__.GetOrderStateRequest]) (*connect_go.Response[__.OrderState], error)
	// Метод получения списка активных заявок по счёту.
	GetOrders(context.Context, *connect_go.Request[__.GetOrdersRequest]) (*connect_go.Response[__.GetOrdersResponse], error)
	// Метод изменения выставленной заявки.
	ReplaceOrder(context.Context, *connect_go.Request[__.ReplaceOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error)
}

// NewOrdersServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrdersServiceHandler(svc OrdersServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder", connect_go.NewUnaryHandler(
		"/tinkoff.public.invest.api.contract.v1.OrdersService/PostOrder",
		svc.PostOrder,
		opts...,
	))
	mux.Handle("/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder", connect_go.NewUnaryHandler(
		"/tinkoff.public.invest.api.contract.v1.OrdersService/CancelOrder",
		svc.CancelOrder,
		opts...,
	))
	mux.Handle("/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState", connect_go.NewUnaryHandler(
		"/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrderState",
		svc.GetOrderState,
		opts...,
	))
	mux.Handle("/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders", connect_go.NewUnaryHandler(
		"/tinkoff.public.invest.api.contract.v1.OrdersService/GetOrders",
		svc.GetOrders,
		opts...,
	))
	mux.Handle("/tinkoff.public.invest.api.contract.v1.OrdersService/ReplaceOrder", connect_go.NewUnaryHandler(
		"/tinkoff.public.invest.api.contract.v1.OrdersService/ReplaceOrder",
		svc.ReplaceOrder,
		opts...,
	))
	return "/tinkoff.public.invest.api.contract.v1.OrdersService/", mux
}

// UnimplementedOrdersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrdersServiceHandler struct{}

func (UnimplementedOrdersServiceHandler) PostOrder(context.Context, *connect_go.Request[__.PostOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tinkoff.public.invest.api.contract.v1.OrdersService.PostOrder is not implemented"))
}

func (UnimplementedOrdersServiceHandler) CancelOrder(context.Context, *connect_go.Request[__.CancelOrderRequest]) (*connect_go.Response[__.CancelOrderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tinkoff.public.invest.api.contract.v1.OrdersService.CancelOrder is not implemented"))
}

func (UnimplementedOrdersServiceHandler) GetOrderState(context.Context, *connect_go.Request[__.GetOrderStateRequest]) (*connect_go.Response[__.OrderState], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tinkoff.public.invest.api.contract.v1.OrdersService.GetOrderState is not implemented"))
}

func (UnimplementedOrdersServiceHandler) GetOrders(context.Context, *connect_go.Request[__.GetOrdersRequest]) (*connect_go.Response[__.GetOrdersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tinkoff.public.invest.api.contract.v1.OrdersService.GetOrders is not implemented"))
}

func (UnimplementedOrdersServiceHandler) ReplaceOrder(context.Context, *connect_go.Request[__.ReplaceOrderRequest]) (*connect_go.Response[__.PostOrderResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tinkoff.public.invest.api.contract.v1.OrdersService.ReplaceOrder is not implemented"))
}
