# 使用官方Go 1.21映像作為建構環境
FROM golang:1.21-alpine as builder

# 設置工作目錄
WORKDIR /app

# 複製go.mod和go.sum文件
COPY go.mod ./
COPY go.sum ./

# 下載依賴
RUN go mod download

# 複製整個專案源代碼到容器中
COPY . .

# 編譯Go程序，指定執行檔路徑
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jubobe ./cmd/api/main.go

# 使用scratch映像作為運行環境
FROM scratch

# 從builder階段複製編譯好的應用程序
COPY --from=builder /app/jubobe .

# 指定啟動命令
CMD ["./jubobe"]