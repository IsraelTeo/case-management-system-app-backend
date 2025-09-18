CREATE TABLE expediente (
    idexpediente UUID NOT NULL,
    estado INT NOT NULL,
    numero_expediente VARCHAR(100) NOT NULL,
    materia VARCHAR(250) NOT NULL,
    idsede UUID NOT NULL,
    idempresa UUID NOT NULL,
    idcliente UUID NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT sede_fk_idsede FOREIGN KEY (idsede) REFERENCES sede(idsede) ON DELETE CASCADE,
    CONSTRAINT empresa_fk_idempresa FOREIGN KEY (idempresa) REFERENCES empresa(idempresa) ON DELETE CASCADE,
    CONSTRAINT cliente_fk_idcliente FOREIGN KEY (idcliente) REFERENCES cliente(idcliente) ON DELETE CASCADE,
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    CONSTRAINT usuario_fk_usuario_actualizacion FOREIGN KEY (usuario_actualizacion) REFERENCES usuario(idusuario),
    CONSTRAINT expediente_pk_idexpediente PRIMARY KEY (idexpediente)
);