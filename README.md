# clean-harbor

清理harbor仓库中的镜像tag。

通过对镜像的tag按照构建时间的先后顺序排列，只保留最新的keepNum个镜像，其余的全部删除。

支持一键清除所有。

需要harbor的定时清理任务配合。

# 1. 二进制执行

## 1.1 打包二进制
```shell
go build 
#或者
make build
```


## 1.2 查看帮助

```shell
./harbor-clean -h
harbor-clean 用于清理harbor的仓库中的tag，以释放存储资源

Usage:
  harbor-clean [flags]

Examples:
harbor-clean --url https://harbor.abc.com  --user admin --password FJLSDfdso3489X --projectName abc-web --keepNum 200

Flags:
  -h, --help                 help for harbor-clean
  -k, --keepNum int          保留的tag数目，例如50
  -p, --password string      密码
  -P, --projectName string   项目名，all表示所有项目
  -u, --url string           例如：https://harbor.abc.com
  -U, --user string          用户名，例如：admin
```
## 1.3 手动执行清理
Linux服务器上操作
```shell
./harbor-clean --url ** --user ** --password ** --projectName ** --keepNum **
```

## 1.4 crontab定时任务执行

```shell
# crontab -l
0 2 * 7 * /root/clean-harbor -url https://harbor.abc.com --user ** --password ** --projectName * --keepNum 100 >> /var/log/harbor-clean`date "+%Y-%m-%d-%H:%M:%S"`.log
```


# 2. docker执行


## 2.1 制作镜像
```shell
# 做镜像,比如
docker build . -t harbor.abc.com/devops/harbor-clean:v1
```
## 2.2 docker run形式执行
```
# 执行清理
docker run harbor.abc.com/devops/harbor-clean:v1 /data/harbor-clean --url ** --user ** --password ** --projectName ** --keepNum **
```
## 2.3 k8s中CronJob形式执行

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: harbor-clean
spec:
  schedule: "*/1 * * * *"
  successfulJobsHistoryLimit: 1
  failedJobsHistoryLimit: 1
  jobTemplate:
    spec:
      template:
        spec:
          restartPolicy: Never
          #imagePullSecrets:
          #- name: regsecret
          containers:
          - name: harbor-clean
            image: harbor.abc.com/devops/harbor-clean:v1
            args:
            - "/bin/sh"
            - "-c"
            - "./harbor-clean --url ** --user ** --password ** --projectName ** --keepNum **"
```


# 3. 输出内容
正常输入如下类似格式：
```shell
当前tag: 50  ，保留tag: 100  of xmc2-lexue/dispatcher-service            ,无需删除!
当前tag: 19  ，保留tag: 100  of xmc2-lexue/engine-audio-process          ,无需删除!
当前tag: 38  ，保留tag: 100  of xmc2-lexue/engine-image-process          ,无需删除!
当前tag: 38  ，保留tag: 100  of xmc2-lexue/engine-metric-service         ,无需删除!
当前tag: 26  ，保留tag: 100  of xmc2-lexue/engine-pipeline-manager       ,无需删除!
当前tag: 26  ，保留tag: 100  of xmc2-lexue/engine-video-extract          ,无需删除!
当前tag: 37  ，保留tag: 100  of xmc2-lexue/media-access                  ,无需删除!
当前tag: 22  ，保留tag: 100  of xmc2-lexue/media-gateway                 ,无需删除!
当前tag: 327 ，保留tag: 100 of xmc2-chongwen/meta-service                ,开始执行删除
     删除image: xmc2-chongwen/meta-service:dev-514-d002e80ede215776037549fb5ee8760d112cfc24, 创建时间为: 2020-02-05 15:05:21.501707432 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-513-46ffd7014d1be9e3ad7f29c287d7f3eaa0f7770b, 创建时间为: 2020-02-05 14:09:23.348353517 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-512-84d201964e7a54b1ed56ab77aa65e4e88b6c9e14, 创建时间为: 2020-02-05 10:27:31.599784047 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-511-147e4a0a1d9b896577c76f666d7218250aa7e32d, 创建时间为: 2020-01-19 09:39:35.06098945 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-509-c44f5b30b6f9535a7dab84e824d68d863542037a, 创建时间为: 2020-01-18 12:57:42.526076891 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-508-c44f5b30b6f9535a7dab84e824d68d863542037a, 创建时间为: 2020-01-18 12:56:14.892997573 +0000 UTC
     删除image: xmc2-chongwen/meta-service:tag-v1.0.6, 创建时间为: 2020-01-17 15:20:33.013336278 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-506-479c12916a973fd2924c94a92b2f0833271e7684, 创建时间为: 2020-01-17 11:48:25.162196274 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-502-a770bd1ffea663b7e1fdce0452c7daacab46883f, 创建时间为: 2020-01-17 01:52:53.846231006 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-498-a770bd1ffea663b7e1fdce0452c7daacab46883f, 创建时间为: 2020-01-17 01:52:53.846231006 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-500-a770bd1ffea663b7e1fdce0452c7daacab46883f, 创建时间为: 2020-01-17 01:52:53.846231006 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-504-a770bd1ffea663b7e1fdce0452c7daacab46883f, 创建时间为: 2020-01-16 15:23:42.689613615 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-496-a770bd1ffea663b7e1fdce0452c7daacab46883f, 创建时间为: 2020-01-16 15:23:15.677063574 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-494-0251343e9ffb12e04cf3b7f7c074374c8539f064, 创建时间为: 2020-01-16 12:02:37.801550463 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-492-5f182821c22a16913a99e992d197cde44d8aba8b, 创建时间为: 2020-01-16 10:42:30.47080662 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-491-5f182821c22a16913a99e992d197cde44d8aba8b, 创建时间为: 2020-01-16 09:13:48.846308245 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-489-feabbc1c7c123147f69a39dbe35e837ec4c2b878, 创建时间为: 2020-01-16 08:47:24.127053689 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-487-2481da139453309268390f3a5319190c8063d928, 创建时间为: 2020-01-16 08:39:44.163760885 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-484-69244f205bff6312ddd9284fbd41b7023067f5b0, 创建时间为: 2020-01-16 08:17:27.456364882 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-483-69244f205bff6312ddd9284fbd41b7023067f5b0, 创建时间为: 2020-01-16 06:34:20.788798053 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-479-79aac507c57394e4eb2cc4accd0dfbe252f1c7e1, 创建时间为: 2020-01-16 02:15:55.670666982 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-478-79aac507c57394e4eb2cc4accd0dfbe252f1c7e1, 创建时间为: 2020-01-15 15:54:32.24282212 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-481-79aac507c57394e4eb2cc4accd0dfbe252f1c7e1, 创建时间为: 2020-01-15 15:54:32.24282212 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-476-5bd399297f990a82eb05b5ae5cf305efea166616, 创建时间为: 2020-01-15 10:13:59.413985129 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-472-5bd399297f990a82eb05b5ae5cf305efea166616, 创建时间为: 2020-01-15 10:00:06.497089626 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-474-5bd399297f990a82eb05b5ae5cf305efea166616, 创建时间为: 2020-01-15 10:00:06.497089626 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-470-257e1a7ec741f7515b0963d7e5abc50b90f4e467, 创建时间为: 2020-01-15 08:43:36.897120442 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-469-3b9c477d46c670be96142e1ab3dcfb9e69c50e8c, 创建时间为: 2020-01-15 07:53:05.481295186 +0000 UTC
     删除image: xmc2-chongwen/meta-service:dev-467-3b9c477d46c670be96142e1ab3dcfb9e69c50e8c, 创建时间为: 2020-01-15 07:52:39.637509898 +0000 UTC
```




