package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vincen320/user-service-graphql/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
	//BUAT TABLE BARU SBG RELASI DAN GET JGUA [DONE] (CHILDS)
	//coba lihat dokumentasi yang operation name yang HeroAndFriends --> WILL DO (ongoing)
	//BUAT UNTUK MUTATIONS (baca lagi dokumentasi)
}
