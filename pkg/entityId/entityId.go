package entityId

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const SALT = "pm-app"

type Encoder struct{}

func (e *Encoder) Encode(id int, entityName string) string {
	idHex := fmt.Sprintf("%016X", id)
	entityHash := e.getEntityHash(entityName)
	checkSum := e.getCheckSum(idHex, entityHash)
	idHex = reverseString(idHex)
	base := splitStringByLength(idHex, 4)
	return strings.ToUpper(checkSum + strings.Join(base, "-") + entityHash)
}

func (e *Encoder) Decode(entityId string, entityName string) (int64, error) {
	entityId = strings.ToLower(entityId)
	//checkSum := entityId[:4]

	entityHash := entityId[len(entityId)-4:]

	if e.getEntityHash(entityName) != entityHash {
		return 0, errors.New(entityId)
	}

	idHex := strings.ReplaceAll(entityId[4:len(entityId)-4], "-", "")
	idHex = reverseString(idHex)

	decodedBytes, err := hex.DecodeString(idHex)
	if err != nil {
		return 0, err
	}

	if len(decodedBytes) != 8 {
		fmt.Println("error here")
		return 0, errors.New("Invalid decoded ID length")
	}

	decoded, _ := hex.DecodeString(idHex)
	etc := binary.BigEndian.Uint64(decoded)
	var l int64 = int64(etc)
	checkSum := entityId[:4]

	if e.getCheckSum(idHex, entityHash) != checkSum {
		return 0, errors.New("Checksum not match" + entityId)
	}

	return l, nil
}

func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

func (e *Encoder) getEntityHash(entityName string) string {
	hash := md5.Sum([]byte(entityName))
	return hex.EncodeToString(hash[:2]) // Change to [:2] for first 4 bytes
}

func (e *Encoder) getCheckSum(entityIdHex string, entityHash string) string {
	data := entityIdHex + "-" + entityHash + "-" + SALT
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[2:4]) // Change to [2:4] for bytes 5-8
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func splitStringByLength(input string, length int) []string {
	var result []string
	for i := 0; i < len(input); i += length {
		end := i + length
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}
