# golang 基礎映像檔版本設定
FROM golang:1.18
# 工作目錄(請同專案根目錄)
# 如該資料夾於容器中不存在將自動創建
WORKDIR /mygame
# 複製整個資料夾到工作目錄
COPY . .
# 編譯 main.go 檔成 app
RUN go build -o app
# 安裝 curl
RUN apk add curl
# 檔案儲存位置
# VOLUME [ "/mygame" ]
# 對外開放 80(API) 50051(gRPC專用)
EXPOSE 80 50051
# 程式進入點
ENTRYPOINT ["./app"]
# 在專案根目錄，包成 mygame 映像檔
# docker build -t mygame . 
# 容器啟動需設置環境變數:
# @ SERVICE 服務名稱
# @ ENV 環境
# 執行容器帶入環境變數
# 本地的 4400 對應到容器的 80
# 本地的 4401 對應到容器的 50051
# docker run -d --name mygame -e SERVICE=mygame -e ENV=docker -p 4400:80 -p 4401:50051 mygame
