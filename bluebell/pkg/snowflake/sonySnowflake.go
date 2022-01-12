package snowflake

/*
@author RandySun
@create 2022-01-12-8:16
*/
import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineId uint16
)

func getMachineId() (uint16, error) {
	return sonyMachineId, nil

}

// 需传入当前的机器ID
func Init1(startTime string, machineId uint16) (err error) {
	sonyMachineId = machineId
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineId,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

//GenId 生成id
func GenId1() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return id, err

}
func main() {
	if err := Init1("2022-01-12", 1); err != nil {
		fmt.Printf("init failed err:%#v\n", err)
	}

	id, _ := GenId1()
	fmt.Println(id)

}
