CREATE TABLE invoices (
  id SERIAL PRIMARY KEY,
  create_date timestamp without time zone,
  create_user_id int REFERENCES users (id),
  due_date timestamp without time zone,
  total numeric(10,2) CHECK (total >= 0)
);
