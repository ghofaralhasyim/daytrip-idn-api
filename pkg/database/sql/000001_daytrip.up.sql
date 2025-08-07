CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(50) NULL,
    image VARCHAR(255) NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    package_id INT NOT NULL,
    message TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE banners (
    id SERIAL PRIMARY KEY,
    desktop_image VARCHAR(255),
    mobile_image VARCHAR(255),
    cta VARCHAR(255),
    cta_url VARCHAR(255),
    title VARCHAR(255),
    description VARCHAR(255)
);

CREATE TABLE web_settings (
    id SERIAL PRIMARY KEY,
    key VARCHAR(255) NOT NULL,
    value TEXT,
    description VARCHAR(255),
    type VARCHAR(50)
);

CREATE TABLE activities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    image VARCHAR(255)
);

CREATE TABLE yachts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    capacity VARCHAR(255),
    speed VARCHAR(255),
    deck VARCHAR(255),
    bedroom VARCHAR(255),
    bathroom VARCHAR(255),
    other_facilities JSON,
    image VARCHAR(255)
);

CREATE TABLE destinations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);