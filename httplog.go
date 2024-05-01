package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	logBuffer *bytes.Buffer
	status    int
}

func (rw *ResponseWriter) Write(msg []byte) (int, error) {
	if _, err := rw.logBuffer.Write(msg); err != nil {
		log.Fatal("unable to write logBuffer:", err)
	}
	return rw.ResponseWriter.Write(msg)
}

func (rw *ResponseWriter) WriteHeader(newStatus int) {
	rw.status = newStatus
	rw.ResponseWriter.WriteHeader(newStatus)
}

func logEncode(b []byte) string {
	return fmt.Sprintf("%q", string(b))
}

func stringifyRequest(r *http.Request) string {
	buf := new(bytes.Buffer)
	r.Write(buf)
	return logEncode(buf.Bytes())
}

func stringifyResponse(res ResponseWriter) string {
	return logEncode(res.logBuffer.Bytes())
}

func Mux(handler http.Handler) http.Handler {
	requestLog := log.New(os.Stdout, "[Request] ", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lmsgprefix)
	responseLog := log.New(os.Stdout, "[Response] ", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lmsgprefix)
	return LogMux(handler, requestLog, responseLog)
}

func LogMux(handler http.Handler, requestLog, responseLog *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := stringifyRequest(r)
		requestLog.Println(request)

		rw := ResponseWriter{w, bytes.NewBuffer([]byte{}), http.StatusOK}
		start := time.Now()
		handler.ServeHTTP(&rw, r)
		elapsedTime := time.Since(start)
		responseLog.Printf("Execution time: %v, Status: %d %s, Response: %s, Request: %s",
			elapsedTime, rw.status, http.StatusText(rw.status), stringifyResponse(rw), request)
	})
}
