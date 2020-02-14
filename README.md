# How to connect oracle database with go

## Dependicies

Make sure that you got libclntsh installed

```
    sudo apt-get install libaio1 libaio-dev
```

## how to run

```
go get ./...
go run main.go
```

## sql

```sql
    CREATE TABLE salgrade
    (
        grade NUMBER,
        losal NUMBER,
        hisal NUMBER
    )

```
```sql
INSERT INTO SALGRADE VALUES (1,700,1200);
INSERT INTO SALGRADE VALUES (2,1201,1400);
INSERT INTO SALGRADE VALUES (3,1401,2000);
INSERT INTO SALGRADE VALUES (4,2001,3000);
INSERT INTO SALGRADE VALUES (5,3001,9999);
```


**That should result in the table bello 👍 👍**

| grade | losal | hisal |
| ----- | ----- | ----- |
| 1     | 700   | 1200  |
| 2     | 1201  | 1400  |
| 3     | 1401  | 2000  |
| 4     | 2001  | 3000  |
| 6     | 3001  | 9999  |
