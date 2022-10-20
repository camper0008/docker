DROP DATABASE IF EXISTS example_db;

CREATE DATABASE example_db;

USE example_db;

CREATE TABLE clicks(
	id INT PRIMARY KEY AUTO_INCREMENT,
    clicked TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE USER 'example_user'@'%' IDENTIFIED BY 'example_password';

GRANT ALL PRIVILEGES ON example_db.* TO 'example_user'@'%';

FLUSH PRIVILEGES;
