package main

import (
	"fmt"
	"math/big"
)

type BigProc struct {
	bigProc      big.Int
	bigProcFloat big.Float
}

func main() {

	//var newBigInt1 = big.NewInt(1 << 210) (overflows)
	bigInt1 := "1645504557321206042154969182557350504982735865633579863348609024"
	//var newBigInt2 = big.NewInt(1 << 220) (overflows)
	bigInt2 := "1684996666696914987166688442938726917102321526408785780068975640576"

	newBigInt1 := new(big.Float)
	newBigInt1.SetString(bigInt1)

	newBigInt2 := new(big.Float)
	newBigInt2.SetString(bigInt2)

	fmt.Println(newBigInt1, newBigInt2)

	myBig := new(BigProc)
	fmt.Println(myBig.bigSum(newBigInt1, newBigInt2))
	fmt.Println(myBig.bigSub(newBigInt1, newBigInt2))
	fmt.Println(myBig.bigDiv(newBigInt1, newBigInt2))
	fmt.Println(myBig.bigMul(newBigInt1, newBigInt2))
}

func (b *BigProc) bigSum(first, second *big.Float) *big.Float {
	return b.bigProcFloat.Add(first, second)
}

func (b *BigProc) bigSub(first, second *big.Float) *big.Float {
	return b.bigProcFloat.Sub(first, second)
}

func (b *BigProc) bigDiv(first, second *big.Float) *big.Float {
	return b.bigProcFloat.Quo(first, second)
}

func (b *BigProc) bigMul(first, second *big.Float) *big.Float {
	return b.bigProcFloat.Mul(first, second)
}
