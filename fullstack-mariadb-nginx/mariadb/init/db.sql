USE example_db;

CREATE TABLE clicks(
	id INT PRIMARY KEY AUTO_INCREMENT,
    clicked TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);
