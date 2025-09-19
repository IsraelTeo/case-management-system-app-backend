CREATE TABLE rol (
    idrol UUID,
    rol VARCHAR(200) NOT NULL,
    descripcion TEXT,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT rol_uq_rol UNIQUE (rol),
    CONSTRAINT rol_pk_idrol PRIMARY KEY (idrol)
);