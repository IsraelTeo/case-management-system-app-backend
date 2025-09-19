CREATE TABLE expediente (
    idexpediente UUID,
    estado INT NOT NULL,
    numero_expediente VARCHAR(100) NOT NULL,
    materia VARCHAR(250) NOT NULL,
    idsede UUID,
    idempresa UUID,
    idcliente UUID,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT sede_fk_idsede FOREIGN KEY (idsede) REFERENCES sede(idsede) ON DELETE CASCADE,
    CONSTRAINT empresa_fk_idempresa FOREIGN KEY (idempresa) REFERENCES empresa(idempresa) ON DELETE CASCADE,
    CONSTRAINT cliente_fk_idcliente FOREIGN KEY (idcliente) REFERENCES cliente(idcliente) ON DELETE CASCADE,
    CONSTRAINT expediente_pk_idexpediente PRIMARY KEY (idexpediente)
);