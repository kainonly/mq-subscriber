<?php

namespace mqtt\trigger;


use GuzzleHttp\Client;

final class MQTT
{
    private $httpClient;

    /**
     * MQTT constructor.
     * @param $base_uri
     */
    public function __construct($base_uri)
    {
        $this->httpClient = new Client([
            'base_uri' => $base_uri,
            'timeout' => 2.0
        ]);
    }

    /**
     * trigger function
     * @param $topic
     * @param $message
     * @param array $options
     * @return array|mixed
     */
    public function trigger($topic, $message, $options = [])
    {
        try {
            $response = $this->httpClient->post('/', [
                'json' => [
                    'topic' => $topic,
                    'message' => $message,
                    'options' => (object)$options,
                ]
            ]);

            return json_decode($response->getBody()->getContents(), true);
        } catch (\Exception $e) {
            return [
                'error' => 1,
                'msg' => $e->getMessage()
            ];
        }
    }
}