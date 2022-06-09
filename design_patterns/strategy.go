package main

import "fmt"

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtector(user string, password string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  password,
		hashAlgorithm: hash,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}
	pwrdProtector := NewPasswordProtector("jose", "gmail", sha)
	pwrdProtector.Hash()
	pwrdProtector.SetHashAlgorithm(md5)
	pwrdProtector.Hash()
}
