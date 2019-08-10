package repositories

import (
	"database/sql"
	"go-web/app/models"
	"go-web/core/ext"
)

type IBookRepository interface {
	Conn() error
	Insert(*models.Book) (int64, error)
	Delete(int64) bool
	Update(*models.Book) error
	SelectByKey(int64) (*models.Book, error)
	SelectAll() ([]*models.Book, error)
	SelectAllWithInfo() (map[int]map[string]string, error)
}

//func NewBookMangerRepository(table string,sql *sql.DB) IBookRepository  {
//	//return &BookMangerRepository{table:"book"}
//}
type BookMangerRepository struct {
	table string
}

var mysqlConn *sql.DB

func init() {
	if mysqlConn == nil {
		mysql, err := ext.NewMysqlConn()
		if err != nil {
			//return err
		}
		mysqlConn = mysql
	}
}

//func Conn() (err error)  {
//	if mysqlConn ==nil {
//		mysql,err := ext.NewMysqlConn()
//		if err !=nil {
//			return err
//		}
//		mysqlConn=mysql
//	}
//
//	return nil
//}
func (b *BookMangerRepository) Insert(book *models.Book) (productID int64, err error) {

	sql := "INSERT " + b.table + " set userID=?,productID=?,orderStatus=?"
	stmt, err := mysqlConn.Prepare(sql)
	if err != nil {
		return
	}

	result, err := stmt.Exec(book.ID, book.Name)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

//func (b *BookMangerRepository)Delete (int64) (isOk bool)  {
//
//}
//func (b *BookMangerRepository)Update (*models.Book) error  {
//
//}
//func (b *BookMangerRepository)SelectByKey(int64) (*models.Book,error)  {
//
//}
//func (b *BookMangerRepository) SelectAll()([]*models.Book,error)  {
//
//}
//func (b *BookMangerRepository)SelectAllWithInfo (map[int]map[string]string,error)  {
//
//}
