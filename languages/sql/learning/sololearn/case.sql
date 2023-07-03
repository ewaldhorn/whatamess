SELECT firstname, lastname,
CASE
    WHEN age >= 65 THEN 'Senior' 
END AS category
FROM Customers;

SELECT firstname, lastname,
CASE
  WHEN age >= 65 THEN 'Senior'
  WHEN age >= 25 AND age < 65 THEN 'Adult'
END AS category
FROM Customers;

SELECT firstname, lastname,
CASE
  WHEN age >= 65 THEN 'Senior'
  WHEN age >= 25 AND age < 65 THEN 'Adult'
  ELSE 'Youth'
END AS category
FROM Customers;

select firstname, lastname, salary,
case
when salary < 1501 then salary * 0.1
when salary > 1500 and salary < 2001 then salary * 0.2
else salary * 0.3
end as tax
from employees
order by lastname asc;