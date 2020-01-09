package config

import (
	"time"
)

// ###################### Context: main ####################################################
var configFile = configEntry{
	name:         "config-file",
	bindFlag:     true,
	defaultValue: "",
	usage:        "Specifies the full path and name of the configuration file for sokar.",
}

var dryRun = configEntry{
	name:         "dry-run",
	bindFlag:     true,
	bindEnv:      true,
	defaultValue: false,
	usage: "If true, then sokar won't execute the planned scaling action. Only scaling\n" +
		"actions triggered via ScaleBy end-point will be executed.",
}

var port = configEntry{
	name:         "port",
	bindFlag:     true,
	bindEnv:      true,
	defaultValue: 11000,
	usage:        "Port where sokar is listening.",
}

// ###################### Context: scaler ####################################################
var scaMode = configEntry{
	name:         "sca.mode",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "nomad-job",
	usage: "Scaling target mode is either job based, instance-based or data-center\n" +
		"(worker/ instance) based scaling. In data-center (dc) mode the nomad workers\n" +
		"will be scaled. In job mode the number of allocations for this job will be adjusted.",
}

var scaWatcherInterval = configEntry{
	name:         "sca.watcher-interval",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "5s",
	usage: "The interval the Scaler will check if the scalingObject count still matches\n" +
		"the desired state.",
}

// ###################### Context: scaler AWS EC2 ############################################
var scaAWSEC2Profile = configEntry{
	name:         "sca.aws-ec2.profile",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "",
	usage: "This parameter represents the name of the aws profile that shall be used to\n" +
		"access the resources to scale the data-center. This parameter is optional. If it\n" +
		"is empty the instance where sokar runs on has to have enough permissions to access\n" +
		"the resources (ASG) for scaling. In this case the AWSRegion parameter has to be\n" +
		"specified as well.",
}

var scaAWSEC2Region = configEntry{
	name:         "sca.aws-ec2.region",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "eu-central-1",
	usage: "This is an optional parameter and is regarded only if the parameter\n" +
		"AWSProfile is empty. The AWSRegion has to specify the region in which the\n" +
		"data-center to be scaled resides in.",
}

var scaAWSEC2ASGTagKey = configEntry{
	name:         "sca.aws-ec2.asg-tag-key",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "scale-object",
	usage: "This parameter specifies which tag on an AWS AutoScalingGroup shall be used\n" +
		"to find the ASG that should be automatically scaled.",
}

// ###################### Context: scaler Nomad ###############################################
var scaNomadDataCenterAWSProfile = configEntry{
	name:         "sca.nomad.dc-aws.profile",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "",
	usage: "This parameter represents the name of the aws profile that shall be used to\n" +
		"access the resources to scale the data-center. This parameter is optional. If it\n" +
		"is empty the instance where sokar runs on has to have enough permissions to access\n" +
		"the resources (ASG) for scaling. In this case the AWSRegion parameter has to be\n" +
		"specified as well.",
}

var scaNomadDataCenterAWSRegion = configEntry{
	name:         "sca.nomad.dc-aws.region",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "eu-central-1",
	usage: "This is an optional parameter and is regarded only if the parameter\n" +
		"AWSProfile is empty. The AWSRegion has to specify the region in which the data-center\n" +
		"to be scaled resides in.",
}

var scaNomadDataCenterAWSInstanceTerminationTimeout = configEntry{
	name:         "sca.nomad.dc-aws.instance-termination-timeout",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: time.Minute * 10,
	usage: "The maximum time the instance termination will be monitored before assuming\n" +
		"that this action (instance termination due to downscale) failed.",
}

var scaNomadModeServerAddress = configEntry{
	name:         "sca.nomad.server-address",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "",
	usage:        "Specifies the address of the nomad server.",
}

// ###################### Context: scale-object ####################################################
var scaleObjectName = configEntry{
	name:         "scale-object.name",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "",
	usage:        "The name of the object to be scaled.",
}

var scaleObjectMin = configEntry{
	name:         "scale-object.min",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: 1,
	usage:        "The minimum count of the object to be scaled.",
}

var scaleObjectMax = configEntry{
	name:         "scale-object.max",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: 10,
	usage:        "The maximum count of the object to be scaled.",
}

// ###################### Context: CapacityPlanner#########################################
var capScaleSchedule = configEntry{
	name:         "cap.scale-schedule",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "",
	usage: "Specifies time ranges within which it is ensured that the ScaleObject is scaled\n" +
		"to at least min and not more than max. The min/ max values specified in this\n" +
		"schedule have lower priority than the --scale-object.min/ --scale-object.max.\n" +
		"This means the sokar will ensure that the --scale-object.min/ --scale-object.max\n" +
		"are not violated no matter what is specified in the schedule.",
}

var capDownScaleCoolDown = configEntry{
	name:         "cap.down-scale-cooldown",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: time.Second * 20,
	usage:        "The time sokar waits between downscaling actions at min.",
}

var capUpScaleCoolDown = configEntry{
	name:         "cap.up-scale-cooldown",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: time.Second * 10,
	usage:        "The time sokar waits between upscaling actions at min.",
}

var capConstantModeEnable = configEntry{
	name:         "cap.constant-mode.enable",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: true,
	usage: "Enable/ disable the constant mode of the CapacityPlanner. Only one of the\n" +
		"modes can be enabled at the same time.",
}

var capConstantModeOffset = configEntry{
	name:         "cap.constant-mode.offset",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: uint(1),
	usage: "The constant offset value that should be used to increment/ decrement the\n" +
		"count of the scale-object. Only values > 0 are valid.",
}

var capLinearModeEnable = configEntry{
	name:         "cap.linear-mode.enable",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: false,
	usage: "Enable/ disable the linear mode of the CapacityPlanner. Only one of the modes\n" +
		"can be enabled at the same time.",
}

var capLinearModeScaleFactorWeight = configEntry{
	name:         "cap.linear-mode.scale-factor-weight",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: 0.5,
	usage: "This weight is used to adjust the impact of the scaleFactor during capacity\n" +
		"planning in linear mode.",
}

// ###################### Context: Logging ################################################
var loggingStructured = configEntry{
	name:         "logging.structured",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: false,
	usage:        "Use structured logging or not.",
}

var loggingUXTS = configEntry{
	name:         "logging.unix-ts",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: false,
	usage:        "Use Unix-Timestamp representation for log entries.",
}

var loggingNoColor = configEntry{
	name:         "logging.no-color",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: false,
	usage:        "If true colors in log out-put will be disabled.",
}

// ###################### Context: ScaleAlertAggregator ###################################
var saaAlertExpirationTime = configEntry{
	name:         "saa.alert-expiration-time",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: time.Minute * 10,
	usage: "Defines after which time an alert will be pruned if he did not get updated\n" +
		"again by the ScaleAlertEmitter, assuming that the alert is not relevant any more.",
}

var saaNoAlertDamping = configEntry{
	name:         "saa.no-alert-damping",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: 1.0,
	usage: "Damping used in case there are currently no alerts firing\n" +
		"(neither down- nor upscaling).",
}

var saaUpThresh = configEntry{
	name:         "saa.up-thresh",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: 10.0,
	usage:        "Threshold for a upscaling event.",
}
var saaDownThresh = configEntry{
	name:         "saa.down-thresh",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: -10.0,
	usage:        "Threshold for a downscaling event.",
}

var saaEvalCylce = configEntry{
	name:         "saa.eval-cycle",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: time.Second * 1,
	usage: "Cycle/ frequency the ScaleAlertAggregator evaluates the weights of the\n" +
		"currently firing alerts.",
}

var saaEvalPeriodFactor = configEntry{
	name:         "saa.eval-period-factor",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: 10,
	usage: "EvaluationPeriodFactor is used to calculate the evaluation period\n" +
		"(evaluationPeriod = evaluationCycle * evaluationPeriodFactor)",
}

var saaCleanupCylce = configEntry{
	name:         "saa.cleanup-cycle",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: time.Second * 60,
	usage:        "Cycle/ frequency the ScaleAlertAggregator removes expired alerts.",
}

var saaScaleAlerts = configEntry{
	name:         "saa.scale-alerts",
	bindEnv:      true,
	bindFlag:     true,
	defaultValue: "",
	usage:        "Cycle/ frequency the ScaleAlertAggregator removes expired alerts.",
}

var configEntries = []configEntry{
	configFile,
	port,
	dryRun,
	scaleObjectName,
	scaleObjectMin,
	scaleObjectMax,
	capScaleSchedule,
	capDownScaleCoolDown,
	capUpScaleCoolDown,
	loggingStructured,
	loggingUXTS,
	loggingNoColor,
	saaNoAlertDamping,
	saaUpThresh,
	saaDownThresh,
	saaEvalCylce,
	saaEvalPeriodFactor,
	saaCleanupCylce,
	saaScaleAlerts,
	saaAlertExpirationTime,
	scaMode,
	scaWatcherInterval,
	scaAWSEC2Profile,
	scaAWSEC2Region,
	scaAWSEC2ASGTagKey,
	scaNomadDataCenterAWSProfile,
	scaNomadDataCenterAWSRegion,
	scaNomadDataCenterAWSInstanceTerminationTimeout,
	scaNomadModeServerAddress,
	capConstantModeEnable,
	capConstantModeOffset,
	capLinearModeEnable,
	capLinearModeScaleFactorWeight,
}
