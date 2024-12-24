# Connection pool

- Simple connection pool, in golang for mysql db connect.
- This is a simple prototype to understand concept of connection pool and how they work.
- I'll be attaching the steps to execute this codebase soon.
- I am also thinking if we can use Makefile to automate the setup process hmm let's see.



## docker command you may want execute to setup mysql container.

```sh
$ docker run --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<your-root-password> -d mysql

$ mysql -u root -ppassword -h 127.0.0.1 -P 3306

$ create database <database-name> # recordings
```


## run go program with env variable

```sh
$ DBUSER=<username> DBPASS=<password> go run main.go 
```


## learning go channel

* Go channel -> puts the event on channel and <- pulls the event from channel.
* While adding back the connection we were pulling from channel which resulted in we in-defintely waiting for channle to free which will never happen since we are not releasing the resource.
* So Put back the connection should send an event to channel by doing `pool.channel <- struct{}{}` which means one place is free.

