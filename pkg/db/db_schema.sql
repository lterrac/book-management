DROP DATABASE IF EXISTS book_management;
CREATE DATABASE book_management;
USE book_management;

DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
	`title` VARCHAR(50) NOT NULL DEFAULT '',
	`isbn` BIGINT(13) NOT NULL,
	`author` VARCHAR(30) NOT NULL DEFAULT '',
	`published_date` DATE,
	`edition` TINYINT unsigned zerofill,
	`description` TEXT,
	`genre` VARCHAR(15),
	KEY `title` (`title`) USING HASH,
    KEY `author` (`author`) USING HASH,
    KEY `dates` (`published_date`) USING BTREE,
    KEY `genre` (`genre`) USING HASH,
	PRIMARY KEY (`isbn`)
);

DROP TABLE IF EXISTS `collections`;

CREATE TABLE `collections` (
	`name` VARCHAR(30) NOT NULL,
	`description` TEXT,
	`creation_date` DATE,
	KEY `creation_date` (`creation_date`) USING BTREE,
	PRIMARY KEY (`name`)
);

DROP TABLE IF EXISTS `collections_members`;

CREATE TABLE `collection_members` (
	`collection_name` VARCHAR(30) NOT NULL,
	`book_isbn` BIGINT(13) NOT NULL,
	KEY `collection_name` (`collection_name`) USING HASH,
	PRIMARY KEY (`collection_name`,`book_isbn`)
);