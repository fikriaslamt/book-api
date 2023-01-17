host : https://book-api-production.up.railway.app/

fitur : 
- login(POST) : /api/login
  use json format : 
  {
    "username": "admin",
    "password": "admin"
  }
  
- logout     : /api/logout

- getBooks    : /api/book
- getBookByID : /api/book/:id
  use header with key id

crud with login auth

- addBook(POST) : /api/book/add
  use json format :
  {
    "title" : "tes3",
    "author" : "tes3",
    "category" : "tes3"
  }
  
- updateBook(PUT) : /api/book/update/:id
   use json format,header with key id :
   
  {
    "title" : "tes3",
    "author" : "tes3",
    "category" : "tes3"
  }
  
- deletBook(DELETE) : /api/book/delete/:id
  header with key id

  
 
