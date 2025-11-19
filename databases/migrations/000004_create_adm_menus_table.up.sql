CREATE TABLE adm_menu (
    id_menu VARCHAR(15) NOT NULL PRIMARY KEY,
    id_parent VARCHAR(15),
    menu_name VARCHAR(50) NOT NULL,
    controller VARCHAR(50),
    class_icon VARCHAR(30),
    menu_sort BIGINT,
    is_active BOOLEAN DEFAULT TRUE
);