CREATE TABLE book_borrow_logs (
      id integer AUTO_INCREMENT,
      book_id integer NOT NULL,
      status varchar(10) NOT NULL,
      created_at datetime,
      updated_at datetime,
      PRIMARY KEY(id)
);