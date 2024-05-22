CREATE TABLE phone (
    id UUID PRIMARY KEY,
    phone_name VARCHAR(50),
    color VARCHAR(50),
    price VARCHAR(100),
    ram int,
    memory int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
