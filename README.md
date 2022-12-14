Necesita 3 servidores para ejecutarse:
- Servidor 1 (Supermercado y Cliente)
- Servidor 2 (Proveedor)
- Servidor 3 (Despacho)

Para la generación de Scheduled Tasks en el servidor 2 se ocupo el siguiente comando usando la base de datos _db\_invetario_:

```
$mysql> CREATE EVENT IF NOT EXISTS restock
	ON SCHEDULE EVERY 1 MINUTE
	DO UPDATE inventario SET cantidad_disponible=(FLOOR(RAND()*10)+1) WHERE cantidad_disponible=0;
```

Para la generación de Scheduled Tasks en el servidor 3 se ocupo el siguiente comando usando la base de datos _db\_despachos_:

```
$mysql> CREATE EVENT IF NOT EXISTS state_update
	ON SCHEDULE EVERY 1 MINUTE
	DO UPDATE despacho SET estado=(
		SELECT
		CASE WHEN estado='EN_TRANSITO' THEN 'ENTREGADO'
		WHEN estado='RECIBIDO' THEN 'EN_TRANSITO'
		ELSE estado
		LIMIT 1);
```

❗️**Servidor 1:** Para el servidor 1 se deben correr simultáneamente main.go y server/server.go.

❗️**Servidor 2:** Para el servidor 2 se debe correr server.py.

❗️**Servidor 3:** Para el servidor 3 se deben correr simultáneamente server.go y consumer.go.

# Equipo:
- Rafael Aros Soto
- Franco Cabezas Poblete
- Paulina Vega Rivera