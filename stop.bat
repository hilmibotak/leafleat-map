@echo off
echo ========================================
echo   CTF CYBER MAP - STOP ALL SERVERS
echo ========================================
echo.

echo Stopping Backend Server...
taskkill /FI "WindowTitle eq CTF Backend*" /T /F >nul 2>nul

echo Stopping Frontend Server...
taskkill /FI "WindowTitle eq CTF Frontend*" /T /F >nul 2>nul

REM Stop any processes on port 8080 and 8000
echo Killing processes on port 8080 and 8000...
for /f "tokens=5" %%a in ('netstat -aon ^| findstr :8080 ^| findstr LISTENING') do taskkill /F /PID %%a >nul 2>nul
for /f "tokens=5" %%a in ('netstat -aon ^| findstr :8000 ^| findstr LISTENING') do taskkill /F /PID %%a >nul 2>nul

echo.
echo All servers stopped successfully!
echo.
pause
