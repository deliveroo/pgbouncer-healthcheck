package main

//import (
//	"context"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//type mockSystemAPI struct {
//	mock.Mock
//	sql.DB
//}
//
//func (d *mockSystemAPI) Info(ctx context.Context) (types.Info, error) {
//	args := d.Called(ctx)
//
//	return args.Get(0).(types.Info), args.Error(1)
//}
//
//func TestHealthyDockerDeviceMapper(t *testing.T) {
//	info := types.Info{
//		Driver: "devicemapper",
//	}
//
//	m := new(mockSystemAPI)
//	m.On("Info", mock.MatchedBy(func(input interface{}) bool {
//		_, ok := input.(context.Context)
//		return ok
//	})).Return(info, nil)
//
//	assert.True(
//		t,
//		checkDockerHealth(context.Background(), m),
//		"expected healthy docker for devicemapper",
//	)
//}
//
//// For our purposes, if the docker daemon is using the overlay2 storage driver
//// something went wrong and we should not treat this host as healthy
//func TestUnHealthyDockerOverlay2(t *testing.T) {
//	info := types.Info{
//		Driver: "overlay2",
//	}
//
//	m := new(mockSystemAPI)
//	m.On("Info", mock.MatchedBy(func(input interface{}) bool {
//		_, ok := input.(context.Context)
//		return ok
//	})).Return(info, nil)
//
//	assert.False(
//		t,
//		checkDockerHealth(context.Background(), m),
//		"expected unhealthcy docker for overlay2",
//	)
//}
//
//func TestEnoughMetadataAvailbe(t *testing.T) {
//	info := types.Info{
//		Driver:       "devicemapper",
//		DriverStatus: [][2]string{{"Metadata Space Available", "90MB"}},
//	}
//
//	m := new(mockSystemAPI)
//	m.On("Info", mock.MatchedBy(func(input interface{}) bool {
//		_, ok := input.(context.Context)
//		return ok
//	})).Return(info, nil)
//
//	assert.True(
//		t,
//		checkDockerHealth(context.Background(), m),
//		"expected enough metadata space available",
//	)
//}
//
//func TestNotEnoughMetadataAvailable15m(t *testing.T) {
//	info := types.Info{
//		Driver:       "devicemapper",
//		DriverStatus: [][2]string{{"Metadata Space Available", "15MB"}},
//	}
//
//	m := new(mockSystemAPI)
//	m.On("Info", mock.MatchedBy(func(input interface{}) bool {
//		_, ok := input.(context.Context)
//		return ok
//	})).Return(info, nil)
//
//	assert.False(
//		t,
//		checkDockerHealth(context.Background(), m),
//		"expected not enough metadata space available",
//	)
//
//}
//
//func TestNotEnoughMetadataAvailable100k(t *testing.T) {
//	info := types.Info{
//		Driver:       "devicemapper",
//		DriverStatus: [][2]string{{"Metadata Space Available", "100KB"}},
//	}
//
//	m := new(mockSystemAPI)
//	m.On("Info", mock.MatchedBy(func(input interface{}) bool {
//		_, ok := input.(context.Context)
//		return ok
//	})).Return(info, nil)
//
//	assert.False(
//		t,
//		checkDockerHealth(context.Background(), m),
//		"expected not enough metadata space available",
//	)
//
//}
