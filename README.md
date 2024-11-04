# Worker Pool
Программа на Go, реализующая примитивный worker-pool с возможностью динамически добавлять и удалять воркеры для обработки входных данных из канала. 

## Описание 
Данный worker pool позволяет обрабатывать строки, поступающие в канал, с использованием воркеров. Воркеры могут быть добавлены и удалены динамически. Задачи распределяются между активными воркерами, которые выводят на экран номер воркера и обрабатываемую задачу. 

## Функциональные возможности
- Добавление воркеров: Возможность динамического добавления воркеров. 
- Удаление воркеров: Остановка работы воркеров.
- Добавление задач: Отправка задач (строк) в канал для обработки воркерами.

## Установка 
1. Убедитесь, что Go установлен. Загрузите и установите Go с [официального сайта](https://golang.org/dl/). 
2. Склонируйте репозиторий или скопируйте исходный код в файл main.go.


# P.S
Код покрыт тестами. Они лежат в папке main_test.go