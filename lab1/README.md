# Лабораторная работа 1. Работа с HTTP протоколом.
## Задача
Передать http GET запрос напрямую из любого приложения (программного кода), но не из браузера. То есть передать «без мусора», только то что надо. <br>

Сервер: http://109.167.241.225/http_example/ <br>
Порт: 8001 <br>
GET запрос: give_me_five <br>
Параметры запроса: <br>
wday — день недели (вс =1) <br>
student — номер в таблице <br>

Заголовки: <br>
REQUEST_AGENT = ITMO student <br>
COURSE = Protocols <br>

В заголовках больше ничего не должно быть. <br>
Сервер присылает ответ — структура из числа и строки. <br>
Строка = «бинарное» представление ещё одного числа. <br>

700/TCP,UDP EPP (Extensible Provisioning Protocol)[23] — используется для управления регистрационной информацией DNS Официально <br>

Пример запроса без параметров (2 возврата строки в конце, поскольку 2 параметра): <br>
GET /http_example/give_me_five?student=30&wday=5 HTTP/1.1

Неправильный ответ на этот запрос: <br>
```HTTP/1.1 409 Conflict 
Connection: Keep-Alive 
Server: Embedthis-http 
Content-Type: text/html 
Cache-Control: no-cache 
Date: Ср, 08 фев 2023 20:14:51 GMT 
Content-Length: 0 
Keep-Alive: timeout=60, max=199 
```
Ответ на правильный запрос: <br>
```
HTTP/1.1 200 OK
Connection: Keep-Alive
Server: Embedthis-http
Content-Type: application/json; charset=utf-8
Cache-Control: no-cache
Date: Ср, 08 фев 2023 20:28:20 GMT
Content-Length: 109
Keep-Alive: timeout=60, max=199

{"workResult":{"value":1008,"strMessage":"0001000000000000000100000000000000010000000000000001000000000000"}}
```