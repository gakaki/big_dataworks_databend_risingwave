SELECT
    orders.order_number,
    users.username,
    products.product_name,
    orders.total_amount
FROM
    orders
        JOIN users ON orders.user_id = users.id
        JOIN products ON orders.product_id = products.id
    LIMIT 1;