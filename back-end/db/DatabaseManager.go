package db

import (
	"back-end/dateExtractor"
	"back-end/model"
)

func establishConnection() {
	if DB == nil {
		MakeDbConnection()
	}
}

func GetUserByUsername(username string) (*model.User, error) {
	establishConnection()
	rows, err := DB.Query("SELECT * FROM service_users WHERE username = ?",
		username)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var user model.User
	if rows.Next() {
		err = rows.Scan(&user.User_Id, &user.Username,
			&user.Password, &user.Email,
			&user.Name, &user.Created_at)
	}
	if err != nil || (user == model.User{}) {
		return nil, err
	}
	return &user, err
}

func GetUrlsByUserId(user_id uint) ([]model.URL, error) {
	establishConnection()
	rows, err := DB.Query("SELECT * FROM urls WHERE user_id = ?", user_id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var urls []model.URL
	for rows.Next() {
		var temp model.URL
		err = rows.Scan(&temp.Id, &temp.Short_url,
			&temp.Original_url, &temp.Name,
			&temp.Created_at, &temp.User_id)
		urls = append(urls, temp)
	}
	if err != nil {
		return nil, err
	}
	var temp []model.URL
	for _, url := range urls {
		if !dateExtractor.IsLinkExpired(url.Created_at) {
			temp = append(temp, url)
		}
	}
	return temp, err
}

func GetUrlsByUsername(username string) ([]model.URL, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	urls, err := GetUrlsByUserId(user.User_Id)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	establishConnection()
	row, err := DB.Query("SELECT * FROM service_users WHERE email = ?", email)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	var user model.User
	if row.Next() {
		err = row.Scan(&user.User_Id, &user.Username,
			&user.Password, &user.Email,
			&user.Name, &user.Created_at)
	}
	if err != nil || (user == model.User{}) {
		return nil, err
	}
	return &user, nil
}

func GetUrlByShortForm(shortForm string) (*model.URL, error) {
	establishConnection()
	row, err := DB.Query("SELECT * FROM urls WHERE short_url = ?", shortForm)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	var temp model.URL
	if row.Next() {
		err = row.Scan(&temp.Id, &temp.Short_url,
			&temp.Original_url, &temp.Name,
			&temp.Created_at, &temp.User_id)
	}
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func IsOriginalUrlAvailable(originalUrl string, username string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false
	}
	row, err := DB.Query("SELECT * FROM urls WHERE original_url = ? AND user_id = ?",
		originalUrl, user.User_Id)
	defer row.Close()
	if err != nil {
		return false
	}
	var temp model.URL
	if row.Next() {
		err = row.Scan(&temp.Id, &temp.Short_url,
			&temp.Original_url, &temp.Name,
			&temp.Created_at, &temp.User_id)
	}
	if (temp == model.URL{}) || err != nil {
		return false
	}
	return true
}

func IsShortUrlAvailable(shortUrl string, username string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false
	}
	row, err := DB.Query("SELECT * FROM urls WHERE short_url = ? AND user_id = ?",
		shortUrl, user.User_Id)
	defer row.Close()
	if err != nil {
		return false
	}
	var temp model.URL
	if row.Next() {
		err = row.Scan(&temp.Id, &temp.Short_url,
			&temp.Original_url, &temp.Name,
			&temp.Created_at, &temp.User_id)
	}
	if (temp == model.URL{}) || err != nil {
		return false
	}
	return true
}

func GetUrlByOriginal(original string, username string) (*model.URL, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	row, err := DB.Query("SELECT * FROM urls WHERE original_url = ? AND user_id = ?",
		original, user.User_Id)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	var temp model.URL
	if row.Next() {
		err = row.Scan(&temp.Id, &temp.Short_url,
			&temp.Original_url, &temp.Name,
			&temp.Created_at, &temp.User_id)
	}
	if err != nil {
		return nil, err
	}
	return &temp, nil
}

func InsertNewUrl(newUrl model.URL) error {
	establishConnection()
	insert, err :=
		DB.Query("INSERT INTO urls(short_url, original_url, name, user_id) VALUES (?, ?, ?, ?)",
			newUrl.Short_url, newUrl.Original_url, newUrl.Name, newUrl.User_id)
	defer insert.Close()
	if err != nil {
		return err
	}
	return nil
}

func InsertNewUser(newUser model.User) error {
	establishConnection()
	insert, err := DB.Query("INSERT INTO service_users(username, password, email, name) VALUES(?, ?, ?, ?)",
		newUser.Username, newUser.Password, newUser.Email, newUser.Name)
	defer insert.Close()
	if err != nil {
		return err
	}
	return nil
}
