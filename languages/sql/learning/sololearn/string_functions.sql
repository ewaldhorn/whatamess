SELECT CONCAT(firstname, lastname) AS name
FROM Customers;

SELECT CONCAT(firstname, ' ', lastname) AS name
FROM Customers;

SELECT LOWER(firstname) FROM Customers;

SELECT UPPER(firstname) FROM Customers;

SELECT SUBSTRING(firstname, 1, 3) 
FROM Customers;

-- Select two characters, from position 3
SELECT SUBSTRING(firstname, 3, 2) FROM Customers;

SELECT firstname, lastname, 
REPLACE(city, 'New York', 'NY') 
FROM Customers

SELECT CONCAT(
   SUBSTRING(firstname, 1, 1), 
   '. ', 
   UPPER(lastname)) AS name
FROM Customers


SELECT CONCAT(
    firstname,'.',lastname,'@company.com'
) AS email
FROM Employees
ORDER BY email;

SELECT LOWER(CONCAT(
    firstname,'.',lastname,'@company.com'
)) AS email
FROM Employees
ORDER BY email;