# Protobuf Files
## Description
This project contains all the protobuf files for their related service. In each folder, exist the `.proto` file where is defined the services contracts.

## How to work
When you modify any proto file inside a folder, it's necessary to compile it in the desired language.

In this case, we are working with go files.

````bash
protoc -Idevice_service --go_out=. --go_opt=module=protobuf-files --go-grpc_out=. --go-grpc_opt=module=protobuf-files device_service/device.proto
````
