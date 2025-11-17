@echo off
cd /d "C:\Gogin"

echo ============================
echo Starting Auto Push Script...
echo ============================

:: ============================
:: Generate commit message with date (ddmmyy)
:: ============================
for /f "tokens=2 delims==." %%a in ('wmic os get localdatetime /value') do set datetime=%%a
set year=%datetime:~2,2%
set month=%datetime:~4,2%
set day=%datetime:~6,2%
set commitmsg=%day%%month%%year%

echo Generated commit message: "%commitmsg%"
echo ============================

:: ============================
:: Stage and Commit
:: ============================
git add .
git commit -m "%commitmsg%"
echo Complete commit: "%commitmsg%"
echo ============================

:: ============================
:: Push to branch
:: ============================
set branch=main
echo Pushing to branch "%branch%" ...
git push origin %branch%
echo Complete push to branch "%branch%"
echo ============================

echo Push Completed at %time%
echo ============================
pause
