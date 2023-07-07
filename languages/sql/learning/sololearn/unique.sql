ALTER TABLE Customers
ADD UNIQUE (lastname);

INSERT INTO Customers (firstname, lastname, city, age)
VALUES
('demo', 'Anderson', 'London', 24);


```

The Primary key uniquely identifies each record of a table. It is usually set as an auto increment integer.

Foreign keys are used to create relationships between tables. They refer to the primary key in other tables.

A table can have multiple foreign keys, but only one single primary key.

The UNIQUE constraint is used to make values in a column unique.
```
