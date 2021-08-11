CREATE TABLE people(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    bdate date
);

drop table people;

INSERT INTO people(name, bdate) VALUES ('Aziz', '2021-09-10') ;
INSERT INTO people(name, bdate) VALUES ('DDD', '2021-09-10') ;


Select to_char(bdate, 'DD MM YY') FROM people ;


SELECT name, date FROM  people;


CREATE TABLE ppp (
    name text
);

insert into ppp(name) VALUES ('AAAA');

CREATE TABLE ddd (
    tt timestamp  WITHOUT TIME ZONE
);

INSERT INTO ddd