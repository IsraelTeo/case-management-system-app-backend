CREATE TABLE documento (
    iddocumento UUID NOT NULL,
    numero_documento VARCHAR(100) NOT NULL,
    numero_digitalizacion VARCHAR(100) NOT NULL,
    numero_notificacion VARCHAR(100),
    tipo INT NOT NULL,
    estado INT NOT NULL,
    idexpediente UUID NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT expediente_fk_idexpediente FOREIGN KEY (idexpediente) REFERENCES expediente(idexpediente) ON DELETE CASCADE,
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    CONSTRAINT usuario_fk_usuario_actualizacion FOREIGN KEY (usuario_actualizacion) REFERENCES usuario(idusuario),
    CONSTRAINT documento_pk_iddocumento PRIMARY KEY (iddocumento)
);