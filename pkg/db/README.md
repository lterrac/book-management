# Database design
In the database books and collections are stored within these tables:

## Books
Used to store the Book resource using `isbn` as primary key. To ease the filtering operations, secondary data structure are built using hashing for `title`, `author`, `dates` and `genre`. `published_date` field has a BTREE index to ease the search within a range of dates.
```
CREATE TABLE `books` (
	`title` VARCHAR(50) NOT NULL DEFAULT '',
	`isbn` BIGINT(13) NOT NULL DEFAULT '',
	`author` VARCHAR(30) NOT NULL DEFAULT '',
	`published_date` DATETIME DEFAULT '',
	`edition` TINYINT unsigned zerofill NOT NULL DEFAULT '',
	`description` TEXT,
	`genre` VARCHAR(15),
	KEY `title` (`title`) USING HASH,
    KEY `author` (`author`) USING HASH,
    KEY `dates` (`published_date`) USING BTREE,
    KEY `genre` (`genre`) USING HASH,
	PRIMARY KEY (`isbn`)
);
```
## Collections
Used to store the Collection resource using `name` as primary key. `creation_date` field has a BTREE index to ease the search within a range of dates.
```
CREATE TABLE `collections` (
	`name` VARCHAR(30) NOT NULL,
	`description` TEXT DEFAULT '',
	`creation_date` DATETIME DEFAULT '',
	KEY `creation_date` (`creation_date`) USING BTREE,
	PRIMARY KEY (`name`)
);
```
## Collection Members
Keeps track of the mapping between a collection and the books contained. It is implemented as a separate table since the relation is many to many. An index on `collection_name` is used to speed up the search of the books contained in a community.
```
CREATE TABLE `collection_members` (
	`collection_name` VARCHAR(30) NOT NULL,
	`book_isbn` BIGINT(13) NOT NULL,
	KEY `collection_name` (`collection_name`) USING HASH,
	PRIMARY KEY (`collection_name`,`book_isbn`)
);
```