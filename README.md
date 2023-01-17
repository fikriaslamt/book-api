host : https://book-api-production.up.railway.app/

fitur : 
- login(POST) : /api/login
  use json format : 
  ```json
    {
      "username": "admin",
      "password": "admin"
    }
  ```

  
- logout     : /api/logout

- getBooks    : /api/book
- getBookByID : /api/book/:id
  use header with key id

crud with login auth

- addBook(POST) : /api/book/add
  use json format :
  ```json
    {
      "title" : "",
      "author" : "",
      "category" : ""
    }
  ```

  
- updateBook(PUT) : /api/book/update/:id
   use json format,header with key id :
  ```json
    {
      "title" : "",
      "author" : "",
      "category" : ""
    }
  ```
  
- deletBook(DELETE) : /api/book/delete/:id
  header with key id

  
 
