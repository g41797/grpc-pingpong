<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Pb;

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
     * @param \Pb\CreateStationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CreateStation(\Pb\CreateStationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.AdapterService/CreateStation',
        $argument,
        ['\Pb\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Pb\DestroyStationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function DestroyStation(\Pb\DestroyStationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/pb.AdapterService/DestroyStation',
        $argument,
        ['\Pb\Status', 'decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function Produce($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/pb.AdapterService/Produce',
        ['\Pb\Status','decode'],
        $metadata, $options);
    }

    /**
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\BidiStreamingCall
     */
    public function Consume($metadata = [], $options = []) {
        return $this->_bidiRequest('/pb.AdapterService/Consume',
        ['\Pb\ConsumeResponse','decode'],
        $metadata, $options);
    }

}
