package main

import (
	// "fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ToReturn struct {
	Url   string  `json:"url"`
}


func main() {
	// fmt.Println("fdfdfdf")
	app := fiber.New();
	app.Get("/baka", GetBaka)
	app.Get("/img", GetImg)
	app.Get("/dice", GetDice)


	app.Listen(":3000")
}


func GetBaka(c *fiber.Ctx) error {
	f, err := os.OpenFile("file.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("http://127.0.0.1:3000/baka")


	randInt := GetRandomInt(12)
	url := "https://cdn.catboys.com/baka/baka_" + strconv.Itoa(randInt) + ".gif"
	ret := ToReturn{url}

	return c.JSON(ret)
}

func GetDice(c *fiber.Ctx) error {
	f, err := os.OpenFile("file.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("http://127.0.0.1:3000/dice")


	
	randInt := GetRandomInt(6)
	randInt++
	url := "https://cdn.catboys.com/dice/" + strconv.Itoa(randInt) + ".png"
	ret := ToReturn{url}

	return c.JSON(ret)
}

func GetImg(c *fiber.Ctx) error {
	f, err := os.OpenFile("file.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("http://127.0.0.1:3000/img")


	
	randInt := GetRandomInt(301)
	url := "https://cdn.catboys.com/images/image_" + strconv.Itoa(randInt) + ".jpg"
	ret := ToReturn{url}

	return c.JSON(ret)
}


func GetRandomInt( max int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}