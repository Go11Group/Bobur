 create table course(course_id uuid default gen_random_uuid(), course_name varchar(20));


INSERT INTO course (course_name) VALUES ('Mathematics');
INSERT INTO course (course_name) VALUES ('Physics');
INSERT INTO course (course_name) VALUES ('Chemistry');
INSERT INTO course (course_name) VALUES ('Biology');
INSERT INTO course (course_name) VALUES ('Computer Science');

CREATE TABLE student (
    id UUID DEFAULT gen_random_uuid(),
    name VARCHAR(20),
    lastname VARCHAR(20),
    phone VARCHAR(13),
    age INT
);


INSERT INTO student (name, lastname, phone, age) VALUES ('John', 'Doe', '12345678901', 20);
INSERT INTO student (name, lastname, phone, age) VALUES ('Jane', 'Smith', '23456789012', 22);
INSERT INTO student (name, lastname, phone, age) VALUES ('Alice', 'Johnson', '34567890123', 19);
INSERT INTO student (name, lastname, phone, age) VALUES ('Bob', 'Williams', '45678901234', 21);
INSERT INTO student (name, lastname, phone, age) VALUES ('Charlie', 'Brown', '56789012345', 23);
INSERT INTO student (name, lastname, phone, age) VALUES ('David', 'Jones', '67890123456', 20);
INSERT INTO student (name, lastname, phone, age) VALUES ('Eve', 'Garcia', '78901234567', 22);
INSERT INTO student (name, lastname, phone, age) VALUES ('Frank', 'Miller', '89012345678', 19);
INSERT INTO student (name, lastname, phone, age) VALUES ('Grace', 'Davis', '90123456789', 21);
INSERT INTO student (name, lastname, phone, age) VALUES ('Hank', 'Martinez', '01234567890', 23);

