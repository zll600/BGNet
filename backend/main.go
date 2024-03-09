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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var dataPaths = [2]string{
	"/home/andy/Code/github/BGNet/web/data/Img/",
	"/home/andy/Code/github/BGNet/web/data/GT/",
}

const savePath = "/home/andy/Code/github/BGNet/web/result/"

func main() {
	r := gin.Default()
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
