package xhttp

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type HandleFunc func() HandleInterface

type HandleInterface interface {
	Handle() Response
}

type Response struct {
	Status int       `json:"status"`
	Img    string    `json:"img"`
	Data   *DataUtil `json:"data"`
}

type DataUtil struct {
	Ent map[string]interface{}
	Ext map[string]interface{}
}

func AuthReqWrap(handleReq HandleFunc) func(c *gin.Context) {
	fun := func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return
		}
		handle := handleReq()
		if c.Request.Method == "POST" {
			_ = json.Unmarshal(body, handle)
		}
		resp := handle.Handle()
		c.JSON(200, resp)
		// todo
	}
	return fun
}

func ReqExec(c *gin.Context) {

}
