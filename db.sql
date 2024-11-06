CREATE TABLE categories (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE products (
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2),
    category_id VARCHAR(50),
    image_url VARCHAR(50),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);