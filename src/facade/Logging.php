<?php

namespace think\logging\facade;

use think\logging\FactoryLogging;
use think\Facade;

/**
 * Class Logging
 * @method static void push($namespace, array $raws = []) 信息收集推送
 * @package think\logging
 */
final class Logging extends Facade
{
    protected static function getFacadeClass()
    {
        return FactoryLogging::class;
    }
}
