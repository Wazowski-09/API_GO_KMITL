package dashboard

import (
	"time"
)

type heartbeatModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputHeartbeat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type inputBuyer struct {
	DateFrom       time.Time `json:"date_from"`
	DateTo         time.Time `json:"date_to"`
	BuyerCompanyID int       `json:"buyer_company_id"`
}

type outputDashboard struct {
	DocumentType   string  `json:"document_type"`
	DocumentStatus int     `json:"document_status"`
	CountDoc       int     `json:"count_doc"`
	ActiveCount    int     `json:"active_count"`
	Total          float64 `json:"total"`
}

type inputDashboard struct {
	DateTo   time.Time `json:"date_to"`
	DateFrom time.Time `json:"date_from"`
}
