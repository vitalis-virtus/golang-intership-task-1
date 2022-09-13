Задача 1:
Создать приложение в котором TCP клиент будет отправлять команду на сервер, сервер обрабатывать команду и отдавать ответ.  Команда может быть любая.
Вспомогательная информацию для разработки:
-  https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/
   Дополнительная функциональность:
- сделать gracefull shutdown (в отдельной ветке)

run server: go run tcpS.go 1234


run client: go run tcpC.go 127.0.0.1:1234
