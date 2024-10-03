package service

import (
	"go.uber.org/ratelimit"
	"sync"
)

type RateLimiter struct {
	UserLimits map[int]ratelimit.Limiter
	Mu         sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		UserLimits: make(map[int]ratelimit.Limiter),
	}
}

func (u *RateLimiter) GetUserLimiter(userID int, rate int) ratelimit.Limiter {
	u.Mu.Lock()
	defer u.Mu.Unlock()

	l, ok := u.UserLimits[userID]
	if !ok {
		// Создаем новый лимитер для пользователя
		l = ratelimit.New(rate)
		u.UserLimits[userID] = l
	}

	return l
}
