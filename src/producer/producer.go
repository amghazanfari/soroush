package producer

import (
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Envelop struct {
	Message   string `json:"message"`
	TopicName string `json:"topicName"`
}

func ProduceMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var envelop Envelop
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		json.Unmarshal([]byte(body), &envelop)

		err = os.MkdirAll("/tmp/messages/"+envelop.TopicName, os.ModePerm)
		if err != nil {
			panic(err)
		}
		f, err := os.OpenFile("/tmp/messages/"+envelop.TopicName+"/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		err = binary.Write(f, binary.LittleEndian, []byte(envelop.Message))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("message wrote"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("metod is not allowed"))
	}

}
