create table users
(
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    name varchar(100), 
    email varchar(100), 
         timestamp, 
    password varchar(100), 
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at bigint DEFAULT 0
);


create table courses
(
    course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    title  varchar(100), 
    description text, 
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at bigint DEFAULT 0
);


create table lessons
(
    lesson_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id uuid REFERENCES courses(course_id) ON DELETE CASCADE,
    title varchar(100), 
    content text, 
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at bigint DEFAULT 0
);


create table enrollments 
(
    enrollment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid REFERENCES users(user_id) ON DELETE CASCADE,
    course_id uuid REFERENCES courses(course_id) ON DELETE CASCADE,
    enrollment_date timestamp,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at bigint DEFAULT 0
);

-- 1 - task
select c.course_id, c.title, c.description from
enrollments as e left join courses as c on c.course_id=e.course_id  
where e.user_id=id;


-- 2 - task
select c.course_id, l.lesson_id, l.title, l.content from lessons as l 
left join courses as c on c.course_id=l.course_id
where l.course_id=id;

--  3 - task

select u.name, u.email, u.user_id from enrollments as e 
inner join users as u on u.user_id=e.user_id 
where e.course_id=id;


--  4 - task 
select u.user_id, u.name, u.email from users as u 
where where true

