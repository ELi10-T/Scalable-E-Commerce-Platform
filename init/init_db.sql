-- This is init for the postgres

-- Schema for Users

CREATE TABLE user_table (
    id SERIAL PRIMARY KEY, -- unique id for each user, auto assigned by the db
    username VARCHAR(50) UNIQUE NOT NULL, -- Unique
    email    VARCHAR(255) UNIQUE NOT NULL, -- email id
    password_hash TEXT NOT NULL, -- need to hash before storing
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- TODO: Check on how to update this intrinsically
    -- is_active BOOLEAN DEFAULT TRUE, -- Track whether the user is active
    -- role VARCHAR DEFAULT 'user' -- role based access control
);

-- carts table
CREATE TABLE carts_table (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_table(id) ON DELETE CASCADE
);

-- Create the cart_items table
CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INT NOT NULL,
    product_id VARCHAR NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE,
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product_catalog(id) ON DELETE CASCADE
);

-- Create the product_catalog table
CREATE TABLE product_catalog (
    id SERIAL PRIMARY KEY,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

-- Questions : TODO
-- 1. Cart_Item quantity should not exceed the Product_Catalog quantity -- What steps should be taken?
