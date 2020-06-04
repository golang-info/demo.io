- mysql install:

  https://github.com/k8s-info/k8s_base
  
- mysql configuration:
```bash 
-- Create a new `webdb` database.
CREATE DATABASE webdb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE webdb;

CREATE USER 'webuser'@'localhost';
GRANT SELECT, INSERT, UPDATE, CREATE, ALTER ON webdb.* TO 'webuser'@'localhost';

ALTER USER 'webuser'@'localhost' IDENTIFIED BY 'webpassword';

-------
set password for 'webuser'@'localhost' = password('webpassword’);

mysql> grant all privileges on *.* to 'webuser'@'%' identified by 'webpassword' with grant option;
Query OK, 0 rows affected (0.01 sec)

mysql> flush privileges;
Query OK, 0 rows affected (0.00 sec)
--------

-- Create a 'posts' table
create table posts (
id INTEGER not null primary key auto_increment,
title VARCHAR(100) NOT null,
content text not null,
created datetime not null,
expires datetime not null
);

-- Add an index on the created column.
create index idx_posts_created on posts(created);

INSERT INTO posts (title, content, created, expires) values (
'net http package',
'Package http provides HTTP client and server implementations.',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(),INTERVAL 365 DAY)
);

INSERT INTO posts (title, content, created, expires) VALUES (
'database sql package',
'Package sql provides a generic interface around SQL (or QQL-LIKE) database. The sql package must be used in conjunction with a database driver',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(),INTERVAL 365 DAY)
);

SELECT * from posts;

```