package awsEc2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/thomasobenaus/sokar/aws"
)

// AdjustScalingObjectCount will scale the AutoScalingGroup to the desired count (amount of instances)
func (c *Connector) AdjustScalingObjectCount(autoScalingGroup string, min uint, max uint, from uint, to uint) error {
	sess, err := c.createSession()
	if err != nil {
		return err
	}
	autoScalingIF := c.autoScalingFactory.CreateAutoScaling(sess)

	asgQuery := aws.AutoScalingGroupQuery{
		AsgIF:    autoScalingIF,
		TagKey:   c.tagKey,
		TagValue: autoScalingGroup,
	}

	asgName, err := asgQuery.GetAutoScalingGroupName()
	if err != nil {
		return err
	}

	size := int64(to)
	minSize := int64(min)
	maxSize := int64(max)

	input := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: &asgName,
		MinSize:              &minSize,
		MaxSize:              &maxSize,
		DesiredCapacity:      &size,
	}

	_, err = autoScalingIF.UpdateAutoScalingGroup(input)
	if err != nil {
		return err
	}

	c.log.Info().Msgf("Adjusted min=max=desiredCapacity of %s from %d to %d.", asgName, from, size)
	return nil
}

// GetScalingObjectCount will return the count of the nomad workers
func (c *Connector) GetScalingObjectCount(datacenter string) (uint, error) {
	sess, err := c.createSession()
	if err != nil {
		return 0, err
	}
	autoScalingIF := c.autoScalingFactory.CreateAutoScaling(sess)

	asgQuery := aws.AutoScalingGroupQuery{
		AsgIF:    autoScalingIF,
		TagKey:   c.tagKey,
		TagValue: datacenter,
	}
	min, desired, max, err := asgQuery.GetScaleNumbers()
	if err != nil {
		return 0, err
	}
	c.log.Info().Msgf("Current scale numbers min=%d, desired=%d, max=%d.", min, desired, max)
	return desired, nil
}

// IsScalingObjectDead will return if the nomad workers of the actual data-center are still available.
func (c *Connector) IsScalingObjectDead(datacenter string) (bool, error) {

	// Currently the nomad worker is assumed to be dead in case the according AutoScalingGroup can't be found.
	sess, err := c.createSession()
	if err != nil {
		return true, err
	}
	autoScalingIF := c.autoScalingFactory.CreateAutoScaling(sess)
	asgs, err := aws.GetAutoScalingGroups(autoScalingIF)
	if err != nil {
		return true, err
	}

	asg := aws.FilterAutoScalingGroupByTag(c.tagKey, datacenter, asgs)

	if asg == nil {
		c.log.Warn().Msgf("AutoScalingGroup with %s=%s not found", c.tagKey, datacenter)
		return true, nil
	}

	c.log.Debug().Msgf("AutoScalingGroup with %s=%s found", c.tagKey, datacenter)
	return false, nil
}

// createSession wrapper function to create the aws session.
// Internally based on the awsprofile set or not the appropriate method is used.
func (c *Connector) createSession() (*session.Session, error) {
	var session *session.Session
	var err error
	if len(c.awsProfile) > 0 {
		c.log.Debug().Msgf("Create aws session based on profile '%s'", c.awsProfile)

		if c.fnCreateSessionFromProfile == nil {
			return nil, fmt.Errorf("fnCreateSessionFromProfile is nil")
		}

		session, err = c.fnCreateSessionFromProfile(c.awsProfile)
	} else {
		c.log.Debug().Msg("Create aws session based on default credentials")

		if c.fnCreateSession == nil {
			return nil, fmt.Errorf("fnCreateSession is nil")
		}

		session, err = c.fnCreateSession(c.awsRegion)
	}
	return session, err
}
