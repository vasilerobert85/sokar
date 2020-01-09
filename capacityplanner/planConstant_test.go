package capacityplanner

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PlanConstant(t *testing.T) {
	// GIVEN
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	metrics, _ := NewMockedMetrics(mockCtrl)

	// WHEN
	cap, err := New(metrics)

	// THEN
	require.NotNil(t, cap)
	require.NoError(t, err)

	// verify default behavior ... constant planning
	assert.Equal(t, uint(0), cap.planConstant(-1, 0, 1))
	assert.Equal(t, uint(0), cap.planConstant(-1, 1, 1))
	assert.Equal(t, uint(0), cap.planConstant(-1, 1, 2))
	assert.Equal(t, uint(3), cap.planConstant(-1, 5, 2))

	assert.Equal(t, uint(0), cap.planConstant(1, 0, 0))
	assert.Equal(t, uint(1), cap.planConstant(1, 0, 1))
	assert.Equal(t, uint(2), cap.planConstant(1, 1, 1))
	assert.Equal(t, uint(3), cap.planConstant(1, 1, 2))
	assert.Equal(t, uint(7), cap.planConstant(1, 5, 2))

	assert.Equal(t, uint(0), cap.planConstant(0, 0, 1))
	assert.Equal(t, uint(2), cap.planConstant(0, 2, 1))
}
