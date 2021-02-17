package documents

import (
	"time"
)

//type heartbeatModel struct {
//	Message  string    `json:"message"`
//	DateTime time.Time `json:"date_time"`
//}

//type inputHeartbeat struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}

//type documentDetailReq struct {
//	DocumentNumber  string `json:"document_number"`
//	SellerCompanyId int64  `json:"seller_company_id"`
//	BuyerCompanyId  int64  `json:"buyer_company_id"`
//	DocumentType    int64  `json:"document_type"`
//}
type documentDetail struct {
	DetailId     int64   `json:"detail_id"`
	Sku          string  `json:"sku"`
	ProductName  string  `json:"product_name"`
	Quantity     int64   `json:"quantity"`
	Unit         string  `json:"unit"`
	PricePerUnit float64 `json:"price_per_unit"`
	TotalPrice   float64 `json:"total_price"`
}
//type documentDetailList struct {
//	documentDetail []documentDetail
//}
type documentHeader struct {
	DocumentNumber    string    `json:"document_number"`
	DocumentStatus    string    `json:"document_status"`
	CreateDate        time.Time `json:"create_date"`
	ExpireDate        time.Time `json:"expire_date"`
	DeliveryDate      time.Time `json:"delivery_date"`
	Net               float64   `json:"net"`
	VatPercent        int       `json:"vat_percent"`
	VatValue          float64   `json:"vat_value"`
	Discount          float64   `json:"discount"`
	GrandTotal        float64   `json:"grand_total"`
	Currency          string    `json:"currency"`
	DeliveryAddress   string    `json:"delivery_address"`
	CompanyNameSeller string    `json:"company_name_seller"`
	ShortNameSeller   string    `json:"short_name_seller"`
	BranchSeller      string    `json:"branch_seller"`
	PhoneNumberSeller string    `json:"phone_number_seller"`
	DistrictSeller    string    `json:"district_seller"`
	SubDistrictSeller string    `json:"sub_district_seller"`
	ZipCodeSeller     string    `json:"zip_code_seller"`
	ProvinceSeller    string    `json:"province_seller"`
	CompanyNameBuyer  string    `json:"company_name_buyer"`
	ShortNameBuyer    string    `json:"short_name_buyer"`
	BranchBuyer       string    `json:"branch_buyer"`
	PhoneNumberBuyer  string    `json:"phone_number_buyer"`
	DistrictBuyer     string    `json:"district_buyer"`
	SubDistrictBuyer  string    `json:"sub_district_buyer"`
	ZipCodeBuyer      string    `json:"zip_code_buyer"`
	ProvinceBuyer     string    `json:"province_buyer"`
}
type documentDetailRes struct {
	Detail []documentDetail
	Header documentHeader
}

const(
	//Document Status
	None = 0
	Draft = 1
	PendingApprove = 3
	Modify = 5
	PendingApproveAfterSubmit = 7
	ForModifyAfter = 9
	Submitted = 11
	Cancelled = 13
	Accept = 15
	CancelledAfterSubmit = 17

	//Document Type
	PurchaseOrder = 110
	GoodsReceive = 120
	Invoice = 210

	//Company Type
	Seller = 1001
	Buyer = 1003
)

type inputBuyerGR struct {
	LoginCompanyId           int64     `json:"login_company_id"`
	FreeSearchDocumentNumber string    `json:"free_search_document_number"`
	FilterCurrency           []string  `json:"filter_currency"`
	FilterCompanyId          []int64   `json:"filter_company_id"`
	FilterDocumentStatus     []int64   `json:"filter_document_status"`
	FilterCreateDateFrom     time.Time `json:"filter_create_date_from"`
	FilterCreateDateTo       time.Time `json:"filter_create_date_to"`
	FilterExpireDateFrom     time.Time `json:"filter_expire_date_from"`
	FilterExpireDateTo       time.Time `json:"filter_expire_date_to"`
	FilterDeliveryDateFrom   time.Time `json:"filter_delivery_date_from"`
	FilterDeliveryDateTo     time.Time `json:"filter_delivery_date_to"`
	SortBy                   string    `json:"sort_by"`
	SortType                 int64     `json:"sort_type"`
	PagingIndex              int64     `json:"paging_index"`
	PagingSize               int64     `json:"paging_size"`
}

type outputBuyerGR struct {
	Sequence       int64     `json:"sequence"`
	DocumentId     int64     `json:"document_id"`
	DocumentNumber string    `json:"document_number"`
	CompanyId      int64     `json:"company_id"`
	CompanyName    string    `json:"company_name"`
	CreateDate     time.Time `json:"create_date"`
	ExpireDate     time.Time `json:"expire_date"`
	DeliveryDate   time.Time `json:"delivery_date"`
	GrandTotal     float64   `json:"grand_total"`
	Currency       string    `json:"currency"`
	DocumentStatus string    `json:"document_status"`
}

type GRCount struct {
	DocumentCount int64   `json:"document_count"`
	GrandTotal    float64 `json:"grand_total"`
}

type docListGR struct {
	Header		GRCount	`json:"header"`
	Detail		[]outputBuyerGR	`json:"detail"`

}
type CompanyInfo struct {
	Id          int    `json:"id"`
	CompanyName string `json:"company_name"`
	ShortName   string `json:"short_name"`
	CompanyType int    `json:"company_type"`
	Branch      string `json:"branch"`
	PhoneNumber string `json:"phone_number"`
	District    string `json:"district"`
	SubDistrict string `json:"sub_district"`
	ZipCode     string `json:"zip_code"`
	Province    string `json:"province"`
}
type messageError struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
//type documentListReq struct {
//	DocumentNumber    string   `json:"document_number"`
//	DocumentStatus    []int64  `json:"document_status"`
//	SupplierName      string   `json:"supplier_name"`
//	Currency          []string `json:"currency"`
//	CreateDateStart   string   `json:"create_date_start"`
//	CreateDateEnd     string   `json:"create_date_end"`
//	ExpireDateStart   string   `json:"expire_date_start"`
//	ExpireDateEnd     string   `json:"expire_date_end"`
//	DeliveryDateStart string   `json:"delivery_date_start"`
//	DeliveryDateEnd   string   `json:"delivery_date_end"`
//	DocumentType      int64    `json:"document_type"`
//	OrderBy           string   `json:"order_by"`
//}

//type documentList struct {
//	Sequence       int64     `json:"sequence"`
//	DocumentNumber string    `json:"document_number"`
//	SupplierName   string    `json:"supplier_name"`
//	CreateDate     time.Time `json:"create_date"`
//	ExpireDate     time.Time `json:"expire_date"`
//	DeliveryDate   time.Time `json:"delivery_date"`
//	Total          float64   `json:"total"`
//	Currency       string    `json:"currency"`
//	DocumentStatus int       `json:"document_status"`
//}
//type PoCount struct {
//	Count         int64   `json:"count"`
//	DocumentType  string  `json:"document_type"`
//	SumGrandTotal float64 `json:"sum_grand_total"`
//}

//type documentListRes struct {
//	PoCount      PoCount        `json:"po_count"`
//	DocumentList []documentList `json:"document_list"`
//}
