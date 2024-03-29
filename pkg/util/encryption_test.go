package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncryption(t *testing.T) {

	Convey("When getting encryption key", t, func() {

		key, err := encryptionKeyToBytes("secret", "salt")
		So(err, ShouldBeNil)
		So(len(key), ShouldEqual, 32)

		key, err = encryptionKeyToBytes("a very long secret key that is larger then 32bytes", "salt")
		So(err, ShouldBeNil)
		So(len(key), ShouldEqual, 32)
	})

	Convey("When decrypting basic payload", t, func() {
		encrypted, encryptErr := Encrypt([]byte("smartems"), "1234")
		decrypted, decryptErr := Decrypt(encrypted, "1234")

		So(encryptErr, ShouldBeNil)
		So(decryptErr, ShouldBeNil)
		So(string(decrypted), ShouldEqual, "smartems")
	})

}
