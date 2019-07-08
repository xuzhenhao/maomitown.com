package auth

import "golang.org/x/crypto/bcrypt"

// Encrypt 对密码进行bcrypt加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 比对加密后的密码和当前密码是否一致
func Compare(hashedPwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
