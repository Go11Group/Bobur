 CREATE TABLE Students (
    student_id INT PRIMARY KEY,
    student_name VARCHAR(100)
);

CREATE TABLE Courses (
    course_id INT PRIMARY KEY,
    course_name VARCHAR(100)
);

CREATE TABLE Enrollments (
    student_id INT,
    course_id INT,
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES Students(student_id),
    FOREIGN KEY (course_id) REFERENCES Courses(course_id)
);

SELECT * FROM Students;
 student_id | student_name 
------------+--------------
          1 | Alice
          2 | Bob


SELECT * FROM Courses;
 course_id | course_name 
-----------+-------------
       101 | Math
       102 | Science

SELECT * FROM Enrollments;
 student_id | course_id 
------------+-----------
          1 |       101
          1 |       102
          2 |       101


SELECT s.student_name, c.course_name
FROM Students s
JOIN Enrollments e ON s.student_id = e.student_id
JOIN Courses c ON e.course_id = c.course_id;
 student_name | course_name 
--------------+-------------
 Alice        | Math
 Alice        | Science
 Bob          | Math
