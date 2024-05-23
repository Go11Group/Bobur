sudo -u postgres psql
[sudo] password for bobur: 
could not change directory to "/home/bobur": Permission denied
psql (14.11 (Ubuntu 14.11-0ubuntu0.22.04.1))
Type "help" for help.

postgres=# \l
postgres=# \l
                                  List of databases
   Name    |  Owner   | Encoding |   Collate   |    Ctype    |   Access privileges   
-----------+----------+----------+-------------+-------------+-----------------------
 bobur     | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | 
 postgres  | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | 
 template0 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/postgres          +
           |          |          |             |             | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/postgres          +
           |          |          |             |             | postgres=CTc/postgres
(4 rows)

postgres=# \c bobur 
You are now connected to database "bobur" as user "postgres".
bobur=# \dt
          List of relations
 Schema |  Name   | Type  |  Owner   
--------+---------+-------+----------
 public | student | table | postgres
(1 row)

bobur=# create table student(id serial primary key, name varchar(50), phone varchar(16), course int, age int);
CREATE TABLE

bobur-# \d student
                                   Table "public.student"
 Column |         Type          | Collation | Nullable |               Default               
--------+-----------------------+-----------+----------+-------------------------------------
 id     | integer               |           | not null | nextval('student_id_seq'::regclass)
 name   | character varying(50) |           |          | 
 phone  | character varying(16) |           |          | 
 course | integer               |           |          | 
 age    | integer               |           |          | 
Indexes:
    "student_pkey" PRIMARY KEY, btree (id)

INSERT INTO student (name, phone, course, age) VALUES
('bobur', '+998938532273', 1, 19),
('diyor', '+998900062916', 1, 19),
('bekzod', '+998935485234', 0, 18),
('dilshod', '+998936467961', 0, 19),
('husan', '+998942255334', 4, 23),
('jonibek', '+998888415777', 1, 19),
('shaymardon', '+998938128877', 1, 19);
INSERT 0 7

INSERT INTO student (name, phone, course, age) VALUES
('jasurbek', '+998937651555', 0, 15),
('javlonbek', '+998937655', 0, 13),
('davlatbek', '+9989145632', 0, 15);
INSERT 0 3

select * from student;
 id |    name    |     phone     | course | age 
----+------------+---------------+--------+-----
  1 | bobur      | +998938532273 |      1 |  19
  2 | diyor      | +998900062916 |      1 |  19
  3 | bekzod     | +998935485234 |      0 |  18
  4 | dilshod    | +998936467961 |      0 |  19
  5 | husan      | +998942255334 |      4 |  23
  6 | jonibek    | +998888415777 |      1 |  19
  7 | shaymardon | +998938128877 |      1 |  19
  8 | jasurbek   | +998937651555 |      0 |  15
  9 | javlonbek  | +998937655    |      0 |  13
 10 | davlatbek  | +9989145632   |      0 |  15
(10 rows)

-- ~ ~ ~ ~ update qilish ~ ~ ~ ~
UPDATE student SET course = 2 WHERE name = 'diyor';
UPDATE 1

select * from student;
 id |    name    |     phone     | course | age 
----+------------+---------------+--------+-----
  1 | bobur      | +998938532273 |      1 |  19
  3 | bekzod     | +998935485234 |      0 |  18
  4 | dilshod    | +998936467961 |      0 |  19
  5 | husan      | +998942255334 |      4 |  23
  6 | jonibek    | +998888415777 |      1 |  19
  7 | shaymardon | +998938128877 |      1 |  19
  8 | jasurbek   | +998937651555 |      0 |  15
  9 | javlonbek  | +998937655    |      0 |  13
 10 | davlatbek  | +9989145632   |      0 |  15
  2 | diyor      | +998900062916 |      2 |  19
(10 rows)

-- buyerda set dan keyin ustunni nommini berib unga yangi qiymat qo`shamiz
-- va shart beramiz yani qaysi ism yoki id di shu songa degan va h.k 
-- keyin o`zi bilib oladi nimani o`zgartirishni

-- tabelda gi malumotni delete qilish

delete from student where name = 'javlonbek';
DELETE 1

 select * from student;
 id |    name    |     phone     | course | age 
----+------------+---------------+--------+-----
  1 | bobur      | +998938532273 |      1 |  19
  3 | bekzod     | +998935485234 |      0 |  18
  4 | dilshod    | +998936467961 |      0 |  19
  5 | husan      | +998942255334 |      4 |  23
  6 | jonibek    | +998888415777 |      1 |  19
  7 | shaymardon | +998938128877 |      1 |  19
  8 | jasurbek   | +998937651555 |      0 |  15
 10 | davlatbek  | +9989145632   |      0 |  15
  2 | diyor      | +998900062916 |      2 |  19
(9 rows)

-- where dan keyin hohlagan filtini besak bo`ladi

-- group by qilish 

select name, count(*) from student group by name;
    name    | count 
------------+-------
 shaymardon |     1
 husan      |     1
 jasurbek   |     1
 diyor      |     1
 bobur      |     1
 jonibek    |     1
 bekzod     |     1
 davlatbek  |     1
 dilshod    |     1
(9 rows)

select course, count(*) from student group by course;
 course | count 
--------+-------
      4 |     1
      0 |     4
      2 |     1
      1 |     3
(4 rows)

-- nimani chiqarishi sorasangiz group by qilishda ham shu narsalarni 
-- yozish kerak 

-- Masalan :
select course, name, id, count(*) from student group by course, name, id;
 course |    name    | id | count 
--------+------------+----+-------
      4 | husan      |  5 |     1
      0 | dilshod    |  4 |     1
      0 | davlatbek  | 10 |     1
      1 | jonibek    |  6 |     1
      2 | diyor      |  2 |     1
      1 | shaymardon |  7 |     1
      1 | bobur      |  1 |     1
      0 | jasurbek   |  8 |     1
      0 | bekzod     |  3 |     1
(9 rows)

-- Agar shunaqa yozmasangiz
select course, name, id, count(*) from student group by course, name;
ERROR:  column "student.id" must appear in the GROUP BY clause or be used in an aggregate function
LINE 1: select course, name, id, count(*) from student group by cour...
-- bunaqa xatolik beradi yana id ni bermagansan deydi


-- order by 

SELECT *
FROM student
ORDER BY age;
 id |    name    |     phone     | course | age 
----+------------+---------------+--------+-----
 10 | davlatbek  | +9989145632   |      0 |  15
  8 | jasurbek   | +998937651555 |      0 |  15
  3 | bekzod     | +998935485234 |      0 |  18
  7 | shaymardon | +998938128877 |      1 |  19
  2 | diyor      | +998900062916 |      2 |  19
  4 | dilshod    | +998936467961 |      0 |  19
  6 | jonibek    | +998888415777 |      1 |  19
  1 | bobur      | +998938532273 |      1 |  19
  5 | husan      | +998942255334 |      4 |  23
(9 rows)
 
-- bu hozir age bo`yicha tartibab berdi 
-- Agar oxiriga DESC qo`ysak kamayish tartibida tartiblab beradi

-- group by bn order by ni farqi 
-- GROUP BY ma'lumotlar to'plamlarini yaratish uchun
-- ORDER BY esa tartiblash uchun ishlatilinadi


-- join - jadval larni birlashtirib ular ustida amallar bajarish uchun ishlatiladi
-- join larni 4 xil turi bor ular inner join, left join, right joint va full join

-- inner join - bu ikkita jadvar oradidagi bir xilliklarni olib beradi
-- agar bir xillik bo`lmasa chiqarmaydi

-- left join - birinchi chap jadvaldagi ni olib keyin o`ng jadvaldagini olib 
-- tekshirib ko`radi ikkalasidagilar mos kelmasa null qaytaradi

-- righ joinda esa left joinga teskari

-- full join - ikki jadvallardagi ma'lumotlar birlashtiriladi. Agar moslik 
--mavjud bo'lmasa, NULL qiymat beriladi.




