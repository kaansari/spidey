CREATE TABLE IF NOT EXISTS accounts (
  id CHAR(27) PRIMARY KEY,
  name VARCHAR(24) NOT NULL,
  email VARCHAR(56),
  social_id VARCHAR(100),
  photoURL VARCHAR(100),
  locationURL VARCHAR (100),
  createtime DATE,
  last_updated DATE
);
