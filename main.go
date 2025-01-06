package main

import (
	"fmt"
	"math/rand"
	// "time"
	"hash"
	"github.com/spaolacci/murmur3"
	"github.com/google/uuid"
)

var hashfns []hash.Hash32

func init(){
	for i:=0; i<100; i++{
		hashfns = append(hashfns, murmur3.New32WithSeed(rand.Uint32()))
	}
}

func murmurhash(key string, size int32, hashFnIdx int) int32{

	//write the data to the hasher
	hashfns[hashFnIdx].Write([]byte(key))
	//Get the resulting hahs value
	var result = hashfns[hashFnIdx].Sum32()%uint32(size)
	hashfns[hashFnIdx].Reset()

	return int32(result)
}

type BloomFilter struct{
	filter []uint8
	size int32
}

func NewBloomFilter(size int32) *BloomFilter{
	return &BloomFilter{
		filter: make([]uint8, size),
		size: size,
	}
}

func (b *BloomFilter) Add(key string, numHashfns int){
	for i:=0; i<numHashfns; i++{
		idx := murmurhash(key, b.size, i)
		aidx := idx/8
		bidx := idx%8
		b.filter[aidx] = b.filter[aidx] | (1 << bidx)
	}	
	// fmt.Println("wrote", key, "index", idx)
}

func (b *BloomFilter) Exists(key string, numHashfns int) bool{
	exist := true
	for i:=0; i<numHashfns; i++{
		idx := murmurhash(key, b.size, i)
		aidx := idx/8
		bidx := idx%8
		if !(b.filter[aidx] & (1 << bidx) > 0){
			return false
		}
	}
	return exist
}

func (b *BloomFilter) Print(){
	fmt.Println(b.filter)
}

func main(){

	dataset := make([]string, 0)
	database_exists := make(map[string]bool)
	database_notexists := make(map[string]bool)

	for i:=0; i<500; i++{
		u := uuid.New()
		dataset = append(dataset, u.String())
		database_exists[u.String()] = true
	}

	for i:=0; i<500; i++{
		u := uuid.New()
		dataset = append(dataset, u.String())
		database_notexists[u.String()] = false
	}

	for i:=1; i<len(hashfns); i++{
		bloom := NewBloomFilter(int32(10000))
		
		for key, _ := range database_exists{
			bloom.Add(key, i)
		}

		truePositive := 0
		falsePositive := 0
		
		for _, key :=  range dataset{
			exists := bloom.Exists(key, i)
			if exists{
				if _, ok := database_exists[key]; ok{
					truePositive+=1
				}
				
				if _, ok := database_notexists[key]; ok{
					falsePositive+=1
				}
			}
		}
		fmt.Println((float64(falsePositive)/float64(len(dataset))))
	}

}
