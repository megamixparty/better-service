package customer

import (
	"errors"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateCustomer(t *testing.T) {
	type args struct {
		req CustomerDetail
	}
	tests := []struct {
		name    string
		args    args
		wantID  int
		wantErr bool
	}{
		{
			name: "success",
			args: args{CustomerDetail{
				Name:  "Linda",
				Phone: "08988877064",
				Addresses: []Address{
					{
						Address: "Jalan kenangan indah No. 10",
						ZipCode: "12345",
					},
				},
			}},
			wantID:  1,
			wantErr: false,
		},
		{
			name: "failed",
			args: args{CustomerDetail{
				Name:  "Linda",
				Phone: "08988877064",
				Addresses: []Address{
					{
						Address: "Jalan kenangan indah No. 10",
						ZipCode: "12345",
					},
				},
			}},
			wantID:  0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectBegin()
			prepCustomer := mock.ExpectQuery("^INSERT INTO customers \\(name, phone\\) VALUES(.+)")
			prepCustomerAddress := mock.ExpectExec("^INSERT INTO customer_addresses \\(customer_id, address, zipcode\\) VALUES(.+)")

			if tt.wantErr {
				prepCustomer.WillReturnError(errors.New("Error from mysql"))
			} else {
				prepCustomer.WillReturnRows(mock.NewRows([]string{"id"}).FromCSVString(strconv.Itoa(tt.wantID)))
				prepCustomerAddress.WillReturnResult(sqlmock.NewResult(2, 1))
			}
			mock.ExpectCommit()

			gotID, err := CreateCustomer(db, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("CreateCustomer() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}
