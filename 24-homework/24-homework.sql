NSERT INTO book (id, name, page, author_name) VALUES
(1, 'Alpomish', 200, 'Anonim'),
(2, 'Otgan kunlar', 320, 'Abdulla Qodiriy'),
(3, 'Mehrobdan Chayon', 280, 'Abdulla Qodiriy'),
(4, 'Shum bola', 150, 'Gafur Gulom'),
(5, 'Yulduzli tunlar', 450, 'Pirimqul Qodirov'),
(6, 'Qutlugqon', 360, 'Abdulla Qodiriy'),
(7, 'Sarvqomat dilbarim', 300, 'Abdulla Qodiriy'),
(8, 'Kecha va kunduz', 400, 'Cholpon'),
(9, 'Qorakoz majnun', 250, 'Mirkarim Osim'),
(10, 'Diydor', 500, 'Asqad Muxtor');

INSERT INTO author (id, name) VALUES
(1, 'Anonim'),
(2, 'Abdulla Qodiriy'),
(3, 'Gafur Gulom'),
(4, 'Pirimqul Qodirov'),
(5, 'Cholpon'),
(6, 'Mirkarim Osim'),
(7, 'Asqad Muxtor');

SELECT * FROM book;
 id |        name        | page |   author_name    
----+--------------------+------+------------------
  1 | Alpomish           |  200 | Anonim
  2 | Otgan kunlar       |  320 | Abdulla Qodiriy
  3 | Mehrobdan Chayon   |  280 | Abdulla Qodiriy
  4 | Shum bola          |  150 | Gafur Gulom
  5 | Yulduzli tunlar    |  450 | Pirimqul Qodirov
  6 | Qutlugqon          |  360 | Abdulla Qodiriy
  7 | Sarvqomat dilbarim |  300 | Abdulla Qodiriy
  8 | Kecha va kunduz    |  400 | Cholpon
  9 | Qorakoz majnun     |  250 | Mirkarim Osim
 10 | Diydor             |  500 | Asqad Muxtor


SELECT * FROM author;
 id |       name       
----+------------------
  1 | Anonim
  2 | Abdulla Qodiriy
  3 | Gafur Gulom
  4 | Pirimqul Qodirov
  5 | Cholpon
  6 | Mirkarim Osim
  7 | Asqad Muxtor


SELECT book.name, book.page, author.name 
FROM book INNER JOIN author ON book.author_name = author.name 
AND book.id = authOr.id;
     name     | page |      name       
--------------+------+-----------------
 Alpomish     |  200 | Anonim
 Otgan kunlar |  320 | Abdulla Qodiriy

SELECT book.name, book.page, author.name 
FROM book LEFT JOIN author ON book.author_name = author.name 
AND book.id = authOr.id;
        name        | page |      name       
--------------------+------+-----------------
 Alpomish           |  200 | Anonim
 Otgan kunlar       |  320 | Abdulla Qodiriy
 Qutlugqon          |  360 | 
 Yulduzli tunlar    |  450 | 
 Qorakoz majnun     |  250 | 
 Mehrobdan Chayon   |  280 | 
 Shum bola          |  150 | 
 Sarvqomat dilbarim |  300 | 
 Diydor             |  500 | 
 Kecha va kunduz    |  400 | 

SELECT book.name, book.page, author.name 
FROM book RIGHT JOIN author ON book.author_name = author.name 
AND book.id = authOr.id;
     name     | page |       name       
--------------+------+------------------
 Alpomish     |  200 | Anonim
 Otgan kunlar |  320 | Abdulla Qodiriy
              |      | Gafur Gulom
              |      | Pirimqul Qodirov
              |      | Cholpon
              |      | Mirkarim Osim
              |      | Asqad Muxtor


SELECT book.name, book.page, author.name 
FROM book FULL JOIN author ON book.author_name = author.name 
AND book.id = authOr.id;
        name        | page |       name       
--------------------+------+------------------
 Alpomish           |  200 | Anonim
 Otgan kunlar       |  320 | Abdulla Qodiriy
                    |      | Gafur Gulom
                    |      | Pirimqul Qodirov
                    |      | Cholpon
                    |      | Mirkarim Osim
                    |      | Asqad Muxtor
 Qutlugqon          |  360 | 
 Yulduzli tunlar    |  450 | 
 Qorakoz majnun     |  250 | 
 Mehrobdan Chayon   |  280 | 
 Shum bola          |  150 | 
 Sarvqomat dilbarim |  300 | 
 Diydor             |  500 | 
 Kecha va kunduz    |  400 | 
