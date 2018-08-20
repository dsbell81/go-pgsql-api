CREATE TABLE users(
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  email VARCHAR (100) UNIQUE NOT NULL,
  password text NOT NULL,
  created TIMESTAMP NOT NULL,
  modified TIMESTAMP NOT NULL,
  last_login TIMESTAMP
);
