package capacityPlanner

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {

	// err - empty
	err := validate(CapacityPlanner{})
	assert.Error(t, err)

	// err - two modes
	constMode := ConstantMode{}
	linearMode := LinearMode{}
	err = validate(CapacityPlanner{constantMode: &constMode, linearMode: &linearMode})
	assert.Error(t, err)

	// err - invalid linear mode
	linearMode.ScaleFactorWeight = 0
	err = validate(CapacityPlanner{linearMode: &linearMode})
	assert.Error(t, err)

	// success
	linearMode.ScaleFactorWeight = 1
	err = validate(CapacityPlanner{linearMode: &linearMode})
	assert.NoError(t, err)
}

func TestCooldowns(t *testing.T) {
	capa, err := New(WithDownScaleCooldown(time.Second * 123))
	assert.NoError(t, err)
	assert.NotNil(t, capa)
	assert.Equal(t, time.Second*123, capa.downScaleCooldownPeriod)

	capa, err = New(WithUpScaleCooldown(time.Second * 234))
	assert.NoError(t, err)
	assert.NotNil(t, capa)
	assert.Equal(t, time.Second*234, capa.upScaleCooldownPeriod)
}

func Test_New(t *testing.T) {
	capa, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, capa)

	capa, err = New(UseLinearMode(0))
	assert.Error(t, err)
	assert.Nil(t, capa)

	capa, err = New(UseConstantMode(0))
	assert.Error(t, err)
	assert.Nil(t, capa)
}

func Test_Plan_ModeLinear(t *testing.T) {
	capa, err := New(UseLinearMode(0.5))
	require.NoError(t, err)
	require.NotNil(t, capa)

	assert.Equal(t, uint(10), capa.Plan(0, 10))

	assert.Equal(t, uint(1), capa.Plan(1, 0))
	assert.Equal(t, uint(15), capa.Plan(1, 10))
	assert.Equal(t, uint(2), capa.Plan(0.5, 1))
	assert.Equal(t, uint(13), capa.Plan(0.5, 10))

	assert.Equal(t, uint(0), capa.Plan(-1, 0))
	assert.Equal(t, uint(5), capa.Plan(-1, 10))
	assert.Equal(t, uint(0), capa.Plan(-0.5, 1))
	assert.Equal(t, uint(8), capa.Plan(-0.5, 10))
}

func Test_Plan_ModeConstant(t *testing.T) {
	capa, err := New()
	require.NotNil(t, capa)
	require.NoError(t, err)
	assert.Equal(t, uint(0), capa.Plan(-1, 0))
	assert.Equal(t, uint(0), capa.Plan(-1, 1))
	assert.Equal(t, uint(0), capa.Plan(0, 0))
	assert.Equal(t, uint(2), capa.Plan(0, 2))
	assert.Equal(t, uint(1), capa.Plan(1, 0))
	assert.Equal(t, uint(2), capa.Plan(1, 1))

	capa, err = New(UseConstantMode(2))
	require.NotNil(t, capa)
	require.NoError(t, err)
	assert.Equal(t, uint(0), capa.Plan(-1, 1))
	assert.Equal(t, uint(3), capa.Plan(-1, 5))
	assert.Equal(t, uint(3), capa.Plan(1, 1))
	assert.Equal(t, uint(7), capa.Plan(1, 5))
}

func Test_IsCoolingDown(t *testing.T) {
	capa, err := New(WithDownScaleCooldown(time.Second*20), WithUpScaleCooldown(time.Second*10))
	require.NoError(t, err)
	require.NotNil(t, capa)

	lastScale := time.Now()
	result := capa.IsCoolingDown(lastScale, false)
	assert.True(t, result)

	result = capa.IsCoolingDown(lastScale, true)
	assert.True(t, result)

	// Upscaling
	lastScale = time.Now().Add(time.Second * -11)
	result = capa.IsCoolingDown(lastScale, false)
	assert.False(t, result)

	lastScale = time.Now().Add(time.Second * -9)
	result = capa.IsCoolingDown(lastScale, false)
	assert.True(t, result)

	// Downscaling
	lastScale = time.Now().Add(time.Second * -21)
	result = capa.IsCoolingDown(lastScale, true)
	assert.False(t, result)

	lastScale = time.Now().Add(time.Second * -19)
	result = capa.IsCoolingDown(lastScale, true)
	assert.True(t, result)
}
