CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  invoice_id int REFERENCES invoices (id),
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  user_id int REFERENCES users (id)
);
