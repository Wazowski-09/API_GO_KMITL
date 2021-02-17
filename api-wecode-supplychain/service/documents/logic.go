package documents

//func checkHeartbeat() (result heartbeatModel, err error) {
//
//	result.Message = "Pong"
//	result.DateTime = time.Now()
//
//	//err = logHeartbeat(result)
//	//if err != nil {
//	//	return
//	//}
//
//	return
//}

//func getDocumentDetail(request documentDetailReq) (result documentDetailRes, err error) {
//	var detail []documentDetail
//	detail, err = getDocDetail(request)
//	if err != nil {
//		return
//	}
//	log.Infof("documentDetail : %+v", detail)
//
//	var header documentHeader
//	header, err = getDocHeader(request)
//	if err != nil {
//		return
//	}
//	log.Infof("documentDetail : %+v", header)
//
//	result = documentDetailRes{
//		detail,
//		header,
//	}
//
//	return
//}

//func getPODocumentList(request documentListReq) (result documentListRes, err error) {
//	var list []documentList
//	list, err = getPODocList(request)
//	if err != nil {
//		return
//	}
//	log.Infof("documentList : %+v", list)
//
//	var count PoCount
//	count, err = getPODocListCount(request)
//	if err != nil {
//		return
//	}
//	log.Infof("documentListCount : %+v", list)
//
//	result = documentListRes{
//		count,
//		list,
//		// header,
//	}
//
//	return
//}

func getGRDocumentList(request inputBuyerGR) (result docListGR,err error){
	//var gr int = GoodsReceive

	result, err = buyerDocGR( request)
	if err != nil {
		return
	}
	return
}

