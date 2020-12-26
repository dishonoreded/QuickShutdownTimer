package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const layout ="2006-01-02 15:04"
func main(){
	args:=os.Args
	if len(args)<2{
		fmt.Printf("请输入必要的参数\n-a:在n个小时后关机(支持小数点后1位)\n-t 在指定时间（格式：%s）关机\n-c:取消关机",layout)
		return
	}
	t:=args[1]
	var execArg []string
	var d int64
	switch t {
	case "-a":
		if args[2]==""{
			fmt.Println("输入值无效(exp:1.5)")
			return
		}
		f,err:=strconv.ParseFloat(args[2],32)
		if err!=nil{
			fmt.Printf("-a的值输入无效,err:%v",err)
			return
		}
		d=int64(f*time.Hour.Seconds())
		execArg=[]string{"-s","-t",strconv.Itoa(int(d))}
	case "-t":
		if args[2]==""||args[3]==""{
			fmt.Printf("值无效(exp:%s)\n",layout)
			return
		}
		loc, _ := time.LoadLocation("Local")
		inputTime, err := time.ParseInLocation(layout, args[2]+" "+args[3],loc )
		if err!=nil{
			fmt.Printf("-t的值输入无效,err:%v",err)
			return
		}
		n:=time.Now().Unix()
		if n>inputTime.Unix(){
			err=fmt.Errorf("这里没有时光机(输入的关机时间小于当前时间)")
			fmt.Printf("-t的值输入无效,err:%v",err)
			return
		}
		d=inputTime.Unix()-n
		execArg=[]string{"-s","-t",strconv.Itoa(int(d))}
	case "-c":
		execArg=[]string{"-a"}
	default:
		fmt.Printf("请输入正确的时间类型：-a/-t\n -a:在n个小时后关机(支持小数点后1位)\n -t 在指定时间（格式：%s）关机)\n",layout)
		return
	}
	c:=exec.Command("shutdown",execArg...)
	b,err:=c.CombinedOutput()
	if err!=nil{
		if err.Error()==fmt.Sprint("exit status 1116"){
			fmt.Print("你没有需要取消的关机任务")
			return
		}
		log.Fatal(err)
		return
	}
	fmt.Printf("%s",b)

}