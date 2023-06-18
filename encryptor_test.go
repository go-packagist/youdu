package youdu

import "testing"

func TestEncryptor(t *testing.T) {
	encryptor := NewEncryptor("ydDC01991D231D55C4F94407522A782280D4", "corpSecret")

	encrypted, err := encryptor.Encrypt("hello")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(encrypted)

	decrypted, err := encryptor.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(decrypted)
}
