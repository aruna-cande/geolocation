package service

import (
	"Geolocation/internal/pkg/geolocation/service/mock"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"path"
	"runtime"
	"testing"
)

func getCsvData() (*http.Response, error){
	csvData := `i_addres,country_code,country,city,latitude,longitude,mystery_value
	200.106.141.15,SI,Nepal,DuBuquemouth,-84.87503094689836,7.206435933364332,7823011346
	160.103.7.140,CZ,Nicaragua,New Neva,-68.31023296602508,-37.62435199624531,7301823115
	70.95.73.73,TL,Saudi Arabia,Gradymouth,-49.16675918861615,-86.05920084416894,2559997162
	,PY,Falkland Islands (Malvinas),,75.41685191518815,-144.6943217219469,0
	125.159.20.54,LI,Guyana,Port Karson,-78.2274228596799,-163.26218895343357,1337885276`

	t := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(csvData)),
	}

	return t, nil
}

func Test_ImportGeolocationData(t *testing.T) {
	type test struct {
		dumpFile string
		valid int64
		invalid int64
	}
	tests := []test{
		{dumpFile: "/testResources/data_dump.csv", valid: 5, invalid:   0},
		{dumpFile: "/testResources/data_dump_invalid_empty_country.csv", valid: 4, invalid:   1},
		{dumpFile : "/testResources/data_dump_invalid_ip.csv", valid: 3, invalid:   2},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepository(ctrl)

	for _, test := range tests{
		//restCliet := mock.NewMockHTTPClient(ctrl)
		//request, _ := http.NewRequest(http.MethodPost, "url", nil)
		repo.EXPECT().AddGeolocation(gomock.Any()).Return(test.valid, nil)

		_, filename, _, _ := runtime.Caller(0)
		csvTestFile := path.Join(path.Dir(filename), test.dumpFile)
		srv := NewImporterService(repo)
		statistics, err := srv.ImportGeolocationData(csvTestFile)
		assert.Nil(t, err)
		assert.Equal(t, statistics.Accepted, test.valid)
		assert.Equal(t, statistics.Discarded, test.invalid)
	}

}