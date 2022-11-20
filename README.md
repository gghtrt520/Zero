# Go-zero project

# Dockerfile
# goctl docker xxx.go

# Deployment
# goctl kube deploy -replicas 3 -requestCpu 200 -name user-rpc -namespace sase -image user/rpc:v1 -o user-rpc.yml -port 8888 --serviceAccount sase-sa
# goctl kube deploy -nodePort 32009 -replicas 3 -requestCpu 200 -name user-api -namespace sase -image user/api:v1 -o user-api.yml -port 8889 --serviceAccount sase-sa