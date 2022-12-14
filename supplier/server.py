import grpc
from concurrent import futures
import mysql.connector
import products_pb2
import products_pb2_grpc


class Products(products_pb2_grpc.ProductsServicer):
    def __init__(self,conn):
        self.conn = conn
    
    def SendProduct(self, request, context):
        cursor = self.conn.cursor()
        cursor.execute("SELECT * FROM inventario WHERE id_producto = %s",(request.id_producto,))
        result = cursor.fetchone()
        print("Se solicito {} unidad(es) del producto con id {}".format(request.cantidad_disponible, request.id_producto))

        if(not result):
            # Agregar request.id_product a la base de datos, + 5 unidades
            sql = "INSERT INTO inventario (id_producto, nombre, cantidad_disponible) VALUES (%s, %s, 5)"
            values = (request.id_producto,request.nombre)
            cursor.execute(sql, values)
            self.conn.commit()
        elif(int(request.cantidad_disponible) > result[2]):
            # request.id_producto + 5 unidades
            sql = "UPDATE inventario SET cantidad_disponible = 5 WHERE id_producto = %s"
            values = (request.id_producto,)
            cursor.execute(sql, values)
            self.conn.commit()
        else:
            # cantidad request.id_producto en base de datos - request.cantidad_disponible
            sql = "UPDATE inventario SET cantidad_disponible = %s WHERE id_producto = %s"
            values = (result[2]-int(request.cantidad_disponible),request.id_producto)
            cursor.execute(sql, values)
            self.conn.commit()
        cursor.Close()
        response = products_pb2.SendProductResponse(id_producto=request.id_producto,cantidad_disponible=request.cantidad_disponible)
        return response

def main():
    conn = mysql.connector.connect(
        host="localhost",
        user="admin",
        password="12345678",
        database="db_inventario"
    )

    server = grpc.server(futures.ThreadPoolExecutor())
    products_pb2_grpc.add_ProductsServicer_to_server(Products(conn), server)

    server.add_insecure_port('[::]:9000')
    server.start()

    print("gRPC server working")
    server.wait_for_termination()
    conn.Close()

if __name__ == "__main__":
    main()