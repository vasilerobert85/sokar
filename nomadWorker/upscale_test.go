package nomadWorker

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mock_aws "github.com/thomasobenaus/sokar/test/aws"
)

func TestUpscale(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	asgFactory := mock_aws.NewMockAutoScalingFactory(mockCtrl)
	asgIF := mock_aws.NewMockAutoScaling(mockCtrl)

	key := "datacenter"
	connector, err := New("http://nomad.io", "profile")
	require.NotNil(t, connector)
	require.NoError(t, err)

	connector.autoScalingFactory = asgFactory

	// error, no numbers
	asgFactory.EXPECT().CreateAutoScaling(gomock.Any()).Return(asgIF)
	asgIF.EXPECT().DescribeAutoScalingGroups(gomock.Any()).Return(nil, nil)
	err = connector.upscale("invalid", 0, 10, 5)
	assert.Error(t, err)

	// no error
	asgFactory.EXPECT().CreateAutoScaling(gomock.Any()).Return(asgIF)
	minCount := int64(1)
	desiredCount := int64(123)
	maxCount := int64(3)
	autoScalingGroups := make([]*autoscaling.Group, 0)
	tagVal := "private-services"
	asgName := "my-asg"
	var tags []*autoscaling.TagDescription
	td := autoscaling.TagDescription{Key: &key, Value: &tagVal}
	tags = append(tags, &td)
	asgIn := autoscaling.Group{
		Tags:                 tags,
		AutoScalingGroupName: &asgName,
		MinSize:              &minCount,
		MaxSize:              &maxCount,
		DesiredCapacity:      &desiredCount,
	}
	autoScalingGroups = append(autoScalingGroups, &asgIn)
	output := &autoscaling.DescribeAutoScalingGroupsOutput{AutoScalingGroups: autoScalingGroups}
	asgIF.EXPECT().DescribeAutoScalingGroups(gomock.Any()).Return(output, nil)

	minCount = int64(1)
	desiredCount = int64(5)
	maxCount = int64(10)
	input := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: &asgName,
		MinSize:              &minCount,
		MaxSize:              &maxCount,
		DesiredCapacity:      &desiredCount,
	}
	asgIF.EXPECT().UpdateAutoScalingGroup(input)
	err = connector.upscale(tagVal, uint(minCount), uint(maxCount), uint(desiredCount))
	assert.NoError(t, err)
}
