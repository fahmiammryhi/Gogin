CREATE TABLE adm_menu_roles (
    id_menu_role BIGINT AUTO_INCREMENT PRIMARY KEY,
    id_role BIGINT NOT NULL,
    id_menu VARCHAR(15) NOT NULL,
    is_view BOOLEAN DEFAULT TRUE,
    is_insert BOOLEAN DEFAULT TRUE,
    is_edit BOOLEAN DEFAULT TRUE,
    is_delete BOOLEAN DEFAULT TRUE,
    is_print BOOLEAN DEFAULT TRUE,
    is_approve BOOLEAN DEFAULT TRUE,
    is_active BOOLEAN DEFAULT TRUE,
    CONSTRAINT fk_menu_roles_role
        FOREIGN KEY (id_role) REFERENCES adm_roles(id_role)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_menu_roles_menu
        FOREIGN KEY (id_menu) REFERENCES adm_menu(id_menu)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
