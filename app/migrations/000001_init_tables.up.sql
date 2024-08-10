-- Таблица users
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     login VARCHAR(50) UNIQUE NOT NULL,
                                     password BYTEA NOT NULL,
                                     salt VARCHAR(50) NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_login ON users(login);

-- Тип данных для типов информации
CREATE TYPE types_information AS ENUM ('login', 'binary', 'text');

-- Таблица entities
CREATE TABLE IF NOT EXISTS entities (
                                        id SERIAL PRIMARY KEY,
                                        user_id BIGINT NOT NULL,
                                        type types_information NOT NULL,
                                        title VARCHAR(255) NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_entities_user_id ON entities(user_id);

-- Таблица login_data
CREATE TABLE IF NOT EXISTS login_data (
                                          id SERIAL PRIMARY KEY,
                                          user_id BIGINT NOT NULL,
                                          username VARCHAR(255) NOT NULL,
                                          password BYTEA NOT NULL,
                                          url VARCHAR(2048),
                                          notes TEXT,
                                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                          FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_login_data_user_id ON login_data(user_id);

-- Таблица binary_data
CREATE TABLE IF NOT EXISTS binary_data (
                                           id SERIAL PRIMARY KEY,
                                           user_id BIGINT NOT NULL,
                                           path VARCHAR(255),
                                           notes TEXT,
                                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_binary_data_user_id ON binary_data(user_id);

-- Таблица text_data
CREATE TABLE IF NOT EXISTS text_data (
                                         id SERIAL PRIMARY KEY,
                                         user_id BIGINT NOT NULL,
                                         text TEXT,
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_text_data_user_id ON text_data(user_id);

-- Таблица credit_card_data
CREATE TABLE IF NOT EXISTS credit_card_data (
                                                id SERIAL PRIMARY KEY,
                                                user_id BIGINT NOT NULL,
                                                card_number BYTEA NOT NULL,
                                                cardholder_name VARCHAR(255) NOT NULL,
                                                expiration_date DATE NOT NULL,
                                                cvv BYTEA NOT NULL,
                                                notes TEXT,
                                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_credit_card_data_user_id ON credit_card_data(user_id);