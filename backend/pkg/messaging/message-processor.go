package consumerinterface

type MessageProcessor interface {
    Process(body []byte) error
    GetQueueName() string
}
