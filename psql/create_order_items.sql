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
