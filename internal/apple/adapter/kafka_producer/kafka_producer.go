package kafka_producer

type KafkaProducer struct{}

func New() (*KafkaProducer, error) {
	return &KafkaProducer{}, nil
}

func (p *KafkaProducer) Close() {
}
