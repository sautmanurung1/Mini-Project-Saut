# Mini Project Make API Discussion Forum Like Google Classroom

Tools :
- Golang
- MySQL
- GORM (Golang ORM)
- Echo (Framework/library)
- JWT (JSON Web Token)

Structure Project 
```
├── app
│   └──main.go
├── domains
│   ── Answare
│      ├── entities.go
│      ├── handler.go
│      └── repository.go
│   ── Assignment
│      ├── entities.go
│      ├── handler.go
│      └── repository.go
│   ── Auth
│      ├── entities.go
│      ├── handler.go
│      └── repository.go
│   ── Discussions
│      ├── entities.go
│      ├── handler.go
│      └── repository.go
│   ── Question
│      ├── entities.go
│      ├── handler.go
│      └── repository.go
│   ── Role
│      ├── entities.go
│      ├── handler.go
│      └── repository.go
├── infrastructure
│   ── Database
│      └── database.go
│   ── Http
│      ├── Middleware
│      ├── student_jwt.go
│      └── teacher_jwt.go
│   ── Server
│      └── echo.go
├── infrastructure
│   ── Answare
│      ├── handler.go
│      └── routes.go
│   ── Assignment
│      ├── handler.go
│      └── routes.go
│   ── Auth
│      ├── handler.go
│      └── routes.go
│   ── Discussions
│      ├── handler.go
│      └── routes.go
│   ── Question
│      ├── handler.go
│      └── routes.go
│   ── Role
│      ├── handler.go
│      └── routes.go
├── repository
│   ── Answare
│      └── answare.go
│   ── Assignment
│      └── assignment.go
│   ── Auth
│      └── auth.go
│   ── Discussions
│      └── discussions.go
│   ── Question
│      └── question.go
│   ── Role
│      └── role.go
```