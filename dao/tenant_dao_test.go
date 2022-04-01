package dao

import (
	"testing"

	"github.com/RedHatInsights/sources-api-go/internal/testutils"
	"github.com/RedHatInsights/sources-api-go/internal/testutils/fixtures"
	"github.com/RedHatInsights/sources-api-go/model"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

// TestGetOrCreateTenantIDEbsNumberCreate tests that when the EBS account number is not found, a new tenant is created.
func TestGetOrCreateTenantIDEbsNumberCreate(t *testing.T) {
	testutils.SkipIfNotRunningIntegrationTests(t)
	SwitchSchema("tenant_tests")

	accountNumber := "98765"
	identityStruct := identity.Identity{
		AccountNumber: accountNumber,
	}

	tenantDao := GetTenantDao()

	id, err := tenantDao.GetOrCreateTenantID(&identityStruct)
	if err != nil {
		t.Errorf(`Error getting or creating the tenant. Want nil error, got "%s"`, err)
	}

	var tenant model.Tenant
	err = DB.
		Debug().
		Model(&model.Tenant{}).
		Where(`id = ?`, id).
		First(&tenant).
		Error

	if err != nil {
		t.Errorf(`error fetching the tenant. Want nil error, got "%s"`, err)
	}

	want := accountNumber
	got := tenant.ExternalTenant

	if want != got {
		t.Errorf(`unexpected tenant fetched. Want EBS number "%s", got "%s"`, want, got)
	}

	DropSchema("tenant_tests")
}

// TestGetOrCreateTenantIDEbsNumberFind tests that when the EBS account number is found, the associated tenant id is
// returned.
func TestGetOrCreateTenantIDEbsNumberFind(t *testing.T) {
	testutils.SkipIfNotRunningIntegrationTests(t)
	SwitchSchema("tenant_tests")

	identityStruct := identity.Identity{
		AccountNumber: fixtures.TestTenantData[0].ExternalTenant,
	}

	tenantDao := GetTenantDao()

	id, err := tenantDao.GetOrCreateTenantID(&identityStruct)
	if err != nil {
		t.Errorf(`Error getting or creating the tenant. Want nil error, got "%s"`, err)
	}

	var tenant model.Tenant
	err = DB.
		Debug().
		Model(&model.Tenant{}).
		Where(`id = ?`, id).
		First(&tenant).
		Error

	if err != nil {
		t.Errorf(`error fetching the tenant. Want nil error, got "%s"`, err)
	}

	want := fixtures.TestTenantData[0].ExternalTenant
	got := tenant.ExternalTenant

	if want != got {
		t.Errorf(`unexpected tenant fetched. Want EBS number "%s", got "%s"`, want, got)
	}

	DropSchema("tenant_tests")
}

// TestGetOrCreateTenantIDOrgIdCreate tests that when the OrgId is not found, a new tenant is created.
func TestGetOrCreateTenantIDOrgIdCreate(t *testing.T) {
	testutils.SkipIfNotRunningIntegrationTests(t)
	SwitchSchema("tenant_tests")

	orgId := "1239875"
	identityStruct := identity.Identity{
		OrgID: orgId,
	}

	tenantDao := GetTenantDao()

	id, err := tenantDao.GetOrCreateTenantID(&identityStruct)
	if err != nil {
		t.Errorf(`Error getting or creating the tenant. Want nil error, got "%s"`, err)
	}

	var tenant model.Tenant
	err = DB.
		Debug().
		Model(&model.Tenant{}).
		Where(`id = ?`, id).
		First(&tenant).
		Error

	if err != nil {
		t.Errorf(`error fetching the tenant. Want nil error, got "%s"`, err)
	}

	want := orgId
	got := tenant.OrgID

	if want != got {
		t.Errorf(`unexpected tenant fetched. Want EBS number "%s", got "%s"`, want, got)
	}

	DropSchema("tenant_tests")
}

// TestGetOrCreateTenantIDOrgIdFind tests that when the OrgId is found, the associated tenant id is returned.
func TestGetOrCreateTenantIDOrgIdFind(t *testing.T) {
	testutils.SkipIfNotRunningIntegrationTests(t)
	SwitchSchema("tenant_tests")

	identityStruct := identity.Identity{
		OrgID: fixtures.TestTenantData[0].OrgID,
	}

	tenantDao := GetTenantDao()

	id, err := tenantDao.GetOrCreateTenantID(&identityStruct)
	if err != nil {
		t.Errorf(`Error getting or creating the tenant. Want nil error, got "%s"`, err)
	}

	var tenant model.Tenant
	err = DB.
		Debug().
		Model(&model.Tenant{}).
		Where(`id = ?`, id).
		First(&tenant).
		Error

	if err != nil {
		t.Errorf(`error fetching the tenant. Want nil error, got "%s"`, err)
	}

	want := fixtures.TestTenantData[0].OrgID
	got := tenant.OrgID

	if want != got {
		t.Errorf(`unexpected tenant fetched. Want EBS number "%s", got "%s"`, want, got)
	}

	DropSchema("tenant_tests")
}
