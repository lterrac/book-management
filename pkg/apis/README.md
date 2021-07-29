# Rest APIS

## Introduction
The book management software offers a set of RESTful APIs to communicate with external clients.

## Input values
The management system accepts data in JSON format written in the request body.
The software allows to manage two kind of objects:
- `book`:
```
{
    "title": string,
    "author": string,
    "isbn": string,
    "published_date": string (format: "MM-DD-YYYY"),
    "edition": int,
    "description": string,
    "genre": string,
}
```
- `collection`
```
{
    "name": string,
    "description": string,
    "creation_date": string (format: "MM-DD-YYYY"),
    "books": []string,
}
```

## Return values
There are two types of standard return types:
- Standard return value
- Error

### Standard return value
For standard operation the following JSON is returned:
```
{
    "status": "success",
    "code": 200,
    "metadata": {} //used for specific action information
}
```

### Error
If something goes wrong the following JSON is returned:
```
{
    "status": "error",
    "code": 400,
    "metadata": {} //more details about the error
}
```

## Filtering
GET queries supports filtering operations to retrieve a specific subset of resources. Filtering is currently implemented for both `books` and `collections`.
The filtering query is structured as follows:
```
?filter=FIELD-OPERATION-VALUE
```
The following logical operators are supported:
- `equals(eq)`
- `not equals(ne)` (implemented only server side)
- `and(and)`: used to concatenate more filters
  
For both `books` and `collections` if the identifier field is specified, all the other filters will be ignored.