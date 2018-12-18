# sqlflowserver
The gRPC proxy server of SQL engines

## Build and Test

According to https://bit.ly/2BCxGwH, we are supposed to commit the protoc-generated Go source files.  So, to build and test, we can run the following command:

```bash
protoc --go_out=plugins=grpc:. *.proto && go test -v
```
