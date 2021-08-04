package hub_helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
	"x-msa-observer/helper"
	"x-msa-observer/store/mongo/model"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitHelper(H HandlerForHelper) Helper { return &help{HandlerForHelper: H} }

func (h *help) AuthGuard(r *http.Request) (*model.UserModel, error) {
	if v := r.Context().Value(helper.CtxKeyValue); v != nil {
		if user, ok := v.(*model.UserModel); ok {
			return user, nil
		} else if err, ok := v.(error); ok {
			return nil, err
		}
	}
	return nil, fmt.Errorf("context is nil")
}

func (h *help) GenerateJWT(id primitive.ObjectID) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &model.UserClaims{ID: id}).SignedString([]byte(helper.Secret))
}

func (h *help) Hash(data string) string {
	hash1 := sha256.New()
	hash2 := md5.New()
	hash1.Write([]byte(data))
	hash2.Write([]byte(data))
	hash1.Write([]byte(string(hash1.Sum(nil)) + string(hash2.Sum(nil))))
	return hex.EncodeToString(hash1.Sum(nil))
}

func (h *help) KeyGenerate() string {
	i := 7
	buf := [8]byte{}
	unix := uint64(time.Now().UnixNano())
	for unix >= 0xff {
		buf[i] = byte(unix & 0xff)
		unix >>= 7
		i--
	}
	return hex.EncodeToString(buf[:])
}
