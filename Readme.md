## Compose sample application

### Use with Docker Development Environments

You can open this sample in the Dev Environments feature of Docker Desktop version 4.12 or later.

[Open in Docker Dev Environments <img src="../open_in_new.svg" alt="Open in Docker Dev Environments" align="top"/>](https://open.docker.com/dashboard/dev-envs?url=https://github.com/docker/awesome-compose/tree/master/django)

## Follow Steps On This Video

https://www.youtube.com/watch?v=OeJyjBoA-UA

## If Wan't To Run

cd app_golang
go get .
go run .

### If Can't Activate venv

Get-ExecutionPolicy
Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy Unrestricted -Force

### Django application in dev mode

Project structure:

```
.
├── compose.yaml
├── app
    ├── Dockerfile
    ├── requirements.txt
    └── manage.py

```

[_compose.yaml_](compose.yaml)

```
services:
  web:
    build: app
    ports:
      - '8000:8000'
```

## Deploy with docker compose

```
$ docker compose up -d
Creating network "django_default" with the default driver
Building web
Step 1/6 : FROM python:3.11-alpine
...
...
Status: Downloaded newer image for python:3.11-alpine
Creating django_web_1 ... done

```

## Expected result

Listing containers must show one container running and the port mapping as below:

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                    NAMES
3adaea94142d        django_web          "python3 manage.py r…"   About a minute ago   Up About a minute   0.0.0.0:8000->8000/tcp   django_web_1
```

After the application starts, navigate to `http://localhost:8000` in your web browser:

Stop and remove the containers

```
$ docker compose down
```

## PhpMyAdmin

running and the port mapping as below:

```
user: root
password: root

http://localhost:8080
```

After the application starts, navigate to `http://localhost:8080` in your web browser:

Stop and remove the containers

```
$ docker compose down
```

## Database Schema

### Table: users

| Field | Type         | Null | Key | Default | Extra          |
| ----- | ------------ | ---- | --- | ------- | -------------- |
| ID    | int(11)      | NO   | PRI |         | AUTO_INCREMENT |
| Fname | varchar(255) | YES  |     | NULL    |                |
| Lname | varchar(255) | YES  |     | NULL    |                |
| Email | varchar(255) | YES  | UNI | NULL    |                |
| Pass  | varchar(255) | YES  |     | NULL    |                |

### Table: workspace

| Field    | Type    | Null | Key | Default | Extra          |
| -------- | ------- | ---- | --- | ------- | -------------- |
| ID       | int(11) | NO   | PRI |         | AUTO_INCREMENT |
| Users_ID | int(11) | YES  | UNI | NULL    |                |

### Table: project

| Field        | Type         | Null | Key | Default | Extra |
| ------------ | ------------ | ---- | --- | ------- | ----- |
| ID           | int(11)      | NO   | PRI |         |       |
| Project_name | varchar(255) | YES  |     | NULL    |       |
| LDate_Time   | datetime     | YES  |     | NULL    |       |
| Project_Path | varchar(255) | YES  |     | NULL    |       |
| Screen_Img   | varchar(255) | YES  |     | NULL    |       |
| Workspace_ID | int(11)      | YES  | MUL | NULL    |       |
