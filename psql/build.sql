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
--------------------------------------------------------------------------------
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
CREATE TABLE invoices (
  id SERIAL PRIMARY KEY,
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  due_date timestamp without time zone,
  total numeric(10,2) CHECK (total >= 0)
);
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  invoice_id int REFERENCES invoices (id),
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  user_id int REFERENCES users (id)
);
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  amount numeric(10,2),
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  description text,
  order_id int REFERENCES orders (id),
  product_option_id int REFERENCES product_options (id),
  quantity int
);
