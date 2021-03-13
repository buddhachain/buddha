package rc

import (
	"fmt"

	"github.com/buddhachain/buddha/apiserver/factory/config"
	"github.com/rongcloud/server-sdk-go/v3/sdk"
)

func RCUserRegister(id, name, cid string) (string, error) {
	rc := sdk.NewRongCloud(config.RCConfig().AppKey, config.RCConfig().AppSecret)
	portraitURI := fmt.Sprintf("%s/%s", config.IpfsExplorer(), cid)
	user, err := rc.UserRegister(id, name, portraitURI)
	if err != nil {
		return "", err
	}
	return user.Token, nil
}
