package main

import (
	"flag"
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"os"
	_ "time"
)

var (
	uri   = flag.String("uri", "tcp://192.158.30.209:1883", "MQTT URI")
	cid   = flag.String("cid", "ian-client", "ClientID")
	qos   = flag.Int("qos", 0, "QoS value [0|1|2]")
	topic = flag.String("topic", "go-mqtt/sample", "Topic on broker")
	mode  = flag.String("mode", "pub", "Mode of operation, [pub|sub]")
	msg   = flag.String("msg", "This is a test", "Message for [pub] mode")
	count = flag.Int("count", 500, "Number of messages for [pub] mode")
)

var end_of_test string = "END-OF-TEST"

func init() {
	flag.Parse()
}

//define a function for the default message handler
var default_msg_handler MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[DEFAULT] Received msg [%s] on topic [%s]\n", msg.Payload(), msg.Topic())
}

func test_pub() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output
	opts := MQTT.NewClientOptions().AddBroker(*uri)
	opts.SetCleanSession(false)
	opts.SetDefaultPublishHandler(default_msg_handler)
	opts.SetClientID(*cid)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//Publish messages to topic at qos and wait for the receipt
	//from the server after sending each message
	for i := 0; i < *count; i++ {
		text := fmt.Sprintf("%s #%d", *msg, i)
		token := c.Publish(*topic, byte(*qos), false, text)
		token.Wait()
		fmt.Printf("Published msg qos[%d] topic[%s] payload[%s]\n",
			*qos, *topic, text)
	}
	// Publish the end-of-test message
	token := c.Publish(*topic, byte(*qos), false, end_of_test)
	token.Wait()

	c.Disconnect(250)
}

func test_sub() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker(*uri)
	opts.SetCleanSession(false)
	opts.SetDefaultPublishHandler(default_msg_handler)
	opts.SetClientID(*cid)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var test_finished chan int
	test_finished = make(chan int, 1)
	var msg_handler MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received msg topic[%s] qos[%d] dup[%t] payload[%s]\n",
			msg.Topic(),
			msg.Qos(),
			msg.Duplicate(),
			msg.Payload(),
		)
		if string(msg.Payload()) == end_of_test {
			test_finished <- 1
		}
	}

	//subscribe to the topic and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe(*topic, byte(*qos), msg_handler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Wait for end-of-test
	<-test_finished

	//unsubscribe from topic
	if token := c.Unsubscribe(*topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

func main() {
	if *mode == "pub" {
		fmt.Printf("Testing publish\n")
		test_pub()
	}
	if *mode == "sub" {
		fmt.Printf("Testing subscribe\n")
		test_sub()
	}
}
