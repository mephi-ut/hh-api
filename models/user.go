package models

//go:generate reform

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"hash"
	"strconv"
	"strings"
)

//reform:users
type user struct {
	Id           int     `reform:"id,pk"`
	Nickname     string  `reform:"nickname"`
	Email        *string `reform:"email"`
	PasswordHash string  `reform:"password_hash" sql_size:"255"`
}

type UserI interface {
	GetId() int
	GetNickname() string
	GetEmail() *string
}

var (
	defaultHashFunct func() hash.Hash
)

const (
	SALT_SIZE   int = 8
	HASH_ROUNDS int = 79
)

func init() {
	defaultHashFunct = sha512.New
}

func NewUser() *user {
	return &user{}
}

func (user user) GetId() int {
	return user.Id
}

func (user user) GetNickname() string {
	return user.Nickname
}

func (user user) GetEmail() *string {
	return user.Email
}

func HashPassword(password string) string {
	return CryptoHash([]byte(password), nil, nil)
}

func genSalt() (salt []byte) {
	salt = make([]byte, SALT_SIZE)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	return
}

func CryptoHash(input []byte, salt []byte, hashfn func() hash.Hash) string {
	if hashfn == nil {
		hashfn = defaultHashFunct
	}

	if len(salt) == 0 {
		salt = genSalt()
	}

	return "sha512," + strconv.Itoa(HASH_ROUNDS) + "," + hex.EncodeToString(salt) + "," + hex.EncodeToString(pbkdf2.Key(input, salt, HASH_ROUNDS, 64, hashfn))
}

func ParseCryptoHash(fullhash string) (hashfn func() hash.Hash, salt []byte, hashself []byte) {
	var err error
	words := strings.Split(fullhash, ",")

	switch words[0] {
	case "sha512":
		hashfn = sha512.New
	default:
		panic(fmt.Errorf("Unknown hash algorithm: \"%s\"", words[0]))
	}

	salt, err = hex.DecodeString(words[2])
	if err != nil {
		panic(err)
	}

	hashself, err = hex.DecodeString(words[3])
	if err != nil {
		panic(err)
	}

	return hashfn, salt, hashself
}

func (user user) CheckPassword(password []byte) bool {
	if user.PasswordHash == "" {
		return false
	}

	hashfunct, salt, _ := ParseCryptoHash(user.PasswordHash)
	passwordHash := CryptoHash(password, salt, hashfunct)

	return passwordHash == user.PasswordHash
}
