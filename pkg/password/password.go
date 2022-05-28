package password

// hash Password 非对称加密
import (
	"golang.org/x/crypto/bcrypt"
)

// 这个函数的功能是将密码加密，返回加密后的密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// 这个函数的功能是将密码与加密后的密码进行比较，如果一致，返回true，否则返回false
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func main() {
// 	password := "secret"
// 	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

// 	fmt.Println("Password:", password)
// 	fmt.Println("Hash:    ", hash)

// 	match := CheckPasswordHash(password, hash)
// 	fmt.Println("Match:   ", match)
// }
