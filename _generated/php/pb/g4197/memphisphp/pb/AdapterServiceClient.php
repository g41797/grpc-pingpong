<?php
// GENERATED CODE -- DO NOT EDIT!

namespace g4197\memphisphp\pb;

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
     * @param \g4197\memphisphp\pb\CreateStationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CreateStation(\g4197\memphisphp\pb\CreateStationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.AdapterService/CreateStation',
        $argument,
        ['\g4197\memphisphp\pb\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \g4197\memphisphp\pb\DestroyStationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function DestroyStation(\g4197\memphisphp\pb\DestroyStationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.AdapterService/DestroyStation',
        $argument,
        ['\g4197\memphisphp\pb\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function Produce($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/pb.AdapterService/Produce',
        ['\g4197\memphisphp\pb\Status','decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function Consume($metadata = [], $options = []) {
        return $this->_bidiRequest('/pb.AdapterService/Consume',
        ['\g4197\memphisphp\pb\ConsumeResponse','decode'],
        $metadata, $options);
    }

}
