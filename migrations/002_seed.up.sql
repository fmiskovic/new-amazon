-- Seed accounts table
INSERT INTO accounts (id, created_at, updated_at, email, full_name, date_of_birth, location, gender) VALUES
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'john.doe@example.com', 'John Doe', '1980-01-01', 'New York', 1),
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'jane.smith@example.com', 'Jane Smith', '1990-02-02', 'Los Angeles', 2);

-- Seed items table
INSERT INTO items (id, created_at, updated_at, title, description, price) VALUES
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Item 1', 'Description for Item 1', 10.00),
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Item 2', 'Description for Item 2', 20.00),
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Item 3', 'Description for Item 3', 30.00);

-- Seed orders table
INSERT INTO orders (id, created_at, updated_at, account_id) VALUES
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, (SELECT id FROM accounts WHERE email = 'john.doe@example.com')),
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, (SELECT id FROM accounts WHERE email = 'jane.smith@example.com'));

-- Seed order_items table
INSERT INTO order_items (id, created_at, updated_at, order_id, item_id, quantity) VALUES
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, (SELECT id FROM orders WHERE account_id = (SELECT id FROM accounts WHERE email = 'john.doe@example.com')), (SELECT id FROM items WHERE title = 'Item 1'), 2),
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, (SELECT id FROM orders WHERE account_id = (SELECT id FROM accounts WHERE email = 'john.doe@example.com')), (SELECT id FROM items WHERE title = 'Item 2'), 1),
(uuid_generate_v4(), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, (SELECT id FROM orders WHERE account_id = (SELECT id FROM accounts WHERE email = 'jane.smith@example.com')), (SELECT id FROM items WHERE title = 'Item 3'), 3);
