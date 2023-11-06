package main

// #include "main.h"
import "C"
import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/KoLLlaka/bongocat/internal/config/json"
	"github.com/KoLLlaka/bongocat/internal/handler"
)

var positionChan chan handler.Position = make(chan handler.Position)
var keyBoardChan chan string = make(chan string)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	// config
	cfg := config.New()
	config.ParseConfig("./static/config/config.json", &cfg)
	keyboardSetMap := config.KeyboardSetToMap(cfg.Keyboard)

	// handler Bongocat
	mux := http.NewServeMux()
	server := handler.New(mux, cfg, positionChan, keyBoardChan)
	server.Start("/bongocat")

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Printf("server start on port :%s\n", cfg.Server.Port)
	fmt.Printf("please follow to http://%s:%s/bongocat on your browser\n", cfg.Server.Host, cfg.Server.Port)
	go http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, mux)

	go func() {
		for {
			posX := C.fx()
			posY := C.fy()

			cursor := handler.Position{
				X: float64(posX),
				Y: float64(posY),
			}

			positionChan <- cursor
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			keyNum := int(C.f())

			if keyValue, ok := keyboardSetMap[keyNum]; ok {
				keyBoardChan <- keyValue
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	<-sc
	fmt.Println("Stopping...")
}
