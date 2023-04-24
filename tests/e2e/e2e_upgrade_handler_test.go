package e2e_test

import (
	"fmt"

	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/tests/util"
	committeetypes "github.com/kava-labs/kava/x/committee/types"
)

// TestUpgradeHandler can be used to run tests post-upgrade. If an upgrade is enabled, all tests
// are run against the upgraded chain. However, this file is a good place to consolidate all
// acceptance tests for a given set of upgrade handlers.
func (suite IntegrationTestSuite) TestUpgradeHandler() {
	suite.SkipIfUpgradeDisabled()
	fmt.Println("An upgrade has run!")
	suite.True(true)

	beforeUpgradeCtx := util.CtxAtHeight(suite.UpgradeHeight - 1)
	afterUpgradeCtx := util.CtxAtHeight(suite.UpgradeHeight)

	lendWithdrawPerm := "/kava.committee.v1beta1.CommunityPoolLendWithdrawPermission"
	cdpRepayDebtPerm := "/kava.committee.v1beta1.CommunityCDPRepayDebtPermission"
	cdpWithdrawPerm := "/kava.committee.v1beta1.CommunityCDPWithdrawCollateralPermission"

	// check stability committee permissions before upgrade to ensure it starts without them
	res, err := suite.Kava.Committee.Committee(
		beforeUpgradeCtx, &committeetypes.QueryCommitteeRequest{CommitteeId: app.MainnetStabilityCommitteeId},
	)
	suite.NoError(err)
	var beforeCommittee committeetypes.Committee
	err = suite.Kava.EncodingConfig.InterfaceRegistry.UnpackAny(res.Committee, &beforeCommittee)
	suite.NoError(err)

	suite.False(suite.committeeHasPermissionWithTypeUrl(beforeCommittee, lendWithdrawPerm))
	suite.False(suite.committeeHasPermissionWithTypeUrl(beforeCommittee, cdpRepayDebtPerm))
	suite.False(suite.committeeHasPermissionWithTypeUrl(beforeCommittee, cdpWithdrawPerm))

	// check stability committee permission after upgrade to ensure it gets them
	res, err = suite.Kava.Committee.Committee(
		afterUpgradeCtx, &committeetypes.QueryCommitteeRequest{CommitteeId: app.MainnetStabilityCommitteeId},
	)
	suite.NoError(err)
	var afterCommittee committeetypes.Committee
	err = suite.Kava.EncodingConfig.InterfaceRegistry.UnpackAny(res.Committee, &afterCommittee)
	suite.NoError(err)

	suite.True(suite.committeeHasPermissionWithTypeUrl(afterCommittee, lendWithdrawPerm))
	suite.True(suite.committeeHasPermissionWithTypeUrl(afterCommittee, cdpRepayDebtPerm))
	suite.True(suite.committeeHasPermissionWithTypeUrl(afterCommittee, cdpWithdrawPerm))
}

// committeeHasPermissionWithTypeUrl iterates the permissions of the committee looking for the desired type url
func (suite *IntegrationTestSuite) committeeHasPermissionWithTypeUrl(c committeetypes.Committee, typeUrl string) bool {
	mc, success := c.(*committeetypes.MemberCommittee)
	if !success {
		panic("failed to cast committee to member committee")
	}
	for _, p := range mc.Permissions {
		fmt.Println(p.TypeUrl)
		if p.TypeUrl == typeUrl {
			return true
		}
	}
	return false
}
