CREATE TABLE IF NOT EXISTS `payments` (
    id INT NOT NULL AUTO_INCREMENT,
    order_id INT NOT NULL,
    amount INT NOT NULL,
    payment_type ENUM('cash', 'visa') NOT NULL DEFAULT 'visa',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
)
