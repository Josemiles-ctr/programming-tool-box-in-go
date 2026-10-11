package main

import "goProgramming/algorithms"

func main() {
	// otai := goBasics.Person{Name: "Otai", Age: 21}
	// otai.Greet()
	// otai.Eat("pizza")

	tiktok :=&algorithms.Node{Data: "ticktok.com"}

	cache := &algorithms.Cache{Capacity: 4}
	cache.InitializeCache()
	cache.PrintLinkedList()
	cache.AddCacheEntry(tiktok)
	cache.PrintLinkedList()
}
