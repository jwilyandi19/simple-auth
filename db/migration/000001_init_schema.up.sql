CREATE TABLE users (
    user_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_name varchar(255) NOT NULL,
    user_full_name varchar(255),
    user_password varchar(1024) NOT NULL,
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);