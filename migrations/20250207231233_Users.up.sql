CREATE TABLE Users (
                       id SERIAL PRIMARY KEY,
                       Email VARCHAR(255) NOT NULL,
                       Password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMP DEFAULT NULL
);