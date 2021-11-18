package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/toky03/oracle-swap-demo/model"
)

func readSymbol() (model.UnCurrencySymbol, error) {
	err, file := readFile("symbol.json")
	if err != nil {
		return model.UnCurrencySymbol{}, err
	}
	var symbol model.UnCurrencySymbol
	err = json.Unmarshal(file, &symbol)
	if err != nil {
		log.Printf("failure reading symbol %s", err)
	}
	return symbol, nil
}

func readWallets() (map[string]string, error) {
	err, walletFiles := readDirectory()
	if err != nil {
		log.Printf("failure reading directory %s", err)
		return map[string]string{}, err
	}
	if len(walletFiles) == 0 {
		return  map[string]string{}, errors.New("No .cid files found in direcotry.")
	}
	wallets := make(map[string]string, 0)
	for _, file := range walletFiles {
		var content []byte
		regex := regexp.MustCompile(`^\d*$`)
		if regex.Match([]byte(file)) {
			err, content = readFile("W" + file + ".cid")
		} else {
			err, content = readFile(file + ".cid")
		}

		wallets[file] = string(content)
	}
	return wallets, nil
}

func readDirectory() (err error, wallets []string) {
	cidRootDirectory := os.Getenv("CID_ROOT")
	if cidRootDirectory == "" {
		cidRootDirectory = "../uniswap-pab/"
	}
	file, err := os.Open(cidRootDirectory)
	defer file.Close()
	list, err := file.Readdirnames(0)
	wallets = make([]string, 0, len(list))
	re := regexp.MustCompile(`W(\d*)\.cid`)
	for _, name := range list {
		result := re.FindAllStringSubmatch(name, 1)

		if len(result) > 0 {
			var fileName string
			if result[0][1] != "" {
				fileName = result[0][1]
			}
			wallets = append(wallets, fileName)
		}

	}
	return err, wallets

}

func readFile(filename string) (error, []byte) {
	cidRootDirectory := os.Getenv("CID_ROOT")
	if cidRootDirectory == "" {
		cidRootDirectory = ".."
	}
	data, err := ioutil.ReadFile(cidRootDirectory + "/" + filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return err, []byte("")
	}
	return nil, data
}
