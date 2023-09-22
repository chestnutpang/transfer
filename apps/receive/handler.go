package receive

import (
	"cts/plugin/coalminePlugin"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

func Receive(c *gin.Context) {
	params, err := io.ReadAll(c.Request.Body)
	if err != nil {

	}
	var p interface{}
	err = json.Unmarshal(params, &p)
	if err != nil {

	}
	coalminePlugin.CoalmineFunc(p)
	c.JSON(200, gin.H{})
}
