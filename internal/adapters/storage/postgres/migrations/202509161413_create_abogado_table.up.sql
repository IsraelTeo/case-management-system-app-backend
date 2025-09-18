CREATE TABLE abogado (
    idabogado UUID NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    apellido VARCHAR(45) NOT NULL,
    dni VARCHAR(45) NOT NULL,
    numero_colegiatura VARCHAR(45) NOT NULL,
    telefono_celular VARCHAR(45),
    telefono_oficina VARCHAR(45),
    telefono_domicilio VARCHAR(45),
    direccion_domicilio VARCHAR(150),
    email_personal VARCHAR(45) NOT NULL,
    email_oficina VARCHAR(45) NOT NULL,
    fecha_nacimiento DATE,
    idempresa UUID,                                                                        
    fecha_registro TIMESTAMP NOT NULL,
    fecha_actualizacion TIMESTAMP NOT NULL,
    usuario_registro UUID NOT NULL,
    usuario_actualizacion UUID NOT NULL,
    CONSTRAINT abogado_uq_dni UNIQUE (dni),
    CONSTRAINT abogado_uq_numero_colegiatura UNIQUE (numero_colegiatura),
    CONSTRAINT abogado_uq_telefono_celular UNIQUE (telefono_celular),
    CONSTRAINT abogado_uq_telefono_oficina UNIQUE (telefono_oficina),
    CONSTRAINT abogado_uq_telefono_domicilio UNIQUE (telefono_domicilio),
    CONSTRAINT abogado_uq_email_personal UNIQUE (email_personal),
    CONSTRAINT abogado_uq_email_oficina UNIQUE (email_oficina),
    CONSTRAINT usuario_fk_usuario_registro FOREIGN KEY (usuario_registro) REFERENCES usuario(idusuario),
    CONSTRAINT usuario_fk_usuario_actualizacion FOREIGN KEY (usuario_actualizacion) REFERENCES usuario(idusuario),
    CONSTRAINT empresa_fk_idempresa FOREIGN KEY (idempresa) REFERENCES empresa(idempresa) ON DELETE SET NULL,
    CONSTRAINT abogado_pk_idabogado PRIMARY KEY (idabogado)
);



