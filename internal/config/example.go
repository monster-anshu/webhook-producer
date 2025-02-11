package config

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func CreateExample() {

	writer := GetKafkaWriter()
	defer writer.Close()

	const LIMIT = 10

	var messagesArray [LIMIT]kafka.Message
	nextDomain := roundRobin()

	for i := 0; i < LIMIT; i++ {
		testMessage := fmt.Sprintf(`{
            "hm": "POST",
            "url": "%s",
            "hdr": {
                "authorization": "auth",
                "content-type": "application/json"
            },
            "u_id": 4663,
            "s_wid": 69,
            "w_type": 2,
            "cid_num": "7505064723",
            "call_num": "7505064723",
            "re": 0,
            "ac": "67aaf83f245de",
            "c_id": "1739257917.28918",
            "ch": "clicktocall",
            "sip_d": "10.104.80.167",
            "ct": "2025-02-11T07:11:57Z",
            "pyld": {
                "uuid": "67aaf83f245de",
                "call_to_number": "918580993919",
                "caller_id_number": "+919240251731",
                "start_stamp": "2025-02-11 12:41:57",
                "answer_agent_number": "+911000001039",
                "call_id": "1739257917.28918",
                "billing_circle": {
                    "operator": "BSNL/MTNL - Central Gov.",
                    "circle": "Bihar & Jharkhand"
                },
                "call_status": "Answered by agent",
                "direction": "click_to_call",
                "customer_no_with_prefix ": "918580993919",
                "async": 1,
                "get_call_id": 1,
                "custom_identifier": "3dd5c548-3df5-4ea4-a1f4-f44a275958dd"
            },
            "partition": 0
            }`,
			nextDomain(),
		)
		messagesArray[i] = kafka.Message{
			Key:   []byte("Test"),
			Value: []byte(testMessage),
			Topic: os.Getenv("KAFKA_CONSUMER_TOPIC"),
		}
	}

	writer.WriteMessages(context.Background(), messagesArray[:]...)

}
