package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var courses = map[int64]string{
	1: "Introduction to programming",
	2: "Introduction to algorithms",
	3: "Data structures",
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go to /courses/description"))
}

func CourseDescHandler(w http.ResponseWriter, r *http.Request) {
	courseNumber, err := strconv.ParseInt(r.URL.Query().Get("course_id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	if val, found := courses[courseNumber]; found {
		w.Write([]byte(val))
	}
}

func StartServer() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", IndexHandler)
	//mux.HandleFunc("/courses/description", CourseDescHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	cwd, _ := os.Getwd()
	logFile := filepath.Join(cwd, ".log")
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	//defer file.Close()
	logger.SetOutput(file)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go to /sum"))
	})

	mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		x, errX := strconv.Atoi(r.URL.Query().Get("x"))
		if errX != nil {
			logrus.WithError(errX).Error("error")
			return
		}

		y, errY := strconv.Atoi(r.URL.Query().Get("y"))
		if errY != nil {
			logrus.WithError(errY).Error("error")
			return
		}

		var result int
		temp := math.MaxInt - x
		if y >= temp {
			logger.WithFields(logrus.Fields{"x": x, "y": y}).Warning("Sum overflows int")
			result = -1
		} else {
			result = x + y
		}
		w.Write([]byte(strconv.Itoa(result)))
	})

	port := "8080"
	logWithPort := logrus.WithFields(logrus.Fields{
		"port": port,
	})
	logWithPort.Info("Starting a web-server on port")
	logWithPort.Fatal(server.ListenAndServe())

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
