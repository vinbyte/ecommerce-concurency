package helpers

import (
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

// InitLogger used for initiate the logger
func (h *Helpers) InitLogger() {
	logAppendMode := os.Getenv("LOG_APPEND_MODE")
	if logAppendMode == "" {
		logAppendMode = "console"
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	if logAppendMode == "file" {
		//path log
		pathName := os.Getenv("LOG_PATH")
		filename := path.Join(pathName, "ecommerce-app-%Y%m%d.log")
		// maxAge default 3 month
		maxAge := h.logMaxAge
		writer, errWriter := rotatelogs.New(
			filename,
			rotatelogs.WithMaxAge(maxAge),
		)
		if errWriter != nil {
			log.Fatalf("Failed to Initialize Log File %s", errWriter)
		}
		log.SetOutput(writer)
	} else {
		log.SetOutput(os.Stdout)
	}
}

// CustomRequestLogger is custom log form for every incoming http request
func (h *Helpers) CustomRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		log.WithFields(log.Fields{
			"ip":      c.ClientIP(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"proto":   c.Request.Proto,
			"status":  c.Writer.Status(),
			"latency": time.Since(startTime),
			"ua":      c.Request.UserAgent(),
		}).Info()
	}
}
