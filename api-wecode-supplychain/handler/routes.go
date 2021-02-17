package handler

import (
	"net/http"

	"api-wecode-supplychain/service/dashboard"
	"api-wecode-supplychain/service/documents"
	"api-wecode-supplychain/service/onboard"
	"api-wecode-supplychain/service/ping"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {

	ping := ping.NewEndpoint()
	dashboard := dashboard.NewEndpoint()
	documents := documents.NewEndpoint()
	onboard := onboard.NewEndpoint()

	r.transaction = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			Endpoint:    ping.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET Prams",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodGet,
			Pattern:     "/ping/:name",
			Endpoint:    ping.PingGetParamsEndpoint,
		},
		{
			Name:        "Ping Pong : POST Prams+Body",
			Description: "Ping Pong : Heartbeat",
			Method:      http.MethodPost,
			Pattern:     "/ping/:name",
			Endpoint:    ping.PingPostParamsAndBodyEndpoint,
		},
	}

	txDashboard := []route{
		{
			Name:        "Buyer : POST Dashboard",
			Description: "Buyer : Data",
			Method:      http.MethodPost,
			Pattern:     "/getbuyer",
			Endpoint:    dashboard.GetBuyerDashboard,
		},
		{
			Name:        "GetdataDashboard : POST Prams+Body",
			Description: "GetdataDashboard : GG",
			Method:      http.MethodPost,
			Pattern:     "/getdashboard",
			Endpoint:    dashboard.GetDashboard,
		},
	}

	txDocument := []route{
		//{
		//	Name:     "Ping Pong : GET",
		//	Method:   http.MethodGet,
		//	Pattern:  "/ping",
		//	Endpoint: documents.PingGetEndpoint,
		//},
		//{
		//	Name:        "Ping Pong : GET Prams",
		//	Description: "Ping Pong : Data",
		//	Method:      http.MethodGet,
		//	Pattern:     "/ping/:name",
		//	Endpoint:    documents.PingGetParamsEndpoint,
		//},
		{
			Name:        "GET DocumentHeader GR",
			Description: "GET DocumentHeader GR",
			Method:      http.MethodPost,
			Pattern:     "/gr/searchDocumentsList",
			Endpoint:    documents.GetGoodsRECEIPTDocListEndpoint,
		},
		//{
		//	Name:        "GET DocumentDetail Endpoint",
		//	Description: "GET DocumentDetail Endpoint",
		//	Method:      http.MethodPost,
		//	Pattern:     "/getDocumentDetail",
		//	Endpoint:    documents.GetDocumentDetailEndpoint,
		//},
		//{
		//	Name:        "GET DocumentListPO Endpoint",
		//	Description: "GET DocumentListPO Endpoint",
		//	Method:      http.MethodPost,
		//	Pattern:     "/getDocumentList",
		//	Endpoint:    documents.GettPurchesOrderDocListEndpoint,
		//},
	}

	txOnboard := []route{
		{
			Name:        "Get Company Endpoint",
			Description: "Get Company Info By Id",
			Method:      http.MethodGet,
			Pattern:     "/getCompanyInfo/:id",
			Endpoint:    onboard.GetCompanyEndpoint,
		},
	}

	ro := gin.New()

	//ro.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"POST", "GET"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/documents")
	for _, e := range txDocument {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/dashboard")
	for _, e := range txDashboard {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/onboard")
	for _, e := range txOnboard {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
