package leetcode

/*
 * Least Recently Used (LRU) cache
 * 
 * GET(key) - get value of key if key exists in cache
 * SET(key, value) - set or insert value if key is not already present
 * 
 * When cache reaches capacity, it should invalidate least recently used
 * item before inserting new item.
 */
 
 
 /*PLAN
  * Cache:
  * 	internal data structure: map[int](int,int)
  * 		key = i suppose this represents the location in memory that the
  * 			user is querying for.
  * 		value =  (the value the user is looking for, num times accessed
  * 	get: 
  * 	set:
  * 
  */ 


type Cache struct {
	

}
