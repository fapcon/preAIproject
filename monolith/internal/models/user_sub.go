package models

type UserSub struct {
	userID    int
	subUserID int
}

func (s *UserSub) SetUserID(userID int) *UserSub {
	s.userID = userID
	return s
}

func (s *UserSub) GetUserID() int {
	return s.userID
}

func (s *UserSub) SetSubscriberID(userID int) *UserSub {
	s.subUserID = userID
	return s
}

func (s *UserSub) GetSubscriberID() int {
	return s.userID
}
