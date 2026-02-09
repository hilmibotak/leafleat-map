@echo off
echo ========================================
echo   CTF CYBER MAP - SETUP & START
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Go is not installed or not in PATH!
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo [1/4] Setting up Backend...
cd backend
if not exist go.sum (
    echo Installing Go dependencies...
    go mod tidy
)

echo.
echo [2/4] Starting Backend Server...
start "CTF Backend" cmd /k "go run main.go"
timeout /t 3 >nul

cd ..

echo.
echo [3/4] Starting Frontend Server...
echo.
echo Choose frontend server option:
echo   1. Simple HTTP Server (Python)
echo   2. PingBox (if installed)
echo   3. Live Server (VS Code extension)
echo.
set /p choice="Enter choice (1-3): "

if "%choice%"=="1" (
    where python >nul 2>nul
    if %errorlevel% equ 0 (
        echo Starting Python HTTP Server on port 8000...
        start "CTF Frontend" cmd /k "python -m http.server 8000"
        set FRONTEND_URL=http://localhost:8000
    ) else (
        echo [ERROR] Python not found!
        pause
        exit /b 1
    )
) else if "%choice%"=="2" (
    where pingbox >nul 2>nul
    if %errorlevel% equ 0 (
        echo Starting PingBox Server...
        start "CTF Frontend" cmd /k "pingbox serve -p 8000 -d ."
        set FRONTEND_URL=http://localhost:8000
    ) else (
        echo [ERROR] PingBox not found! Please install from:
        echo https://github.com/noz0/pingbox
        pause
        exit /b 1
    )
) else if "%choice%"=="3" (
    echo Please start Live Server from VS Code
    echo Right-click on index.html and select "Open with Live Server"
    set FRONTEND_URL=http://127.0.0.1:5500
    pause
) else (
    echo Invalid choice!
    pause
    exit /b 1
)

timeout /t 3 >nul

echo.
echo [4/4] Setup Complete!
echo ========================================
echo.
echo   Backend API: http://localhost:8080
echo   Frontend:    %FRONTEND_URL%
echo.
echo ========================================
echo.
echo Opening browser...
timeout /t 2 >nul

REM Open browser
start "" "%FRONTEND_URL%"

echo.
echo Press any key to stop all servers...
pause >nul

REM Kill servers
taskkill /FI "WindowTitle eq CTF Backend*" /T /F >nul 2>nul
taskkill /FI "WindowTitle eq CTF Frontend*" /T /F >nul 2>nul

echo.
echo All servers stopped.
pause
