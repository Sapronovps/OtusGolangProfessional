//go:generate easyjson -all stats.go
package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

type users []User

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

func getUsers(r io.Reader) (result users, err error) {
	res := make(users, 0, 100000)
	sc := bufio.NewScanner(r)
	var user User

	for sc.Scan() {
		if err := user.UnmarshalJSON(sc.Bytes()); err != nil {
			return res, err
		}
		res = append(res, user)
	}
	return res, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		parts := strings.SplitN(user.Email, "@", 2)
		if len(parts) != 2 {
			continue // пропускаем некорректный email
		}

		domainPart := strings.ToLower(parts[1])
		if strings.HasSuffix(domainPart, "."+domain) {
			result[domainPart]++
		}
	}
	return result, nil
}
