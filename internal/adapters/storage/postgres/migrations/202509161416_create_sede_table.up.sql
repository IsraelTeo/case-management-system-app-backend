CREATE TABLE sede (
    idsede UUID NOT NULL,
    nombre VARCHAR(200) NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    idinstitucion UUID NOT NULL,
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    CONSTRAINT usuario_fk_usuario_actualizacion FOREIGN KEY (usuario_actualizacion) REFERENCES usuario(idusuario),
    CONSTRAINT institucion_fk_idinstitucion FOREIGN KEY (idinstitucion) REFERENCES institucion(idinstitucion) ON DELETE CASCADE,
    CONSTRAINT sede_pk_idsede PRIMARY KEY (idsede)
);