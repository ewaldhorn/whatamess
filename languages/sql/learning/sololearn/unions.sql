SELECT firstname, lastname, age FROM Customers
UNION
SELECT firstname, lastname, age FROM Contacts;

SELECT firstname, lastname, age FROM Customers
UNION ALL
SELECT firstname, lastname, age FROM Contacts;

SELECT firstname, lastname, age, city FROM Customers
UNION
SELECT firstname, lastname, age, NULL FROM Contacts;

SELECT firstname, lastname, age FROM Customers
WHERE age > 30
UNION
SELECT firstname, lastname, age FROM Contacts
WHERE age < 25;

select name, year from books where year > 1900
union
select name, 2022 from new
order by name;
