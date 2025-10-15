CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name TEXT NOT NULL, /*текст взял чтоб с запасом было :)*/
    price NUMERIC(10) NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL
);

INSERT INTO subscriptions (service_name, price, user_id, start_date) VALUES     /*немного тестовых данных*/
('Yandex Plus', 399, '60601fee-2bf1-4721-ae6f-7636e79a0cba', '2025-07-01'),
('Spotify', 599, '60601fee-2bf1-4721-ae6f-7636e79a0cba', '2025-08-17'),
('Netflix', 899, 'f1e8b7d3-1b44-4a12-8a0f-3b2b2bbcb888', '2025-06-15'),
('Apple Music', 169, 'c7709f76-d527-4217-a2de-c0f05cc15c7a', '2025-05-29');