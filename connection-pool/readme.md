# Connection pool

- Simple connection pool, in golang for mysql db connect.
- This is a simple prototype to understand concept of connection pool and how they work.
- I'll be attaching the steps to execute this codebase soon.
- I should also use docker to kind do things. 
- I am also thinking if we can use Makefile to automate the setup process hmm let's see.



## docker command you may want execute to setup mysql container.

```sh
$ docker run --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<your-root-password> -d mysql

$ mysql -u root -ppassword -h 127.0.0.1 -P 3306
```
