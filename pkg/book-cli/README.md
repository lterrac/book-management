# CLI
This section includes the cli commands used to interact with the book management software.

## Usage
Every command has the following syntax:

```
book-cli <COMMAND> <TYPE> [OPTIONS] [ -f FILE-PATH | OBJECT]
```

## Commands
The commands available are:
- `create`:    create a new instance of either a book or a collection
- `get`:       retrieve object instance
- `update`:   update an object instance
- `delete`:    delete an object instance

## General flags
Up to this moment the only flag that can be used with every command is `host` which allows to specify the book-server host

# Create command
Create adds a new resource in the database. Up to this moment, only JSON representation of the resource is allowed. The object definition can be written both directly on the command line or supplying the path where the object definition is stored (using the `-f` flag).

## flags
- `-f, --file`: specify the file path containing the object definition

## examples
- Create a new Book:
```
book-cli create book '{"title": "Romeo and Juliet", "author": "William Shakespeare" ...}'
```
- Create a new Book defined in file `book.json`:  
```
book-cli create book -f book.json
```

# Get command
Get command is used to retrieve a resource. The default command schema is:
```
book-cli get <RESOURCE_TYPE> <RESOURCE_NAME>
```
where `<RESOURCE_TYPE>` can be found in [types section](../apis/README.md#input-values) and `<RESOURCE_NAME>` is the identifier of the resource which is the `"isbn"` field for `books` and `"name"` field for `collections`.  
It is possibile to specify some filters and combine them together to retrieve a subset of objects.
In particular, for `book` resource the following filters are available:
- `--title`: the title of the book
- `--author`: the book author
- `--genre`: the book genre
- `--dates`: a range of pubblication dates written using the following format `"start_date:end_date"` where dates are "MM-DD-YYYY"
- `--all`: retrieves all resources  
Instead, `collections` resource has the following filters:
- `--dates`: a range of creation dates written using the following format `"start_date:end_date"` where dates are "MM-DD-YYYY"
- `--all`: retrieves all resources  

## examples
- Get a book using its unique isbn:
```
book-cli get book 9780671722852
```
- Get all books with the same title:
```
book-cli get book --title "Romeo and Juliet"
```
- Get a set of books with the same author:
```
book-cli get book --author "William Shakespeare"
```
- Get the books pubblished in 1996:
```
book-cli get book --dates "1996-01-01-to-1996-31-12"
```
- Get all books:
```
book-cli get book --all
```

# Update command
As for `create` command, the update of a resource can be done supplying a file path or the new resource definition directly using the command line.

## flags
- `-f, --file`: specify the file path containing the object definition

## examples
- Update a Book:
```
book-cli update book '{"title": "Romeo and Juliet", "author": "William Shakespeare" ...}'
```
- Create a new Book defined in file `book.json`:  
```
book-cli update book -f book.json
```
# Delete command
The `delete` command allows to delete a single or a subset of books. All the filtering options of [get](#get-command) command can be reused for this command.

## examples
- Delete a book using its unique isbn:
```
book-cli delete book 9780671722852
```
- Delete all book with the same title:
```
book-cli delete book --title "Romeo and Juliet"
```
- Delete a set of books with the same author:
```
book-cli delete book --author "William Shakespeare"
```
- Delete the books pubblished in 1996:
```
book-cli delete book --dates "1996-01-01-to-1996-31-12"
```
- Delete all books:
```
book-cli delete book --all
```