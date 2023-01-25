# go-zap-logger

基於zap可動態變更相關設置，並應用檔案大小切割日誌文件

1. 安裝相關package
```go
go get -u go.uber.org/zap

go get -u github.com/natefinch/lumberjack

go get -u github.com/gin-gonic/gin
```

2. 新建logger並動態設置
```go

    log1 := logger.New()
	log1.Error("log1:before submit")
	//新增project name
	log1.Config.SetProjectName("log2")
	//送出設定
	log1.SubmitConfig()
	log1.Error("log2:set project name")
	//將text改變json訊息
	log1.Config.SetJSONFormat(true)
	log1.SubmitConfig()
	log1.Error("log3: change to json")
```

輸出結果:
```bash
2023-01-25 20:36:11	ERROR	go-zap-logger/mian.go:13	log1:before submit
2023-01-25 20:36:11	ERROR	log2	go-zap-logger/mian.go:16	log2:set project name
{"Level":"ERROR","Time":"2023-01-25 20:36:11","Project":"log2","Caller":"go-zap-logger/mian.go:19","Message":"log3: change to json"}
```

3. logger接收gin默認日誌並自訂想收集的訊息
main.go
```go
//新增gin引擎
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
	}), log1.GinLogger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
```
logger.go
```go
func (l *Logger) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		l.Config.SetProjectName("gin")
		l.SubmitConfig()

		cost := time.Since(start)
		l.logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
```
輸出結果:
```bash
{"Level":"INFO","Time":"2023-01-25 21:41:42","Project":"gin","Caller":"gin@v1.8.2/context.go:173","Message":"/ping{status 11 200  <nil>} {method 15 0 GET <nil>} {path 15 0 /ping <nil>} {query 15 0  <nil>} {ip 15 0 127.0.0.1 <nil>} {user-agent 15 0 Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 <nil>} {errors 15 0  <nil>} {cost 8 0  <nil>}"}
```