

# FMI API ASSESSMENT

### Technology Stack:
| Teknologi   | Version | Link |
| ----------- | ---------------- | ------------------- |
| Golang      | v1.19 or later   | [Go Download](https://go.dev/dl)  |
| Go Echo Framework     | v3.3.10     | [Echo Installation](https://echo.labstack.com/guide/#installation) | 
| MySql | v5.7 or later |  |
| Docker | v24.0.6 or later |  |
| Docker compose| v12.21.0 or later |  |
<br>

### To Do
    install docker dan docker-compose
    Install postman
    Install git
    clone repo [https://github.com/dirgadm/fmi-assessment]

### Running Server
| Command   | Description | Link |
| ----------- | ---------------- | ------------------- |
| docker compose up -d      | Running compose yaml file in background side, and then do the migration to the mysql. the file path is in .database.sql. Also, it build a few **Store Procedure**|
| docker compose down      | Terminate all docker image that running in background|
| go run main.go     | Running server in 9090 port. This configuration is provided from /.config.json | 
<br>

## Endpoint Testing 
    - available in `./doc/FMI.postman_collection.json` and ready to import to postman
    - base_url: http://localhost:9090/v1

## Technical Documentation
- There are file upload.md, attendance.md and user.md in folder .doc
- There are consist of Technical specification for each endpoint in this API, and also can be used as API contract between Backend & FrontEnd

## Fungsional Requirement
1. [POST] Registration(email, password, name, phone)user object
2. [POST] Login(email, password string)user object
3. [POST] Upload Photo(userId int, file string) file string
4. [POST] Attendance(userId int, longitude, latitude float64) message string




