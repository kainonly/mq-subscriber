# mqtt-trigger-sdk

MQTT Trigger PHP SDK for [mqtt-trigger](https://github.com/kainonly/mqtt-trigger)

### Initialization

Setup

```shell
composer require kain/mqtt-trigger-sdk
```

Usage

```php
$mqtt = new MQTT('http://localhost:3000');
$req = $mqtt->trigger('erp.order.create', 'L2-ccq123456');
```