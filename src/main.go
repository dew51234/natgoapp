package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/testgoroutine", callGoRoutine)
	e.POST("/users2", testPost2)
	e.Logger.Fatal(e.Start(GetPort()))
}

func GetPort() string {
	return_port := os.Getenv("PORT")
	if return_port == "" {
		return_port = "1234"
	}
	return ":" + return_port
}

func callGoRoutine(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	res := testgoRoutine(u)
	return c.String(http.StatusOK, res)
}

func testgoRoutine(u *User) string {
	c := make(chan string)
	c2 := make(chan string)
	go f(c, u.Name)
	go f2(c2, u.Email)
	nsg := <-c
	msg := <-c2
	return nsg + msg
}

func f(c chan string, n string) {
	var ff string
	for i := 0; i < 10; i++ {
		ff += ("chan" + strconv.Itoa(i) + ":" + n)
	}
	c <- ff
}

func f2(c chan string, ip string) {
	time.Sleep(time.Second * 3)
	c <- ip
}

func testPost2(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
