package nomadWorker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnector(t *testing.T) {

	// AWSProfile and AWSRegion not specified
	connector, err := New("http://nomad.io", "")
	assert.Nil(t, connector)
	assert.Error(t, err)

	// NomadSrvAddr missing
	connector, err = New("", "profile")
	assert.Nil(t, connector)
	assert.Error(t, err)

	// Success
	connector, err = New("http://nomad.io", "profile")
	assert.NotNil(t, connector)
	assert.NoError(t, err)
}
