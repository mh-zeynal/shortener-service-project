package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"url/internal/db"
	"url/internal/model"
	"url/internal/urlGenerator"
)

func HandleShortener(c echo.Context) error {
	temp := model.Input{}
	err := json.NewDecoder(c.Request().Body).Decode(&temp)
	if !db.IsConnectionStablished() {
		db.MakeDatabase("root", "m@96@s97", "gamedatabase")
	}
	if err != nil {
		fmt.Println("no json passed to handler")
	}
	if db.IsLongAvailable(temp.IntendedUrl) {
		return c.String(http.StatusOK, "this url already exists:\n" +
			"localhost:9090/" + db.GetRowViaLongUrl(temp.IntendedUrl).Short)
	}
	short := urlGenerator.GenerateShortUrl()
	db.AddNewRow(short, temp.IntendedUrl)
	return c.String(http.StatusOK, "shortened URL: localhost:9090/" + short)
}

func HandleRedirects(c echo.Context) error {
	shortUrl := c.Param("s")
	if !db.IsConnectionStablished() {
		db.MakeDatabase("root", "m@96@s97", "gamedatabase")
	}
	temp := db.GetRowViaShortUrl(shortUrl)
	return c.Redirect(http.StatusMovedPermanently, temp.Long)
}
