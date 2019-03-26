# wooshyaApp
* website for international student job seeking

* designed by ST Wang and ZZ Zhang

### init setting:
* you have to install mysql and go get the package below:

```bash
$ go get github.com/go-sql-driver/mysql

$ go get golang.org/x/crypto/bcrypt
```
* then create the table in mysql:

```sql
CREATE TABLE stdusers(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120),
    mailaddress VARCHAR(50),
    collegename VARCHAR(50),
    degree VARCHAR(50),
    department VARCHAR(50),
    major VARCHAR(50),
    graduatedate(50),
    lastlogindate DATE
);

create table companyuser(
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120),
    mailaddress VARCHAR(50),
    companyname VARCHAR(50),
    telephonenumber VARCHAR(50),
    lastlogindate DATE
)
```
