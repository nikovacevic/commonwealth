CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text UNIQUE NOT NULL,
  phone text UNIQUE NOT NULL,
  organization text,
  password_hash text NOT NULL,
  create_date timestamp without time zone DEFAULT now()
);
