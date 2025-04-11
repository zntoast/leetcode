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

# 570. 至少有5名直接下属的经理
WITH temp AS (
	SELECT
		managerId 
	FROM
		`employee` 
	GROUP BY
		managerId 
HAVING
	count( managerId )  >= 5)
	
SELECT name from employee WHERE id in ( SELECT managerId FROM temp )

# 1045. 买下所有产品的客户
WITH temp AS ( SELECT DISTINCT customer_id, product_key FROM Customer )
SELECT
    customer_id,
    count( customer_id ) AS num 
FROM
	temp 
GROUP BY
	customer_id 
HAVING
	( SELECT count( 1 ) FROM Product ) = num


# 262. 行程和用户
WITH banned_users AS (
    SELECT 
        users_id,
        role 
    FROM `users` 
    WHERE banned = "Yes"
),
valid_trips AS (
    SELECT 
        request_at,
        COUNT(*) AS total_trips,
        SUM(status IN ("cancelled_by_driver", "cancelled_by_client")) AS cancelled_trips
    FROM trips
    WHERE 
        client_id NOT IN (SELECT users_id FROM banned_users WHERE role = "client")
        AND driver_id NOT IN (SELECT users_id FROM banned_users WHERE role = "driver")
        AND request_at BETWEEN '2013-10-01' AND '2013-10-03'
    GROUP BY request_at
)
SELECT 
    request_at AS Day,
    ROUND(
        IFNULL(cancelled_trips / NULLIF(total_trips, 0), 0),
        2
    ) AS "Cancellation Rate"
FROM valid_trips
ORDER BY request_at;