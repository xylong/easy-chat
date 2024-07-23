goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c -style go_zero

goctl model mysql ddl -src="./deploy/sql/social.sql" -dir="./apps/social/socialModels/" -c