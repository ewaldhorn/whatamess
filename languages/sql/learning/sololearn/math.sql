SELECT firstname, lastname, age / 2 AS discount
FROM Customers;

SELECT firstname, lastname, weight/(height*height) AS bmi
FROM Customers;

SELECT MAX(age) FROM Customers;

SELECT MIN(age) FROM Customers;

SELECT AVG(age) FROM Customers;

SELECT city, COUNT(*) AS ct 
FROM Customers
GROUP BY city
ORDER BY ct DESC;

SELECT nickname, (SELECT MAX(score) WHERE nickname=nickname) AS best
FROM Scores
GROUP BY nickname
ORDER BY best DESC;