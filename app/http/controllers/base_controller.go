package controllers

import (
	"fmt"
	"goblog/pkg/logger"
	"net/http"

	"gorm.io/gorm"
)

// BaseController 基础控制器
type BaseController struct {
}

// ResposeForSQLError 处理 SQL 错误并返回
func (bc BaseController) ResposeForSQLError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		// 3.1 数据未找到
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 文章未找到")
	} else {
		// 3.2 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}
}
