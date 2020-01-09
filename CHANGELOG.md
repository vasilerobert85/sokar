# Changelog

## v0.0.12 (2019-12-31)

- Capacity Planning: With [#110](https://github.com/ThomasObenaus/sokar/issues/110) the new feature **Scheduled Scaling** was added to sokar. For more details read [ScheduledScaling.md](doc/ScheduledScaling.md).

## v0.0.11 (2019-11-03)

- BugFixes:
  - [#106 Downscaling of AWS instances fails (Throttling: Rate exceeded)](https://github.com/ThomasObenaus/sokar/issues/106)
  - [#108 Instance Downscaling does not complete](https://github.com/ThomasObenaus/sokar/issues/108)
- Config: With [#108](https://github.com/ThomasObenaus/sokar/issues/108) the new configuration parameter `instance-termination-timeout` was introduced.

## v0.0.10 (2019-10-13)

- Config: With [#103](https://github.com/ThomasObenaus/sokar/issues/103) the deprecated parameter `sca.nomad.mode` was removed. That parameter was replaced by `sca.mode` in a previous release.
- Build: With [#99](https://github.com/ThomasObenaus/sokar/issues/99) sokar is build based on golang 1.13.1 and the docker image is based on alpine 3.10.
- With [#98](https://github.com/ThomasObenaus/sokar/issues/98) sokar won't modify any scaleObject automatically any more. Now the ScaleObjectWatcher is disabled in dry-run mode. For more information about dry-run mode see [DryRunMode.md](doc/DryRunMode.md).
- Refactor: With [#95](https://github.com/ThomasObenaus/sokar/issues/95) the code was refactored to use the functional option approach instead of config structs.

## v0.0.9 (2019-09-24)

- With [#90](https://github.com/ThomasObenaus/sokar/issues/90) sokar gets a real downscaling for nomad workers running on AWS.
- This means:
  1.  The nomad worker with least running allocations, CPU and memory will be selected as candidate.
  2.  This nomad worker will be drained.
  3.  Then the according instance is terminated and the AWS ASG's desired capacity is reduced accordingly.
- ScalingTarget: With [#92](https://github.com/ThomasObenaus/sokar/issues/92) sokar can scale up/ down AWS EC2 instances that are managed by an AWS AutoScalingGroup. Therefore the ASG just have to be tagged with `scale-object=<some_name>` and sokar has to run in scaler mode `aws-ec2` (`--sca.mode=aws-ec2`). Then if sokar is introduced to scale the scale-object `some_name` (`--scale-object.name=some_name`) he would adjust the number of instances created by this ASG based on the scaling alerts he gets.
- Config: With [#92](https://github.com/ThomasObenaus/sokar/issues/92) the config parameter `--sca.nomad.mode` is deprecated and replaced by `--sca.mode`. Furthermore the possible values for this parameter are now: `nomad-job` instead of `job`, `nomad-dc` instead of `dc` and `aws-ec2`. The old values `job` and `dc` are deprecated.
- Config: With [#88](https://github.com/ThomasObenaus/sokar/issues/88) the interval sokar checks if someone externally (e.g. a deplyoment) has changed the scale of the scale-object is now configurable (`--sca.watcher-interval`).

## v0.0.8 (2019-09-12)

- Fix: With [#86](https://github.com/ThomasObenaus/sokar/issues/86) the issue that the scale watchers desired scale value could be changed unwanted due to side-effects was fixed.
- Config: Removed deprecated/ unused config-parmeters:
  - `--mode` was replaced by `--sca.nomad.mode`
  - `--server-address` was replaced by `--sca.nomad.server-address`.
- Testing: Service-/ Integration Test added. Can be called with `make test.service`.

## v0.0.7 (2019-06-10)

- CapacityPlanner: With [#49](https://github.com/ThomasObenaus/sokar/issues/49), beside the constant mode, a new mode for the CapacityPlanner, the linear mode, was added. Per default, the constant mode will be still used though. For more details see [CapacityPlanner](capacityplanner/README.md).
- Config: With [#68](https://github.com/ThomasObenaus/sokar/issues/68) two new config parameters for scaling a nomad data-center on AWS where added. These parameters are:
  - `--sca.nomad.dc-aws.profile`
  - `--sca.nomad.dc-aws.region`
- Config: With [#68](https://github.com/ThomasObenaus/sokar/issues/68) some config parameters where marked as deprecated and will be removed in next release. These are:
  - `--sca.mode` is deprecated, instead `--sca.nomad.mode` should be used
  - `--nomad.server-address` is deprecated, instead `--sca.nomad.server-address` should be used
- Config: With [#71](https://github.com/ThomasObenaus/sokar/issues/71) some deprecated config parameters where removed. These are:
  - `--dummy-scaling-target` was removed
  - `--job.name` was replaced by `--scale-object.name`
  - `--job.min` was replaced by `--scale-object.min`
  - `--job.max` was replaced by `--scale-object.max`

## v0.0.6 (2019-05-27)

- Data-Center Scaling: With [#65](https://github.com/ThomasObenaus/sokar/issues/60) and [#60](https://github.com/ThomasObenaus/sokar/issues/65) the functionality to scale not only jobs but also a datacenter was added. In the currently implementation an up-scaling and a **hard down-scaling is implemented**. This means **during down-scaling NO node draining will be applied**, the datacenter instances are directly terminated.
- Config: With [#64](https://github.com/ThomasObenaus/sokar/issues/64) the following config-parameters where replaced by new ones and are thus deprecated from now on:
  - `--dummy-scaling-target` is replaced by `--sca.mode=dc`. With `--sca.mode` the scaler mode can be configured to scale either a job (`--sca.mode=job`) or a datacenter (`--sca.mode=dc`).
  - `--job.name` is replaced by `--scale-object.name`
  - `--job.min` is replaced by `--scale-object.min`
  - `--job.max` is replaced by `--scale-object.max`
- Docu: With [#54](https://github.com/ThomasObenaus/sokar/issues/54) metrics documentation was added to [Metrics.md](https://github.com/ThomasObenaus/sokar/blob/master/Metrics.md).
- Fix: With [#57](https://github.com/ThomasObenaus/sokar/issues/57) an issue was fixed that sokar did ignored the down-scaling threshold.

## v0.0.5 (2019-05-12)

- Config: The expiration time used to prune scale alerts that are not updated for a long time is now configurable (--saa.alert-expiration-time, default: 10m).
- API: The scale-by end-points are enabled for dry-run mode only. This means in case sokar runs not in dry-run mode these and-points will respond an error and won't do any action.
- Robustness: Sokar now immediately adjusts the scale count at startup in case the deployed scale violates the min/ max bounds of the job.
- Robustness: Sokar now immediately adjusts the scale count to its expected count in case another source has adjusted the job count in between.

## v0.0.4 (2019-04-22)

- CI/CD: Sokar is now dockerized and available on [docker hub](https://hub.docker.com/r/thobe/sokar). It can be pulled via `docker pull thobe/sokar`.
- Config: A http end-point was added that provides the configuration used by sokar. It can be requested via GET request at `/api/config`.
- General: A http end-point was added that provides the build information of sokar. It can be requested via GET request at `/api/build`.

  - Example response:

    ```json
    {
      "version": "v0.0.3-17-ga9e9d18",
      "build_time": "2019-04-22_16-36-49",
      "revision": "a9e9d18_dirty",
      "branch": "f/34_configuration_endpoint"
    }
    ```

## v0.0.3 (2019-04-14)

- Cli/ Config: All parameters can now be configured/ set via cli, environment-variable or config-file. [#38](https://github.com/ThomasObenaus/sokar/issues/38)
- Metrics: Added basic metrics to all components. [#26](https://github.com/ThomasObenaus/sokar/issues/26)
- Dry-Run: A dry run mode was added. This can be used to test sokars decisions. [#31](https://github.com/ThomasObenaus/sokar/issues/31)
- Scale-By: A scale-by end-point was added. It can be used to manually scale (relative based on percentage or value). [#27](https://github.com/ThomasObenaus/sokar/issues/27)
