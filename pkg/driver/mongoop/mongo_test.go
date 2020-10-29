package mongoop

import "testing"

var config = MongoConfig{URI: "mongodb://root:abc123456@192.168.214.20:31767/test?authSource=admin"}

func TestConn(t *testing.T) {
	_, err := config.Conn()
	if err != nil {
		t.Errorf("error occured: %+v", err)
	}

}
