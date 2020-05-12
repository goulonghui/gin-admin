/*
Package app 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger */
package app

// @title gin-admin
// @version 6.2.0
// @description Basic development framework based on gin + gorm/mongo + wire.
// @schemes http https
// @basePath /
// @contact.name LyricTian
// @contact.email tiannianshou@gmail.com
