select e.name from Employee e
LEFT JOIN Employee e1 on e.managerId = e1.id and e.salary > e1.salary


select count(p1.email) from Person

# 177. 第N高的薪水
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
DECLARE M INT; 
    SET M = N-1; 
  RETURN (
      SELECT DISTINCT salary
      FROM Employee
      ORDER BY salary DESC
      LIMIT M, 1
  );
END

# 178. 分数排名
SELECT
	t1.score,
	t2.rank_id as `rank`
FROM
	Scores t1
	LEFT JOIN (
	SELECT
		score,
		DENSE_RANK() OVER ( ORDER BY score DESC ) AS rank_id 
	FROM
		( SELECT DISTINCT score FROM Scores ) t 
	ORDER BY
	score DESC) t2 ON t1.score = t2.score 
ORDER BY
	t1.score DESC

# 180. 连续出现的数字
SELECT DISTINCT
    l1.Num AS ConsecutiveNums
FROM
    Logs l1,
    Logs l2,
    Logs l3
WHERE
    l1.Id = l2.Id - 1
    AND l2.Id = l3.Id - 1
    AND l1.Num = l2.Num
    AND l2.Num = l3.Num
;


