# x-msa-core

Генирация gRPC

генирация моделей
``` bash
protoc -I=./ --go_out=./ ./proto/**.proto
```
генирация интерфейсов
``` bash
protoc -I=./ --go-grpc_out=./ ./proto/**.proto
```