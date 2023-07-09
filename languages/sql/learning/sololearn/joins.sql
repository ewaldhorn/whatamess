SELECT firstname, lastname, city, number, type
FROM Customers JOIN PhoneNumbers
ON Customers.id = PhoneNumbers.customer_id;

SELECT C.firstname, C.lastname, C.city, PN.number, PN.type
FROM Customers AS C JOIN PhoneNumbers AS PN
ON C.id = PN.customer_id;

SELECT C.firstname, C.lastname, C.city, PN.number, PN.type
FROM Customers AS C LEFT JOIN PhoneNumbers AS PN
ON C.id = PN.customer_id;

SELECT C.firstname, C.lastname, C.city, PN.number, PN.type
FROM PhoneNumbers AS PN RIGHT JOIN Customers AS C
ON C.id = PN.customer_id
ORDER BY C.id;

select a.name as name, count(b.id) as books
from Books as b
right join Authors as a
on b.author_id = a.id
group by a.name
order by books desc;
