<?php

declare (strict_types=1);

namespace think\logging\service;

use think\logging\FactoryLogging;
use think\Service;

final class LoggingService extends Service
{
    public function register()
    {
        $this->app->bind('logging', function () {
            $appid = $this->app->config
                ->get('app.app_id');
            $config = $this->app->config
                ->get('queue.logging');

            return new FactoryLogging(
                $appid,
                $config['exchange'],
                !empty($config['router_key']) ? $config['router_key'] : '',
                !empty($config['virualhost']) ? $config['virualhost'] : '/'
            );
        });
    }
}