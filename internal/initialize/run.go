package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()
	r.Run()
}
