USE db;

-- Main table
CREATE TABLE IF NOT EXISTS items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    location CHAR(16),
    item_type INT,
    quantity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE items
ADD INDEX idx_location (location);

ALTER TABLE items
ADD INDEX idx_created_at_id (created_at, id);