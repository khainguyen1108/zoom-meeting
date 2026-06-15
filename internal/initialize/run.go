package initialize

import (
	"ROOM-MEETING/global"
	StringUtils "ROOM-MEETING/pkg/utils"
)

func Run() {
	r := InitRouter()
	r.Run(StringUtils.GetServerPort(global.Config.Server.Port))
}
