module grpc-gateway-demo

go 1.16

require (
	github.com/felixge/httpsnoop v1.0.2
	github.com/gin-gonic/gin v1.7.7
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2
	github.com/soheilhy/cmux v0.1.5
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210816183151-1e6c022a8912 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4 // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0 // indirect
)

replace google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4 => github.com/donetkit/go-genproto v0.0.0-20220415075338-07e796c0f121