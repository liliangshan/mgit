@echo off
setlocal enabledelayedexpansion

:: 设置版本号
set VERSION=1.0.16

:: 设置输出目录
set OUTPUT_DIR=build

:: 创建输出目录
if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

:: 定义目标平台
set PLATFORMS[0]=windows-amd64
set PLATFORMS[1]=linux-amd64
set PLATFORMS[2]=darwin-amd64
set PLATFORMS[3]=darwin-arm64

:: 定义对应的 GOOS 和 GOARCH
set OS[0]=windows
set OS[1]=linux
set OS[2]=darwin
set OS[3]=darwin

set ARCH[0]=amd64
set ARCH[1]=amd64
set ARCH[2]=amd64
set ARCH[3]=arm64

:: 清理旧文件
echo 清理旧文件...
del /q /s %OUTPUT_DIR%\*
rmdir /s /q i18nbak
xcopy i18n i18nbak\ /E /I /Y
del /q /s i18nbak\messages.go
:: 遍历平台进行编译
for /l %%i in (0,1,3) do (
    set PLATFORM=!PLATFORMS[%%i]!
    set GOOS=!OS[%%i]!
    set GOARCH=!ARCH[%%i]!
    
    echo.
    echo 正在编译 !PLATFORM!...
    
    :: 设置环境变量
    set GOOS=!GOOS!
    set GOARCH=!GOARCH!
    set CGO_ENABLED=0
    
    :: 创建输出子目录
    if not exist %OUTPUT_DIR%\!PLATFORM! mkdir %OUTPUT_DIR%\!PLATFORM!
    
    :: 编译（添加 -trimpath 参数以移除本地路径信息）
    if "!GOOS!"=="windows" (
        go build -trimpath -o %OUTPUT_DIR%\!PLATFORM!\mgit.exe
    ) else (
        go build -trimpath -o %OUTPUT_DIR%\!PLATFORM!\mgit
    )
    
    :: 复制 README.md
    copy README.md %OUTPUT_DIR%\!PLATFORM!\
    ::复制i18n文件夹
    xcopy i18nbak %OUTPUT_DIR%\!PLATFORM!\i18n\ /E /I /Y
    :: 创建版本文件
    echo !VERSION!> %OUTPUT_DIR%\!PLATFORM!\version.txt
    
    if errorlevel 1 (
        echo 编译 !PLATFORM! 失败
    ) else (
        echo 编译 !PLATFORM! 成功
    )
)
rmdir /s /q i18nbak
echo.
echo 编译完成！输出目录：%OUTPUT_DIR%
echo 版本号：%VERSION%

:: 列出生成的文件
echo.
echo 生成的文件：
dir /s /b %OUTPUT_DIR%

endlocal 