# Kava 13 Upgrade Instructions

## Software Version and Key Dates

- The version of `kava` for Kava 13 is v0.22.0
- The Kava 12 chain will be shutdown with a `SoftwareUpgradeProposal` that activates at approximately 15:00 UTC on May, 10th 2023.

### On the day of the upgrade

**The kava chain is expected to halt at block height **PLACEHOLDER**. Do not stop your node and begin the upgrade before the upgrade height, or you may go offline and be unable to recover until after the upgrade!**

**Make sure the kava process is stopped before proceeding and that you have backed up your validator**. Failure to backup your validator could make it impossible to restart your node if the upgrade fails.

**Ensure you are using golang 1.20.x and not aa different version.** Golang 1.19 and below may cause app hash mismatches!

1. Update to v0.22.0

```sh
  # check go version - look for 1.20!
  go version
  # go version go1.20.3 linux/arm64

  # in the `kava` folder
  git fetch
  git checkout v0.22.0

  # Note: Golang 1.20 must be installed before this step
  make install

  # verify versions
  kava version --long
  # name: kava
  # server_name: kava
  # version: 0.21.0
  # commit: <commit placeholder>
  # build_tags: netgo ledger,
  # go: go version go1.20.3 linux/arm64
  # build_deps:
  #  ...
  # cosmos_sdk_version: v0.46.11

  # Restart node -
  kava start
```

### Risks

As a validator, performing the upgrade procedure on your consensus nodes carries a heightened risk of double-signing and being slashed. The most important piece of this procedure is verifying your software version and genesis file hash before starting your validator and signing.

The riskiest thing a validator can do is discover that they made a mistake and repeat the upgrade procedure again during the network startup. If you discover a mistake in the process, the best thing to do is wait for the network to start before correcting it. If the network is halted and you have started with a different genesis file than the expected one, seek advice from a Kava developer before resetting your validator.

### Recovery

Prior to applying the Kava 13 upgrade, validators are encouraged to take a full data snapshot at the upgrade height before proceeding. Snap-shotting depends heavily on infrastructure, but generally this can be done by backing up the .kava directory.

It is critically important to back-up the .kava/data/priv_validator_state.json file after stopping your kava process. This file is updated every block as your validator participates in consensus rounds. It is a critical file needed to prevent double-signing, in case the upgrade fails and the previous chain needs to be restarted.

In the event that the upgrade does not succeed, validators and operators must downgrade back to v0.21.x of the Kava software and restore to their latest snapshot before restarting their nodes.

### Coordination

If the kava 13 chain does not launch by May 11th, 2023 at 00:00 UTC, the launch should be considered a failure. In the event of launch failure, coordination will occur in the [Kava discord](https://discord.com/invite/kQzh3Uv).