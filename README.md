# sqlflowserver
The gRPC proxy server of SQL engines

## Build

We have included the precompiled protobuf.

In case you need to rebuild after modify the `sqlflow.proto`, run

```bash
cd server && go generate
```
