基本操作
==
**重命名，移表**
ALTER TABLE 旧表名 RENAME TO 新表名;	

**增加表中的列**
ALTER TABLE 表名 ADD COLUMN 列名 数据类型 [列的属性];

**指定位置**
ALTER TABLE 表名 ADD COLUMN 列名 列的类型 [列的属性] FIRST;
ALTER TABLE 表名 ADD COLUMN 列名 列的类型 [列的属性] AFTER 指定列名;

**删除列**
ALTER TABLE 表名 DROP COLUMN 列名;	

**修改信息**
ALTER TABLE 表名 MODIFY 列名 新数据类型 [新属性];	
ALTER TABLE 表名 CHANGE 旧列名 新列名 新数据类型 [新属性];	

**修改位置**
ALTER TABLE 表名 MODIFY 列名 列的类型 列的属性 FIRST;
ALTER TABLE 表名 MODIFY 列名 列的类型 列的属性 AFTER 指定列名;	

**插入值**
INSERT INTO 表名(列1, 列2, ...) VALUES(列1的值，列2的值, ...);	
INSERT INTO 表名(列1,列2, ...) VAULES(列1的值，列2的值, ...), (列1的值，列2的值, ...), (列1的值，列2的值, ...), ...;	

**设置默认值**
列名 列的类型 DEFAULT 默认值

小细节
--
AUTO_INCREMENT
SELECT * FROM
desc
show
mysql -h localhost -u root -p

**查询**
========
**查询**
SELECT 列名 FROM 表名;	
SELECT 列名 [AS] 列的别名 FROM 表名;	*改名输出查询*

**排序**
ORDER BY 列名 ASC|DESC	*ASC和DESC指的是排序方向。ASC是指按照指定列的值进行由小到大进行排序，也叫做升序，DESC是指按照指定列的值进行由大到小进行排序，也叫做降序，中间的|表示这两种方式只能选一个。*

WHERE
IN/NOT IN	eg:a IN (b1, b2, ...)
IS/IS NOT NULL
AND
OR
()

**聚集函数**
如果将函数以函数调用的形式放在查询列表中，那么会为表中符合WHERE条件的每一条记录调用一次该函数。

**模糊查询**
LIKE/NOT LIKE
*通配符：*% 、_、

**操作符**
DIV	a DIV b	除法，取商的整数部分
XOR	a XOR b	a和b有且只有一个为真，表达式为真

**处理函数**

	SELECT CONCAT('学号为', number, '的学生在《', subject, '》课程的成绩是：', score) AS 成绩描述 FROM student_score;

COUNT函数使用来统计行数的，它有下边两种使用方式：
1.COUNT(*)：对表中行的数目进行计数，不管列的值是不是NULL。
2.COUNT(列名)：对特定的列进行计数，会忽略掉该列为NULL的行。

**隐式类型转换**
1.把操作数类型转换为适合操作符计算的相应类型。
2.将函数参数转换为该函数期望的类型。
3.存储数据时，把某个值转换为某个列需要的类型。

**分组**
GROUP BY *分组列*
非分组模式会报错、可用ONLY_FULL_GROUP_BY的SQL模式
HAVING *针对分组的条件*

	SELECT subject, AVG(score) FROM student_score GROUP BY subject having subject = '母猪的产后护理';
 	SELECT subject, AVG(score) FROM student_score GROUP BY subject HAVING MAX(score) > 98;
嵌套分组： 用，隔开	
1.如果存在多个分组列，也就是嵌套分组，聚集函数将作用在最后的那个分组列上。
2.如果查询语句中存在WHERE子句和ORDER BY子句，那么GROUP BY子句必须出现在WHERE子句之后，ORDER BY子句之前。
3.WHERE子句和HAVING子句的区别：WHERE子句在分组前进行过滤，作用于每一条记录，WHERE子句过滤掉的记录将不包括在分组中。而HAVING子句在数据分组后进行过滤，作用于整个分组。
4.GROUP BY子句后也可以跟随表达式(但不能是聚集函数)。
5.非分组列不能单独出现在检索列表中(可以被放到聚集函数中)。

	SELECT department, major, COUNT(*) FROM student_info GROUP BY department, major;

**查询语句顺序**
SELECT [DISTINCT] 查询列表
[FROM 表名]
[WHERE 布尔表达式]
[GROUP BY 分组列表 ]
[HAVING 分组过滤条件]
[ORDER BY 排序列表]
[LIMIT 开始行, 限制条数]

**（子查询）**
标量子查询
列子查询
行子查询

	 SELECT * FROM student_score WHERE (number, subject) = (SELECT number, '母猪的产后护理' FROM student_info LIMIT 1);
表子查询
EXISTS和NOT EXISTS子查询
不相关子查询和相关子查询

	SELECT number, name, id_number, major FROM student_info WHERE EXISTS (SELECT * FROM student_score WHERE student_score.number = student_info.number);
对同一个表的子查询

	SELECT * FROM student_score WHERE subject = '母猪的产后护理' AND score > (SELECT AVG(score) FROM student_score WHERE subject = '母猪的产后护理');














储存程序
==






















