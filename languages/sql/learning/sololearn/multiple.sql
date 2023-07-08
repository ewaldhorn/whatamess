SELECT firstname, lastname, city, number, type 
FROM Customers, PhoneNumbers
WHERE Customers.id = PhoneNumbers.customer_id;

SELECT Customers.firstname, Customers.lastname, Customers.city, PhoneNumbers.number, PhoneNumbers.type
FROM Customers, PhoneNumbers
WHERE Customers.id = PhoneNumbers.customer_id;

select Books.name, Books.year, Authors.name as author 
from Books, Authors
where Books.author_id = Authors.id 
order by Authors.name asc, Books.year asc;
