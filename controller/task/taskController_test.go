package task

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetAllTasks(t *testing.T) {

	//db := repository.GetInstance()
	//ur := userRepo.NewSqliteRepository(db)
	//tr := taskRepo.NewSqliteRepository(db)
	//
	//userService := service.NewUserService(ur)
	//taskService := service.NewTaskService(tr, userService)
	//taskController := NewTaskController(taskService)
	//
	//// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	//// pass 'nil' as the third parameter.
	//req, err := http.NewRequest("GET", "/tasks", nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(getAllTasks(taskController))
	//
	//// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	//// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, req)
	//
	//// Check the status code is what we expect.
	//if status := rr.Code; status != http.StatusOK {
	//	t.Errorf("handler returned wrong status code: got %v want %v",
	//		status, http.StatusOK)
	//}
	//
	//// Check the response body is what we expect.
	//expected := `{"alive": true}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}

	log.Print("haiii")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	open := sqlite.Open("gorm.db")
	db, err := gorm.Open(open, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Print("oops")
		log.Fatal(err)
	}

	log.Print("its open now")

	if db != nil {
		t.Error("handler returned unexpected body: got want")
	}
}
