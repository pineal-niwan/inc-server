package main

import (
	"time"
)

func (s *IncService) getIncNumberByKey(key string) (uint32, error) {
	time.Sleep(20 * time.Millisecond)
	return 0, nil
}
