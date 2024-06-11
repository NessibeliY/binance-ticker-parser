# binance-ticker-parser
Данный проект парсер тикеров с биржи Бинанс. 
Для запроса цен используется эндпоинт GET /api/v3/ticker/price https://binance-docs.github.io/apidocs/spot/en/#symbol-price-ticker



1. Запустить программу можно таким путем
```bash
make build
make run
```
2. Или таким
```bash
go run ./cmd/parser/
```




Входные данные:
Конфиг файл в формате yaml
```yaml
symbols:
 - "LTCBTC"
 - "BTCUSDT"
 - "EOSUSDT"
# max_workers - ограничивает количество параллельно запущенных воркеров, которые делают запросы на биржу
max_workers: 2
```


Выходные данные:
Вот пример, что должна выводить программа:
```
EOSUSDT price:111
BTCUSDT price:222
LTCBTC price:444
EOSUSDT price:111
BTCUSDT price:222
workers requests total: 6
LTCBTC price:333 changed
EOSUSDT price:222 changed
BTCUSDT price:222
LTCBTC price:333
EOSUSDT price:111 changed
BTCUSDT price:222
LTCBTC price:333
EOSUSDT price:111
workers requests total: 15
BTCUSDT price:222
LTCBTC price:333
EOSUSDT price:111
BTCUSDT price:222
```