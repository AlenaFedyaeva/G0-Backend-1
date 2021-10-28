# G0-Backend-1
g b cource

# Подготовка репозитория согласно "Лучшим практикам GO"

1) Линтеры
 - Устанавливаем golangci-lint
 - Настраиваем линтеры - пример .golangci.yaml 
 - запускаем golangci-lint run ./path_to_dir/

2) Настраиваем пре-коммит хуки
 - Устанавливаем pre-commit согласно https://pre-commit.com/#install
 - Настравиваем pre-commit: создаем файл  .pre-commit-config.yaml, 
 - Запускаем "pre-commit install" для установки локальных хуков.

3) Задание Github Actions
 - В своем Github-репозитории создайте папку .github/workflows.
 - для запуска golangci-lint  файл .github/workflows/lint.yml 
 - для запуска pre-commit .github/workflows/pre-commit.yml 

4) Makefile или Taskfile
 - вызов make или go-task





# Links: 
1) обработка ошибок API https://medium.com/nuances-of-programming/%D0%BE%D0%B1%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%BA%D0%B0-%D0%BE%D1%88%D0%B8%D0%B1%D0%BE%D0%BA-api-%D0%B2-%D0%B2%D0%B5%D0%B1-%D0%BF%D1%80%D0%B8%D0%BB%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D0%B8-%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D1%83%D1%8F-axios-932e9d66a526
2) Шпаргалка по работе с сетью в Go, емко и с примерами
http://tumregels.github.io/Network-Programming-with-Go/
 
3) Многопользовательская игра на Go через telnet
https://habr.com/ru/post/330512/

4) Документация к утилите make: https://www.gnu.org/software/make/ 
5) Пример работы с утилитой make для Go-разработчиков: https://blog.gopheracademy.com/advent-2017/make/ 
6) Статья “Эффективное использование make”: http://www.opennet.ru/docs/RUS/gnumake/ 
7) Готовим сборку Go-приложения в продакшн: https://habr.com/ru/post/337158/ 



