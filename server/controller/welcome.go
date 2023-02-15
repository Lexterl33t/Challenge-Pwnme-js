package controller

import (
	b64 "encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Token struct {
	Token string `json:"token"`
}

const VALID_KEY = "0x5df1347da0cdd31a2d1c77d8ce5e4bdd65ee8effd949b841ad9a882597c97fbe"
const MAGIC_NUMBER_CLIENT = 0x13371337

var types map[string]string = map[string]string{
	"0": "Send valid request to get flag",
	"1": "1_l1ke_0bfusc4t1on",
}

type ParsedToken struct {
	Alphabet  string
	Timestamp string
	Key       string
	Axiome    string
	Type      string
	Magic     string
}

func is_alpha_uppercase(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	return false
}

func is_str_alpha_uppercase(str string) bool {
	for _, c := range str {
		if !is_alpha_uppercase(byte(c)) {
			return false
		}
	}

	return true
}

func parse_alphabet(alphabet string) (string, error) {
	var alphabet_decode_string string = ""

	alphabet_decoded, err := b64.StdEncoding.DecodeString(alphabet)
	if err != nil {
		return "", err
	}

	alphabet_wine := strings.Split(string(alphabet_decoded), ".")

	for _, letter := range alphabet_wine {
		letter_dec, err := b64.StdEncoding.DecodeString(letter)
		if err != nil {
			return "", err
		}

		alphabet_decode_string += string(letter_dec)
	}

	return alphabet_decode_string, nil

}

func parse_timestamp(timestamp string) (string, error) {
	timeDec, err := b64.StdEncoding.DecodeString(timestamp)
	if err != nil {
		return "", err
	}

	return string(timeDec), nil
}

func parse_key(key string) (string, error) {
	keyDec, err := b64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	return string(keyDec), nil
}

func parse_axiome(axiome string) (string, error) {
	axiomeDec, err := b64.StdEncoding.DecodeString(axiome)
	if err != nil {
		return "", err
	}

	return string(axiomeDec), nil
}

func parse_type(type_enc string) (string, error) {
	typeDec, err := b64.StdEncoding.DecodeString(type_enc)
	if err != nil {
		return "", err
	}

	return string(typeDec), nil
}

func parse_magic(magic string) (string, error) {
	magicDec, err := b64.StdEncoding.DecodeString(magic)
	if err != nil {
		return "", err
	}

	return string(magicDec), nil
}

func parse_token(token string) (ParsedToken, error) {

	tokens := strings.Split(token, ".")

	if len(tokens) != 6 {
		return ParsedToken{}, errors.New("Syntax error token")
	}

	alphabet := tokens[0]
	timestamp := tokens[1]
	key := tokens[3]
	axiome := tokens[2]
	typereq := tokens[4]
	magic := tokens[5]

	alphabet_string, err := parse_alphabet(alphabet)
	if err != nil {
		return ParsedToken{}, err
	}

	timestamp_dec, err := parse_timestamp(timestamp)
	if err != nil {
		return ParsedToken{}, err
	}

	key_dec, err := parse_key(key)
	if err != nil {
		return ParsedToken{}, err
	}

	axiome_dec, err := parse_axiome(axiome)
	if err != nil {
		return ParsedToken{}, err
	}

	type_dec, err := parse_type(typereq)
	if err != nil {
		return ParsedToken{}, err
	}

	magic_dec, err := parse_type(magic)
	if err != nil {
		return ParsedToken{}, err
	}

	return ParsedToken{
		Alphabet:  alphabet_string,
		Timestamp: timestamp_dec,
		Key:       key_dec,
		Axiome:    axiome_dec,
		Type:      type_dec,
		Magic:     magic_dec,
	}, nil

}

func check_alphabet(alphabet string) error {
	if len(alphabet) != 26 {
		return errors.New("Unknow alphabet token")
	}

	if !is_str_alpha_uppercase(alphabet) {
		return errors.New("Unknow alphabet")
	}

	return nil
}

func check_timestamp(timestamp string) error {

	timed, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return err
	}

	tm := time.Unix(timed, 0)

	time_now := time.Now()
	if tm.Year() != time_now.Year() {
		return errors.New("invalid timestamp")
	}
	return nil
}

func check_key(key string) error {
	if key != VALID_KEY {
		return errors.New("Unknow key")
	}

	return nil
}

func check_axiome(axiome string) error {
	num, err := strconv.Atoi(axiome)
	if err != nil {
		return err
	}

	num = num >> 1

	if (num^MAGIC_NUMBER_CLIENT) < 0 || (num^MAGIC_NUMBER_CLIENT) > 100 {
		return errors.New("Invalid axiome")
	}

	return nil

}

func check_magic(magic string) error {
	if magic != "token" {
		return errors.New("Invalid token")
	}

	return nil
}

func check_token(p_token ParsedToken) (string, error) {
	if p_token == (ParsedToken{}) {
		return "", errors.New("Where is fucking token bro ?")
	}

	if err := check_alphabet(p_token.Alphabet); err != nil {
		return "", err
	}

	if err := check_timestamp(p_token.Timestamp); err != nil {
		return "", err
	}

	if err := check_key(p_token.Key); err != nil {
		return "", err
	}

	if err := check_axiome(p_token.Axiome); err != nil {
		return "", err
	}

	typed, ok := types[p_token.Type]
	if !ok {
		return "", errors.New("invalid type")
	}

	if err := check_magic(p_token.Magic); err != nil {
		return "", nil
	}

	return typed, nil

}

func Welcome(ctx *fiber.Ctx) error {
	token := new(Token)

	if ctx.Get("User-Agent") != "l33t_Akeur" {
		ctx.Context().SetStatusCode(403)
		ctx.JSON(map[string]string{
			"error": "Invalid Header",
		})
		return nil
	}

	if err := ctx.BodyParser(token); err != nil {
		return err
	}

	if token.Token == "" {
		ctx.Context().SetStatusCode(403)
		ctx.JSON(map[string]string{
			"error": "Invalid token",
		})

		return nil
	}

	if len(token.Token) != 305 {
		ctx.Context().SetStatusCode(403)
		ctx.JSON(map[string]string{
			"error": "Invalid token",
		})

		return nil
	}

	parsed_token, err := parse_token(token.Token)
	if err != nil {
		ctx.Context().SetStatusCode(403)
		ctx.JSON(map[string]string{
			"error": err.Error(),
		})
		return nil
	}

	res, err := check_token(parsed_token)
	if err != nil {
		ctx.Context().SetStatusCode(403)
		ctx.JSON(map[string]string{
			"error": err.Error(),
		})
		return nil
	}

	ctx.JSON(map[string]string{
		"message": res,
	})
	return nil
}
