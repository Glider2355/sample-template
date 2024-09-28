# ベースイメージ
FROM golang:1.19

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルをコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係を取得
RUN go mod download

# アプリケーションソースをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o /myapp

# アプリケーションの実行
CMD ["/myapp"]
