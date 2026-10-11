package main

import "goProgramming/algorithms"

func main() {
	// otai := goBasics.Person{Name: "Otai", Age: 21}
	// otai.Greet()
	// otai.Eat("pizza")

	tiktok :=&algorithms.Node{Data: "ticktok.com"}

	josemilies :=&algorithms.Node{Data: "josemiles.com"}

	cache := &algorithms.Cache{Capacity: 4}
	cache.InitializeCache()
	cache.PrintLinkedList()
	println("\n After Adding tiktok ")
	cache.AddCacheEntry(tiktok)
	cache.PrintLinkedList()

	println("\n After Adding josemilies ")
	cache.AddCacheEntry(josemilies)
	cache.PrintLinkedList()

	println("\n Trying to add a tiktok Duplicate")
	cache.AddCacheEntry(tiktok)
	cache.PrintLinkedList()

	
}
