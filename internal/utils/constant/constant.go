package constant

import "time"

const (
	DEVELOPMENT_MODE = "development"
	STAGING_MODE     = "staging"
	PRODUCTION_MODE  = "production"

	DefaultConnectionTimeout = 30 * time.Second

	RoleSuperAdmin     = "superadmin"
	RoleAdminScreening = "admin_screening"
	RoleAdminVendor    = "admin_vendor"
	RoleUMKM           = "pelaku_usaha"

	OTPStatusActive = "active"
	OTPStatusUsed   = "used"
)
