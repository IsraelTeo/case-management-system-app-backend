CREATE TABLE sede (
    idsede UUID,
    nombre VARCHAR(200) NOT NULL,
    idinstitucion UUID,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT institucion_fk_idinstitucion FOREIGN KEY (idinstitucion) REFERENCES institucion(idinstitucion) ON DELETE CASCADE,
    CONSTRAINT sede_pk_idsede PRIMARY KEY (idsede)
);