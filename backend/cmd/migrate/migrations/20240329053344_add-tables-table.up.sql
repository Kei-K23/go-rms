CREATE TABLE IF NOT EXISTS `restaurant_tables` (
    id INT NOT NULL AUTO_INCREMENT,
    table_number INT NOT NULL,
    capacity INT NOT NULL,
    status ENUM('occupied', 'vacant') DEFAULT 'vacant',
    restaurant_id INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants`(`id`) ON DELETE CASCADE
)
