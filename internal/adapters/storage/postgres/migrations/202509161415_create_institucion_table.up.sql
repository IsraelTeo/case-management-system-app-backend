CREATE TABLE institucion (
    idinstitucion UUID,
    nombre VARCHAR(200) NOT NULL,
    descripcion VARCHAR(255) NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT institucion_pk_idinstitucion PRIMARY KEY (idinstitucion)
);