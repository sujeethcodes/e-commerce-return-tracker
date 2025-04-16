CREATE TABLE return_products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id VARCHAR(100),
    order_id VARCHAR(100),
    account_id VARCHAR(100),
    reason VARCHAR(100),
    return_date DATE,
    flag VARCHAR(100),
    status VARCHAR(100)
);
