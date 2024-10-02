@echo off
:: Запрашиваем имя коммита
set /p commitName="Введите имя коммита (обязательно): "

if "%commitName%"=="" (
    echo Ошибка: Имя коммита обязательно!
    exit /b 1
)

set /p commitComment="Введите комментарий (опционально): "

git add .

if "%commitComment%"=="" (
    git commit -m "%commitName%"
) else (
    git commit -m "%commitName%" -m "%commitComment%"
)

git push

echo Коммит и пуш завершены.
pause
