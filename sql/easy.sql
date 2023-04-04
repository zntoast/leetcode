select e.name from Employee e
LEFT JOIN Employee e1 on e.managerId = e1.id and e.salary > e1.salary


select count(p1.email) from Person