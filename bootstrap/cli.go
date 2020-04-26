package bootstrap

import "flag"

//var name string
//var age int
//var married bool
//var delay time.Duration
//flag.StringVar(&name, "env", "", "姓名")
//flag.IntVar(&age, "age", 18, "年龄")
//flag.BoolVar(&married, "married", false, "婚否")
//flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

type Args struct {
	Env string
}

func GetArgs() Args {
	args := Args{}
	flag.StringVar(&args.Env, "env", "", "环境")

	flag.Parse()
	return args
}