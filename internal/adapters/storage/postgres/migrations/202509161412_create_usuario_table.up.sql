CREATE TABLE usuario (
    idusuario UUID NOT NULL,
    usuario VARCHAR(200) NOT NULL,
    password VARCHAR(45) NOT NULL,
    email VARCHAR(100) NOT NULL,
    telefono VARCHAR(45),
    idrol UUID NOT NULL,
    idempresa UUID NOT NULL,
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT usuario_uq_usuario UNIQUE (usuario),
    CONSTRAINT usuario_uq_email UNIQUE (email),
    CONSTRAINT usuario_uq_telefono UNIQUE (telefono),
    CONSTRAINT rol_fk_idrol FOREIGN KEY (idrol) REFERENCES rol(idrol),
    CONSTRAINT empresa_fk_idempresa FOREIGN KEY (idempresa) REFERENCES empresa(idempresa) ON DELETE SET NULL,
    CONSTRAINT usuario_pk_idusuario PRIMARY KEY (idusuario)
);