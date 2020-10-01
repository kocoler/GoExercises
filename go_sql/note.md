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



索引
==
** 1、分类：单列索引(普通索引，唯一索引，主键索引)、组合索引、全文索引、空间索引**

 1.1、单列索引：一个索引只包含单个列，但一个表中可以有多个单列索引。 这里不要搞混淆了。
　　　　   1.1.1、普通索引：MySQL中基本索引类型，没有什么限制，允许在定义索引的列中插入重复值和空值，纯粹为了查询数据更快一点。
　　　　   1.1.2、唯一索引：索引列中的值必须是唯一的，但是允许为空值，
　　　　   1.1.3、主键索引：是一种特殊的唯一索引，不允许有空值。

 1.2、组合索引
　　在表中的多个字段组合上创建的索引，只有在查询条件中使用了这些字段的左边字段时，索引才会被使用，使用组合索引时遵循最左前缀集合。这个如果还不明白，等后面举例讲解时在细说

1.3、全文索引
　　全文索引，只有在MyISAM引擎上才能使用，只能在CHAR,VARCHAR,TEXT类型字段上使用全文索引，介绍了要求，说说什么是全文索引，就是在一堆文字中，通过其中的某个关键字等，就能找到该字段所属的记录行，比如有"你是个靓仔，靓女 ..." 通过靓仔，可能就可以找到该条记录。这里说的是可能，因为全文索引的使用涉及了很多细节，我们只需要知道这个大概意思，如果感兴趣进一步深入使用它，那么看下面测试该索引时，会给出一个博文，供大家参考。

 1.4、空间索引
　　空间索引是对空间数据类型的字段建立的索引，MySQL中的空间数据类型有四种，GEOMETRY、POINT、LINESTRING、POLYGON。在创建空间索引时，使用SPATIAL关键字。要求，引擎为MyISAM，创建空间索引的列，必须将其声明为NOT NULL。


创建索引

	CREATE [UNIQUE|FULLTEXT|SPATIAL] [INDEX|KEY] 索引名称 ON 表名(创建索引的字段名[length])[ASC|DESC]

	CREATE INDEX BkBookNameIdx ON book(bookname);    //示例
	
	ALTER TABLE 表名 ADD INDEX 表名（创建索引的字段名[length])
	
	ALTER TABLE book ADD INDEX BkNameIdx(bookname(30));  //create或alter
	
	//创建联合索引
	CREATE TABLE index_demo(
    c1 INT,
    c2 INT,
    c3 CHAR(1),
    PRIMARY KEY(c1),
    INDEX idx_c2_c3 (c2, c3)
	);


查看一张表中所创建的索引
	
	SHOW INDEX FROM 表名\G

删除索引

	ALTER TABLE 表名 DROP INDEX 索引名。
	DROP INDEX 索引名 ON 表名；




B+tree 索引
==
**聚簇索引**
1. 使用记录主键值的大小进行记录和页的排序，这包括三个方面的含义：

+ 页内的记录是按照主键的大小顺序排成一个单向链表。

+ 各个存放用户记录的页也是根据页中用户记录的主键大小顺序排成一个双向链表。

+ 存放目录项记录的页分为不同的层次，在同一层次中的页也是根据页中目录项记录的主键大小顺序排成一个双向链表。

2. B+树的叶子节点存储的是完整的用户记录。

+ 所谓完整的用户记录，就是指这个记录中存储了所有列的值（包括隐藏列）。

我们把具有这两种特性的B+树称为聚簇索引，所有完整的用户记录都存放在这个聚簇索引的叶子节点处。这种聚簇索引并不需要我们在MySQL语句中显式的使用INDEX语句去创建（后边会介绍索引相关的语句），InnoDB存储引擎会自动的为我们创建聚簇索引。另外有趣的一点是，在InnoDB存储引擎中，聚簇索引就是数据的存储方式（所有的用户记录都存储在了叶子节点），也就是所谓的索引即数据，数据即索引。

**二级索引**
这个B+树与上边介绍的聚簇索引有几处不同：

+ 使用记录c2列的大小进行记录和页的排序，这包括三个方面的含义：

    + 页内的记录是按照c2列的大小顺序排成一个单向链表。
    + 各个存放用户记录的页也是根据页中记录的c2列大小顺序排成一个双向链表。
    + 存放目录项记录的页分为不同的层次，在同一层次中的页也是根据页中目录项记录的c2列大小顺序排成一个双向链表。
    + B+树的叶子节点存储的并不是完整的用户记录，而只是c2列+主键这两个列的值。
   + 目录项记录中不再是主键+页号的搭配，而变成了c2列+页号的搭配。

因为这种按照非主键列建立的B+树需要一次回表操作才可以定位到完整的用户记录，所以这种B+树也被称为二级索引（英文名secondary index），或者辅助索引。

**联合索引**






