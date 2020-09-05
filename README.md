# Project Layout
API is designed with Clean Architecture.  

![Clean Architecture](https://miro.medium.com/max/772/1*B7LkQDyDqLN3rRSrNYkETA.jpeg) 

# Directory
Directory structre of api follows Standard Go Project Layout.  
https://github.com/golang-standards/project-layout

```
├── api - OpenAPI/Swagger specs, JSON schema files, protocol definition files.
├── build - Packaging and Continuous Integration.
├── cmd - Main applications for this project.
├── configs - Configuration files.
├── deployments - Confiuration files for deployment
├── internal - Private application and library code.
│   ├── domain - Domain layer
│   ├── infrastracture - Infrastracture layer
│   ├── interface - Interface layer
│   ├── registry - DI container
│   └── usecase - Usecase files
├── pkg - Library code that's ok to use by external applications.
└── test - Testing files. 
```

This is an example of implementation of Clean Architecture in Go (Golang) projects.
https://github.com/bxcodec/go-clean-arch

Rule of Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :
 * Models Layer
 * Repository Layer
 * Usecase Layer
 * Delivery Layer

 # Run
 cd tools && ./start_api.sh


 # Golang Reference:
 https://www.bookstack.cn/read/gin-en/html-rendering
 https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/