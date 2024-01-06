<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Memphisphp\Pb;

/**
 */
class AdapterServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Memphisphp\Pb\CreateStationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CreateStation(\Memphisphp\Pb\CreateStationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.AdapterService/CreateStation',
        $argument,
        ['\Memphisphp\Pb\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Memphisphp\Pb\DestroyStationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function DestroyStation(\Memphisphp\Pb\DestroyStationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.AdapterService/DestroyStation',
        $argument,
        ['\Memphisphp\Pb\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function Produce($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/pb.AdapterService/Produce',
        ['\Memphisphp\Pb\Status','decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function Consume($metadata = [], $options = []) {
        return $this->_bidiRequest('/pb.AdapterService/Consume',
        ['\Memphisphp\Pb\ConsumeResponse','decode'],
        $metadata, $options);
    }

}
