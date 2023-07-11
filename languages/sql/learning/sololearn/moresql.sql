SELECT * FROM Customers;

SELECT * FROM PhoneNumbers;

SELECT C.id, C.firstname, C.lastname, PN.number, PN.type
FROM Customers AS C LEFT JOIN PhoneNumbers AS PN
ON C.id = PN.customer_id
ORDER BY C.id;

SELECT C.id, COUNT(PN.number) AS count
FROM Customers AS C LEFT JOIN PhoneNumbers AS PN
ON C.id = PN.cus;

SELECT AVG(count) FROM
(SELECT C.id, COUNT(PN.number) AS count
FROM Customers AS C LEFT JOIN PhoneNumbers AS PN
ON C.id = PN.customer_id
GROUP BY C.id) AS Numbers;

