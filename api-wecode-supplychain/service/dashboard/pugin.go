package dashboard

import (
	mssql "api-wecode-supplychain/database/mssql"
)

func logHeartbeat(hb heartbeatModel) (err error) { //sql
	if err = mssql.DB.Table("document_header").Save(hb).Error; err != nil {
		return
	}
	return
}

func buyerDashboard(request inputBuyer) (result []outputDashboard, err error) {
	if err = mssql.DB.Select("document_type ,document_status, COUNT(document_status) as count_doc , COUNT(*) as active_count , ROUND(SUM(grand_total),2) as total").
		Table("document_header AS dh").
		Where("(document_status in ('1','3','5','7','9','11')) and currency = 'THB' and document_type ='110' and buyer_company_id = ? and create_date BETWEEN ? and ?", request.BuyerCompanyID, request.DateFrom, request.DateTo).
		Group("document_type,document_status").
		Order("document_type").
		Find(&result).Error; err != nil {
		return
	}
	return
}

func getDataDashboard(request inputDashboard) (result []outputDashboard, err error) {

	if err = mssql.DB.Select("document_type,document_status , COUNT(*) as count_doc ,SUM(grand_total) as total").
		Table("document_header AS dh").
		//Joins("LEFT JOIN config AS cf ON c.company_type = cf.status_code").
		Where("((document_type = '110' and document_status in ('3','5','9','11','13')) or (document_type = '210' and document_status in ('1','5','9','11','13'))) and currency = 'THB' and create_date BETWEEN ? and ? ", request.DateTo, request.DateFrom).
		Group("document_type,document_status").Order("document_type").
		Find(&result).Error; err != nil {
		return
	}

	return
}
