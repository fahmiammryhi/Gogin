CREATE TABLE adm_users (
    id_user BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    id_role BIGINT NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(250),
    refresh_token_expired DATETIME,
    is_active BOOLEAN DEFAULT TRUE,
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_users_roles 
        FOREIGN KEY (id_role) REFERENCES adm_roles(id_role)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);