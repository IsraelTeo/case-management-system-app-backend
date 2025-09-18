CREATE TABLE expedientexabogado (
    idexpediente UUID NOT NULL,
    idabogado UUID NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    CONSTRAINT expediente_fk_idexpediente FOREIGN KEY (idexpediente) REFERENCES expediente(idexpediente) ON DELETE CASCADE,
    CONSTRAINT abogado_fk_idabogado FOREIGN KEY (idabogado) REFERENCES abogado(idabogado) ON DELETE CASCADE,
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    PRIMARY KEY (idexpediente, idabogado)
);