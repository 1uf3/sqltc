## sqltc - a simple sql file type checker

[![sqlint](https://github.com/lufeee/sqltc/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/lufeee/sqltc/actions/workflows/go.yml)

### About

sqltc is a simple command-line sql variable type checker whitch reads your sql file.

Now, Only avaiable one file.

### Installation

```sh
go install github.com/lufeee/sqltc/cmd/sqltc
```

### Usage

```sql
-- +migrate Up
CREATE DATABASE IF NOT EXISTS test;
CREATE TABLE IF NOT EXISTS test.testdata(
  name VARCHAR(255) NOT NULL,
  info VARCHAR(200) NOT NULL,
  PRIMARY KEY (name)
);
-- +migrate Down
```

Execute
```sh
sqltc -file ../create_table.sql
```

Response
```
[{"Name":"name","Type":"VARCHAR","IsNULL":true},{"Name":"info","Type":"VARCHAR","IsNULL":true}]
```

### License 

MIT license.

<hr>
