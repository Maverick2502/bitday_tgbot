CREATE TABLE people(
                       id BIGSERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       bdate date
);

drop table people;


-- в  тотже день поздравить
select name,bdate from people where bdate > current_timestamp  order by bdate  LIMIT 1;





DELETE FROM people where  id =  4;



SELECT name, date FROM  people;


CREATE TABLE ppp (
    name text
);

insert into ppp(name) VALUES ('AAAA');







                                                    --м --d
INSERT INTO people(name, bdate)VALUES ('Рашид', '2021-02-11') ;
INSERT INTO people(name, bdate)VALUES ('Рушод', '2021-03-02') ;
INSERT INTO people(name, bdate)VALUES ('Махина', '2021-03-27') ;
INSERT INTO people(name, bdate)VALUES ('Диана', ' 2021-04-01') ;
INSERT INTO people(name, bdate)VALUES ('Аминчон', ' 2021-05-01') ;
INSERT INTO people(name, bdate)VALUES ('Шамсия', ' 2021-05-15') ;
INSERT INTO people(name, bdate)VALUES ('Михаил', ' 2021-06-17') ;
INSERT INTO people(name, bdate)VALUES ('Тахмина', ' 2021-06-24') ;
INSERT INTO people(name, bdate)VALUES ('Нафиса', ' 2021-07-13');
INSERT INTO people(name, bdate)VALUES ('Садриддин', ' 2021-08-19');
INSERT INTO people(name, bdate)VALUES ('Азиз', ' 2021-09-10');
INSERT INTO people(name, bdate)VALUES ('Камила', ' 2021-09-11');
INSERT INTO people(name, bdate)VALUES ('Беназир', ' 2021-10-22');
INSERT INTO people(name, bdate)VALUES ('Динара', ' 2021-10-24');
INSERT INTO people(name, bdate)VALUES ('Манижа', ' 2021-11-12');
INSERT INTO people(name, bdate)VALUES ('Хушбахт', ' 2021-11-14');


UPDATE people SET  bdate = '2021-10-24' WHERE name = 'Динара';
UPDATE people SET  bdate = '2021-03-16' WHERE name = 'Рушод';


