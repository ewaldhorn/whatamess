CREATE TABLE Customers (
  id int NOT NULL AUTO_INCREMENT,
  firstname varchar(255),
  lastname varchar(255),
  PRIMARY KEY (id)
);

CREATE TABLE PhoneNumbers (
  id int NOT NULL AUTO_INCREMENT,
  customer_id int NOT NULL,
  number varchar(55),
  type varchar(55),
  PRIMARY KEY (id),
  FOREIGN KEY (customer_id) REFERENCES Customers(id)
);