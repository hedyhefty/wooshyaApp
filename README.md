# wooshyaApp
* Website for international student career in japan.

* Designed by ST Wang and ZZ Zhang.

### init setting:
* You have to install mysql and go get the package below:

```bash
$ go get github.com/go-sql-driver/mysql

$ go get golang.org/x/crypto/bcrypt
```
* Then create the table in mysql:

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
    lastlogindate DATETIME
);
```

```sql
CREATE TABLE cpyusers(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120),
    mailaddress VARCHAR(50),
    companyname VARCHAR(50),
    category VARCHAR(50),
    description VARCHAR(500),
    telephonenumber VARCHAR(50),
    lastlogindate DATETIME
);
```
```sql
CREATE TABLE jobs(
    jid INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    cpyid INT,
    jtitle VARCHAR(50),
    jdescribe VARCHAR(500),
    jsalary VARCHAR(500),
    jlocation VARCHAR(50),
    jotherdetails VARCHAR(500),
    releasedate DATETIME,
    startdate DATETIME,
    deadline DATETIME
);
```
```sql
CREATE TABLE news(
    news_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    cpy_id INT,
    news_title VARCHAR(50),
    news_content VARCHAR(500),
    release_date DATETIME
);
```
```sql
CREATE TABLE application(
    jid INT,
    stdid INT,
    applydate DATETIME,
    primary key(jid,stdid)
);
```