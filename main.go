package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/haron1996/fb/0/utils"
)

// // Mutex to prevent concurrent execution
// var mu sync.Mutex
// var isRunning bool

// // Function to run every 30 seconds if not already running
// func runEvery30Seconds() {
// 	// Load East Africa Time location
// 	location, err := time.LoadLocation("Africa/Nairobi")
// 	if err != nil {
// 		fmt.Println("Error loading location:", err)
// 		return
// 	}

// 	fmt.Println("Location:", location)

// 	// Get the current time in EAT
// 	currentTime := time.Now().In(location)
// 	hour := currentTime.Hour()

// 	// Check if the time is one of the specified hours
// 	if hour == 0 || hour == 7 || hour == 10 || hour == 12 || hour == 15 || hour == 19 {
// 		// Lock to check and update the running state
// 		mu.Lock()

// 		if isRunning {
// 			// If already running, release lock and skip execution
// 			fmt.Println("Function is already running, skipping this cycle:", currentTime)
// 			mu.Unlock()
// 			return
// 		}

// 		// Mark the function as running
// 		isRunning = true
// 		mu.Unlock()

// 		// Run the function
// 		fmt.Println("Function started at", currentTime)

// 		root := "/home/kwandapchumba/Pictures/PHONES/ALL"

// 		// login
// 		browser, page := utils.Login()
// 		if browser == nil || page == nil {
// 			fmt.Println("Browser or page is nil")
// 			return
// 		}

// 		// get phones
// 		phones, err := utils.GetPhones(root)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		totalPhones := len(phones)

// 		fmt.Println("Total Phones:", totalPhones)

// 		r := rand.New(rand.NewSource(time.Now().UnixNano()))

// 		// shuffle phones
// 		r.Shuffle(totalPhones, func(i, j int) {
// 			phones[i], phones[j] = phones[j], phones[i]
// 		})

// 		// list phones
// 		err = utils.ListPhones(browser, page, phones)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		fmt.Println("All Phones Have Been Listed")

// 		// Mark as not running after completion
// 		mu.Lock()
// 		isRunning = false
// 		mu.Unlock()
// 	} else {
// 		fmt.Println("Current time is outside the specified window:", currentTime)
// 	}
// }

func main() {

	// // Create a ticker that ticks every 30 seconds
	// ticker := time.NewTicker(30 * time.Second)
	// defer ticker.Stop()

	// // Run the ticker in a goroutine
	// go func() {
	// 	for range ticker.C {
	// 		runEvery30Seconds()
	// 	}
	// }()

	// // Block the main thread
	// select {}

	browser, page := utils.Login()
	if browser == nil || page == nil {
		fmt.Println("Browser or page is nil")
		return
	}

	root := "/home/kwandapchumba/Pictures/SIMU"

	items, err := utils.GetItems(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	utils.ListItems(browser, page, items)

	// for {

	// 	// shuffle phones
	// 	r.Shuffle(totalPhones, func(i, j int) {
	// 		phones[i], phones[j] = phones[j], phones[i]
	// 	})

	// 	// login
	// 	browser, page := utils.Login()
	// 	if browser == nil || page == nil {
	// 		fmt.Println("Browser or page is nil")
	// 		return
	// 	}

	// 	utils.ListPhones(browser, page, phones)

	// 	time.Sleep(30 * time.Second)
	// }
}
