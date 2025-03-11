@echo off



rem 设置构建配置

rem DB_TYPE 数据库类型，可选项有 mysql, sqlserver, sqlite, postgres
set DB_TYPE=mysql

rem DSN 数据库连接串
set DSN=gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local

rem SQL_DIR SQL文件目录
set SQL_DIR=gorm.sql

rem OUTPUT_DIR 模型生成目录
set OUTPUT_DIR=service



rem 打印构建开始信息
echo "model-build: build model for %OUTPUT_DIR% with cwgo"

rem 检查模型生成目录是否存在
if not exist "%OUTPUT_DIR%" (
    mkdir "%OUTPUT_DIR%"
    echo "model-build: create dir %OUTPUT_DIR%"
)

rem 切换到模型生成目录
cd /d "%OUTPUT_DIR%"

rem 执行 cwgo 命令生成模型
cwgo  model --db_type %DB_TYPE% --dsn "%DSN%" --sql_dir "../sql/%SQL_DIR%"

rem 整理 Go 模块依赖
go mod tidy
