# Приложение для идентификации сотрудников компании на КПП

## Краткое описание проекта

Приложение для идентификации сотрудников компании на КПП. Разрабатываемый программный продукт позволяет сотруднику СБ быстро и качественно организовывать мероприятия по обеспечению безопасности на объектах.

## Краткое описание предметной области

Предметной областью является сфера безопасности (в частности, организация безопасности каких-либо объектов)

## Краткий анализ аналогичных решений

Решение                                      | Поиск сотрудников по различным параметрам | Хранение документа, удостоверяющего личность | Возможность саморегистрации сотрудником | 
-------------------------------------------- |-------------------------------------------|----------------------------------------------|-----------------------------------------
[ИНФОСТАРТ](https://infostart.ru)         | -                                         | -                                            | +                                       |
[PASS24.online](https://pass24online.ru)   | -                                         | -                                            | +                                       |
[Малленом Системс](https://www.mallenom.ru) | -                                         | +                                            | -                                       |
Реализуемое решение                          | +                                         | +                                            | +                                       |

## Краткое обоснование целесообразности и актуальности проекта

В современном мире такая система является очень актуальной и востребованной, так как всегда необходимо поддерживать безопасность в абсолютно разных областях и местах.

## Краткое описание акторов (ролей)

В проекте присутствуют следующие роли:

- Гость - неавторизованный сотрудник компании
- Сотрудник компании - пользователь, зарегистрированный в сервисе
- Сотрудник СБ - сотрудник компании, обеспечивающий безопасность в рамках компании

## Use-Case

![Use-Case диаграмма](diagrams/use-case.svg)

## ER-диаграмма

![ER-диаграмма](diagrams/ER.svg)

## Диаграмма БД

![Диаграмма БД](diagrams/db.jpg)

## Сценарии использования

- Регистрация сотрудника в приложении:
    1) Сотрудник переходит на сайт
    2) Сотрудник открывает форму для регистрации и регистрируется
    3) Пользователь перенаправляется на страницу со своей информационной карточкой и заполняет ее
- Поиск сотрудника по номеру телефона:
    1) Сотрудник СБ переходит на сайт
    2) Сотрудник СБ переходит на страницу входа и авторизовывается
    3) Сотрудник СБ перенаправляется на страницу со списком всех сотрудников
    4) Сотрудник СБ вводит в поисковую строку, расположенную на той же странице, номер телефона человека
    5) Сотрудник СБ нажимает на кнопку "Поиск" для осуществления поиска
    6) На экране отображается информация о найденном сотруднике или информация о том, что сотрудник не найден
- Получение списка всех сотрудников, отсортированных в порядке убывания по ФИО:
    1) Сотрудник СБ переходит на сайт
    2) Сотрудник СБ переходит на страницу входа и авторизовывается
    3) Сотрудник СБ перенаправляется на страницу со списком всех сотрудников
    4) Сотрудник СБ выбирает поле для сортировки "ФИО"
    5) Сотрудник СБ выбирает порядок сортировки "по убыванию"
    6) На экране отображается список сотрудников

## Формализация ключевых бизнес-процессов

![Процесс входа для работы с информационной карточкой](diagrams/BPMN1.svg)
![Процесс входа для поиска информации о пользователе](diagrams/BPMN2.svg)
![Процесс входа для поиска информации о нескольких сотрудниках](diagrams/BPMN3.svg)

## Разбиение на компоненты

![Верхнеуровневое разбиение на компоненты](diagrams/updated_components.svg)

## Экраны будущего приложения

![demo1](demo/1.png)
![demo2](demo/2.png)
![demo3](demo/3.png)
![demo4](demo/4.png)
![demo5](demo/5.png)
![demo6](demo/6.png)
![demo7](demo/7.png)
![demo8](demo/8.png)

## ApacheBenchmark для одного хоста

```shell
Server Software:        
Server Hostname:        localhost
Server Port:            8081

Document Path:          /api/v2/profile
Document Length:        141 bytes

Concurrency Level:      100
Time taken for tests:   2.678 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      2650000 bytes
HTML transferred:       1410000 bytes
Requests per second:    3734.73 [#/sec] (mean)
Time per request:       26.776 [ms] (mean)
Time per request:       0.268 [ms] (mean, across all concurrent requests)
Transfer rate:          966.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       2
Processing:    15   26   1.9     26      40
Waiting:       15   26   1.9     26      40
Total:         17   26   1.9     26      42

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     27
  75%     27
  80%     28
  90%     29
  95%     30
  98%     32
  99%     32
 100%     42 (longest request)
```

## ApacheBenchmark для трех хостов с балансировкой

```shell
Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        nginx
Server Hostname:        127.0.0.1
Server Port:            80

Document Path:          /api/v2/profile
Document Length:        141 bytes

Concurrency Level:      100
Time taken for tests:   0.256 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      3690001 bytes
HTML transferred:       1410000 bytes
Requests per second:    39133.74 [#/sec] (mean)
Time per request:       2.555 [ms] (mean)
Time per request:       0.026 [ms] (mean, across all concurrent requests)
Transfer rate:          14101.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       3
Processing:     1    2   0.6      2      13
Waiting:        1    2   0.6      2      10
Total:          2    2   0.8      2      13

Percentage of the requests served within a certain time (ms)
  50%      2
  66%      2
  75%      2
  80%      2
  90%      2
  95%      4
  98%      6
  99%      6
 100%     13 (longest request)
```