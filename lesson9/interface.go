package main 
import "fmt"

type Os struct{
	Name string
}

type Linux struct{
	Os
}
type Windows struct{
	Os
}
type MacOs struct{
	Os
}

type OsInterface interface{
	GetName() string
}

func (o Os) GetName()string{
	return o.Name
} 
//Overadding

func (o MacOs) GetName()string{
	return "MacBook" + o.Os.GetName()
}

func Run(os OsInterface){
	fmt.Println("Loading "+ os.GetName()+"...")
}

func main(){
	Run(Linux{Os{Name: "Ubuntu 22.04"}})
	Run(MacOs{Os{Name: "Pro M1"}})
	Run(Windows{Os{Name: "Windows 11 Pro"}})
}