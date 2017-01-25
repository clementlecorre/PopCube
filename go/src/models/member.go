package models

import (
	u "utils"
)

// Member describe the associtive table member between USER, CHANNEL, and ROLE
type Member struct {
	MemberID uint64  `gorm:"primary_key;column:idMember;AUTO_INCREMENT" json:"-"`
	User     User    `gorm:"column:user; not null;" json:"-`
	Channel  Channel `gorm:"column:channel; not null;" json:"-"`
	Role     Role    `gorm:"column:role; ForeignKey:IDRole;" json:"-"`
}

// IsValid check validity of member object
func (member *Member) IsValid() *u.AppError {
	if member.User == (User{}) {
		return u.NewLocAppError("Member.IsValid", "model.member.user.app_error", nil, "")
	}
	if member.Channel == (Channel{}) {
		return u.NewLocAppError("Member.IsValid", "model.member.channel.app_error", nil, "")
	}
	return nil
}
