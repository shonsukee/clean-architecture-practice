package main

import (
    "log"
    "net/http"
)

func main() {
    // Wireで依存関係を自動注入
    userHandler, err := Wire()
    if err != nil {
        log.Fatal(err)
    }

    // ルーティングの設定
    mux := http.NewServeMux()
    mux.HandleFunc("/users", userHandler.CreateUser)
    mux.HandleFunc("/users/", userHandler.GetUser)

    // サーバーの起動
    log.Println("Server starting on :8080...")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}