CREATE TABLE product_options (
  id SERIAL PRIMARY KEY,
  cost numeric(10,2) CHECK (cost >= 0),
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  description text,
  is_active boolean,
  price numeric(10,2) CHECK (price >= 0),
  product_id int REFERENCES products (id)
);
