Para compilar los protobuffers se debe usar el siguiente comando:

´´´
$> protoc -I=protobuf/ --go_out=plugins=grpc:protobuf/  protobuf/products.proto
´´´