package repository

import (
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// const for code
const (
	SuccessCode        = 200
	ContinueCode       = 100
	UndefinedCode      = 500
	BadRequestCode     = 400
	NotFoundCode       = 404
	UnauthorizedCode   = 401
	PendingCode        = 451
	DuplicateEntryCode = 462
)

// const for Status
const (
	SuccessStatus   = "SUCCESS"
	PendingStatus   = "PENDING"
	FailedStatus    = "FAILED"
	UndefinedStatus = "FAILED"
	ContinueStatus  = "CONTINUE"
	ErrorStatus     = "ERROR"
)

var (
	PendingErr        = Error.NewError(PendingCode, PendingStatus, "Transaksi Sedang Diproses. Jika transaksi gagal dana Anda akan dikembalikan ke saldo OttoCash")
	UndefinedErr      = Error.NewError(UndefinedCode, ErrorStatus, "Terjadi Kesalahan Pada Server")
	ContinueErr       = Error.NewError(ContinueCode, ContinueStatus, "Silahkan Lanjutkan ke Tahap Berikutnya")
	UnauthorizedErr   = Error.NewError(UnauthorizedCode, FailedStatus, "Sesi Anda Telah Habis")
	NotFoundErr       = Error.NewError(NotFoundCode, FailedStatus, "Data Tidak Ditemukan")
	BadRequestErr     = Error.NewError(BadRequestCode, FailedStatus, "Format Request Salah")
	DuplicateEntryErr = Error.NewError(DuplicateEntryCode, FailedStatus, "Data Sudah ada")
)
var listError = []error{
	PendingCode:        PendingErr,
	ContinueCode:       ContinueErr,
	UndefinedCode:      UndefinedErr,
	UnauthorizedCode:   UnauthorizedErr,
	NotFoundCode:       NotFoundErr,
	BadRequestCode:     BadRequestErr,
	DuplicateEntryCode: DuplicateEntryErr,
}

func SetError(code int, message ...string) error {
	if code == SuccessCode {
		return nil
	}
	defaultMessage := UndefinedErr.Error()
	if code >= len(listError) {
		m := defaultMessage
		if len(message) > 0 && message[0] != "" {
			m = message[0]
		}
		return Error.NewError(UndefinedCode, ErrorStatus, m)
	}
	errFromList := listError[code]
	if errFromList != nil {
		if len(message) > 0 {
			m := message[0]
			if he, ok := errFromList.(*Error.ApplicationError); ok {
				if m == "" {
					m = he.Message
				}
				return Error.NewError(he.ErrorCode, he.Status, m)
			} else {
				return errFromList
			}
		} else {
			return errFromList
		}
	} else {
		m := defaultMessage
		if len(message) > 0 && message[0] != "" {
			m = message[0]
		}
		return Error.NewError(UndefinedCode, ErrorStatus, m)
	}
}

// Convert and Mapping  MySql Error
func HandleSqlError(err error) error {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			println("1")
			return Error.NewError(NotFoundCode, FailedStatus)
		}
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			println("2")
			return MapMysqlError(mysqlErr)
		} else if pqErr, okp := err.(*pq.Error); okp {
			println("3")
			return MapPostgresError(pqErr)
		} else {
			println("4")
			return err
		}
	}
	return nil
}

// Mapping  MySql Error
func MapMysqlError(mySqlErr *mysql.MySQLError) (err error) {
	switch mySqlErr.Number {
	case 1062: // MySQL code for duplicate entry
		return DuplicateEntryErr
	default:
		return UndefinedErr
	}
}

func MapPostgresError(pqErr *pq.Error) (err error) {
	println(pqErr.Code.Name())
	switch pqErr.Code.Name() {
	case "unique_violation":
		fmt.Println("Error: Unique constraint violation")
	case "foreign_key_violation":
		fmt.Println("Error: Foreign key violation")
	case "null_value_not_allowed":
		fmt.Println("Error: Null value not allowed")
	case "check_violation":
		fmt.Println("Error: Check constraint violation")
	case "string_data_right_truncation":
		fmt.Println("Error: String data right truncation")
	default:
		fmt.Printf("PostgreSQL Error [%s]: %s\n", pqErr.Code, pqErr.Message)
	}

	return
}
