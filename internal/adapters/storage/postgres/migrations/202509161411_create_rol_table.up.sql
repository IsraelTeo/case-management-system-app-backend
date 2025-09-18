CREATE TABLE rol (
    idrol UUID NOT NULL,
    rol VARCHAR(200) NOT NULL,
    descripcion TEXT,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT rol_uq_rol UNIQUE (rol),
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    CONSTRAINT usuario_fk_usuario_actualizacion FOREIGN KEY (usuario_actualizacion) REFERENCES usuario(idusuario),

    CONSTRAINT rol_pk_idrol PRIMARY KEY (idrol)
);