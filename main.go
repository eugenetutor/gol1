package main

import (
	"fmt"
	"time"
)

func main() {
	// if змінна := вираз; умова{
	// 	 істина
	// }

	if today := time.Now().Weekday(); today == time.Saturday || today == time.Sunday {
		fmt.Println("Це вихідний")
	} else if today == time.Friday{
		fmt.Println("Пʼятниця! Скоро вихідні")
	}else{
		fmt.Println("Робочий день :((", today)
	}
}
