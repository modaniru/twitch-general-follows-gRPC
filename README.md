
# Twitch general follows
A project **rewritten** in gRPC to train this protocol

## Run Locally

Clone the project

~~~bash
  git clone https://github.com/modaniru/twitch-general-follows-gRPC
~~~

Go to the project directory

~~~bash
  cd twitch-general-follows-gRPC
~~~

Create .env file

~~~bash
  touch .env
~~~

Write secrets in .env ([more](https://github.com/modaniru/twitch-general-follows-gRPC#environment-variables))

~~~bash
  TWITCH_CLIENT_ID=your twitch client id
  TWITCH_CLIENT_SECRET=your twitch client secret
~~~

If you can run "make" commands

~~~bash
  make
~~~

Else: \
Install dependencies

~~~bash
go mod download
~~~

Start the server

~~~bash
go run src/main.go
~~~

the server will run on **8080** port\
You can change port in *configuration/config.yaml* file

## Docker
run from **Docker Hub**
~~~bash
docker run -p 8080:8080 -e TWITCH_CLIENT_ID=clientId -e TWITCH_CLIENT_SECRET=clientSecert modaniru/tgf
~~~
or
~~~bash
docker run -p 8080:8080 --env-file path modaniru/tgf
~~~
**build** and run docker container
~~~bash
docker build --name imageName
docker run --env-file path -p 8080:8080 imageName
~~~

## Environment variables

~~~bash
  TWITCH_CLIENT_ID=id // your twitch apps client id
  TWITCH_CLIENT_SECRET=secret // your twitch apps client secret
  PORT=80 // application running port (optional, default: 8080)
~~~

## Tasks
- [x] Remove yaml configuration
- [ ] SOLID (refactoring)
- [ ] CI/CD
- [ ] Tests