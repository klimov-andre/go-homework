package metrics

import "expvar"

var (
	GatewayTotalAddRequests = expvar.NewInt("gateway.totalAddRequests")
	GatewayTotalUpdateRequests = expvar.NewInt("gateway.totalUpdateRequests")
	GatewayTotalDeleteRequests = expvar.NewInt("gateway.totalDeleteRequests")
	GatewayTotalGetOneRequests = expvar.NewInt("gateway.totalGetOneRequests")
	GatewayTotalListRequests = expvar.NewInt("gateway.totalListRequests")
)

