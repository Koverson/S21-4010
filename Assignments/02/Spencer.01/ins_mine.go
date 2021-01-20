package mine

import (
	"encoding/hex"
	"fmt"
	"runtime"

	block "github.com/Univ-Wyo-Education/S20-4010/Assignments/02/block"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/02/hash"
)

// MineBlock implements a proof of work mining system where the first 4 digits (2 bytes) of the hex value are 0.
// Difficulty can be increaesed by requiring more digits to be 0 or by requring some other pattern to apear in
// the resulting hash.
func InstructorImplementationMineBlock(bk *block.BlockType, difficulty string) {
	// Pseudo-Code
	//
	// 1. Use an infinite loop to:
	//		1. Searilize the data from the block for hasing, Call block.SerializeForSeal to do this.
	//		2. Calculate the hash of the data, Call hash.HashOf to do this. This is the slow part.  What would happen if we
	//         replaced the software with a hash calculator on a graphics card where you could run 4096 hahes at once?
	//		   What would happen if we replaced the graphics card with an ASIC - so you had dedicated hardware to do
	//         the hash and you could run 4 billion hashes a second?
	//		3. Convert the hash (it is []byte) to a hex string.  Use the hex.EncodeToString standard go library function.
	//		4. `fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\r", theHashAsAString, bk.Nonce)`
	//			`\r` will overwrite the same line instead of advancing to the next.
	//		5. See if the first 4 characters of the hash are 0's. - if so we have met the work criteria.
	//			In go this is `if theHashAsAString[0:4] == "0000" {`.  This is create a slice, 4 long from
	//			character 0 with length of 4, then compare that to the string `"0000"`.
	//			- Set the blcoks "Seal" to the hash
	//			- `fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\n", theHashAsAString, bk.Nonce)`
	//			  `\n` will overwrite the same *and then advance* to the next line.
	//			- return
	//		5. Increment the Nonce in the block, and...
	//		6. Back to the top of the loop for another try at finding a seal for this block.
	//
	// For the genesis block, when I do this it requires 54586 trips through the loop to calculate the
	// proof of work.
	//
	for {
		data := block.SerializeForSeal(bk)
		hh := hash.HashOf(data)
		hhStr := hex.EncodeToString(hh)
		if runtime.GOOS != "windows" {
			fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\r", hhStr, bk.Nonce)
		}
		if len(difficulty) == 0 || hhStr[0:len(difficulty)] == difficulty {
			bk.Seal = hh
			fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\n", hhStr, bk.Nonce)
			return
		}
		bk.Nonce++
	}
}