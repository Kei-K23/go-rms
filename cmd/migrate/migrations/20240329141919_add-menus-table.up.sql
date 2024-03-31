CREATE TABLE IF NOT EXISTS `menus` (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    category_id INT NOT NULL,
    price INT NOT NULL,
    restaurant_id INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`) ON DELETE CASCADE,
    UNIQUE KEY `unique_menu_number_per_restaurant` (`name`, `restaurant_id`)
)
