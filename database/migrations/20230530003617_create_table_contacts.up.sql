CREATE TABLE IF NOT EXISTS contacts (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL UNIQUE,
    phone_number varchar(255) NOT NULL UNIQUE,
    slug varchar(255) NOT NULL UNIQUE,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB;