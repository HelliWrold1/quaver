# 根据thrift生成脚手架代码
kitex_gen_user:
	kitex --thrift-plugin validator -module github.com/HelliWrold1/quaver idl/user.thrift
kitex_gen_comment:
	kitex --thrift-plugin validator -module github.com/HelliWrold1/quaver idl/comment.thrift