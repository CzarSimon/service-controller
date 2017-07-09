CREATE TABLE PROJECT(
  NAME VARCHAR(50) PRIMARY KEY,
  FOLDER VARCHAR(255),
  SWARM_TOKEN VARCHAR(100),
  IS_ACTIVE BOOLEAN,
  NETWORK VARCHAR(60),
  MASTER_TOKEN VARCHAR(260)
);

CREATE TABLE NODE(
  PROJECT VARCHAR(50),
  IP VARCHAR(50),
  OS VARCHAR(10) DEFAULT 'linux',
  IS_MASTER BOOLEAN,
  FOREIGN KEY (PROJECT) REFERENCES PROJECT(NAME),
  PRIMARY KEY (PROJECT, IP)
);
