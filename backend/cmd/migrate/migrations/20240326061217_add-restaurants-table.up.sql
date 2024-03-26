CREATE TABLE IF NOT EXISTS restaurants (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `open_hours` TIMESTAMP NOT NULL,
    `close_hours` TIMESTAMP NOT NULL,
    `cuisine_type` VARCHAR(255) NOT NULL,
    `access_token` VARCHAR(255) NOT NULL,
    `user_id` INT NOT NULL,
    `capacity` INT,

    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
)