package main

import (
	"fmt"
)

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}
	err2 := a.track(em.status)
	if err2 != nil {
		return fmt.Errorf("error updating user bounce: %w", err2)
	}
	return nil
}
