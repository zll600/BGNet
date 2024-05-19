package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var dataPaths = [2]string{
	"/home/andy/Code/github/BGNet/web/data/Img/",
	"/home/andy/Code/github/BGNet/web/data/GT/",
}

const savePath = "/home/andy/Code/github/BGNet/web/result/"

var fileName string

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func promhttpServer() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func main() {
	go promhttpServer()

	r := gin.Default()
	// r.RedirectTrailingSlash = false
	r.MaxMultipartMemory = 8 << 20 // 20MiB

	r.Use(gin.Logger())
	r.Use(RequestLoggerMiddleware())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/uploadv2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for file_idx, file := range files {
			slog.Info("lgz debug", "PWD: ", os.Getenv("PWD"), "file: ", file.Filename)
			if strings.HasSuffix(file.Filename, ".jpg") {
				fileName = file.Filename
			}
			c.SaveUploadedFile(file, dataPaths[file_idx]+file.Filename)
		}
	})

	// r.GET("/generate", func(c *gin.Context) {
	// 	// ch := make(chan struct{})
	// 	// go func() {
	// 	// 	// JSON body
	// 	// 	body := []byte(`{
	// 	// 		"title": "Post title",
	// 	// 		"body": "Post description",
	// 	// 		"userId": 1
	// 	// 		}`)

	// 	// 	_, err := http.Post("http://127.0.0.1:5000/model", "application/json", bytes.NewBuffer(body))
	// 	// 	if err != nil {
	// 	// 		slog.Error("error when call model", "err", err)
	// 	// 	}
	// 	// 	ch <- struct{}{}
	// 	// }()
	// 	// <-ch
	// 	// getResultImage()
	// 	c.Header("Content-Type", "image/png")
	// 	// c.Header("Access-Control-Allow-Origin", "*")
	// 	c.File(savePath + fileName)
	// })

	r.GET("/generate/", func(c *gin.Context) {
		// getResultImage()
		c.Header("Content-Type", "image/png")
		c.File(savePath + fileName)
	})

	r.POST("/uploadv3", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		if len(files) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "upload not two files",
			})
			return
		}

		for file_idx, file := range files {
			slog.Info("lgz debug", "PWD: ", os.Getenv("PWD"))

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dataPaths[file_idx]+file.Filename)
		}

		getResultImage()
		c.Header("Content-Type", "image/png")
		c.File(savePath + files[1].Filename)
	})

	r.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		if len(files) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "upload not two files",
			})
			return
		}

		for file_idx, file := range files {
			slog.Info("lgz debug", "PWD: ", os.Getenv("PWD"))

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dataPaths[file_idx]+file.Filename)
		}

		getResultImage()
		c.Header("Content-Type", "image/png")
		c.File(savePath + files[1].Filename)
	})

	r.GET("/test", func(c *gin.Context) {
		c.Header("Content-Type", "image/png")
		c.File(savePath + "camourflage_00012.png")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getResultImage() error {
	args := []string{
		"/home/andy/Code/github/BGNet/etest_one.py",
	}
	cmd := exec.Command("python3", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Error when executing python: %s\n", stderr.Bytes())
		return fmt.Errorf("error executing python: %w", err)
	}

	slog.Info("lgz debug", "stdout: ", stdout.Bytes(), "stderr: ", stderr.Bytes())
	return nil
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)
		slog.Debug("log request body", "body", string(body))
		slog.Debug("log request header", "header", c.Request.Header)
		c.Next()
	}
}
