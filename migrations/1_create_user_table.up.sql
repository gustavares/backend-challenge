CREATE TABLE user (
    id binary(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    user_role ENUM('admin', 'technician') NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
);
