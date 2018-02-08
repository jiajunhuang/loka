# Loka

Loka is a simple configuration class, set a default value in your subclass, and if there exist the
same configuration in environment variables, it will be overwrite.

```python
from loka import LokaConfig


class Config(LokaConfig):
    MYSQL_URI = "mysql://root@127.0.0.1:3306"
    LISTEN_PORT = 8080


config = Config()

# below is just for test
print(config.MYSQL_URI, type(config.MYSQL_URI))
print(config.LISTEN_PORT, type(config.LISTEN_PORT))
```

run it:

```bash
$ python test.py 
mysql://root@127.0.0.1:3306 <class 'str'>
8080 <class 'int'>
$ MYSQL_URI="mysql://user@127.0.0.1:9090" LISTEN_PORT=9090 python test.py 
mysql://user@127.0.0.1:9090 <class 'str'>
9090 <class 'int'>
```
