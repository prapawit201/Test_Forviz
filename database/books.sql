CREATE TABLE books (
      id integer AUTO_INCREMENT,
      uuid varchar(255) NOT NULL,
      name varchar(150) NOT NULL,
      author varchar(100) NOT NULL,
      type varchar(50) NOT NULL,
      price double NOT NULL,
      zone varchar(20),
      borrow_count integer,
      is_borrow boolean NOT NULL,
      deleted_at datetime,
      created_at datetime,
      updated_at datetime,
      PRIMARY KEY(id),
      UNIQUE KEY(name)
);