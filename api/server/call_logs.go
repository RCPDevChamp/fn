package server

import (
	"bytes"
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleCallLogGet(c *gin.Context) {
	ctx := c.Request.Context()

	appName := c.MustGet(api.AppName).(string)
	callID := c.Param(api.Call)

	logReader, err := s.LogDB.GetLog(ctx, appName, callID)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	var b bytes.Buffer
	b.ReadFrom(logReader)

	c.String(http.StatusOK, b.String())
}
