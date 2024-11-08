CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    product_id UUID,
    quantity INT,
    customer_id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY,
    name TEXT,
    email TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP

);

CREATE TABLE IF NOT EXISTS products(
    id TIMEUUID ,
    name TEXT,
    description TEXT,
    price INT,
    stock INT,
    category TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (category, id)


);