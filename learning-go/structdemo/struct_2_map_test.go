package structdemo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/fatih/structs"
)

type Server struct {
	Name        string `json:"name,omitempty" defaults:"test"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

func TestStruct2Map(t *testing.T) {

	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}
	s := structs.New(server)
	m := s.Map()
	fmt.Println(m)

	s1 := &Server{}
	val, _ := json.Marshal(s1)
	fmt.Println(string(val))
}
