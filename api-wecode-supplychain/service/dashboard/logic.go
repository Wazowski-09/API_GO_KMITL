package dashboard

import (
	"time"
)

func checkHeartbeat() (result heartbeatModel, err error) {

	result.Message = "Pong"
	result.DateTime = time.Now()

	//err = logHeartbeat(result)
	//if err != nil {
	//	return
	//}

	return
}

func GetDataDashboard(request inputDashboard) (result []outputDashboard, err error) {
	result, err = getDataDashboard(request)
	if err != nil {
		return
	}

	return
}

func BuyerDashboard(request inputBuyer) (result []outputDashboard, err error) {
	result, err = buyerDashboard(request)
	if err != nil {
		return
	}
	return
}
