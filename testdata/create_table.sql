-- +migrate Up
CREATE DATABASE IF NOT EXISTS test;
CREATE TABLE IF NOT EXISTS test.testdata(
  name VARCHAR(255) NOT NULL,
  info VARCHAR(200) NOT NULL,
  PRIMARY KEY (name)
);
-- +migrate Down
