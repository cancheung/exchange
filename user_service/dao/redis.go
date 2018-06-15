package dao

import (
	"digicon/common/encryption"
	cf "digicon/user_service/conf"
	. "digicon/user_service/log"
	"digicon/user_service/tools"
	"github.com/go-redis/redis"
	"time"
)

type RedisCli struct {
	rcon   *redis.Client
	KeyTtl time.Duration
	salt   string
}

func NewRedisCli() *RedisCli {
	addr := cf.Cfg.MustValue("redis", "addr")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		Log.Fatalf("redis connect faild ")
	}
	Log.Infoln(pong)

	ct, err := cf.Cfg.Int64("redis", "ttl")
	if err != nil {
		ct = 30
	}
	return &RedisCli{
		rcon:   client,
		salt:   "mjfdsap832-1##1!",
		KeyTtl: time.Duration(ct) * time.Second,
	}
}

func (s *Dao) GenSecurityKey(phone string) (security_key []byte, err error) {
	security_key = encryption.Gensha256(phone, time.Now().Unix(), s.redis.salt)
	err = s.redis.rcon.Set(tools.GetPhoneTagByLogic(phone, tools.LOGIC_SECURITY), security_key, s.redis.KeyTtl).Err()
	if err != nil {
		Log.Errorln(err.Error())
		return
	}
	return
}

//func (s *Dao) GetSecurityKey(uid int32) (security_key []byte, err error) {
func (s *Dao) GetSecurityKeyByPhone(phone string) (security_key []byte, err error) {
	security_key, err = s.redis.rcon.Get(tools.GetPhoneTagByLogic(phone, tools.LOGIC_SECURITY)).Bytes()
	if err != nil {
		Log.Errorln(err.Error())
		return
	}
	return
}

func (s *Dao) SetSmsCode(phone string, code string, ty int32) (err error) {
	err = s.redis.rcon.Set(tools.GetPhoneTagByLogic(phone, ty), code, 600*time.Second).Err()
	if err != nil {
		Log.Errorln(err.Error())
		return
	}
	return
}

func (s *Dao) GetSmsCode(phone string, ty int32) (code string, err error) {
	code, err = s.redis.rcon.Get(tools.GetPhoneTagByLogic(phone, ty)).Result()
	if err != nil {
		Log.Errorln(err.Error())
		return
	}
	return
}

func (s *Dao) SetTmpGoogleSecertKey(uid int32, code string) (err error) {
	err = s.redis.rcon.Set(tools.GetUserTagByLogic(uid, tools.UID_TAG_GOOGLE_SECERT_KEY), code, 600*time.Second).Err()
	if err != nil {
		Log.Errorln(err.Error())
		return
	}
	return
}

func (s *Dao) GetTmpGoogleSecertKey(uid int32) (key string, err error) {
	key, err = s.redis.rcon.Get(tools.GetUserTagByLogic(uid, tools.UID_TAG_GOOGLE_SECERT_KEY)).Result()
	if err != nil {
		Log.Errorln(err.Error())
		return
	}
	return
}
