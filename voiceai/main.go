package main

import (
	"fmt"
)

func main() {
	/*	appid := "6b4ac245d17744e6b8dda9348e6e8ff5"
		appSecret := "9bcd7652cd3b2df35e24c9a062de5414"
		nonce := "2644e43375024b70b326dc9278274903|1640749853053046"
		timestamp := 1640749853053046
		hash := sha256.New()
		hashElements := []string{appid, appSecret, nonce, strconv.FormatInt(int64(timestamp), 10)}
		sort.Strings(hashElements)
		hash.Write([]byte(strings.Join(hashElements, "")))
		hashSum := hash.Sum(nil)
		signature := hex.EncodeToString(hashSum)
		fmt.Println(signature)
	*/
	var res string
	var err error
	if res, err = Encrypt([]byte(`{"app_id":"1496415d27da4cf4998c6ecb1bb4b812"}`), []byte("1496415d27da4cf4998c6ecb1bb4b814")); err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
