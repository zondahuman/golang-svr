package pb

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Printf("start...............")

	tesst := Test{Label: "hello", Type: 17, Reps: "what"}
	data, err := proto.Marshal(&tesst)
	fmt.Printf("data=", string(data))

}
