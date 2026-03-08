package service

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/aruna-cande/geolocation/pkg/geolocation/service/mock"
)

func TestImporterService_ImportGeolocationData(t *testing.T) {
	tests := []struct {
		name      string
		dumpFile  string
		accepted  int64
		discarded int64
		err       error
	}{
		{name: "DBError", dumpFile: "/testdata/data_dump.csv", accepted: 0, discarded: 5, err: sql.ErrConnDone},
		{name: "DuplicatedData", dumpFile: "/testdata/data_dump_duplicated_data.csv", accepted: 1, discarded: 4, err: nil},
		{name: "InvalidEmptyCountry", dumpFile: "/testdata/data_dump_invalid_empty_country.csv", accepted: 4, discarded: 1, err: nil},
		{name: "InvalidIP", dumpFile: "/testdata/data_dump_invalid_ip.csv", accepted: 3, discarded: 2, err: nil},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepository(ctrl)
	logger := log.New(os.Stderr, "logger: ", log.Ldate)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.EXPECT().AddGeolocation(gomock.Any(), gomock.Any()).Return(tc.accepted, tc.err)

			_, filename, _, _ := runtime.Caller(0)
			csvTestFile := path.Join(path.Dir(filename), tc.dumpFile)
			srv := NewImporterService(repo, logger)
			statistics, err := srv.ImportGeolocationData(context.Background(), csvTestFile)

			assert.Nil(t, err)
			assert.Equal(t, tc.accepted, statistics.Accepted)
			assert.Equal(t, tc.discarded, statistics.Discarded)
		})
	}
}
