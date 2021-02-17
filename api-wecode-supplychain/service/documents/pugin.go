package documents

import (
	mssql "api-wecode-supplychain/database/mssql"
	log "github.com/sirupsen/logrus"
	"time"
)

//func logHeartbeat(hb heartbeatModel) (err error) { //sql
//	if err = mssql.DB.Table("document_header").Save(hb).Error; err != nil {
//		return
//	}
//	return
//}

//func getDocDetail(request documentDetailReq) (result []documentDetail, err error) {
//
//	if err = mssql.DB.Select("dd.id AS detail_id, dd.sku, dd.product_name, dd.quantity, dd.unit, dd.price_per_unit, dd.total_price").
//		Table("document_detail AS dd").
//		Joins("LEFT JOIN document_header AS dh ON dd.document_header_id  = dh.document_id").
//		Where("dh.document_number = ? AND dh.seller_company_id = ? AND dh.buyer_company_id = ? AND dh.document_type = ? AND dh.record_status =1", request.DocumentNumber, request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType).
//		Find(&result).Error; err != nil {
//		return
//	}
//
//	//if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
//	//	return
//	//}
//
//	return
//}

//func getDocHeader(request documentDetailReq) (result documentHeader, err error) {
//
//	if err = mssql.DB.Select("dh.document_number, c.description AS document_status, dh.create_date, dh.expire_date, dh.delivery_date, dh.net, dh.vat_percent, dh.vat_value, dh.discount, dh.grand_total, dh.currency, dh.delivery_address, cs.company_name AS company_name_seller, cs.short_name AS short_name_seller, cs.branch AS branch_seller, cs.phone_number AS phone_number_seller, cs.district AS district_seller, cs.sub_district AS sub_district_seller, cs.zip_code AS zip_code_seller, cs.province AS province_seller, cb.company_name AS company_name_buyer, cb.short_name AS short_name_buyer, cb.branch AS branch_buyer, cb.phone_number AS phone_number_buyer, cb.district AS district_buyer, cb.sub_district AS sub_district_buyer, cb.zip_code AS zip_code_buyer, cb.province AS province_buyer").
//		Table("document_header AS dh").
//		Joins("LEFT JOIN company AS cs on dh.seller_company_id = cs.id LEFT JOIN company AS cb on dh.seller_company_id = cb.id LEFT JOIN config AS c ON dh.document_status = c.status_code").
//		Where("dh.document_number = ? AND dh.seller_company_id = ? AND dh.buyer_company_id = ? AND dh.document_type = ? AND dh.record_status =1", request.DocumentNumber, request.SellerCompanyId, request.BuyerCompanyId, request.DocumentType).
//		Find(&result).Error; err != nil {
//		return
//	}
//
//	//if err = mssql.DB.Raw("SELECT c.id, c.company_name, c.short_name, c.company_type AS company_type_code, cf.description AS company_type_name, c.branch, c.phone_number, c.district, c.sub_district, c.zip_code, c.province FROM company AS c LEFT JOIN config AS cf ON c.company_type = cf.status_code WHERE id = ?", companyId).Find(&result).Error; err != nil {
//	//	return
//	//}
//
//	return
//}

//func buyerDocGR(request inputBuyerGR) (result []outputBuyerGR, err error) {
//	var zeroValueTime time.Time
//	var zeroValueString string
//	var SortBy string
//	//grSubQuery := mssql.DB
//	//grMainQuery := mssql.DB
//	var getCompany []CompanyInfo
//	getCompany, err = getCompanyById(request.LoginCompanyId)
//	log.Infof("company ",getCompany[0].Id)
//	if getCompany[0].CompanyType != 1003 {
//		return
//	}
//
//	//grSubQuery = grSubQuery.Select("ROW_NUMBER() OVER(ORDER BY (SELECT NULL)) AS sequence,dh.document_number,c.company_name,dh.create_date ,dh.expire_date ,dh.delivery_date ,dh.grand_total,dh.currency,dh.document_id,co.description AS document_status,dh.buyer_company_id AS company_id,dh.seller_company_id").
//	//	Table("malar.dbo.document_header AS dh,malar.dbo.company AS c,malar.dbo.config AS co").
//	//	Where("document_type = ? AND co.status_code = dh.document_status AND dh.seller_company_id = c.id AND dh.buyer_company_id = ?",GoodsReceive, request.LoginCompanyId)
//
//	sql := mssql.DB.Select("ROW_NUMBER() OVER(ORDER BY (SELECT NULL)) AS sequence,dh.document_number,c.company_name,dh.create_date ,dh.expire_date ,dh.delivery_date ,dh.grand_total,dh.currency,dh.document_id,co.description AS document_status,dh.buyer_company_id AS company_id,dh.seller_company_id").
//		Table("malar.dbo.document_header AS dh,malar.dbo.company AS c,malar.dbo.config AS co").
//		Where("document_type = ? AND co.status_code = dh.document_status AND dh.seller_company_id = c.id AND dh.buyer_company_id = ?",GoodsReceive, request.LoginCompanyId)
//
//	switch request.SortBy {
//	case "document_number":
//		SortBy = "dh.document_number"
//	case "currency":
//		SortBy = "dh.currency"
//	case "document_status":
//		SortBy = "dh.document_status"
//	case "seller_company_id":
//		SortBy = "dh.seller_company_id"
//	case "create_date":
//		SortBy = "dh.create_date"
//	case "expire_date":
//		SortBy = "dh.expire_date"
//	case "delivery_date":
//		SortBy = "dh.delivery_date"
//	default :
//		SortBy = "dh.document_number"
//	}
//	if request.SortType != 0 && request.SortType != 1 || request.SortBy == zeroValueString{
//		SortBy = "dh.document_number"
//	}
//	if request.SortType == 1 {
//		SortBy += " DESC"
//	}
//
//
//	if SortBy != zeroValueString {
//		sql = sql.Order(SortBy)
//	}
//
//
//	if request.PagingIndex != 0 && request.PagingSize != 0 {
//		sql = sql.Offset((request.PagingSize * request.PagingIndex) - request.PagingSize)
//	}
//	if request.PagingSize != 0 {
//		sql = sql.Limit(request.PagingSize)
//	}
//	if request.FreeSearchDocumentNumber != zeroValueString {
//		sql = sql.Where("document_number LIKE ?", "%"+request.FreeSearchDocumentNumber+"%")
//	}
//	if len(request.FilterCurrency) != 0 {
//		sql = sql.Where("currency IN( ? )", request.FilterCurrency)
//	}
//	if len(request.FilterCurrency) == 0 {
//		sql = sql.Where("currency = 'THB'")
//	}
//	if len(request.FilterCompanyId) != 0 {
//		sql = sql.Where("seller_company_id = ?", request.FilterCompanyId)
//	}
//	if len(request.FilterDocumentStatus) != 0 {
//		sql = sql.Where("document_status IN( ? )", request.FilterDocumentStatus)
//	}
//	if request.FilterCreateDateFrom != zeroValueTime {
//		sql = sql.Where("create_date >= ?", request.FilterCreateDateFrom)
//	}
//	if request.FilterCreateDateTo != zeroValueTime {
//		sql = sql.Where("create_date <= ?", request.FilterCreateDateTo)
//	}
//	if request.FilterExpireDateFrom != zeroValueTime {
//		sql = sql.Where("expire_date >= ?", request.FilterExpireDateFrom)
//	}
//	if request.FilterExpireDateTo != zeroValueTime {
//		sql = sql.Where("expire_date <= ?", request.FilterExpireDateTo)
//	}
//	if request.FilterDeliveryDateFrom != zeroValueTime {
//		sql = sql.Where("delivery_date >= ?", request.FilterDeliveryDateFrom)
//	}
//	if request.FilterDeliveryDateTo != zeroValueTime {
//		sql = sql.Where("delivery_date <= ?", request.FilterDeliveryDateTo)
//	}
//
//	if err = sql.Find(&result).Error; err != nil {
//		return
//	}
//	return
//
//}

func getCompanyById(request int64) (result []CompanyInfo, err error) {

	if err = mssql.DB.Select("*").
		Table("malar.dbo.company").
		Where("id = ?", request).
		Find(&result).Error; err != nil {
		return
	}

	return
}

func buyerDocGR(request inputBuyerGR) (result docListGR, err error) {

		var zeroValueTime time.Time
		var zeroValueString string
		var SortBy string

		var getCompany []CompanyInfo
		getCompany, err = getCompanyById(request.LoginCompanyId)
		log.Infof("company ",getCompany[0].Id)
		if getCompany[0].CompanyType != 1003 {
			return
		}

		switch request.SortBy {
		case "document_number":
			SortBy = "dh.document_number"
		case "currency":
			SortBy = "dh.currency"
		case "document_status":
			SortBy = "dh.document_status"
		case "seller_company_id":
			SortBy = "dh.seller_company_id"
		case "create_date":
			SortBy = "dh.create_date"
		case "expire_date":
			SortBy = "dh.expire_date"
		case "delivery_date":
			SortBy = "dh.delivery_date"
		default :
			SortBy = "dh.document_number"
		}
		//if request.SortType != 0 && request.SortType != 1 || request.SortBy == zeroValueString{
		//	SortBy = "dh.document_number"
		//}
		if request.SortType == 1 {
			SortBy += " DESC"
		}

		grSubQuery := mssql.DB
		grMainQuery := mssql.DB

		grSubQuery = grSubQuery.Select("ROW_NUMBER() OVER(ORDER BY (SELECT NULL)) AS sequence,dh.document_number,c.company_name,dh.create_date ,dh.expire_date ,dh.delivery_date ,dh.grand_total,dh.currency,dh.document_id,co.description AS document_status,dh.buyer_company_id AS company_id,dh.seller_company_id").
			Table("malar.dbo.document_header AS dh,malar.dbo.company AS c,malar.dbo.config AS co").
			//Where("dh.document_type = ? AND co.status_code = dh.document_status AND dh.seller_company_id = c.id ",GoodsReceive)
			Where("dh.document_type = ? AND dh.document_status IN (?,?,?,?,?,?,?,?,?) AND co.status_code = dh.document_status AND dh.buyer_company_id = c.id", GoodsReceive, Draft, PendingApprove, Modify, PendingApproveAfterSubmit, ForModifyAfter,Submitted, Cancelled, Accept, CancelledAfterSubmit).
			Where("dh.buyer_company_id = ?",request.LoginCompanyId)

		if request.FreeSearchDocumentNumber != zeroValueString {
			grSubQuery = grSubQuery.Where("document_number LIKE ?", "%"+request.FreeSearchDocumentNumber+"%")
		}
		if len(request.FilterCurrency) != 0 {
			grSubQuery = grSubQuery.Where("currency IN( ? )", request.FilterCurrency)
		}
		if len(request.FilterCurrency) == 0 {
			grSubQuery = grSubQuery.Where("currency = 'THB'")
		}
		if len(request.FilterCompanyId) != 0 {
			grSubQuery = grSubQuery.Where("seller_company_id = ?", request.FilterCompanyId)
		}
		if len(request.FilterDocumentStatus) != 0 {
			grSubQuery = grSubQuery.Where("document_status IN( ? )", request.FilterDocumentStatus)
		}
		if request.FilterCreateDateFrom != zeroValueTime {
			grSubQuery = grSubQuery.Where("create_date >= ?", request.FilterCreateDateFrom)
		}
		if request.FilterCreateDateTo != zeroValueTime {
			grSubQuery = grSubQuery.Where("create_date <= ?", request.FilterCreateDateTo)
		}
		if request.FilterExpireDateFrom != zeroValueTime {
			grSubQuery = grSubQuery.Where("expire_date >= ?", request.FilterExpireDateFrom)
		}
		if request.FilterExpireDateTo != zeroValueTime {
			grSubQuery = grSubQuery.Where("expire_date <= ?", request.FilterExpireDateTo)
		}
		if request.FilterDeliveryDateFrom != zeroValueTime {
			grSubQuery = grSubQuery.Where("delivery_date >= ?", request.FilterDeliveryDateFrom)
		}
		if request.FilterDeliveryDateTo != zeroValueTime {
			grSubQuery = grSubQuery.Where("delivery_date <= ?", request.FilterDeliveryDateTo)
		}

		grMainQuery = grMainQuery.Raw("SELECT COUNT(grSubQuery.document_number ) AS document_count,SUM(grSubQuery.grand_total) AS grand_total From(?) AS grSubQuery", grSubQuery.SubQuery())
		//txMainQuery = txMainQuery.Raw("SELECT COUNT(txSubQuery.document_id) AS document_id_total ,SUM(txSubQuery.grand_total ) AS sum_grand_total From(?) AS txSubQuery", txSubQuery.SubQuery())
		var err2 error

		if err2 = grMainQuery.Find(&result.Header).Error; err2 != nil{
			return
		}

		if SortBy != zeroValueString {
			grSubQuery = grSubQuery.Order(SortBy)
		}

		if request.PagingIndex != 0 && request.PagingSize != 0 {
			grSubQuery = grSubQuery.Offset((request.PagingSize * request.PagingIndex) - request.PagingSize)
		}
		if request.PagingSize != 0 {
			grSubQuery = grSubQuery.Limit(request.PagingSize)
		}
	    if request.PagingSize == 0 {
		    grSubQuery = grSubQuery.Limit(10)
	    }

	if err = grSubQuery.Find(&result.Detail).Error; err != nil {
			return
		}
		return

	}



//func getPODocList(request documentListReq) (result []documentList, err error) {
//
//	if err = mssql.DB.Select("ROW_NUMBER() OVER(ORDER BY document_number) AS sequence,document_number,c.company_name AS supplier_name,create_date ,expire_date ,delivery_date ,grand_total AS total,currency ,document_status").
//		Table("document_header AS dh,company AS c").
//		Where("dh.document_type = ? AND currency IN ( ? ) AND company_name = ? AND document_status IN ( ? ) AND create_date BETWEEN ? AND ? AND expire_date BETWEEN ? AND ? AND delivery_date BETWEEN ? AND ?", request.DocumentType, request.Currency, request.SupplierName, request.DocumentStatus, request.CreateDateStart, request.CreateDateEnd, request.ExpireDateStart, request.ExpireDateEnd, request.DeliveryDateStart, request.DeliveryDateEnd).
//		Find(&result).Error; err != nil {
//		return
//	}
//	return
//}

//func getPODocListCount(request documentListReq) (result PoCount, err error) {
//
//	if err = mssql.DB.Select("COUNT(document_number) AS count,SUM(grand_total) AS sum_grand_total").
//		Table("document_header AS dh,company AS c ").
//		Where("dh.document_type = ? AND currency IN ( ? ) AND company_name = ? AND document_status IN ( ? ) AND create_date BETWEEN ? AND ? AND expire_date BETWEEN ? AND ? AND delivery_date BETWEEN ? AND ?", request.DocumentType, request.Currency, request.SupplierName, request.DocumentStatus, request.CreateDateStart, request.CreateDateEnd, request.ExpireDateStart, request.ExpireDateEnd, request.DeliveryDateStart, request.DeliveryDateEnd).
//		Find(&result).Error; err != nil {
//		return
//	}
//	return
//}
