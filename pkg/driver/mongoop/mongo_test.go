package mongoop

import "testing"

var config = MongoConfig{URI: "mongodb://mongoadmin:secret@172.18.81.6:27017"}

func TestConn(t *testing.T) {
	_, err := config.Conn()
	if err != nil {
		t.Errorf("error occured: %+v", err)
	}

}
