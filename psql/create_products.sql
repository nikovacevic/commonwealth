CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  cost numeric(10,2) CHECK (cost >= 0),
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  description text,
  is_active boolean,
  name text,
  price numeric(10,2) CHECK (price >= 0)
);
