CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

CREATE TABLE korzinka (
    id SERIAL PRIMARY KEY,
    user_id foreign KEY references users(id),
    product_id foreign KEY references products(id)
)



INSERT INTO users (username, email, password)
VALUES 
    ('Alisher', 'alisher@example.com', 'password123'),
    ('Dilshod', 'dilshod@example.com', 'password123'),
    ('Gulnora', 'gulnora@example.com', 'password123'),
    ('Kamola', 'kamola@example.com', 'password123'),
    ('Jasur', 'jasur@example.com', 'password123');  


INSERT INTO products (name, description, price, stock_quantity)
VALUES 
    ('Olma', 'Shirin va suvli olma', 2.50, 100),
    ('Non', 'Yangi pishirilgan non', 1.00, 50),
    ('Gilos', 'Yangi terilgan gilos', 3.75, 200),
    ('Sut', '1 litr yangi sut', 0.80, 150),
    ('Go''sht', 'Yangi mol go''shti', 5.00, 75);