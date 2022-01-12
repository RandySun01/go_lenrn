package snowflake

/*
@author RandySun
@create 2022-01-12-8:16
*/
import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	// 毫秒值
	snowflake.Epoch = st.UnixNano() / 1000000
	// 机器id
	node, err = snowflake.NewNode(machineId)
	return
}
func GenId() int64 {
	return node.Generate().Int64()

}

//func main() {
//	if err := Init("2022-01-12", 1); err != nil {
//		fmt.Printf("init failed err:%#v\n", err)
//	}
//
//	id := GenId()
//	fmt.Println(id)
//
//}
