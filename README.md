Демон на [Golang](https://golang.org/) (с использованием фреймворка [Beego](https://github.com/astaxie/beego)), рассылающий уведомления в [Telegram](https://telegram.org/) о сообщениях [Vk Callback API](https://vk.com/dev/callback_api).
---
Использование
===
  * Запускать от пользователя имеющего возможность использовать ``sudo`` без пароля
  * В конфигурационном файле ``conf/app.conf`` настроить директорию и имя файла для SQLite базы данных и настроить id бота, token бота, id канала Telegram https://core.telegram.org/bots/api#authorizing-your-bot
  * Требует запущенного socks proxy на 127.0.0.1:9050, привет РКН
