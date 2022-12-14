Las dependencias están especificadas en el archivo _requeriments.txt_

Para compilar los protobuffers se debe usar el siguiente comando:
```bash
$> python3 -m grpc_tools.protoc -I protobufs --python_out=. --grpc_python_out=. protobufs/products.proto 
```