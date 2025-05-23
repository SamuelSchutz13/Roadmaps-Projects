CREATE TABLE expense (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE
);