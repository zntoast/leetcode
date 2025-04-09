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


# 184. 部门工资最高的员工
with temp as (
	SELECT id , name ,salary , departmentId , 
	DENSE_RANK() over( partition by departmentId ORDER BY salary desc ) as ranking 
	from employee
)
SELECT
	t2.`name` AS Department,
	t1.`name` AS Employee,
	salary AS Salary 
FROM
	temp t1
	LEFT JOIN department t2 ON t1.departmentId = t2.id 
WHERE
	t1.ranking IN (1,2,3)


# 550. 游戏玩法分析 IV
WITH p AS ( SELECT player_id, min( event_date ) AS login FROM activity GROUP BY player_id ) 

SELECT
	round( avg( a.event_date IS NOT NULL ), 2 ) fraction 
FROM
	p
	LEFT JOIN activity a ON p.player_id = a.player_id 
	AND DATEDIFF( a.event_date, p.login ) = 1
