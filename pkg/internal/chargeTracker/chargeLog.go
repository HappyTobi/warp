package chargeTracker

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"

	"github.com/HappyTobi/warp/pkg/internal/users"
	"github.com/HappyTobi/warp/pkg/internal/util"
	"github.com/HappyTobi/warp/pkg/internal/warp"
)

func NewChargeLog(request *warp.Request) *ChargeLog {
	return &ChargeLog{request: request}
}

func (cl *ChargeLog) Load(users []*users.User) (*Charges, error) {

	data, err := cl.request.Get()
	if err != nil {
		return nil, err
	}

	return deserialize(data, users)
}

func deserialize(data []byte, users []*users.User) (*Charges, error) {
	userMapping := make(map[int]string, len(users))

	for user := range users {
		userMapping[users[user].Id] = users[user].Username
	}

	charges := &Charges{
		Charges: make([]*Charge, 0),
	}

	reader := bytes.NewReader(data)
	buf := make([]byte, 16)
	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return charges, err
			}
			break
		}

		timestampMin := binary.LittleEndian.Uint32(buf[0:])
		meterStart := binary.LittleEndian.Uint32(buf[4:])
		userId := int(buf[8])
		chargeDurationSec := binary.LittleEndian.Uint32(buf[9:]) & 0x00FFFFFF
		meterEnd := binary.LittleEndian.Uint32(buf[12:])

		charge := &Charge{}

		charge.User = mapUserNameToId(userId, userMapping) //fmt.Sprintf("%d", userId)
		charge.Time = util.TimestampMinutesToDate(timestampMin)
		charge.Duration = util.ChargeDuration(chargeDurationSec)
		charge.PowerMeterStart = math.Float32frombits(meterStart)
		charge.PowerMeterEnd = math.Float32frombits(meterEnd)

		charges.Charges = append(charges.Charges, charge)

		if len(string(buf)) == 0 {
			break
		}
	}

	return charges, nil
}

func mapUserNameToId(userId int, userMapping map[int]string) string {
	username := userMapping[userId]
	if len(username) == 0 {
		return "unknown user"
	}

	return username
}
