package service

import (
	"database/sql"
	"log"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/aruna-cande/geolocation/pkg/geolocation/service/mock"
)

func TestImporterService_ImportGeolocationData(t *testing.T) {
	type test struct {
		dumpFile  string
		accepted  int64
		discarded int64
		err       error
	}
	tests := []test{
		//{dumpFile: "/testdata/data_dump.csv", accepted: 5, discarded: 0, err: nil},
		{dumpFile: "/testdata/data_dump.csv", accepted: 0, discarded: 5, err: sql.ErrConnDone},
		{dumpFile: "/testdata/data_dump_duplicated_data.csv", accepted: 1, discarded: 4, err: nil},
		{dumpFile: "/testdata/data_dump_invalid_empty_country.csv", accepted: 4, discarded: 1, err: nil},
		{dumpFile: "/testdata/data_dump_invalid_ip.csv", accepted: 3, discarded: 2, err: nil},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepository(ctrl)
	logger := log.New(os.Stderr, "logger: ", log.Ldate)

	for _, test := range tests {
		repo.EXPECT().AddGeolocation(gomock.Any()).Return(test.accepted, test.err)

		_, filename, _, _ := runtime.Caller(0)
		csvTestFile := path.Join(path.Dir(filename), test.dumpFile)
		srv := NewImporterService(repo, logger)
		statistics, err := srv.ImportGeolocationData(csvTestFile)

		assert.Nil(t, err)
		assert.Equal(t, statistics.Accepted, test.accepted)
		assert.Equal(t, statistics.Discarded, test.discarded)
	}

}
