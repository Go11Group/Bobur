CREATE TABLE weather1 (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID tipi, PRIMARY KEY, va standart qiymat gen_random_uuid() funksiyasi bilan
    name VARCHAR,       
    date VARCHAR,        
    windS VARCHAR,      
    temperature VARCHAR, 
    rain VARCHAR,       
    sun VARCHAR         
);

CREATE TABLE weather2 (
    city VARCHAR,
    temperature VARCHAR,
    description VARCHAR
);



CREATE TABLE weather3 (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    date VARCHAR,
    city VARCHAR NOT NULL,
    temperature VARCHAR,
    description VARCHAR,
); 


