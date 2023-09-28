package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study_gin/service"
)

// GetName @Tags USER
// @Summary 查询地址
// @Produce json
// @Accept json
// @Param name query string true "名称"
// @Success 200 {object} string "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} string "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /getName [get]
func GetName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Query("name")
		if len(param) == 0 {
			ctx.String(http.StatusBadRequest, "lose param name")
			return
		}
		s := service.GetUserSrv()
		address := s.GetNameService(ctx.Request.Context(), param)
		ctx.String(http.StatusOK, address)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.PostForm("name")
		if len(param) == 0 {
			ctx.String(http.StatusBadRequest, "lose param name")
			return
		}
		s := service.GetUserSrv()
		resp := s.UpdateUserService(ctx.Request.Context(), param)
		ctx.JSON(http.StatusOK, resp)
	}
}
