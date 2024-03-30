CREATE TABLE IF NOT EXISTS `orders` (
    id INT NOT NULL AUTO_INCREMENT,
    table_number INT NOT NULL,
    total_price INT NOT NULL,
    total_quantity INT NOT NULL,
    order_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    restaurant_id INT NOT NULL,
    order_status ENUM('pending', 'completed') NOT NULL DEFAULT 'pending',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants`(`id`) ON DELETE CASCADE
)
