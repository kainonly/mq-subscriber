# mqtt-api-trigger-sdk

MQTT API Trigger PHP SDK for [mqtt-api-trigger](https://github.com/kainonly/mqtt-api-trigger)

### Initialization

Setup

```shell
composer require kain/mqtt-api-trigger-sdk
```

Usage

```php
$mqtt = new MQTT('http://localhost:3000');
$req = $mqtt->trigger('erp.order.create', 'L2-ccq123456');
```
