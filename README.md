
### googleapis

```
file://C:/Users/Microfoft/go/src/googleapis
OR
file://C:/Users/Microfoft/go/src/donetkit/googleapis
```

<img src="2022-04-15_16-11-11.png">



C:\Users\Microfoft\go\pkg\mod\github.com\golang\protobuf@v1.5.2\protoc-gen-go


### protoc

```
protoc \
  -I . \
  -I ${GOPATH}/src \
  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  --go_out=":../generated" \
  --validate_out="lang=go:../generated" \
  example.proto
```
  
```
protoc -I ./ -I C:/Users/Microfoft/AppData/Local/JetBrains/GoLand2022.1/protoeditor -I C:/Users/Microfoft/go/src/donetkit/googleapis --go_out=. proto/hello/hello_world.proto
```

### 安装protoc-gen-validate protoc-gen-validate.exe 并修改名称为 protoc-gen-go.exe
```
 go install github.com/envoyproxy/protoc-gen-validate@latest
```


```
protoc -I ./ -I C:/Users/Microfoft/AppData/Local/JetBrains/GoLand2022.1/protoeditor -I C:/Users/Microfoft/go/src/donetkit/googleapis --go_out=. --validate_out="lang=go:." proto/hello/
hello_world.proto
```

```replace
replace google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 => github.com/donetkit/googleapis v0.0.0-20220415075058-c10cd7222356
``` 