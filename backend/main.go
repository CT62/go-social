package main

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

var db *gorm.DB

type User struct {
    gorm.Model
    Username     string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
}

type Post struct {
    gorm.Model
    UserID   uint   `gorm:"not null"`
    User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    ImageURL string `gorm:"not null"`
    Caption  string
    Title    string
    Likes    []Like `gorm:"foreignKey:PostID"`
}

type Like struct {
    gorm.Model
    UserID uint `gorm:"not null"`
    User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    PostID uint `gorm:"not null;column:post_id"`
}

func Users(c *fiber.Ctx) error {
    var users []User
    db.Find(&users)
    return c.JSON(users)
}

func Signup(c *fiber.Ctx) error {
    var user User
    if err := c.BodyParser(&user); err != nil {
        return err
    }
    usr :=db.Find(&user, "username =",user.Username)
    if usr {

    }
    //err := db.Create(&User{Username:user.Username, PasswordHash:user.PasswordHash})
    return c.JSON(user)
}


func main() {
    app := fiber.New()

    var err error
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{}, &Post{}, &Like{})


    app.Get("/users", Users)
    app.Post("/signup", Signup)

    app.Listen(":3000")
}
