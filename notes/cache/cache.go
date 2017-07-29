package main
import (
	
)

 
type (
	RPC struct {
		cache		map[string]string
		requests	*Requests
		mu			*sync.RWMutex

	}
	
	

)

