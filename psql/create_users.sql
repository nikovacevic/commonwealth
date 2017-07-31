CREATE TABLE users (
  id integer NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL,
  phone text NOT NULL,
  organization text,
  password_hash text NOT NULL,
  create_date timestamp without time zone DEFAULT now()
);
ALTER TABLE ONLY users ADD CONSTRAINT email UNIQUE (email);
ALTER TABLE ONLY users ADD CONSTRAINT phone UNIQUE (phone);
ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id);
