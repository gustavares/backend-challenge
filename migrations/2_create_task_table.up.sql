CREATE TABLE task (
    id binary(36) PRIMARY KEY,
    user_id binary(36) REFERENCES user(id),
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    summary VARCHAR(2500) NOT NULL,
    performed_in TIMESTAMP
);