# habit-tracking

## 二进制启动
```bash
go run main.go
go run main.go -data-path "./data/real"
./bin/habit-tracking  -data-path "./data/real"
```


## 容器启动

```bash
# 编译
make bin

# 构建镜像
docker build -t habit-tracking:test . --progress=plain --no-cache

# 启动容器
docker run  --name habit-tracking  -it -p 8866:8888 -v /root/habit-tracking/data/real:/root/data/real habit-tracking:test
```
