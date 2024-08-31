CREATE TABLE post (
        post_id SERIAL PRIMARY KEY,
        post_name VARCHAR(255) not null,
        likes INT DEFAULT 0
);