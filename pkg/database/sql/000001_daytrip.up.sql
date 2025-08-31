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
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    package_name VARCHAR(255) NOT NULL,
    message TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE banners (
    id SERIAL PRIMARY KEY,
    desktop_image VARCHAR(255),
    mobile_image VARCHAR(255),
    cta VARCHAR(255) NULL,
    cta_url VARCHAR(255) NULL,
    title VARCHAR(255) NULL,
    description VARCHAR(255) NULL
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

CREATE TABLE invitations (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) NOT NULL, 
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    template_id INT,
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    maps_url VARCHAR(255),
    address VARCHAR(255),
    location VARCHAR(255),
    dress_code VARCHAR(255),
    birthday_val INT,
    image VARCHAR(255),
    image1 VARCHAR(255),
    keyPass VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE invitation_response (
    id SERIAL PRIMARY KEY,
    invitation_id INT REFERENCES invitations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    is_attending VARCHAR(10) NOT NULL,
    message TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


-- DATA

INSERT INTO invitation_templates (name, description, invitation_assets_count)
values ('yachts party', '', 0);

INSERT INTO invitations 
(slug, user_id, title, description, template_id, start_date, end_date, maps_url, address, location, dress_code) 
VALUES
('belya-maliha-party', 1, 'belya maliha yatch party', 'Get ready to make waves and celebrate! Wishing you a fantastic time on the water, filled with endless fun, stunning views, and incredible memories.', 1, 
 '2025-08-23 18:00:00+00', '2025-08-23 23:00:00+00', 
 'https://maps.app.goo.gl/u1yquUYur8QWsY3r8', 
 'Jl Raya Pantai Mutiara No 57, Jakarta Utara', 'PANTAI MUTIARA', 'Anything but floral!');