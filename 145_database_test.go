import (
	"database/sql"
	"time"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}


func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	defer.close()
}