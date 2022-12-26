package main

import (
	"fmt"
	iiko_transport "iikoapi/cmd/app"
	dishes "iikoapi/internal/dishes"
)

func main() {
	transport, err := iiko_transport.NewIikoTransport("3e6eea5e-185")
	if err != nil {
		return
	}

	ar := []string{"3b5ba63d-88f2-4174-9c76-1bbc1c2d8e59"}
	menu, err := transport.GetDishesByMenuId(ar, 5704)
	if err != nil {
		fmt.Println(err)
	}

	var d []dishes.Dish

	for _, v := range menu.ItemCategories {
		for _, dish := range v.Items {
			d = append(d, dish)
		}
	}

	for _, m := range d {
		fmt.Println(m.Name)
		fmt.Println(m.ItemSizes[0].Prices[0].Price)
	}
}
