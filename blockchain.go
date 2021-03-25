package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Time          int64       // seconds since (unix) epoch (1970-01-01)
	Data          string
	Prev          string      // use []byte/int256/uint256 ??
	Hash          string      // use []byteint256/uint256 ??
	Nonce         int64       // number used once - lucky (mining) lottery number
}


// bin(ary) bytes and integer number to (conversion) string helpers
func binToStr( bytes []byte ) string {
	return hex.EncodeToString( bytes )
}

func intToStr( num int64 ) string {
	return strconv.FormatInt( num, 10 )
}



func calcHash( data string ) string {
	hashed := sha256.Sum256( []byte(data) )
	return binToStr( hashed[:] )   // note: [:] converts [32]byte to []byte
}

func computeHashWithProofOfWork( data string, difficulty string ) (int64,string) {
	nonce := int64( 0 )
	for {
		hash := calcHash( intToStr(nonce) + data )
		if strings.HasPrefix( hash, difficulty ) {
			return nonce,hash    // bingo! proof of work if hash starts with leading zeros (00)
		} else {
			nonce += 1           // keep trying (and trying and trying)
		}
	}
}


func NewBlock( data string, prev string ) Block {
	t           := time.Now().Unix()
	difficulty  := "000000"
	nonce, hash := computeHashWithProofOfWork( intToStr(t) + prev + data, difficulty )

	return Block { t, data, prev, hash, nonce }
}


func main() {
	b0 := NewBlock( "Hello, Cryptos!", "0000000000000000000000000000000000000000000000000000000000000000" )
	b1 := NewBlock( "Hello, Cryptos! - Hello, Cryptos!", b0.Hash )

	fmt.Println( b0 )
	// {1522691756 Hello, Cryptos!
	//    0000000000000000000000000000000000000000000000000000000000000000
	//    00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
	//    42278}
	fmt.Println( len( b0.Hash ))
	// => 64
	fmt.Println( len( b0.Prev ))
	// => 64

	fmt.Println( b1 )
	// {1522691756 Hello, Cryptos! - Hello, Cryptos!
	//     00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
	//     00009ef5ea432f840c3fb23dbedb5cce4c72e2951a140c1289dda1fedbcd6e99
	//     105106}

	fmt.Println( len( b1.Hash ))
	// => 64
	fmt.Println( len( b1.Prev ))
	// => 64

	blockchain := []Block {b0, b1}
	fmt.Println( blockchain )
	// => [{1522691756 Hello, Cryptos!
	//        0000000000000000000000000000000000000000000000000000000000000000
	//        00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
	//        42278}
	//     {1522691756 Hello, Cryptos! - Hello, Cryptos!
	//        00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
	//        00009ef5ea432f840c3fb23dbedb5cce4c72e2951a140c1289dda1fedbcd6e99
	//        105106}]
}


