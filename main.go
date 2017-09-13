package main

import (
	"runtime"
	"port-forward/web"
	"github.com/VividCortex/godaemon"
)

func main() {
	//app := cli.NewApp()
	//app.Name = "greet"
	//app.Usage = "fight the loneliness!"
	//app.Action = func(c *cli.Context) error {
	//	fmt.Println("Hello friend!")
	//	time.Sleep(time.Second*10)
	//	return nil
	//}
	//app.Run(os.Args)

	godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	runtime.GOMAXPROCS(runtime.NumCPU())

	web.InitConfig().Run(":5555")
}
