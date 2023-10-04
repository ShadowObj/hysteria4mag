go env -w GOARCH=amd64
go env -w GOOS=linux
go env -w GOAMD64=v3
go env -w CGO_ENABLED=0
go build -o hysteria_linux_amd64_avx_v1.3.71 -tags=gpl -ldflags "-s -w -X 'main.appDate=2023-10-04 12:01:01' -X 'main.appVersion=1.3.71'  -X 'main.appCommit=b28f3c0'" -trimpath ./app/cmd/