
package util

import (
	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"

	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/pkg/setting"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
