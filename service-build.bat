@echo off



rem 设置构建配置

rem SERVICE_IDL 服务IDL文件，编写的IDL文件请放在idl目录下
set SERVICE_IDL=user.proto

rem SERVICE_TYPE 服务类型，可选项有 RPC, HTTP
set SERVICE_TYPE=RPC

rem SERVICE_NAME 服务名称
set SERVICE_NAME=user

rem SERVICE_MODULE_NAME 服务模块名称
set SERVICE_MODULE_NAME=user

rem SERVICE_OUTPUT_DIR 服务生成目录
set SERVICE_OUTPUT_DIR=user



rem 打印构建开始信息
echo idl-build: build server for %SERVICE_NAME% with cwgo

rem 检查服务生成目录是否存在
if not exist "%SERVICE_OUTPUT_DIR%" (
    rem 如果目录不存在，则创建它
    mkdir "%SERVICE_OUTPUT_DIR%"
    echo idl-build: create dir %SERVICE_OUTPUT_DIR% for %SERVICE_NAME%
)

rem 切换到服务生成目录
cd /d "%SERVICE_OUTPUT_DIR%"

rem 执行 cwgo 命令启动服务
cwgo server --type %SERVICE_TYPE% --idl ../idl/%SERVICE_IDL% --service %SERVICE_NAME% --module %SERVICE_MODULE_NAME% -I ../idl

rem 将当前目录添加到 Go workspace
go work use .

rem 整理 Go 模块依赖
go mod tidy
