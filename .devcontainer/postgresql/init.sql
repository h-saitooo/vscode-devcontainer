\c devcontainer

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL
);

INSERT INTO users (name, email, password_hash) VALUES
  ('devcontainer', 'devcontainer@example.com', 'devcontainer@135'),
  ('devcontainer2', 'devcontainer2@example.com', 'devcontainer@136'),
  ('devcontainer3', 'devcontainer3@example.com', 'devcontainer@137'),
  ('devcontainer4', 'devcontainer4@example.com', 'devcontainer@138');
