CREATE TABLE IF NOT EXISTS users (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `access_key` VARCHAR(255) NOT NULL UNIQUE,

    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
)
