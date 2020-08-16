# trojan-web
前端控制面板

### 二进制文件生成
#### Mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build ./
 
#### Linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build ./