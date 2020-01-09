package capacityplanner

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PlanLinear(t *testing.T) {
	// GIVEN
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	metrics, _ := NewMockedMetrics(mockCtrl)

	// WHEN
	cap, err := New(metrics, UseLinearMode(0.5))

	// THEN
	require.NotNil(t, cap)
	require.NoError(t, err)

	assert.Equal(t, uint(10), cap.planLinear(0, 10))

	assert.Equal(t, uint(1), cap.planLinear(1, 0))
	assert.Equal(t, uint(15), cap.planLinear(1, 10))
	assert.Equal(t, uint(2), cap.planLinear(0.5, 1))
	assert.Equal(t, uint(13), cap.planLinear(0.5, 10))

	assert.Equal(t, uint(0), cap.planLinear(-1, 0))
	assert.Equal(t, uint(5), cap.planLinear(-1, 10))
	assert.Equal(t, uint(0), cap.planLinear(-0.5, 1))
	assert.Equal(t, uint(8), cap.planLinear(-0.5, 10))
}
