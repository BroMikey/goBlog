package image

import (
	"net/http"

	imagesvc "github.com/BroMikey/goBlog/internal/service/image"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *imagesvc.Service
}

func NewHandler(svc *imagesvc.Service) *Handler {
	return &Handler{svc: svc}
}

// Ping 示例接口：用于验证 image 模块路由/分层是否正常工作
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    h.svc.Ping(),
	})
}
