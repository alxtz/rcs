package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_A(t *testing.T) {
	t.Skip()
	assert.Exactly(
		t,
		4,
		getLengthOfOptimalCompression("a3bc3d", 2),
	)
}

func Test_B(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		2,
		getLengthOfOptimalCompression("a2b5a2", 2),
		getLengthOfOptimalCompression("a2b5a2c2a2", 2),
		// 要是每個 item 都獨特，就不會有這個問題了
		// 如果我再設計一個 DS，記錄所有這樣的東西呢？

		// A: "a2 *5 a2 *2 a2"
		// B: "*2 b5 *6"

		// with: aaa b aaa cc d ee f gg xxx y xxx
		// A: a3 *1 a3 *15
		// 價值 1 單位換 3 單位, 15 單位換 15 單位
		// B: *3 b *18
		// 價值 3 單位換 3 單位, 18 單位換 18 單位
		// C: *7 c2 *13
		// 價值 7 單位換 7 單位, 13 單位換 13 單位
		// D: *9 d *12
		// 價值 9 單位換 9 單位, 12 單位換 12 單位
		// E: *10 e2 *10
		// 價值 10 單位換 10 單位, 12 單位換 12 單位
	)
}

func Test_C(t *testing.T) {
	t.Skip()
	assert.Exactly(
		t,
		3,
		getLengthOfOptimalCompression("a11", 3),
	)
}
