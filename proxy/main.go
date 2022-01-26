package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	gen "grpc-gateway-demo/gen/go/hello"
)

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m:=httpsnoop.CaptureMetrics(handler,writer,request)
		log.Printf("http[%d]-- %s -- %s\n",m.Code,m.Duration,request.URL.Path)
	})
}

func convertToMetadata(m map[string][]string) map[string]string{
	newM:=make(map[string]string,len(m))
	for key,value:=range m{
		newM[key]=value[0]
	}
	return newM
}

var allowedHeaders=map[string]struct{}{
	"X-REQUEST-ID": {},
}

func isHeaderAllowed(s string)( string,bool) {
	if _,isAllowed:=allowedHeaders[s];isAllowed {
		return strings.ToLower(s),true
	}
	return s,false
}

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux:=runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			md:=metadata.New(convertToMetadata(request.Header))
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError:=runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx,mux,marshaler,writer,request,&newError)
		}))
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := gen.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, "localhost:8080", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	// Creating a normal HTTP server
	server:=http.Server{
		Handler: withLogger(mux),
	}

	// creating a listener for server
	l,err:=net.Listen("tcp",":8081")
	if err!=nil {
		log.Fatal(err)
	}
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
