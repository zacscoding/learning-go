package main

// https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/models.go

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type runTest func(db *gorm.DB) error

type TestSet struct {
	test  runTest
	title string
}

func main() {
	testSets := []TestSet{
		{
			test:  testUser(),
			title: "Test user",
		},
	}

	db, err := gorm.Open("mysql", "root:password@(192.168.79.130)/my_database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)
	// setup
	db.DropTable(&Follow{}, &User{})
	db.AutoMigrate(&Follow{}, &User{})

	for _, testSet := range testSets {
		fmt.Println("==========================")
		fmt.Println(testSet.title)
		fmt.Println("Error :", testSet.test(db))
		fmt.Println("--------------------------")
		fmt.Println("==========================")
	}
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email;unique_index"`
	Username string `json:"username" gorm:"column:username;unique_index"`
	Password string `json:"password" gorm:"column:password; not null"`
	Bio      string `json:"bio" gorm:"column:bio;size:1024"`
	Image    string `json:"image"`
}

type Follow struct {
	gorm.Model
	Follower    User
	FollowerID  uint
	Following   User
	FollowingID uint
}

func testUser() runTest {
	return func(db *gorm.DB) error {
		// user1, user2, user3 저장
		u1 := User{
			Email:    "user1@gmail.com",
			Username: "user1",
			Password: "password1",
			Bio:      "user1's profile",
		}
		u2 := User{
			Email:    "user2@gmail.com",
			Username: "user2",
			Password: "password2",
			Bio:      "user2's profile",
		}
		u3 := User{
			Email:    "user3@gmail.com",
			Username: "user3",
			Password: "password3",
			Bio:      "user3's profile",
		}
		// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`username`,`password`,`bio`,`image`) VALUES ('2020-04-27 22:34:23','2020-04-27 22:34:23',NULL,'user1@gmail.com','user1','password1','user1's profile','')
		db.Create(&u1)
		db.Create(&u2)
		db.Create(&u3)

		// follows
		// user1 -> user2
		// user3 -> user2
		// user2 -> user3
		db.Create(&Follow{
			FollowerID:  u2.ID,
			FollowingID: u1.ID,
		})
		db.Create(&Follow{
			FollowerID:  u2.ID,
			FollowingID: u3.ID,
		})
		db.Create(&Follow{
			FollowerID:  u3.ID,
			FollowingID: u2.ID,
		})

		// 1) findUserByEmail
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`email` = 'user1@gmail.com')) ORDER BY `users`.`id` ASC LIMIT 1
		var find1 User
		fmt.Println("## findUserByEmail. email:", u1.Email)
		_ = db.Where(&User{Email: u1.Email}).First(&find1).Error
		b, _ := json.Marshal(find1)
		fmt.Println("result :", string(b))

		// 2) findUserByUsername
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((username)) ORDER BY `users`.`id` ASC LIMIT 1'user2'
		var find2 User
		fmt.Println("## findUserByUsername. username:", u2.Username)
		_ = db.Where("username", u2.Username).First(&find2).Error
		b, _ = json.Marshal(find2)
		fmt.Println("result :", string(b))

		// 3) user2's follower and following
		// 3-1) user2's follower count
		// TODO : refactor
		fmt.Println("## countFollowers")
		var followerCount, followingCount int
		db.Where(&Follow{FollowerID: u2.ID}).Find(&Follow{}).Count(&followerCount)
		db.Model(&Follow{}).Where(&Follow{FollowingID: u2.ID}).Count(&followingCount)
		fmt.Println("==> Followers #", followerCount, ", Following : #", followingCount)

		// TODO :
		// 3-3) user2's followers
		// 3-4) user2's following

		// 4) unfollow
		fmt.Println("## Unfollow user1 to user2")
		db.Where(Follow{
			FollowerID:  u2.ID,
			FollowingID: u1.ID,
		}).Delete(Follow{})
		db.Model(&Follow{}).Where(&Follow{FollowingID: u2.ID}).Count(&followingCount)
		fmt.Println("==> user2's following count after unfollow :", followingCount)

		// 5) check user1 follow to user2
		var count int
		db.Where(Follow{
			FollowerID:  u2.ID,
			FollowingID: u1.ID,
		}).Count(&count)
		if count == 1 {
			fmt.Println("user1 follow user2")
		} else {
			fmt.Println("user1 don't follow user2")
		}
		return nil
	}
}
