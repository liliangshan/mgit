@echo off
setlocal enabledelayedexpansion

:: 设置版本号和应用名称
set VERSION=1.0.17
set APP_NAME=mgit
set BUILD_DIR=build
set RELEASE_DIR=releases

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

:: 创建必要的目录
if not exist %BUILD_DIR% mkdir %BUILD_DIR%
if not exist %RELEASE_DIR% mkdir %RELEASE_DIR%

echo [32m开始构建 %APP_NAME% v%VERSION%[0m

:: 备份i18n文件夹
rmdir /s /q i18nbak 2>nul
xcopy i18n i18nbak\ /E /I /Y
del /q /s i18nbak\messages.go 2>nul

:: 遍历平台进行编译和打包
for /l %%i in (0,1,3) do (
    set PLATFORM=!PLATFORMS[%%i]!
    set GOOS=!OS[%%i]!
    set GOARCH=!ARCH[%%i]!
    
    echo [32m正在编译 !PLATFORM!...[0m
    
    :: 设置环境变量
    set GOOS=!GOOS!
    set GOARCH=!GOARCH!
    set CGO_ENABLED=0
    
    :: 创建目标目录
    set TARGET_DIR=%BUILD_DIR%\%APP_NAME%-%VERSION%-!PLATFORM!
    if not exist !TARGET_DIR! mkdir !TARGET_DIR!
    
    :: 编译（添加 -trimpath 参数）
    if "!GOOS!"=="windows" (
        go build -trimpath -o !TARGET_DIR!\%APP_NAME%.exe
    ) else (
        go build -trimpath -o !TARGET_DIR!\%APP_NAME%
    )
    
    :: 复制必要文件
    xcopy i18nbak !TARGET_DIR!\i18n\ /E /I /Y
    copy /Y README.md !TARGET_DIR!\
    copy /Y LICENSE !TARGET_DIR!\
    echo !VERSION!> !TARGET_DIR!\version.txt
    
    :: 创建发布包
    if "!GOOS!"=="windows" (
        cd %BUILD_DIR%
        "C:\Program Files\7-Zip\7z" a -tzip ..\%RELEASE_DIR%\%APP_NAME%-%VERSION%-!PLATFORM!.zip %APP_NAME%-%VERSION%-!PLATFORM!\*
        cd ..
    ) else if "!GOOS!"=="linux" (
        cd %BUILD_DIR%
        "C:\Program Files\7-Zip\7z" a -tzip ..\%RELEASE_DIR%\%APP_NAME%-%VERSION%-!PLATFORM!.zip %APP_NAME%-%VERSION%-!PLATFORM!\*
        cd ..
    ) else if "!GOOS!"=="darwin" (
        cd %BUILD_DIR%
        "C:\Program Files\7-Zip\7z" a -tzip ..\%RELEASE_DIR%\%APP_NAME%-%VERSION%-!PLATFORM!.zip %APP_NAME%-%VERSION%-!PLATFORM!\*
        cd ..
    )
    
    if errorlevel 1 (
        echo [31m构建 !PLATFORM! 失败[0m
    ) else (
        echo [32m构建 !PLATFORM! 成功[0m
    )
)

:: 清理临时文件
rmdir /s /q i18nbak

echo [32m构建完成！[0m
echo [32m发布文件已保存在 %RELEASE_DIR% 目录下的各平台子目录中[0m


del /q /s %BUILD_DIR%\*
rd /s/q %BUILD_DIR%
endlocal