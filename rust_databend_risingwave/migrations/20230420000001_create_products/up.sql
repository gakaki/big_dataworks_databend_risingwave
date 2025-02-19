CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  sku VARCHAR(100) UNIQUE,
  price DECIMAL(10,2) NOT NULL,
  stock_quantity INT DEFAULT 0,
  category VARCHAR(100),
  brand VARCHAR(100),
  weight DECIMAL(10,2),
  dimensions VARCHAR(100),
  color VARCHAR(50),
  material VARCHAR(100),
  manufacturer VARCHAR(255),
  supplier_id INT,
  min_stock_level INT,
  max_stock_level INT,
  reorder_quantity INT,
  is_active BOOLEAN DEFAULT TRUE,
  tax_rate DECIMAL(5,2) DEFAULT 0,
  discount_percentage DECIMAL(5,2) DEFAULT 0,
  rating DECIMAL(3,2),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_products_price ON products(price);