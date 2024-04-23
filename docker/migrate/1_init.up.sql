create table people (
    id serial primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    patronymic varchar(255)     
);

insert into people (name, surname, patronymic) values
    ('Александров', 'Андрей', 'Олегович'),
    ('Петров', 'Василий', 'Андреевич'),
    ('Смирнов', 'Сергей', 'Владимирович'),
    ('Кузнецов', 'Виталий', 'Дмитриевич'),
    ('Попов', 'Александр', 'Евгеньевич'),
    ('Васильев', 'Давид', 'Михайлович'),
    ('Иванов', 'Артем', 'Сергеевич');


create table cars (
    id serial primary key,
    reg_num varchar(255) not null,
    mark varchar(255) not null,
    model varchar(255) not null,
    year integer not null,
    people_id integer references people(id)
);

insert into cars (reg_num, mark, model, year, people_id) values
    ('X123XX150', 'Lada', 'Vesta', 2002, 1),
    ('Y456YY777', 'Toyota', 'Camry', 2015, 2),
    ('Z789ZZ999', 'BMW', 'X5', 2018, 3),
    ('A111AA111', 'Mercedes', 'E-Class', 2019, 4),
    ('B222BB222', 'Audi', 'A4', 2016, 5),
    ('C333CC333', 'Ford', 'Focus', 2017, 6);