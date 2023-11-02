1, build image: docker build .
2, create pod
3, sudo docker login --username=752623077@qq.com registry.cn-hangzhou.aliyuncs.com
4, docker tag fe8103da32 registry.cn-hangzhou.aliyuncs.com/connor886/httpserver:1.0.0
5, docker push registry.cn-hangzhou.aliyuncs.com/connor886/httpserver:1.0.0
docker pull  registry.cn-hangzhou.aliyuncs.com/connor886/httpserver:1.0.0
docker rm $(docker ps -a | grep registry.cn-hangzhou.aliyuncs.com/connor886/httpserver:1.0.0 | awk '{print $1}' )

https://raw.githubusercontent.com/anjia0532/gcr.io_mirror/master/pull-k8s-image.sh
