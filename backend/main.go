package main

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    _ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func main() {
    var err error

    // データベース接続の設定
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err = sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("データベース接続に失敗しました: %v", err)
    }
    defer db.Close()

    // エンドポイントの設定
    r := mux.NewRouter()
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users", getUsers).Methods("GET")

    // サーバーの起動
    log.Println("サーバーをポート8080で開始します")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func createUser(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    email := r.FormValue("email")

    _, err := db.ExecContext(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "ユーザーが作成されました")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.QueryContext(context.Background(), "SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    users := []map[string]interface{}{}
    for rows.Next() {
        var id int
        var name, email string
        if err := rows.Scan(&id, &name, &email); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        user := map[string]interface{}{
            "id":    id,
            "name":  name,
            "email": email,
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
