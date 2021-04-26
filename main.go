package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/hello":   hello,
	"/headers": headers,
}

func main() {
	for pattern, handler := range handlers {
		http.HandleFunc(pattern, handler)
	}

	// Приложим некоторые усилия, чтобы приложение завершалось с нулевым кодом выхода
	// Это важно для тестов, и в целом приятно
	server := &http.Server{Addr: ":8090", Handler: nil}
	// Запускаем приложение в отдельной горутине
	go func() {
		server.ListenAndServe()
	}()

	// А в текущей ждём сигнала об остановке приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
