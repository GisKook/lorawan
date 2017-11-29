package redis_socket

import (
	"github.com/giskook/lorawan/base"
)

func (r *RedisSocket) UserSignup(uuid, html string, user *base.User) error {
	conn := r.GetConn()
	defer conn.Close()

	key := user.Email + uuid
	err := conn.Send("HSET", key, SIGNUP_EMAIL_FIELD, user.Email)
	if err != nil {
		return err
	}
	err = conn.Send("HSET", key, SIGNUP_ALIAS_FIELD, user.Alias)
	if err != nil {
		return err
	}
	err = conn.Send("HSET", key, SIGNUP_PASSWD_FIELD, user.Passwd)
	if err != nil {
		return err
	}
	err = conn.Send("HSET", key, SIGNUP_TMP_HTML_FIELD, html)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, r.conf.Expire)
	if err != nil {
		return err
	}
	_, err = conn.Do("")
	if err != nil {
		return err
	}

	return nil
}
