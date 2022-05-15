package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/pkg/setting"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/models"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/pkg/util"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/pkg/e"
)


func GetAddress(c *gin.Context) {
	address := com.StrTo(c.Param("address")).String()

    maps := make(map[string]interface{})
    data := make(map[string]interface{})
	maps["from_sources"] = 1
    if address != "" {
        maps["address"] = address
		
    }

    code := e.SUCCESS

    data["lists"] = models.GetAddress(util.GetPage(c), setting.PageSize, maps)
    data["total"] = models.GetAddressTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

func AddAddress(c *gin.Context) {
}


func DeleteAddress(c *gin.Context) {
}
