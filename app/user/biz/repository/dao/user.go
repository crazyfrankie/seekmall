package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id     int    `gorm:"primaryKey,autoIncrement"`
	Phone  string `gorm:"unique"`
	Name   string
	Avatar string
	Ctime  int64
	Utime  int64
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Create(ctx context.Context, u *User) error {
	now := time.Now().Unix()
	u.Ctime = now
	u.Utime = now

	return d.db.WithContext(ctx).Create(&u).Error
}

func (d *UserDao) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Model(&User{}).Where("phone = ?", phone).Find(&u).Error
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (d *UserDao) FindById(ctx context.Context, uid int) (User, error) {
	var u User
	err := d.db.WithContext(ctx).Model(&User{}).Where("id = ?", uid).Find(&u).Error
	if err != nil {
		return User{}, err
	}

	return u, nil
}
