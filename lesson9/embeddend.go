package main
import "fmt"

type Component struct{
	weigth, suger float64
}

type Fruit struct{
	Component
	name string
}

type Vagiatable struct{
	Component
	name string
	color struct{
		R,G,B uint8
	}
}

func (c Component) getSuger()string{
	return fmt.Sprintf("Suger: %fgr",c.suger)
}

type sweet struct{
	name string
	Sugar float64
	weight float64
}

func (s sweet) getSuger()string{
	return fmt.Sprintf("Suger: %fgr",s.Sugar)
}
func (s sweet) getWeight()string{
	return fmt.Sprintf("Weigth: %f",s.weight)
}
func main(){

	var apple = Fruit{
		Component: Component{100,19},
		name: "apple",
	}

	var tomato = Vagiatable{
		Component: Component{200,0.24},
		name: "tomato",
		
	}
	fmt.Println(apple)
	fmt.Println(tomato)


	keks := sweet{
		name: "keks",
		Sugar: 30,
		weight: 70,
	}
	napalion := sweet{
		name: "nampalion",
		Sugar: 40,
		weight: 2100,
	}
	fmt.Println(keks)
	fmt.Println(keks.getSuger())
	fmt.Println(napalion)
	fmt.Println( napalion.getWeight())

}