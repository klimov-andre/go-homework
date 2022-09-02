package metrics

import "expvar"

var (
	GatewayTotalAddRequests = expvar.NewInt("gateway.totalAddRequests")
	GatewayTotalUpdateRequests = expvar.NewInt("gateway.totalUpdateRequests")
	GatewayTotalDeleteRequests = expvar.NewInt("gateway.totalDeleteRequests")
	GatewayTotalGetOneRequests = expvar.NewInt("gateway.totalGetOneRequests")
	GatewayTotalListRequests = expvar.NewInt("gateway.totalListRequests")

	GatewaySuccessAddRequests = expvar.NewInt("gateway.successAddRequests")
	GatewaySuccessUpdateRequests = expvar.NewInt("gateway.successUpdateRequests")
	GatewaySuccessDeleteRequests = expvar.NewInt("gateway.successDeleteRequests")
	GatewaySuccessGetOneRequests = expvar.NewInt("gateway.successGetOneRequests")
	GatewaySuccessListRequests = expvar.NewInt("gateway.successListRequests")

	GatewayUnsuccessfulAddRequests = expvar.NewInt("gateway.unsuccessfulAddRequests")
	GatewayUnsuccessfulUpdateRequests = expvar.NewInt("gateway.unsuccessfulUpdateRequests")
	GatewayUnsuccessfulDeleteRequests = expvar.NewInt("gateway.unsuccessfulDeleteRequests")
	GatewayUnsuccessfulGetOneRequests = expvar.NewInt("gateway.unsuccessfulGetOneRequests")
	GatewayUnsuccessfulListRequests = expvar.NewInt("gateway.unsuccessfulListRequests")

	GatewayInvalidAddRequests = expvar.NewInt("gateway.invalidAddRequests")
	GatewayInvalidUpdateRequests = expvar.NewInt("gateway.invalidUpdateRequests")
	GatewayInvalidDeleteRequests = expvar.NewInt("gateway.invalidDeleteRequests")
	GatewayInvalidGetOneRequests = expvar.NewInt("gateway.invalidGetOneRequests")
	GatewayInvalidListRequests = expvar.NewInt("gateway.invalidListRequests")

)

