package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privateKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), publicKeyLen)
}

func TestSignature_Verify(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()

	msg := []byte("aku mung ngetest")
	sig := privKey.Sign(msg) // sign the true message

	assert.True(t, sig.Verify(pubKey, msg))

	// do invalid message
	assert.False(t, sig.Verify(pubKey, []byte("expect salah")))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.Public()
	address := publicKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "d94b04e84739fa9a491002c4a10f8de1f7479eea7a45e3ab9e91f7792649fa61" // lu harus generate manual
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "11d0ac6e4e2c83ee08a9f9a54074555883afcc7f" // ini juga harus generate manual pake `func (p *PublicKey) Address() Address {}`
	)
	assert.Equal(t, privateKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}
