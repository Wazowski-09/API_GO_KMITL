package documents

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

//รับ INPUT แปลงค่า

//func (ep *Endpoint) PingGetEndpoint(c *gin.Context) { //GET app/ping
//	defer c.Request.Body.Close()
//	log.Infof("Check Heartbeat : PingGetEndpoint")
//
//	//เรียก logic
//	result, err := checkHeartbeat()
//	if err != nil {
//		//return err
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	//return success
//	c.JSON(http.StatusOK, result)
//	return
//}

//func (ep *Endpoint) PingGetParamsEndpoint(c *gin.Context) { //GET app/ping/:name
//	defer c.Request.Body.Close()
//	log.Infof("Check Heartbeat : PingGetPramsEndpoint")
//
//	//ดึงค่าจาก params ชื่อ name
//	name := c.Params.ByName("name") //key name in param
//	log.Infof("Params name : %s", name)
//
//	//เรียก logic
//	result, err := checkHeartbeat()
//	if err != nil {
//		//return err
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//
//	result.Message = fmt.Sprintf("%s [%s]", result.Message, name)
//
//	//return success
//	c.JSON(http.StatusOK, result)
//	return
//}

//func (ep *Endpoint) PingPostParamsAndBodyEndpoint(c *gin.Context) { //POST app/ping/:name
//	defer c.Request.Body.Close()
//	log.Infof("Check Heartbeat : PingPostBodyEndpoint")
//
//	//ดึงค่าจาก params ชื่อ name
//	name := c.Params.ByName("name")
//	log.Infof("Params name : %s", name)
//
//	//ดึงค่าจาก body
//	var request inputHeartbeat //model รับ input จาก body
//	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	log.Infof("Body Name : %s", request.Name)
//	log.Infof("Body Age  : %d", request.Age)
//
//	//เรียก logic
//	result, err := checkHeartbeat()
//	if err != nil {
//		//return err
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//
//	result.Message = fmt.Sprintf("%s ชื่อเล่น [%s] ชื่อเต็ม [%s] บริษัท [%d]", result.Message, name, request.Name, request.Age)
//
//	//return success
//	c.JSON(http.StatusOK, result)
//	return
//}

//func (ep *Endpoint) GetDocumentDetailEndpoint(c *gin.Context) { //POST documents/getDocumentDetail
//	defer c.Request.Body.Close()
//	log.Infof("Get DocumentDetail Endpoint")
//
//	var request documentDetailReq //model รับ input จาก body
//	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	log.Infof("request : %+v", request)
//
//	//เรียก logic
//	result, err := getDocumentDetail(request)
//	if err != nil {
//		//return err
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	//return success
//	c.JSON(http.StatusOK, result)
//	return
//
//}

//func (ep *Endpoint) GettPurchesOrderDocListEndpoint(c *gin.Context) {
//	defer c.Request.Body.Close()
//	log.Infof("Get DocumentList Endpoint")
//
//	var request documentListReq
//	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	log.Infof("request : %+v", request)
//
//	result, err := getPODocumentList(request)
//	if err != nil {
//		//return err
//		c.JSON(http.StatusBadRequest, err)
//		return
//	}
//	//return success
//	c.JSON(http.StatusOK, result)
//	return
//}

func (ep *Endpoint) GetGoodsRECEIPTDocListEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Get DocumentList Endpoint")

	var request inputBuyerGR
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		msg := messageError{
			Status:             http.StatusBadRequest,
			MessageCode:        "400",
			MessageDescription: "Bad Request"}
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	log.Infof("request : %+v", request)

	result, err  := getGRDocumentList(request)
	if err != nil {
		msg := messageError{
			Status:             http.StatusBadRequest,
			MessageCode:        "400",
			MessageDescription: "notFound"}
		//return err
		c.JSON(msg.Status, msg)
		return
	}

		msg := messageError{
			Status:             http.StatusOK,
			MessageCode:        "200",
			MessageDescription: "Found"}
		c.JSON(msg.Status, result)
		return
}
