请在config/app下新建private.yml配置文件,格式为:
```yaml
Postgresql: # Postgresql配置
  DriverName: postgresql # 驱动名
  SourceName: "postgres://root:123456@47.108.173.156:1122/Qianshou?sslmode=disable&pool_max_conns=10"
```
开发者本地测试步骤:
    1.确保go环境go1.17,docker版本20.10.17 
    2.运行docker命令初始化数据库docker run --name chat_postgres_zr --network chat_net -v chat_postgres_zr_data:/var/lib/postgresql/data -v 项目路径/config/postgres/my_postgres.conf:/etc/postgresql/postgresql.conf -p 5432:5432 -e ALLOW_IP_RANGE=0.0.0.0/0 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=chat -d chenxinaz/zhparser -c 'config_file=/etc/postgresql/postgresql.conf'
    3.数据库中运行src/dao/postgres/migration/000001_init_schema.up.sql初始化表
    4.go run main.go运行程序