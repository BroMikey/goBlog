package settings

import (
	"net/http"

	settingssvc "github.com/BroMikey/goBlog/internal/service/settings"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *settingssvc.Service
}

func NewHandler(svc *settingssvc.Service) *Handler {
	return &Handler{svc: svc}
}

// Status 示例接口：用于验证模块路由/分层是否正常工作
func (h *Handler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    h.svc.Status(),
	})
}
