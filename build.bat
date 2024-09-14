@echo off
REM 构建 Windows 的 .exe 文件
echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o turing-resume-windows.exe main.go

REM 构建 Linux 的二进制文件
echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o turing-resume-linux main.go

@REM REM 构建 macOS 的二进制文件
@REM echo Building for macOS...
@REM set GOOS=darwin
@REM set GOARCH=amd64
@REM go build -o turing-resume-macos main.go

echo Build complete!