# 使用方法
# sh model.sh test user
# 再将model目录拖到对应服务的目录里

#生成的表名
table=$2

#表生成的model目录
modeldir=./model

host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=123456

echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${table}" -dir="${modeldir}" --style=goZero -c