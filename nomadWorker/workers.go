package nomadWorker

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// SetScalingObjectCount will scale the nomad workers to the desired count (amount of instances)
func (c *Connector) SetScalingObjectCount(datacenter string, count uint) error {

	sess, err := c.createSession()
	if err != nil {
		return err
	}
	autoScalingIF := c.autoScalingFactory.CreateAutoScaling(sess)

	asgQuery := autoScalingGroupQuery{
		asgIF:    autoScalingIF,
		tagKey:   c.tagKey,
		tagValue: datacenter,
	}

	asgName, err := asgQuery.getAutoScalingGroupName()
	if err != nil {
		return err
	}

	size := int64(count)

	input := &autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: &asgName,
		MaxSize:              &size,
		MinSize:              &size,
		DesiredCapacity:      &size,
	}

	_, err = autoScalingIF.UpdateAutoScalingGroup(input)
	if err != nil {
		return err
	}

	c.log.Info().Msgf("Adjusted min=max=desiredCapacity of %s to %d.", asgName, size)
	return nil
}

// GetScalingObjectCount will return the count of the nomad workers
func (c *Connector) GetScalingObjectCount(datacenter string) (uint, error) {
	sess, err := c.createSession()
	if err != nil {
		return 0, err
	}
	autoScalingIF := c.autoScalingFactory.CreateAutoScaling(sess)

	asgQuery := autoScalingGroupQuery{
		asgIF:    autoScalingIF,
		tagKey:   c.tagKey,
		tagValue: datacenter,
	}
	min, desired, max, err := asgQuery.getScaleNumbers()
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
	asgs, err := getAutoScalingGroups(autoScalingIF)
	if err != nil {
		return true, err
	}

	asg := filterAutoScalingGroupByTag(c.tagKey, datacenter, asgs)

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
