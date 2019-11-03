package model

import (
    "crypto/rand"
    _ "fmt"
    _ "log"

    argon2 "github.com/tvdburgt/go-argon2"
)

type User struct {
    Base
    Email    string `json:"email"`
    Password string `json:"-"`
    Slug     string `json:"slug"`
}

func (u *User) SetPassword(password string) error {
    salt, err := generateRandomBytes(16)
    if err != nil {
        return err
    }

    hashed, err := argon2.HashEncoded(argon2.NewContext(), []byte(password), salt)
    if err != nil {
        return err
    }
    u.Password = hashed
    return nil
}

func (u *User) VerifyPassword(password string) bool {
    match, err := argon2.VerifyEncoded(u.Password, []byte(password))
    if err != nil {
        // log.Println(err)
        return false
    }
    return match
}

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}
