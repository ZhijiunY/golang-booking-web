package main

import (
	"fmt"
	"testing"

	"github.com/ZhijiunY/booking-web/cmd/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}


// func TestRoutes(t *testing.T) {
// 	var app config.AppConfig

// 	// 初始化 app 變數
// 	app = config.AppConfig{}

// 	// 呼叫 routes 函數
// 	mux := routes(&app)

// 	// 確認回傳值的型態是否為 *chi.Mux
// 	switch v := mux.(type) {
// 	case *chi.Mux:
// 		// do nothing
// 	default:
// 		t.Errorf("type is not *chi.Mux, but is %T", v)
// 	}
// }