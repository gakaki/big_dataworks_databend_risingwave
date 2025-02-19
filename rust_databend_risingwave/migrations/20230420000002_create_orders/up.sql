CREATE TABLE orders (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  order_date DATETIME DEFAULT CURRENT_TIMESTAMP,
  status VARCHAR(50),
  shipping_address TEXT,
  payment_method VARCHAR(50),
  total_amount DECIMAL(10,2) NOT NULL,
  discount DECIMAL(10,2) DEFAULT 0,
  tax DECIMAL(10,2) DEFAULT 0,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  tracking_number VARCHAR(100),
  delivery_date DATETIME,
  customer_feedback TEXT,
  is_gift BOOLEAN DEFAULT FALSE,
  gift_message TEXT,
  order_source VARCHAR(50),
  currency VARCHAR(10) DEFAULT 'CNY',
  order_type VARCHAR(50)
);
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_product_id ON orders(product_id);
CREATE INDEX idx_orders_order_date ON orders(order_date);