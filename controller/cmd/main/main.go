package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yochaipagi/dataBaseProject/controller/api"
	"github.com/yochaipagi/dataBaseProject/controller/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: " + err.Error())
	}

	err = database.SetupDB()
	if err != nil {
		log.Fatal("error setting up db: " + err.Error())
	}

	router := api.SetupRouter()
	err = router.Run(":9090")
	if err != nil {
		log.Fatal("error starting the server: " + err.Error())
	}
}

//דברים לעז:
//בקשת פוסט לבנצ׳מארקינג
//איפה בםועל כתוב הבנאצמארק מה בדיוק עשית איפה תהליך שינוי האינדקס אם קיים לפי ההבנה שלנו פשוט שיכפלת את הטבלה והבאת כתבה רנדומלית אנחנו טועים?
//איפה נוצר הקישור בין הוורד לוורד_גרופ
//איפה נמצא הניתוח של הנתונים הסטטיסטים עבור כל מילה
