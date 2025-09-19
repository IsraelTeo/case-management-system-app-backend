CREATE TABLE empresa (
    idempresa UUID,
    nombre VARCHAR(200) NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT empresa_pk_idempresa PRIMARY KEY (idempresa)
);