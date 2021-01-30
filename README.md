<h3 align = "center"> Dockerized Flutter(web) + Golang(Gin) + PostgreSQL(with GORM) + Redis --> Template  </h1>

#

<br>
<p align = "center">
  <i>
    This repository is a fully dockerized application, which by default is a <b>web</b> app, but modyfing Dockerfiles lets to run it on <b> mobile </b> device too. 
    Front-end is written in <b> Flutter </b>, Back-end in <b> Golang with Gin </b> framework. Databases are <b> PostgreSQL (with GORM) </b> and <b> Redis ( Redigo )</b>. Reverse proxy with <b> Nginx </b>. It contains simple CRUD functions for both, Postgresql and Redis. 
  </i>
</p>

#

<br/>

## How to launch it locally?

<br/>

<b> 1. </b> Make sure that you have `docker` and `docker-compose` installed on your machine.

<br/>

<b> 2. </b> Clone the repository to the folder of your choice, by executing: `git clone https://github.com/wzslr321/fggp.git` in the console.

<br/>

<b> 3. </b> Navigate to  `"/"` path of cloned repository. To ensure that you are in a right place, run `ls`. Outpout should look like this:

> LICENSE  README.md  all_lint_rules.yaml  analysis_options.yaml  app  data  docker-compose.yml  env  nginx  server


#### 4.  Run : <b> `docker-compose up --build` </b> in the console.

<br/> 

<b> 5. </b> Wait for the docker installation to complete. It probably is going to take some time, depending on your internet connection.

<br/> 

When the installation will be finished, you should see this output:
![img](https://user-images.githubusercontent.com/59893892/104651401-e1aac080-56b7-11eb-8c85-a449ee4cb6c0.png)

<br/>

#### Now just head to the `http://0.0.0.0/#/` in the browser too see the website!
---

<p align="center">
  Found it useful? Want more updates?
</p>

<p align = "center">
  <b> <i> Show your support by giving a :star: </b> </i>
</p>

---

### Project tree
```bash
-- LICENSE
|-- README.md
|-- all_lint_rules.yaml
|-- analysis_options.yaml
|-- app 
|   |-- Dockerfile
|   |-- README.md
|   |-- android (...) --> build on android
|   |-- assets
|   |   |-- fonts
|   |   `-- images
|   |-- integration_test(...)
|   |-- ios(...) --> build on ios
|   |-- lib
|   |   |-- constants.dart
|   |   |-- main.dart
|   |   |-- models
|   |   |   `-- http_exception.dart
|   |   |-- providers
|   |   |   `-- tasks_provider.dart
|   |   `-- screens
|   |       |-- home_screen
|   |       |   `-- home_screen.dart
|   |       |-- loading_screen.dart
|   |       |-- page_not_found_screen
|   |       |   `-- page_not_found.dart
|   |       `-- tasks_screen
|   |           `-- tasks_screen.dart
|   |-- pubspec.lock
|   |-- pubspec.yaml
|   |-- test
|   |   `-- widget_test.dart
|   `-- web
|       |-- favicon.png
|       |-- icons
|       |-- index.html
|       `-- manifest.json
|-- docker-compose.yml
|-- env
|   |-- flutter.env
|   `-- postgre.env
|-- nginx
|   |-- Dockerfile
|   `-- nginx.conf
`-- server
    |-- Dockerfile
    |-- air.conf
    |-- conf
    |   `-- app.ini
    |-- database
    |   |-- postgres
    |   |   |-- funcs
    |   |   |   `-- posts.go
    |   |   `-- postgres.go
    |   `-- redis
    |       |-- funcs
    |       |   |-- announcements.go
    |       |   `-- majors.go
    |       `-- redis.go
    |-- go.mod
    |-- go.sum
    |-- main.go
    |-- main_test.go
    |-- models
    |   |-- announcement.go
    |   `-- post.go
    |-- routers
    |   |-- api
    |   |   |-- announcements.go
    |   |   `-- posts.go
    |   `-- router.go
    |-- settings
    |   `-- settings.go
```

