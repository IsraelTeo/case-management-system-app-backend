CREATE TABLE empresa (
    idempresa UUID NOT NULL,
    nombre VARCHAR(200) NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    CONSTRAINT usuario_fk_usuario_actualizacion FOREIGN KEY (usuario_actualizacion) REFERENCES usuario(idusuario),
    CONSTRAINT empresa_pk_idempresa PRIMARY KEY (idempresa)
);