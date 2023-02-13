package data

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Thread - struct for Thread
type Thread struct {
	Title      string
	UserID     int
	Likes      int
	Dislikes   int
	ThreadID   int
	ToThreadID int
	Date       string
	Content    string
	Category   string
	Username   string
	Image      string
	Liked      int
}

//ThreadStats - struct for statistic Thread
type ThreadStats struct {
	FromUserID int
	ToThreadID int
	Value      int
}

//CreateTH - Adding Thread to Database
func CreateTH(r *http.Request, ID int, username string) (int, error) {
	var CurTH Thread
	err2 := r.ParseMultipartForm(20 << 20)
	if err2 != nil {
		return 0, err2
	}
	category := ""
	cur := 2
	for i := range r.Form {
		if i == "FileImage" {
			cur++
		}
	}
	if len(r.Form) == cur {
		category = "Oftop"
	}
	for i := range r.Form {
		if i == "title" || i == "comment" || i == "FileImage" {
			continue
		}
		if category != "" {
			category += "$"
		}
		category += i
	}
	Db.Exec("insert into Thread(Title, UserID, Likes, Dislikes, ToThreadID, Date, Content, Category, Username, Image) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", r.FormValue("title"), ID, 0, 0, 0, time.Now().Format("2006-01-02 15:04"), r.FormValue("comment"), category, username, "")
	Db.QueryRow("select * from Thread where ToThreadID = 0").Scan(&CurTH.Title, &CurTH.UserID, &CurTH.Likes, &CurTH.Dislikes, &CurTH.ThreadID, &CurTH.ToThreadID, &CurTH.Date, &CurTH.Content, &CurTH.Category, &CurTH.Username, &CurTH.Image)
	Db.Exec("update Thread set ToThreadID = $1 where ThreadID = $2", CurTH.ThreadID, CurTH.ThreadID)
	err := AddImage("Thread"+strconv.Itoa(CurTH.ThreadID), 0, CurTH.ThreadID, r)
	if err != nil && err.Error() == "No image" {
		err = nil
	} else {
		Db.Exec("update Thread set Image = $1 where ThreadID = $2", "Thread"+strconv.Itoa(CurTH.ThreadID), CurTH.ThreadID)
	}
	return CurTH.ThreadID, err
}

//GetThreadByID - Get Thread By ID
func GetThreadByID(ID int) (thread Thread, err error) {
	thread = Thread{}
	err = Db.QueryRow("select * from Thread where ThreadID = $1", ID).Scan(&thread.Title, &thread.UserID, &thread.Likes, &thread.Dislikes, &thread.ThreadID, &thread.ToThreadID, &thread.Date, &thread.Content, &thread.Category, &thread.Username, &thread.Image)
	if thread.Title == "" {
		err = fmt.Errorf("it's comment %v", err.Error())
	}
	return
}

//GetAllToThreadByID - Get All ToThread By ThreadID
func GetAllToThreadByID(ThreadID int, UserID int) []Thread {
	tmp := []Thread{}
	rows, err := Db.Query("select * from Thread where ToThreadID = $1", ThreadID)
	if err != nil {
		log.Println("err: ", err.Error())
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		cur := Thread{}
		err := rows.Scan(&cur.Title, &cur.UserID, &cur.Likes, &cur.Dislikes, &cur.ThreadID, &cur.ToThreadID, &cur.Date, &cur.Content, &cur.Category, &cur.Username, &cur.Image)
		if err != nil {
			break
		}
		cur.Liked = CheckUserLikedThread(UserID, cur.ThreadID)
		tmp = append(tmp, cur)
	}
	return tmp
}

//CheckUserLikedThread - Did a user liked this thread?
func CheckUserLikedThread(UserID int, ThreadID int) int {
	Stats := ThreadStats{}
	row := Db.QueryRow("select * from ThreadStats where FromUserID = $1 and ToThreadID = $2", UserID, ThreadID)
	err := row.Scan(&Stats.FromUserID, &Stats.ToThreadID, &Stats.Value)
	if err != nil {
		return 0
	}
	return Stats.Value
}

//CreateCommentToThread - Create comment to thread
func CreateCommentToThread(r *http.Request, UserID int, ThreadID int, username string) {
	r.ParseForm()
	Db.Exec("insert into Thread(Title, UserID, Likes, Dislikes, ToThreadID, Date, Content, Category, Username, Image) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", "", UserID, 0, 0, ThreadID, time.Now().Format("2006-01-02 15:04"), r.Form["comment"][0], "", username, "")
}

//UpdateThreadStats - Updates thread statistic
func UpdateThreadStats(ThreadID int, Value int, operation string) {
	Thread, _ := GetThreadByID(ThreadID)
	if operation == "-" {
		if Value == 1 {
			Db.Exec("update Thread set Likes = $1 where ThreadID = $2", Thread.Likes-1, ThreadID)
		} else {
			Db.Exec("update Thread set Dislikes = $1 where ThreadID = $2", Thread.Dislikes-1, ThreadID)
		}
	} else {
		if Value == 1 {
			Db.Exec("update Thread set Likes = $1 where ThreadID = $2", Thread.Likes+1, ThreadID)
		} else {
			Db.Exec("update Thread set Dislikes = $1 where ThreadID = $2", Thread.Dislikes+1, ThreadID)
		}
	}
}

//AddNewValueToThread - function allow to dislike and like comments or threads
func AddNewValueToThread(UserID int, ThreadID int, Value int) {
	Stats := ThreadStats{}
	row := Db.QueryRow("select * from ThreadStats where FromUserID = $1 and ToThreadID = $2", UserID, ThreadID)
	err := row.Scan(&Stats.FromUserID, &Stats.ToThreadID, &Stats.Value)
	if err == nil {
		if Stats.Value == 1 {
			UpdateThreadStats(ThreadID, Stats.Value, "-")
		} else {
			UpdateThreadStats(ThreadID, Stats.Value, "-")
		}
		if Stats.Value != Value {
			UpdateThreadStats(ThreadID, Value, "+")
			Db.Exec("update ThreadStats set Value = $1 where FromUserID = $2 and ToThreadID = $3", Value, UserID, ThreadID)
		} else {
			Db.Exec("delete from ThreadStats where FromUserID = $1 and ToThreadID = $2", UserID, ThreadID)
		}
	} else {
		UpdateThreadStats(ThreadID, Value, "+")
		Db.Exec("insert into ThreadStats(FromUserID,ToThreadID,Value) values($1,$2,$3)", UserID, ThreadID, Value)
	}
}

//GetAllUserCreatedPosts - Get List of Created Threads by User
func GetAllUserCreatedPosts(UserID int) []Thread {
	tmp := []Thread{}
	rows, _ := Db.Query("select * from Thread where UserID = $1", UserID)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		cur := Thread{}
		err := rows.Scan(&cur.Title, &cur.UserID, &cur.Likes, &cur.Dislikes, &cur.ThreadID, &cur.ToThreadID, &cur.Date, &cur.Content, &cur.Category, &cur.Username, &cur.Image)
		if err != nil {
			break
		}
		if cur.Title != "" {
			tmp = append(tmp, cur)
		}
	}
	return tmp
}

//GetAllUserLikedThread - Get List of Liked Threads by User
func GetAllUserLikedThread(UserID int) []Thread {
	tmp := []Thread{}
	rows, _ := Db.Query("select * from Thread where ThreadID > 0")
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		cur := Thread{}
		err := rows.Scan(&cur.Title, &cur.UserID, &cur.Likes, &cur.Dislikes, &cur.ThreadID, &cur.ToThreadID, &cur.Date, &cur.Content, &cur.Category, &cur.Username, &cur.Image)
		cur.Liked = CheckUserLikedThread(UserID, cur.ThreadID)
		if err != nil {
			break
		}
		if cur.Title != "" && cur.Liked == 1 {
			tmp = append(tmp, cur)
		}
	}
	return tmp
}

//GetAllUserLikedComments - Get List of Liked Comments By User
func GetAllUserLikedComments(UserID int) []Thread {
	tmp := []Thread{}
	rows, _ := Db.Query("select * from Thread where ThreadID > 0")
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		cur := Thread{}
		err := rows.Scan(&cur.Title, &cur.UserID, &cur.Likes, &cur.Dislikes, &cur.ThreadID, &cur.ToThreadID, &cur.Date, &cur.Content, &cur.Category, &cur.Username, &cur.Image)
		cur.Liked = CheckUserLikedThread(UserID, cur.ThreadID)
		if err != nil {
			break
		}
		if cur.Title == "" && cur.Liked == 1 {
			tmp = append(tmp, cur)
		}
	}
	return tmp
}

//GetAll - All threads
func GetAll(UserID int) []Thread {
	tmp := []Thread{}
	rows, err := Db.Query("select * from Thread where ThreadID > 0")
	if err != nil {
		return tmp
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	// defer func() {
	// 	if rows != nil {
	// 		rows.Close()
	// 	}
	// }()
	for rows.Next() {
		cur := Thread{}
		err := rows.Scan(&cur.Title, &cur.UserID, &cur.Likes, &cur.Dislikes, &cur.ThreadID, &cur.ToThreadID, &cur.Date, &cur.Content, &cur.Category, &cur.Username, &cur.Image)
		if err != nil {
			break
		}
		cur.Liked = CheckUserLikedThread(UserID, cur.ThreadID)
		if cur.Title != "" {
			tmp = append(tmp, cur)
		}
	}
	return tmp
}
