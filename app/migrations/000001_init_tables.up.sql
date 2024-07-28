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

create type types_information as enum ('login', 'binary', 'text');
-- Таблица entities
CREATE TABLE IF NOT EXISTS entities (
                                        id SERIAL PRIMARY KEY,
                                        user_id BIGINT NOT NULL,
                                        type types_information NOT NULL,
                                        title VARCHAR(255) NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_entities_user_id ON entities(user_id);

-- Таблица login_data
CREATE TABLE IF NOT EXISTS login_data (
                                          id SERIAL PRIMARY KEY,
                                          entry_id BIGINT NOT NULL,
                                          username VARCHAR(255) NOT NULL,
                                          password BYTEA NOT NULL,
                                          url VARCHAR(2048),
                                          notes TEXT,
                                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                          FOREIGN KEY (entry_id) REFERENCES entities(id)
);

CREATE INDEX IF NOT EXISTS idx_login_data_entry_id ON login_data(entry_id);

-- Таблица binary_data
CREATE TABLE IF NOT EXISTS binary_data (
                                           id SERIAL PRIMARY KEY,
                                           entry_id BIGINT NOT NULL,
                                           path VARCHAR(255),
                                           notes TEXT,
                                           FOREIGN KEY (entry_id) REFERENCES entities(id)
);

CREATE INDEX IF NOT EXISTS idx_binary_data_entry_id ON binary_data(entry_id);

-- Таблица text_data
CREATE TABLE IF NOT EXISTS text_data (
                                         id SERIAL PRIMARY KEY,
                                         entry_id BIGINT NOT NULL,
                                         path VARCHAR(255),
                                         notes TEXT,
                                         FOREIGN KEY (entry_id) REFERENCES entities(id)
);

CREATE INDEX IF NOT EXISTS idx_text_data_entry_id ON text_data(entry_id);