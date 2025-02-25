package mocking_test

import (
	"testing"
	"time"

	// blackbox тестирование
	"github.com/VitaminP8/go-practice/start/advanced_testing/mocking"
	"github.com/VitaminP8/go-practice/start/advanced_testing/mocking/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//type GeoMock struct {
//	mock.Mock
//}
//
//func (gm GeoMock) GetCoordsByName(city string) (latitude, longitude float64, err error) {
//	return 0, 0, nil
//}

func TestMainFile(t *testing.T) {
	gm := &mocks.Geocoder{}
	// говорит мокеру что эта функция может быть вызвана в тесте! (без этого будет ошибка)
	gm.On("GetCoordsByName", "Ekaterinburg").Return(45.0, 45.0, nil)
	// более удобный вариант
	//gm.EXPECT().GetCoordsByName("Ekaterinburg2").Return(45, 45, nil)

	rise, set, err := mocking.CalcSunrise("Ekaterinburg", gm)
	require.NoError(t, err)
	assert.Equal(t, time.Date(2000, time.January, 1, 4, 38, 13, 0, time.UTC), rise)
	assert.Equal(t, time.Date(2000, time.January, 1, 13, 28, 2, 0, time.UTC), set)
}
