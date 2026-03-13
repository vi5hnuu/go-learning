package main

import (
	"fmt"
)

const GlobalLimit=100;
const MaxCacheSize int =10*GlobalLimit;
const (
	CacheKeyBook = "book_"
	CacheKeyCd = "cd_"
)

var cache map[string]string;
func main() {
	fmt.Println(CacheKeyBook)
	fmt.Println(CacheKeyCd)
	cache = make(map[string]string);
	setBook("1234-5678","Get ready to go");
	setCD("1234-5678","Get ready to go audio book");
	fmt.Println("Book : ",GetBook("1234-5678"));
	fmt.Println("CD : ",getCD("1234-5678"));
}

func cacheGet(key string)	string{
	return cache[key];
}

func cacheSet(key string, value string) {
	if len(cache)+1>=MaxCacheSize {
		return;
	}
	cache[key] = value;
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn);
}

func setBook(isbn string, name string) {
	cacheSet(CacheKeyBook + isbn, name);
}

func getCD(sku string) string {
	return cacheGet(CacheKeyCd + sku);
}

func setCD(sku string, title string) {
	cacheSet(CacheKeyCd + sku, title);
}
