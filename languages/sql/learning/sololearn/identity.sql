CREATE TABLE Customers (
  id int NOT NULL AUTO_INCREMENT,
  firstname varchar(255),
  lastname varchar(255)
);

INSERT INTO Customers (firstname, lastname, city, age)
VALUES
('demo', 'demo', 'Paris', 52),
('test', 'test', 'London', 21);

ALTER TABLE Customers
 AUTO_INCREMENT=555


INSERT INTO Customers (firstname, lastname, city, age)
VALUES
('test', 'test', 'London', 21)

insert into employees (firstname, lastname, salary)
values ('Wang','Lee',1900), ('Greta', 'Wu', 1200);

select id,firstname,lastname,salary from employees
order by id desc;
