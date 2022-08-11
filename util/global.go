package util

import "sync"

// mutex -> used to solve race condition when insert
var Mutex sync.Mutex
