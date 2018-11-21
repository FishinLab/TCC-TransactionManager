package participant

import "sync"

type Activity struct {
	locker 	sync.Locker
}

func NewActivity() *Activity {
	a := new(Activity)
	return a
}

func (a *Activity)Begin() (bool, error) {
	a.locker.Lock()
	defer a.locker.Unlock()

	// TODO: start a new transaction
	return true, nil
}