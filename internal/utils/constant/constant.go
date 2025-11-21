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

	NotificationSubmitted        = "application_submitted"
	NotificationApproved         = "screening_approved"
	NotificationRejected         = "screening_rejected"
	NotificationRevised          = "screening_revised"
	NotificationFinalApproved    = "final_approved"
	NotificationFinalRejected    = "final_rejected"
	NotificationProgramReminder  = "program_reminder"
	NotificationDocumentRequired = "document_required"
	NotificationGeneralInfo      = "general_info"

	NotificationTitleSubmitted        = "Pengajuan Dikirim"
	NotificationTitleApproved         = "Pengajuan Disetujui pada Tahap Screening"
	NotificationTitleRejected         = "Pengajuan Ditolak pada Tahap Screening"
	NotificationTitleRevised          = "Pengajuan Direvisi pada Tahap Screening"
	NotificationTitleFinalApproved    = "Pengajuan Disetujui pada Tahap Final"
	NotificationTitleFinalRejected    = "Pengajuan Ditolak pada Tahap Final"
	NotificationTitleProgramReminder  = "Pengingat Program"
	NotificationTitleDocumentRequired = "Dokumen Diperlukan"
	NotificationTitleGeneralInfo      = "Informasi Umum"

	NotificationMessageSubmitted        = "Pengajuan Anda telah berhasil dikirim. Silakan tunggu proses screening."
	NotificationMessageApproved         = "Pengajuan Anda telah disetujui pada tahap screening. Silakan menunggu lanjut ke tahap final."
	NotificationMessageRejected         = "Pengajuan Anda telah ditolak pada tahap screening. Karena %s. Silakan periksa kembali data yang Anda kirim."
	NotificationMessageRevised          = "Pengajuan Anda perlu direvisi pada tahap screening. Karena %s. Silakan periksa kembali data yang Anda kirim."
	NotificationMessageFinalApproved    = "Pengajuan Anda telah disetujui pada tahap final. Selamat!"
	NotificationMessageFinalRejected    = "Pengajuan Anda telah ditolak pada tahap final. Karena %s. Silakan periksa kembali data yang Anda kirim."
	NotificationMessageProgramReminder  = "Ingatkan program yang akan datang."
	NotificationMessageDocumentRequired = "Dokumen tambahan diperlukan untuk melanjutkan proses pengajuan."
	NotificationMessageGeneralInfo      = "Informasi umum terkait program atau aplikasi."
)
