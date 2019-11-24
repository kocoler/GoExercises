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

AUTO_INCREMENT
SELECT * FROM
desc
show

**查询**
SELECT 列名 FROM 表名;	
SELECT 列名 [AS] 列的别名 FROM 表名;	*改名输出查询*


