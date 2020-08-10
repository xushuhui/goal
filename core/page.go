package core

import (
	"github.com/gin-gonic/gin"
	"goal/setting"
	"goal/utils"
)

type Pager struct {
	List interface{} `json:"list"`
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

func GetPage(c *gin.Context) int {
	page, _ := utils.StringToInt(c.Query("page"))
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize, _ := utils.StringToInt(c.Query("page_size"))
	if pageSize <= 0 {
		return setting.AppSetting.DefaultPageSize
	}
	if pageSize > setting.AppSetting.MaxPageSize {
		return setting.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
