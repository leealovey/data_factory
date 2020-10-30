package mongoop

import "testing"

// var config = MongoConfig{URI: "mongodb://mongoadmin:secret@172.18.81.6:27017"}
// var config = MongoConfig{URI: "mongodb://root:abc123456@192.168.214.20:31767"}
var config = MongoConfig{URI: "mongodb://192.168.214.199:27017"}

func TestConn(t *testing.T) {
	_, err := config.SendCollections()
	if err != nil {
		t.Errorf("error occured: %+v", err)
	}

}
