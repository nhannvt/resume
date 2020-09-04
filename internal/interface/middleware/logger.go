package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/nhannvt/resume/internal/interface/auth"
	"github.com/gin-gonic/gin"
)

type loggingItems []map[string]interface{}

func Logger() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		items := loggingItems{
			{"k": "time", "v": param.TimeStamp.Format("2006-01-02T15:04:05+09:00")},
			{"k": "host", "v": param.Request.RemoteAddr},
			{"k": "forwardedfor", "v": param.Request.Header.Get("X-Forwarded-For")},
			{"k": "req", "v": param.Path},
			{"k": "status", "v": param.StatusCode},
			{"k": "method", "v": param.Method},
			{"k": "uri", "v": param.Request.RequestURI},
			{"k": "size", "v": param.BodySize},
			{"k": "referer", "v": param.Request.Referer()},
			{"k": "ua", "v": param.Request.UserAgent()},
			{"k": "reqtime_microsec", "v": int64(param.Latency / time.Microsecond)},
			{"k": "cache", "v": param.Request.Header.Get("X-Cache")},
			{"k": "runtime", "v": param.Request.Header.Get("X-Runtime")},
			{"k": "vhost", "v": param.Request.Host},
		}

		clientKey := ""
		if client, ok := param.Keys["APIClient"]; ok {
			clientKey = client.(*auth.Client).Name()
		}
		items = append(items, map[string]interface{}{"k": "client", "v": clientKey})

		if info, ok := param.Keys["APIInfo"]; ok {
			items = append(items, map[string]interface{}{"k": "info", "v": info})
		}

		logtexts := make([]string, len(items))

		for i := 0; i < len(items); i++ {
			logtexts[i] = fmt.Sprintf("%s:%v", items[i]["k"], items[i]["v"])
		}

		return fmt.Sprintf("%s\n", strings.Join(logtexts, "\t"))
	})
}
