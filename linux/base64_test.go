package linux

import "testing"

/**
 * @Author: LFM
 * @Date: 2022/8/18 00:01
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestBase64(t *testing.T) {
	encode := Encode("abcd52833sss")

	t.Logf(encode)

	decode, _ := Decode(encode)

	t.Logf(decode)
}
