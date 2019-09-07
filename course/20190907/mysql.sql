CREATE TABLE tablename (

) ENGINE=INNODB DEFAULT CHARSET utf8mb4;


INSERT INTO tablename VALUES(value1, value2, ..., valuen);
INSERT INTO tablename(col1, col2, ..., coln) VALUES(value1, value2, ..., valuen);


DELETE FROM tablename;
DELETE FROM tablename WHERE 条件;


UPDATE tablename
SET col1=value1, col2=value2, ..., coln=valuen;


UPDATE tablename
SET col1=value1, col2=value2, ..., coln=valuen
WHERE 条件;


SELECT * FROM tablename;
SELECT count(*) FROM tablename;

SELECT col1, col2, ..., coln FROM tablename;
SELECT col1, col2, ..., coln FROM tablename limit 10;
SELECT col1, col2, ..., coln FROM tablename limit 10 offset 10;
SELECT col1, col2, ..., coln FROM tablename WHERE 条件;
SELECT col1, col2, ..., coln FROM tablename WHERE 条件 limit 10;
SELECT col1, col2, ..., coln FROM tablename ORDER BY colx [DESC/ASC];

SELECT col1, col2, ..., coln FROM tablename ORDER BY colx [DESC/ASC], colx2 [DESC/ASC];
SELECT col1, col2, ..., coln FROM tablename WHERE 条件 ORDER BY colx [DESC/ASC];
SELECT col1, col2, ..., coln FROM tablename WHERE 条件 ORDER BY colx [DESC/ASC] limit 10;
SELECT col1, col2, ..., coln FROM tablename WHERE 条件 ORDER BY colx [DESC/ASC] limit 10 offset 100;


name age region score

SELECT age, count(*) FROM user GROUP BY age;

SELECT region, age, count(*) FROM user GROUP BY region, age;

SELECT region, avg(age) FROM user GROUP BY region;

SELECT region, sum(score) FROM user GROUP BY region;

SELECT region, avg(age), sum(score), max(score), min(score) FROM user GROUP BY region;


SELECT region, avg(age), sum(score), max(score), min(score) FROM user GROUP BY region HAVING avg(score) >= 80;

SELECT region, avg(age), avg(score), sum(score), max(score), min(score) FROM user WHERE age < 10 GROUP BY region HAVING avg(score) >= 80;


select * from ? whrere