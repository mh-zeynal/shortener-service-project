package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"back-end/dateExtractor"
	"back-end/db"
	"back-end/model"
	"back-end/urlGenerator"
)

func HandleShortener(c echo.Context) error {
	temp := model.Input{}
	err := json.NewDecoder(c.Request().Body).Decode(&temp)
	if !db.IsConnectionEstablished() {
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
	if !db.IsConnectionEstablished() {
		db.MakeDatabase("root", "m@96@s97", "gamedatabase")
	}
	if !db.IsShortAvailable(shortUrl) {
		return c.String(http.StatusNotFound, "this link is unavailable")
	}
	temp := db.GetRowViaShortUrl(shortUrl)
	if dateExtractor.IsLinkExpired(temp.Date) {
		return c.String(http.StatusForbidden, "this link is expired")
	}
	return c.Redirect(http.StatusSeeOther, temp.Long)
}

func HandleAdminAccess(c echo.Context) error {
	user := c.QueryParam("u")
	pass := c.QueryParam("p")
	if !(user == "mh" && pass == "zz") {
		return c.String(http.StatusBadRequest, "access denied")
	}
	if !db.IsConnectionEstablished() {
		db.MakeDatabase("root", "m@96@s97", "gamedatabase")
	}
	rows := *db.GetAllRows()
	return c.JSON(http.StatusOK, rows)
}
