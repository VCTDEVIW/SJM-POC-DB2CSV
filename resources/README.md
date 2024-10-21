(sqlsrv)

USE test;

CREATE TABLE your_table_name (
rid INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
field1 NVARCHAR(512),
field2 NVARCHAR(512),
inserted_timestamp DATETIME2 DEFAULT GETDATE() NOT NULL
);
