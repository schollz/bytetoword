package bytetoword

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	s := EncodeToString([]byte{1, 2, 3, 4})
	assert.Equal(t, "paff-halsy", s)
	s = EncodeToString([]byte{1, 2, 3, 4, 5})
	assert.Equal(t, "paff-halsy-repays", s)
	b, err := Decode("paff-halsy-repays")
	assert.Nil(t, err)
	assert.Equal(t, []byte{1, 2, 3, 4, 5}, b)
}

func ExampleMd5() {
	hasher := md5.New()
	hasher.Write([]byte("hello, world"))
	s := EncodeToString(hasher.Sum(nil))
	fmt.Println(s)
	b, _ := Decode(s)
	fmt.Println(hex.EncodeToString(b))
	// Output:
	// angili-apepsy-uriel-justs-puir-make-ort-rotted
	// e4d7f1b4ed2e42d15898f4b27b019da4
}
