package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	publicKeyLen  = 32
	privateKeyLen = 64
	seedLen       = 32
	addressLen    = 20
)

type PrivateKey struct {
	privKey ed25519.PrivateKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.privKey
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.privKey, msg),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	buf := make([]byte, publicKeyLen)
	copy(buf, p.privKey[32:])

	return &PublicKey{
		pubKey: buf,
	}
}

func NewPrivateKeyFromString(key string) *PrivateKey {
	b, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLen {
		panic(fmt.Sprintf("expect for %d seed length, but got %d ", seedLen, len(seed)))
	}

	return &PrivateKey{
		privKey: ed25519.NewKeyFromSeed(seed),
	}
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)
	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		panic(err)
	}

	return &PrivateKey{
		privKey: ed25519.NewKeyFromSeed(seed),
	}

}

type PublicKey struct {
	pubKey ed25519.PublicKey
}

func (p *PublicKey) Bytes() []byte {
	return p.pubKey
}

func (p *PublicKey) Address() Address {
	return Address{
		value: p.pubKey[len(p.pubKey)-addressLen:],
	}
}

type Signature struct {
	value []byte
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.pubKey, msg, s.value)
}

type Address struct {
	value []byte
}

func (a *Address) Bytes() []byte {
	return a.value
}

func (a *Address) String() string {
	return hex.EncodeToString([]byte(a.value))
}
