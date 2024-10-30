-- CREATE TABLE categories (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
-- );

-- CREATE TYPE user_role AS ENUM ('admin', 'user');

-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL,
--     username VARCHAR(255) NOT NULL UNIQUE,
--     password VARCHAR(255),
--     role user_role NOT NULL
-- );

-- CREATE TABLE posts (
--     id SERIAL PRIMARY KEY,
--     category_id INT REFERENCES categories (id) NOT NULL,
--     user_id INT REFERENCES users (id) NOT NULL,
--     category_name VARCHAR(255),
--     user_name VARCHAR(255),
--     title VARCHAR(255) NOT NULL,
--     description VARCHAR(255) NOT NULL,
--     image_path VARCHAR(255)
-- );


