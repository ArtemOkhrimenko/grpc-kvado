-- up

create table books (
                       id int(11) auto_increment primary key,
                       name VARCHAR(30) not null
);

create table creators (
                          id int(11) auto_increment primary key,
                          name VARCHAR(30) not null

);


CREATE TABLE `books_creators`
(
    `book_id`    INT(11) NOT NULL,
    `creator_id` INT(11) NOT NULL,
    PRIMARY KEY (`book_id`, `creator_id`),
    INDEX `book_id` (`book_id`),
    INDEX `creator_id` (`creator_id`),
    CONSTRAINT `FK_creators` FOREIGN KEY (`book_id`)
        REFERENCES `creators` (`id`) ON DELETE CASCADE,
    CONSTRAINT `FK_books` FOREIGN KEY (`book_id`)
        REFERENCES `books` (`id`) ON DELETE CASCADE

);

INSERT books(name)
VALUES ('book1');

INSERT creators(name)
VALUES ('author1');

insert books_creators(book_id, creator_id)
values (1, 1);

INSERT books(name)
VALUES ('book2');

INSERT creators(name)
VALUES ('author2');

insert books_creators(book_id, creator_id)
values (2, 2);

INSERT books(name)
VALUES ('book3');

INSERT creators(name)
VALUES ('author3');

insert books_creators(book_id, creator_id)
values (3, 3);

-- down

drop table books_creators;
drop table books;
drop table  creators;