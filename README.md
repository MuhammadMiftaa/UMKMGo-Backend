Kelompok 4

# UMKMGo - Backend {#umkmgo---backend .unnumbered}

08 Desember 2025

![Gambar placeholder](media/image2.jpg){width="6.5in"
height="3.3090912073490815in"}

**Dokumentasi Source Code**  
*(Source Code Documentation)*

Dokumentasi source code adalah kumpulan informasi tertulis yang
menjelaskan cara suatu kode program bekerja, bagaimana struktur
sistemnya, serta bagaimana cara mengembangkan atau menggunakannya.
Dokumentasi ini bisa berupa komentar di dalam kode maupun dokumen
terpisah yang menjelaskan arsitektur, API, alur data, dependensi, hingga
petunjuk instalasi.

[**UMKMGo - Backend 1**](#umkmgo---backend)

> [1. Start Apikasi 15](#start-apikasi)
>
> [a. Prerequisites 15](#prerequisites)
>
> [1. Environment Variables 15](#environment-variables)
>
> [2. Required Services 16](#required-services)
>
> [3. Database Migrations 17](#database-migrations)
>
> [4. MinIO Bucket Setup 17](#minio-bucket-setup)
>
> [b. Running the Application 18](#running-the-application)
>
> [Development Mode 18](#development-mode)
>
> [c. Health Checks 19](#health-checks)
>
> [d. Troubleshooting 19](#troubleshooting)
>
> [Application Won't Start 19](#application-wont-start)
>
> [Database Connection Failed 20](#database-connection-failed)
>
> [Redis Connection Failed 20](#redis-connection-failed)
>
> [MinIO Connection Failed 20](#minio-connection-failed)
>
> [Vault Connection Failed 21](#vault-connection-failed)
>
> [Port Already in Use 21](#port-already-in-use)
>
> [e. Default Credentials 21](#default-credentials)
>
> [2. Inisialisasi Aplikasi 22](#inisialisasi-aplikasi)
>
> [a. Main Entry Point 22](#main-entry-point)
>
> [cmd/api/main.go 22](#cmdapimain.go)
>
> [3. Konfigurasi 23](#konfigurasi)
>
> [a. Environment Configuration 23](#environment-configuration)
>
> [config/env/config.go 23](#configenvconfig.go)
>
> [b. Database Configuration 27](#database-configuration)
>
> [config/db/config.go 27](#configdbconfig.go)
>
> [Database Migrations 28](#database-migrations-1)
>
> [Database Seeders 33](#database-seeders)
>
> [c. Logging Configuration 34](#logging-configuration)
>
> [config/log/logrus.go 34](#configloglogrus.go)
>
> [d. Redis Configuration 37](#redis-configuration)
>
> [config/redis/config.go 37](#configredisconfig.go)
>
> [config/redis/redis.go 38](#configredisredis.go)
>
> [e. MinIO Storage Configuration 39](#minio-storage-configuration)
>
> [config/storage/constant.go 39](#configstorageconstant.go)
>
> [config/storage/minio.go 39](#configstorageminio.go)
>
> [f. Vault Configuration 43](#vault-configuration)
>
> [config/vault/vault.go 43](#configvaultvault.go)
>
> [g. HTTP Router Configuration 47](#http-router-configuration)
>
> [interface/http/router/router.go 47](#interfacehttprouterrouter.go)
>
> [4. API Routes dan Handlers 51](#api-routes-dan-handlers)
>
> [a. Web Authentication Routes 51](#web-authentication-routes)
>
> [b. Mobile Authentication Routes 51](#mobile-authentication-routes)
>
> [c. User Management Routes 52](#user-management-routes)
>
> [d. Permissions Routes 52](#permissions-routes)
>
> [e. Programs Routes 53](#programs-routes)
>
> [f. Applications Routes 54](#applications-routes)
>
> [g. Dashboard Routes 54](#dashboard-routes)
>
> [h. SLA Routes 55](#sla-routes)
>
> [i. News Routes 56](#news-routes)
>
> [j. Mobile Routes 56](#mobile-routes)
>
> [Dashboard 56](#dashboard)
>
> [Programs 57](#programs)
>
> [Profile 57](#profile)
>
> [Documents 57](#documents)
>
> [Applications 58](#applications)
>
> [Notifications 58](#notifications)
>
> [News (Mobile) 59](#news-mobile)
>
> [k. Vault Decrypt Logs Routes 59](#vault-decrypt-logs-routes)
>
> [5. Middleware 60](#middleware)
>
> [a. AuthMiddleware 60](#authmiddleware)
>
> [b. MobileAuthMiddleware 60](#mobileauthmiddleware)
>
> [c. CORSMiddleware 60](#corsmiddleware)
>
> [d. LoggerMiddleware 60](#loggermiddleware)
>
> [6. Service Layer 60](#service-layer)
>
> [a. Users Service 60](#users-service)
>
> [Register 60](#register)
>
> [Login 61](#login)
>
> [UpdateProfile 62](#updateprofile)
>
> [RegisterMobile 62](#registermobile)
>
> [VerifyOTP 63](#verifyotp)
>
> [RegisterMobileProfile 64](#registermobileprofile)
>
> [LoginMobile 65](#loginmobile)
>
> [ForgotPassword 65](#forgotpassword)
>
> [ResetPassword 66](#resetpassword)
>
> [GetAllUsers 66](#getallusers)
>
> [GetUserByID 67](#getuserbyid)
>
> [UpdateUser 67](#updateuser)
>
> [DeleteUser 68](#deleteuser)
>
> [MetaCityAndProvince 68](#metacityandprovince)
>
> [GetListPermissions 68](#getlistpermissions)
>
> [GetListRolePermissions 69](#getlistrolepermissions)
>
> [UpdateRolePermissions 69](#updaterolepermissions)
>
> [b. Programs Service 70](#programs-service)
>
> [GetAllPrograms 70](#getallprograms)
>
> [GetProgramByID 70](#getprogrambyid)
>
> [CreateProgram 71](#createprogram)
>
> [UpdateProgram 71](#updateprogram)
>
> [DeleteProgram 72](#deleteprogram)
>
> [ActivateProgram 72](#activateprogram)
>
> [DeactivateProgram 73](#deactivateprogram)
>
> [c. Applications Service 73](#applications-service)
>
> [GetAllApplications 74](#getallapplications)
>
> [GetApplicationByID 74](#getapplicationbyid)
>
> [ScreeningApprove 75](#screeningapprove)
>
> [ScreeningReject 76](#screeningreject)
>
> [ScreeningRevise 76](#screeningrevise)
>
> [FinalApprove 77](#finalapprove)
>
> [FinalReject 77](#finalreject)
>
> [d. Mobile Service 78](#mobile-service)
>
> [GetDashboard 78](#getdashboard)
>
> [GetTrainingPrograms 79](#gettrainingprograms)
>
> [GetCertificationPrograms 79](#getcertificationprograms)
>
> [GetFundingPrograms 79](#getfundingprograms)
>
> [GetProgramDetail 80](#getprogramdetail)
>
> [GetUMKMProfile 80](#getumkmprofile)
>
> [UpdateUMKMProfile 81](#updateumkmprofile)
>
> [GetUMKMDocuments 81](#getumkmdocuments)
>
> [UploadDocument 82](#uploaddocument)
>
> [CreateTrainingApplication 82](#createtrainingapplication)
>
> [CreateCertificationApplication 83](#createcertificationapplication)
>
> [CreateFundingApplication 84](#createfundingapplication)
>
> [GetApplicationList 84](#getapplicationlist)
>
> [GetApplicationDetail 85](#getapplicationdetail)
>
> [ReviseApplication 85](#reviseapplication)
>
> [GetNotificationsByUMKMID 86](#getnotificationsbyumkmid)
>
> [GetUnreadCount 86](#getunreadcount)
>
> [MarkNotificationsAsRead 87](#marknotificationsasread)
>
> [MarkAllNotificationsAsRead 87](#markallnotificationsasread)
>
> [GetPublishedNews 87](#getpublishednews)
>
> [GetNewsDetail 88](#getnewsdetail)
>
> [e. Dashboard Service 88](#dashboard-service)
>
> [GetUMKMByCardType 88](#getumkmbycardtype)
>
> [GetApplicationStatusSummary 89](#getapplicationstatussummary)
>
> [GetApplicationStatusDetail 89](#getapplicationstatusdetail)
>
> [GetApplicationByType 89](#getapplicationbytype)
>
> [f. SLA Service 90](#sla-service)
>
> [GetSLAScreening 90](#getslascreening)
>
> [GetSLAFinal 90](#getslafinal)
>
> [UpdateSLAScreening 91](#updateslascreening)
>
> [UpdateSLAFinal 91](#updateslafinal)
>
> [ExportApplications 92](#exportapplications)
>
> [ExportPrograms 92](#exportprograms)
>
> [g. News Service 92](#news-service)
>
> [GetAllNews 93](#getallnews)
>
> [GetNewsByID 93](#getnewsbyid)
>
> [CreateNews 93](#createnews)
>
> [UpdateNews 94](#updatenews)
>
> [DeleteNews 95](#deletenews)
>
> [PublishNews 95](#publishnews)
>
> [UnpublishNews 95](#unpublishnews)
>
> [h. Vault Decrypt Log Service 96](#vault-decrypt-log-service)
>
> [GetLogs 96](#getlogs)
>
> [GetLogsByUserID 96](#getlogsbyuserid)
>
> [GetLogsByUMKMID 97](#getlogsbyumkmid)
>
> [7. Repository Layer 98](#repository-layer)
>
> [a. Applications Repository 98](#applications-repository)
>
> [GetAllApplications 98](#getallapplications-1)
>
> [GetApplicationByID 98](#getapplicationbyid-1)
>
> [GetApplicationsByUMKMID 99](#getapplicationsbyumkmid)
>
> [CreateApplication 99](#createapplication)
>
> [UpdateApplication 100](#updateapplication)
>
> [DeleteApplication 100](#deleteapplication)
>
> [CreateApplicationDocuments 101](#createapplicationdocuments)
>
> [GetApplicationDocuments 101](#getapplicationdocuments)
>
> [DeleteApplicationDocuments 101](#deleteapplicationdocuments)
>
> [CreateApplicationHistory 102](#createapplicationhistory)
>
> [GetApplicationHistories 102](#getapplicationhistories)
>
> [GetProgramByID 103](#getprogrambyid-1)
>
> [GetUMKMByUserID 103](#getumkmbyuserid)
>
> [IsApplicationExists 104](#isapplicationexists)
>
> [b. Dashboard Repository 104](#dashboard-repository)
>
> [GetUMKMByCardType 104](#getumkmbycardtype-1)
>
> [GetApplicationStatusSummary 105](#getapplicationstatussummary-1)
>
> [GetApplicationStatusDetail 105](#getapplicationstatusdetail-1)
>
> [GetApplicationByType 106](#getapplicationbytype-1)
>
> [c. Mobile Repository 106](#mobile-repository)
>
> [GetProgramsByType 106](#getprogramsbytype)
>
> [GetProgramDetailByID 107](#getprogramdetailbyid)
>
> [GetUMKMProfileByID 107](#getumkmprofilebyid)
>
> [UpdateUMKMProfile 108](#updateumkmprofile-1)
>
> [UpdateUMKMDocument 108](#updateumkmdocument)
>
> [CreateApplication (Mobile) 109](#createapplication-mobile)
>
> [CreateApplicationDocuments (Mobile)
> 109](#createapplicationdocuments-mobile)
>
> [CreateApplicationHistory (Mobile)
> 110](#createapplicationhistory-mobile)
>
> [CreateTrainingApplication 110](#createtrainingapplication-1)
>
> [CreateCertificationApplication
> 110](#createcertificationapplication-1)
>
> [CreateFundingApplication 111](#createfundingapplication-1)
>
> [GetApplicationsByUMKMID (Mobile)
> 111](#getapplicationsbyumkmid-mobile)
>
> [GetApplicationDetailByID (Mobile)
> 112](#getapplicationdetailbyid-mobile)
>
> [DeleteApplicationDocumentsByApplicationID
> 112](#deleteapplicationdocumentsbyapplicationid)
>
> [GetProgramRequirements 112](#getprogramrequirements)
>
> [GetPublishedNews 113](#getpublishednews-1)
>
> [GetPublishedNewsBySlug 113](#getpublishednewsbyslug)
>
> [IncrementViews 114](#incrementviews)
>
> [d. News Repository 114](#news-repository)
>
> [GetAllNews 114](#getallnews-1)
>
> [GetNewsByID 115](#getnewsbyid-1)
>
> [GetNewsBySlug 115](#getnewsbyslug)
>
> [CreateNews 116](#createnews-1)
>
> [UpdateNews 116](#updatenews-1)
>
> [DeleteNews 117](#deletenews-1)
>
> [IsSlugExists 117](#isslugexists)
>
> [CreateNewsTags 118](#createnewstags)
>
> [DeleteNewsTags 118](#deletenewstags)
>
> [GetNewsTags 118](#getnewstags)
>
> [e. Notification Repository 119](#notification-repository)
>
> [CreateNotification 119](#createnotification)
>
> [GetNotificationsByUMKMID 119](#getnotificationsbyumkmid-1)
>
> [GetUnreadCount 120](#getunreadcount-1)
>
> [MarkAsRead 120](#markasread)
>
> [MarkAllAsRead 121](#markallasread)
>
> [f. OTP Repository 121](#otp-repository)
>
> [CreateOTP 121](#createotp)
>
> [GetOTPByPhone 122](#getotpbyphone)
>
> [GetOTPByTempToken 122](#getotpbytemptoken)
>
> [UpdateOTP 122](#updateotp)
>
> [g. Programs Repository 123](#programs-repository)
>
> [GetAllPrograms 123](#getallprograms-1)
>
> [GetProgramByID 123](#getprogrambyid-2)
>
> [CreateProgram 124](#createprogram-1)
>
> [UpdateProgram 124](#updateprogram-1)
>
> [DeleteProgram 125](#deleteprogram-1)
>
> [CreateProgramBenefits 125](#createprogrambenefits)
>
> [CreateProgramRequirements 125](#createprogramrequirements)
>
> [GetProgramBenefits 126](#getprogrambenefits)
>
> [GetProgramRequirements 126](#getprogramrequirements-1)
>
> [DeleteProgramBenefits 127](#deleteprogrambenefits)
>
> [DeleteProgramRequirements 127](#deleteprogramrequirements)
>
> [h. SLA Repository 128](#sla-repository)
>
> [GetSLAByStatus 128](#getslabystatus)
>
> [UpdateSLA 128](#updatesla)
>
> [GetApplicationsForExport 128](#getapplicationsforexport)
>
> [GetProgramsForExport 129](#getprogramsforexport)
>
> [i. Users Repository 129](#users-repository)
>
> [GetAllUsers 130](#getallusers-1)
>
> [GetUserByID 130](#getuserbyid-1)
>
> [GetUserByEmail 130](#getuserbyemail)
>
> [CreateUser 131](#createuser)
>
> [UpdateUser 131](#updateuser-1)
>
> [DeleteUser 132](#deleteuser-1)
>
> [CreateUMKM 132](#createumkm)
>
> [GetUMKMByPhone 133](#getumkmbyphone)
>
> [GetAllRoles 133](#getallroles)
>
> [GetRoleByID 133](#getrolebyid)
>
> [GetRoleByName 134](#getrolebyname)
>
> [IsRoleExist 134](#isroleexist)
>
> [IsPermissionExist 134](#ispermissionexist)
>
> [GetListPermissions 135](#getlistpermissions-1)
>
> [GetListPermissionsByRoleID 135](#getlistpermissionsbyroleid)
>
> [GetListRolePermissions 136](#getlistrolepermissions-1)
>
> [DeletePermissionsByRoleID 136](#deletepermissionsbyroleid)
>
> [AddRolePermissions 137](#addrolepermissions)
>
> [GetProvinces 137](#getprovinces)
>
> [GetCities 138](#getcities)
>
> [j. Vault Decrypt Logs Repository 138](#vault-decrypt-logs-repository)
>
> [LogDecrypt 138](#logdecrypt)
>
> [GetLogs 139](#getlogs-1)
>
> [GetLogsByUserID 139](#getlogsbyuserid-1)
>
> [GetLogsByUMKMID 140](#getlogsbyumkmid-1)
>
> [8. Unit Test 140](#unit-test)
>
> [a. Users Service Tests 140](#users-service-tests)
>
> [TestNewUsersService 140](#testnewusersservice)
>
> [TestLogin 141](#testlogin)
>
> [TestRegister 142](#testregister)
>
> [TestGetMeta 142](#testgetmeta)
>
> [TestLoginMobile 143](#testloginmobile)
>
> [TestRegisterMobile 144](#testregistermobile)
>
> [TestRegisterMobileProfile 145](#testregistermobileprofile)
>
> [TestForgotPassword 146](#testforgotpassword)
>
> [TestResetPassword 146](#testresetpassword)
>
> [TestVerifyOTP 147](#testverifyotp)
>
> [TestGetAllUsers 148](#testgetallusers)
>
> [TestGetUserByID 149](#testgetuserbyid)
>
> [TestUpdateUser 149](#testupdateuser)
>
> [TestUpdateProfile 150](#testupdateprofile)
>
> [TestDeleteUser 151](#testdeleteuser)
>
> [TestGetListPermissions 152](#testgetlistpermissions)
>
> [TestGetListRolePermissions 152](#testgetlistrolepermissions)
>
> [TestUpdateRolePermissions 153](#testupdaterolepermissions)
>
> [b. Programs Service Tests 154](#programs-service-tests)
>
> [TestNewProgramsService 154](#testnewprogramsservice)
>
> [TestGetAllPrograms 154](#testgetallprograms)
>
> [TestGetProgramByID 155](#testgetprogrambyid)
>
> [TestCreateProgram 156](#testcreateprogram)
>
> [TestUpdateProgram 157](#testupdateprogram)
>
> [TestActivateProgram 157](#testactivateprogram)
>
> [TestDeactivateProgram 158](#testdeactivateprogram)
>
> [TestDeleteProgram 159](#testdeleteprogram)
>
> [c. Applications Service Tests 160](#applications-service-tests)
>
> [TestNewApplicationsService 160](#testnewapplicationsservice)
>
> [TestGetAllApplications 160](#testgetallapplications)
>
> [TestGetApplicationByID 161](#testgetapplicationbyid)
>
> [TestScreeningApprove 162](#testscreeningapprove)
>
> [TestScreeningReject 163](#testscreeningreject)
>
> [TestScreeningRevise 164](#testscreeningrevise)
>
> [TestFinalApprove 165](#testfinalapprove)
>
> [TestFinalReject 165](#testfinalreject)
>
> [d. Dashboard Service Tests 166](#dashboard-service-tests)
>
> [TestNewDashboardService 166](#testnewdashboardservice)
>
> [TestGetUMKMByCardType 167](#testgetumkmbycardtype)
>
> [TestGetApplicationStatusSummary
> 168](#testgetapplicationstatussummary)
>
> [TestGetApplicationStatusDetail 168](#testgetapplicationstatusdetail)
>
> [TestGetApplicationByType 169](#testgetapplicationbytype)
>
> [e. SLA Service Tests 170](#sla-service-tests)
>
> [TestNewSLAService 170](#testnewslaservice)
>
> [TestGetSLAScreening 170](#testgetslascreening)
>
> [TestGetSLAFinal 171](#testgetslafinal)
>
> [TestUpdateSLAScreening 172](#testupdateslascreening)
>
> [TestUpdateSLAFinal 172](#testupdateslafinal)
>
> [TestExportApplications 173](#testexportapplications)
>
> [TestExportPrograms 174](#testexportprograms)
>
> [f. News Service Tests 174](#news-service-tests)
>
> [TestNewNewsService 175](#testnewnewsservice)
>
> [TestGetAllNews 175](#testgetallnews)
>
> [TestGetNewsByID 176](#testgetnewsbyid)
>
> [TestCreateNews 177](#testcreatenews)
>
> [TestUpdateNews 177](#testupdatenews)
>
> [TestDeleteNews 178](#testdeletenews)
>
> [TestPublishNews 179](#testpublishnews)
>
> [TestUnpublishNews 180](#testunpublishnews)
>
> [g. Mobile Service Tests 180](#mobile-service-tests)
>
> [TestNewMobileService 180](#testnewmobileservice)
>
> [TestGetDashboard 181](#testgetdashboard)
>
> [TestGetTrainingPrograms 182](#testgettrainingprograms)
>
> [TestGetCertificationPrograms 182](#testgetcertificationprograms)
>
> [TestGetFundingPrograms 183](#testgetfundingprograms)
>
> [TestGetProgramDetail 184](#testgetprogramdetail)
>
> [TestGetUMKMProfile 185](#testgetumkmprofile)
>
> [TestUpdateUMKMProfile 185](#testupdateumkmprofile)
>
> [TestGetUMKMDocuments 186](#testgetumkmdocuments)
>
> [TestUploadDocument 187](#testuploaddocument)
>
> [TestCreateTrainingApplication 188](#testcreatetrainingapplication)
>
> [TestCreateCertificationApplication
> 189](#testcreatecertificationapplication)
>
> [TestCreateFundingApplication 190](#testcreatefundingapplication)
>
> [TestGetApplicationList 190](#testgetapplicationlist)
>
> [TestGetApplicationDetail 191](#testgetapplicationdetail)
>
> [TestReviseApplication 192](#testreviseapplication)
>
> [TestGetNotificationsByUMKMID 193](#testgetnotificationsbyumkmid)
>
> [TestGetUnreadCount 194](#testgetunreadcount)
>
> [TestMarkNotificationsAsRead 194](#testmarknotificationsasread)
>
> [TestMarkAllNotificationsAsRead 195](#testmarkallnotificationsasread)
>
> [TestGetPublishedNews 196](#testgetpublishednews)
>
> [TestGetNewsDetailBySlug 197](#testgetnewsdetailbyslug)
>
> [h. Vault Decrypt Log Service Tests
> 197](#vault-decrypt-log-service-tests)
>
> [TestNewVaultDecryptLogService 197](#testnewvaultdecryptlogservice)
>
> [TestGetLogs 198](#testgetlogs)
>
> [TestGetLogsByUserID 199](#testgetlogsbyuserid)
>
> [TestGetLogsByUMKMID 200](#testgetlogsbyumkmid)
>
> [9. Model 200](#model)
>
> [a. Base Model 200](#base-model)
>
> [b. User Model 201](#user-model)
>
> [c. Role Model 201](#role-model)
>
> [d. Permission Model 201](#permission-model)
>
> [e. RolePermission Model 202](#rolepermission-model)
>
> [f. Province Model 202](#province-model)
>
> [g. City Model 202](#city-model)
>
> [h. UMKM Model 203](#umkm-model)
>
> [i. Program Model 204](#program-model)
>
> [j. ProgramBenefit Model 204](#programbenefit-model)
>
> [k. ProgramRequirement Model 205](#programrequirement-model)
>
> [Application Model 205](#application-model)
>
> [l. ApplicationDocument Model 206](#applicationdocument-model)
>
> [ApplicationHistory Model 206](#applicationhistory-model)
>
> [m. TrainingApplication Model 206](#trainingapplication-model)
>
> [n. CertificationApplication Model
> 207](#certificationapplication-model)
>
> [o. FundingApplication Model 207](#fundingapplication-model)
>
> [p. SLA Model 208](#sla-model)
>
> [q. Notification Model 208](#notification-model)
>
> [r. News Model 209](#news-model)
>
> [s. NewsTag Model 209](#newstag-model)
>
> [t. OTP Model 210](#otp-model)
>
> [u. VaultDecryptLog Model 210](#vaultdecryptlog-model)
>
> [10. DTO 211](#dto)
>
> [a. Users DTO 211](#users-dto)
>
> [b. OTP DTO 211](#otp-dto)
>
> [c. Permissions DTO 211](#permissions-dto)
>
> [d. RolePermissions DTO 212](#rolepermissions-dto)
>
> [e. RolePermissionsResponse DTO 212](#rolepermissionsresponse-dto)
>
> [f. Province DTO 212](#province-dto)
>
> [g. City DTO 212](#city-dto)
>
> [h. User DTO 213](#user-dto)
>
> [i. RegisterMobile DTO 213](#registermobile-dto)
>
> [j. ResetPasswordMobile DTO 213](#resetpasswordmobile-dto)
>
> [k. UMKMMobile DTO 213](#umkmmobile-dto)
>
> [l. UMKMWeb DTO 214](#umkmweb-dto)
>
> [m. MetaCityAndProvince DTO 214](#metacityandprovince-dto)
>
> [n. Programs DTO 215](#programs-dto)
>
> [o. Applications DTO 215](#applications-dto)
>
> [p. ApplicationDocuments DTO 216](#applicationdocuments-dto)
>
> [q. ApplicationHistories DTO 216](#applicationhistories-dto)
>
> [r. ApplicationDecision DTO 217](#applicationdecision-dto)
>
> [s. TrainingApplicationData DTO 217](#trainingapplicationdata-dto)
>
> [t. CertificationApplicationData DTO
> 217](#certificationapplicationdata-dto)
>
> [u. FundingApplicationData DTO 218](#fundingapplicationdata-dto)
>
> [v. SLA DTO 218](#sla-dto)
>
> [w. ExportRequest DTO 218](#exportrequest-dto)
>
> [x. UMKMByCardType DTO 219](#umkmbycardtype-dto)
>
> [y. ApplicationStatusSummary DTO 219](#applicationstatussummary-dto)
>
> [z. ApplicationStatusDetail DTO 219](#applicationstatusdetail-dto)
>
> [aa. ApplicationByType DTO 219](#applicationbytype-dto)
>
> [ab. UserData DTO 220](#userdata-dto)
>
> [ac. ProgramListMobile DTO 220](#programlistmobile-dto)
>
> [ad. ProgramDetailMobile DTO 221](#programdetailmobile-dto)
>
> [ae. UMKMProfile DTO 221](#umkmprofile-dto)
>
> [af. UpdateUMKMProfile DTO 222](#updateumkmprofile-dto)
>
> [ag. UploadDocumentRequest DTO 222](#uploaddocumentrequest-dto)
>
> [ah. CreateApplicationTraining DTO
> 222](#createapplicationtraining-dto)
>
> [ai. CreateApplicationCertification DTO
> 223](#createapplicationcertification-dto)
>
> [aj. CreateApplicationFunding DTO 223](#createapplicationfunding-dto)
>
> [ak. ApplicationListMobile DTO 224](#applicationlistmobile-dto)
>
> [al. ApplicationDetailMobile DTO 224](#applicationdetailmobile-dto)
>
> [am. DashboardData DTO 224](#dashboarddata-dto)
>
> [an. UMKMDocument DTO 225](#umkmdocument-dto)
>
> [ao. NotificationResponse DTO 225](#notificationresponse-dto)
>
> [ap. NewsRequest DTO 225](#newsrequest-dto)
>
> [aq. NewsResponse DTO 226](#newsresponse-dto)
>
> [ar. NewsListResponse DTO 226](#newslistresponse-dto)
>
> [as. NewsListMobile DTO 227](#newslistmobile-dto)
>
> [at. NewsDetailMobile DTO 227](#newsdetailmobile-dto)
>
> [au. NewsQueryParams DTO 228](#newsqueryparams-dto)
>
> [11. Helper 228](#helper)
>
> [a. Password Management 228](#password-management)
>
> [PasswordValidator 228](#passwordvalidator)
>
> [PasswordHashing 228](#passwordhashing)
>
> [ComparePass 229](#comparepass)
>
> [b. JWT Token Management 229](#jwt-token-management)
>
> [GenerateWebToken 229](#generatewebtoken)
>
> [GenerateMobileToken 230](#generatemobiletoken)
>
> [VerifyToken 230](#verifytoken)
>
> [Email Validation 231](#email-validation)
>
> [EmailValidator 231](#emailvalidator)
>
> [c. Phone Number Management 231](#phone-number-management)
>
> [NormalizePhone 231](#normalizephone)
>
> [DenormalizePhone 232](#denormalizephone)
>
> [d. OTP Management 232](#otp-management)
>
> [GenerateOTP 232](#generateotp)
>
> [e. NIK Validation 233](#nik-validation)
>
> [NIKValidator 233](#nikvalidator)
>
> [f. File and String Utilities 233](#file-and-string-utilities)
>
> [GenerateRequestID 233](#generaterequestid)
>
> [GenerateFileName 233](#generatefilename)
>
> [RandomString 234](#randomstring)
>
> [MaskMiddle 234](#maskmiddle)
>
> [g. QR Code Generation 235](#qr-code-generation)
>
> [GenerateQRCode 235](#generateqrcode)
>
> [h. Email Sending Utilities 235](#email-sending-utilities)
>
> [SMTPInterface 235](#smtpinterface)
>
> [NewSMTPClient 236](#newsmtpclient)
>
> [12. Constants 236](#constants)
>
> [a. Environment Modes 236](#environment-modes)
>
> [b. Connection Settings 237](#connection-settings)
>
> [c. User Roles 237](#user-roles)
>
> [d. OTP Status 238](#otp-status)
>
> [e. Application Status 238](#application-status)
>
> [f. Notification Types 239](#notification-types)
>
> [g. Notification Titles 239](#notification-titles)
>
> [h. Notification Messages 240](#notification-messages)
>
> [i. Document Types 241](#document-types)
>
> [13. Deployments 242](#deployments)
>
> [a. GitLab CI/CD Pipeline 242](#gitlab-cicd-pipeline)
>
> [.gitlab-ci.yml Stages 242](#gitlab-ci.yml-stages)
>
> [b. Docker Configurations 244](#docker-configurations)
>
> [Dockerfile.api 244](#dockerfile.api)
>
> [Dockerfile.migrate 245](#dockerfile.migrate)
>
> [c. Migration Script 245](#migration-script)
>
> [migrate.sh 245](#migrate.sh)

## Start Apikasi  {#start-apikasi}

### Prerequisites

#### **1. Environment Variables** {#environment-variables .unnumbered}

> Create .env file di root project:
>
> \# Server Configuration  
> MODE=development  
> PORT=8080  
> JWT_SECRET_KEY=your_jwt_secret_key_here  
>   
> \# Database Configuration  
> DB_HOST=localhost  
> DB_PORT=5432  
> DB_USER=postgres  
> DB_PASSWORD=your_password  
> DB_NAME=umkmgo_db  
>   
> \# Redis Configuration  
> REDIS_HOST=localhost  
> REDIS_PORT=6379  
>   
> \# MinIO Configuration  
> MINIO_HOST=localhost:9000  
> MINIO_ROOT_USER=minioadmin  
> MINIO_ROOT_PASSWORD=minioadmin  
> MINIO_MAX_OPEN_CONN=10  
> MINIO_USE_SSL=0  
>   
> \# Zoho SMTP Configuration  
> ZOHO_SMTP_HOST=smtp.zoho.com  
> ZOHO_SMTP_PORT=587  
> ZOHO_SMTP_USER=your_email@domain.com  
> ZOHO_SMTP_PASSWORD=your_smtp_password  
> ZOHO_SMTP_SECURE=TLS  
> ZOHO_SMTP_AUTH=true  
>   
> \# Fonnte Configuration (WhatsApp Gateway)  
> FONNTE_TOKEN=your_fonnte_api_token  
>   
> \# Vault Configuration  
> VAULT_ADDR=http://localhost:8200  
> VAULT_ROLE_ID=your_role_id  
> VAULT_SECRET_ID=your_secret_id  
> VAULT_TRANSIT_PATH=transit  
> VAULT_NIK_ENCRYPTION_KEY=nik-key  
> VAULT_KARTU_ENCRYPTION_KEY=kartu-key

#### **2. Required Services** {#required-services .unnumbered}

> **PostgreSQL Database**
>
> *\# Using Docker  
> * docker run -d \\  
> \--name postgres \\  
> -e POSTGRES_PASSWORD=your_password \\  
> -e POSTGRES_DB=umkmgo_db \\  
> -p 5432:5432 \\  
> postgres:15-alpine
>
> **Redis Server**
>
> *\# Using Docker  
> * docker run -d \\  
> \--name redis \\  
> -p 6379:6379 \\  
> redis:7-alpine
>
> **MinIO Object Storage**
>
> *\# Using Docker  
> * docker run -d \\  
> \--name minio \\  
> -p 9000:9000 \\  
> -p 9001:9001 \\  
> -e MINIO_ROOT_USER=minioadmin \\  
> -e MINIO_ROOT_PASSWORD=minioadmin \\  
> minio/minio server /data \--console-address \":9001\"
>
> **HashiCorp Vault**
>
> *\# Using Docker (Development Mode)  
> * docker run -d \\  
> \--name vault \\  
> -p 8200:8200 \\  
> -e VAULT_DEV_ROOT_TOKEN_ID=root \\  
> -e VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200 \\  
> \--cap-add=IPC_LOCK \\  
> vault:latest  
>   
> *\# Setup Transit Engine  
> * export VAULT_ADDR=\'http://localhost:8200\'  
> export VAULT_TOKEN=\'root\'  
>   
> vault secrets enable transit  
> vault write -f transit/keys/nik-key  
> vault write -f transit/keys/kartu-key  
>   
> *\# Setup AppRole  
> * vault auth enable approle  
> vault write auth/approle/role/umkmgo-backend \\  
> token_ttl=1h \\  
> token_max_ttl=4h \\  
> policies=\"default\"  
>   
> *\# Get Role ID and Secret ID  
> * vault read auth/approle/role/umkmgo-backend/role-id  
> vault write -f auth/approle/role/umkmgo-backend/secret-id

#### **3. Database Migrations** {#database-migrations .unnumbered}

> **Using Makefile:**
>
> *\# Run migrations  
> * make migrate-up  
>   
> *\# Run seeders  
> * make seed-up  
>   
> *\# Rollback migration  
> * make migrate-down  
>   
> *\# Reset database (drop all + migrate + seed)  
> * make migrate-reset
>
> **Manual Migration:**
>
> *\# Install Goose  
> * go install github.com/pressly/goose/v3/cmd/goose@latest  
>   
> *\# Run migrations  
> * goose -dir config/db/migrations postgres \"host=localhost port=5432
> user=postgres password=your_password dbname=umkmgo_db
> sslmode=disable\" up  
>   
> *\# Run seeders  
> * goose -dir config/db/seeder postgres \"host=localhost port=5432
> user=postgres password=your_password dbname=umkmgo_db
> sslmode=disable\" up

#### **4. MinIO Bucket Setup** {#minio-bucket-setup .unnumbered}

> *\# Using MinIO Client (mc)  
> * mc alias set local http://localhost:9000 minioadmin minioadmin  
>   
> *\# Create buckets  
> * mc mb local/umkmgo-programs  
> mc mb local/umkmgo-umkms  
> mc mb local/umkmgo-applications  
> mc mb local/umkmgo-news  
>   
> *\# Set public read policy (optional untuk development)  
> * mc anonymous set download local/umkmgo-programs  
> mc anonymous set download local/umkmgo-umkms  
> mc anonymous set download local/umkmgo-applications  
> mc anonymous set download local/umkmgo-news

### Running the Application

#### **Development Mode** {#development-mode .unnumbered}

> **Using Go Run:**
>
> go run cmd/api/main.go
>
> **Using Makefile:**
>
> make run
>
> **Expected Output:**
>
> \[09/Dec/2025:10:30:00 +0700\] INFO: Setup Database Connection Start  
> \[09/Dec/2025:10:30:00 +0700\] INFO: Setup Database Connection
> Success  
> \[09/Dec/2025:10:30:00 +0700\] INFO: Setup Redis Connection Start  
> \[09/Dec/2025:10:30:00 +0700\] INFO: Setup Redis Connection Success  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Setup MinIO Connection Start  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Setup MinIO Connection Success  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Setup Vault Connection Start  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Vault login successful, token
> TTL: 3600 seconds  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Setup Vault Connection Success  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Starting UMKMGo API\...  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Registered route - METHOD: GET,
> PATH: /test  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Registered route - METHOD: GET,
> PATH: /v1/users  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Registered route - METHOD: POST,
> PATH: /v1/users/login  
> \...  
> \[09/Dec/2025:10:30:01 +0700\] INFO: Starting HTTP server on port 8080

### Health Checks

> **API Health Check:**
>
> curl http://localhost:8080/test  
> *\# Response: {\"message\":\"Hello World!\"}*
>
> **Database Connection:**
>
> *\# Check if migrations are applied  
> * psql -h localhost -U postgres -d umkmgo_db -c \"SELECT \* FROM
> schema_migrations LIMIT 5;\"
>
> **Redis Connection:**
>
> *\# Check Redis connection  
> * redis-cli -h localhost -p 6379 ping  
> *\# Response: PONG*
>
> **MinIO Connection:**
>
> *\# List buckets  
> * mc ls local/  
> *\# Should show: umkmgo-programs, umkmgo-umkms, umkmgo-applications,
> umkmgo-news*
>
> **Vault Connection:**
>
> *\# Check Vault status  
> * curl http://localhost:8200/v1/sys/health  
> *\# Response:
> {\"initialized\":true,\"sealed\":false,\"standby\":false}*

### Troubleshooting

#### **Application Won't Start** {#application-wont-start .unnumbered}

> **Check Environment Variables:**
>
> *\# Verify .env file exists  
> * ls -la .env  
>   
> *\# Check for missing variables  
> * go run cmd/api/main.go 2\>&1 **\|** grep \"env is not set\"
>
> **Check Service Availability:**
>
> *\# PostgreSQL  
> * pg_isready -h localhost -p 5432  
>   
> *\# Redis  
> * redis-cli -h localhost -p 6379 ping  
>   
> *\# MinIO  
> * curl http://localhost:9000/minio/health/live  
>   
> *\# Vault  
> * curl http://localhost:8200/v1/sys/health

#### **Database Connection Failed** {#database-connection-failed .unnumbered}

> **Check PostgreSQL logs:**
>
> docker logs postgres
>
> **Verify credentials:**
>
> psql -h localhost -p 5432 -U postgres -d umkmgo_db
>
> **Check firewall:**
>
> sudo ufw status  
> *\# Ensure port 5432 is open*

#### **Redis Connection Failed** {#redis-connection-failed .unnumbered}

> **Check Redis logs:**
>
> docker logs redis
>
> **Test connection:**
>
> redis-cli -h localhost -p 6379  
> \> ping  
> PONG

#### **MinIO Connection Failed** {#minio-connection-failed .unnumbered}

> **Check MinIO logs:**
>
> docker logs minio
>
> **Access MinIO Console:**

- Open browser: http://localhost:9001

- Login dengan MINIO_ROOT_USER dan MINIO_ROOT_PASSWORD

- Verify buckets exist

#### **Vault Connection Failed** {#vault-connection-failed .unnumbered}

> **Check Vault logs:**
>
> docker logs vault
>
> **Verify Vault unsealed:**
>
> export VAULT_ADDR=\'http://localhost:8200\'  
> vault status
>
> **Re-setup AppRole:**
>
> *\# If role doesn\'t exist  
> * vault auth enable approle  
> vault write auth/approle/role/umkmgo-backend token_ttl=1h
> token_max_ttl=4h policies=\"default\"  
> vault read auth/approle/role/umkmgo-backend/role-id  
> vault write -f auth/approle/role/umkmgo-backend/secret-id

#### **Port Already in Use** {#port-already-in-use .unnumbered}

> **Check process using port:**
>
> *\# Linux  
> * sudo lsof -i :8080  
>   
> *\# Kill process  
> * kill -9 \<PID\>
>
> **Change port:**
>
> \# In .env file  
> PORT=8081

### Default Credentials

> **Superadmin Account:**
>
> Email: admin@example.com  
> Password: admin123  
> Role: superadmin
>
> **Admin Screening Account:**
>
> Email: screening1@example.com  
> Password: admin123  
> Role: admin_screening
>
> **Admin Vendor Account:**
>
> Email: vendor1@example.com  
> Password: admin123  
> Role: admin_vendor
>
> **UMKM Account (Sample):**
>
> Email: umkm1@example.com  
> Password: admin123  
> Role: pelaku_usaha

## Inisialisasi Aplikasi  {#inisialisasi-aplikasi}

### Main Entry Point

#### **cmd/api/main.go** {#cmdapimain.go .unnumbered}

> **Fungsi:** Entry point aplikasi yang menginisialisasi semua komponen
> sistem secara berurutan.
>
> **Initialization Sequence:**

1.  **Logger Setup**

- log.SetupLogger()

  - Dijalankan pertama kali sebelum komponen lain

  - Mengkonfigurasi logging berdasarkan environment mode

  - Tujuan: Memastikan logging tersedia untuk semua proses selanjutnya

2.  **Environment Variables Loading**

- env.LoadNative()

  - Load semua environment variables dari .env file

  - Validasi kelengkapan environment variables

  - Return list missing variables jika ada yang belum di-set

  - Tujuan: Memastikan semua konfigurasi tersedia sebelum inisialisasi
    > komponen

3.  **Database Connection**

- db.SetupDatabase(env.Cfg.Database)

  - Establish koneksi ke PostgreSQL database

  - Setup GORM sebagai ORM

  - Tujuan: Menyediakan database connection pool untuk repository layer

4.  **Redis Connection**

- redis.SetupRedisDatabase(env.Cfg.Redis)

  - Establish koneksi ke Redis server

  - Setup connection untuk caching dan session management

  - Tujuan: Menyediakan fast in-memory data store

5.  **MinIO Connection**

- storage.SetupMinio(env.Cfg.Minio)

  - Initialize MinIO client untuk object storage

  - Setup connection pool dan validate accessibility

  - Tujuan: Menyediakan file storage untuk upload dokumen, gambar, dll

6.  **Vault Connection**

- vault.SetupVault(env.Cfg.Vault)

  - Initialize HashiCorp Vault client

  - Login menggunakan AppRole authentication

  - Setup auto token renewal

  - Tujuan: Menyediakan encryption/decryption service untuk data
    > sensitif

7.  **HTTP Router Setup**

- r := router.SetupRouter()  
  > r.Listen(\":\" + env.Cfg.Server.Port)

  - Setup Fiber framework sebagai HTTP server

  - Register semua routes dan middleware

  - Start HTTP server pada port yang dikonfigurasi

  - Tujuan: Menerima dan memproses HTTP requests

8.  **Defer Statement:**

> **defer** log.Info(\"UMKMGo API stopped\")

- Log pesan ketika aplikasi berhenti

- Dieksekusi terakhir sebelum program exit

## Konfigurasi  {#konfigurasi}

### Environment Configuration

#### **config/env/config.go** {#configenvconfig.go .unnumbered}

> **Fungsi:** Centralized configuration management menggunakan
> environment variables.
>
> **Configuration Structs**
>
> **Server Configuration**
>
> **type** Server **struct** {  
> Mode string *// development, staging, production  
> * Port string *// HTTP server port  
> * JWTSecretKey string *// Secret key untuk JWT signing  
> * }
>
> **Environment Variables:**

- MODE: Application mode (development/staging/production)

- PORT: HTTP server port (default: 8080)

- JWT_SECRET_KEY: Secret key untuk JWT token generation

> **Digunakan di:**

- JWT token generation/validation

- Logger level configuration

- HTTP server configuration

> **Database Configuration**
>
> **type** Database **struct** {  
> DBHost string *// PostgreSQL host  
> * DBPort string *// PostgreSQL port  
> * DBUser string *// PostgreSQL username  
> * DBPassword string *// PostgreSQL password  
> * DBName string *// Database name  
> * }
>
> **Environment Variables:**

- DB_HOST: Database server hostname

- DB_PORT: Database server port (default: 5432)

- DB_USER: Database username

- DB_PASSWORD: Database password

- DB_NAME: Database name

> **Digunakan di:**

- Database connection string construction

- GORM initialization

> **Redis Configuration**
>
> **type** Redis **struct** {  
> RHost string *// Redis host  
> * RPort string *// Redis port  
> * }
>
> **Environment Variables:**

- REDIS_HOST: Redis server hostname

- REDIS_PORT: Redis server port (default: 6379)

> **Digunakan di:**

- Redis client initialization

- Caching operations

- OTP storage

> **MinIO Configuration**
>
> **type** Minio **struct** {  
> Host string *// MinIO server endpoint  
> * AccessKey string *// MinIO access key  
> * SecretKey string *// MinIO secret key  
> * MaxOpenConn int *// Maximum connection pool size  
> * UseSSL int *// SSL enabled (1) or disabled (0)  
> * }
>
> **Environment Variables:**

- MINIO_HOST: MinIO server endpoint (e.g., localhost:9000)

- MINIO_ROOT_USER: Access key ID

- MINIO_ROOT_PASSWORD: Secret access key

- MINIO_MAX_OPEN_CONN: Connection pool size

- MINIO_USE_SSL: SSL flag (0 or 1)

> **Digunakan di:**

- MinIO client initialization

- File upload/download operations

> **SMTP Configuration (Zoho)**
>
> **type** ZSMTP **struct** {  
> ZSHost string *// SMTP server host  
> * ZSPort string *// SMTP server port  
> * ZSUser string *// SMTP username (email)  
> * ZSPassword string *// SMTP password  
> * ZSSecure string *// Security protocol (TLS/SSL)  
> * ZSAuth bool *// Authentication enabled  
> * }
>
> **Environment Variables:**

- ZOHO_SMTP_HOST: SMTP server hostname

- ZOHO_SMTP_PORT: SMTP server port (587 for TLS)

- ZOHO_SMTP_USER: Email address for authentication

- ZOHO_SMTP_PASSWORD: Email password

- ZOHO_SMTP_SECURE: Security protocol

- ZOHO_SMTP_AUTH: Authentication flag (true/false)

> **Digunakan di:**

- Email sending operations

- OTP email delivery

- Notification emails

> **Fonnte Configuration (WhatsApp OTP)**
>
> **type** Fonnte **struct** {  
> Token string *// Fonnte API token  
> * }
>
> **Environment Variables:**

- FONNTE_TOKEN: API token untuk Fonnte WhatsApp gateway

> **Digunakan di:**

- WhatsApp OTP sending

- Mobile authentication flow

> **Vault Configuration (HashiCorp Vault)**
>
> **type** Vault **struct** {  
> Addr string *// Vault server address  
> * RoleID string *// AppRole role ID  
> * SecretID string *// AppRole secret ID  
> * TransitPath string *// Transit engine mount path  
> * NIKEncryptionKey string *// Encryption key for NIK  
> * KartuEncryptionKey string *// Encryption key for Kartu Number  
> * }
>
> **Environment Variables:**

- VAULT_ADDR: Vault server URL (e.g., http://localhost:8200)

- VAULT_ROLE_ID: AppRole role ID untuk authentication

- VAULT_SECRET_ID: AppRole secret ID untuk authentication

- VAULT_TRANSIT_PATH: Mount path untuk transit engine (default: transit)

- VAULT_NIK_ENCRYPTION_KEY: Transit key name untuk NIK encryption

- VAULT_KARTU_ENCRYPTION_KEY: Transit key name untuk Kartu Number
  > encryption

> **Digunakan di:**

- Vault client initialization

- Data encryption/decryption operations

- Secure storage of sensitive data

> **LoadNative Function**
>
> **Fungsi:** Load dan validate semua environment variables.
>
> **Process Flow:**

1.  Load .env file jika ada menggunakan godotenv

2.  Lookup setiap environment variable

3.  Parse nilai dengan tipe data yang sesuai (string, int, bool)

4.  Collect missing variables ke dalam slice

5.  Return list missing variables dan error jika ada

> **Return Values:**

- \[\]string: List environment variables yang belum di-set

- error: Error jika file .env gagal di-load

> **Validation:**

- Integer values: Validate dan convert string ke int

- Boolean values: Check string "true"/"false"

- Required fields: Track missing variables

### Database Configuration

#### **config/db/config.go** {#configdbconfig.go .unnumbered}

> **Fungsi:** Setup PostgreSQL database connection menggunakan GORM ORM.
>
> **Global Variables**
>
> **var** DB \*gorm.DB

- Global database instance accessible dari seluruh aplikasi

- Digunakan oleh repository layer untuk database operations

> **SetupDatabase Function**
>
> **Fungsi:** Initialize database connection dengan GORM.
>
> **Parameter:**

- cfg env.Database: Database configuration dari environment variables

> **Process:**

1.  **Build DSN (Data Source Name)**

- dsn := fmt.Sprintf(\"host=%s user=%s password=%s dbname=%s port=%s
  > sslmode=disable\", \...)

  - Format: PostgreSQL connection string

  - Components: host, user, password, dbname, port

  - SSL mode: Disabled untuk development/staging

2.  **Open GORM Connection**

- db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  - Driver: PostgreSQL driver dari gorm.io

  - Config: Default GORM configuration

  - Error handling: Fatal error jika koneksi gagal

3.  **Assign to Global Variable**

- DB = db

  - Set global DB variable untuk digunakan aplikasi

> **Error Handling:**

- Connection failure: Log fatal dan stop application

- Tujuan: Memastikan database accessible sebelum aplikasi running

> **Digunakan di:**

- Repository initialization di router setup

- Semua database operations melalui repository layer

#### **Database Migrations** {#database-migrations-1 .unnumbered}

> **Location:** config/db/migrations/
>
> **Migration Tool:** Goose (github.com/pressly/goose)
>
> **Migration Files:**
>
> **20250906015548_create_user.sql**
>
> **Fungsi:** Create initial schema untuk users, roles, permissions, dan
> master data.
>
> **Objects Created:**

1.  **Custom Types**

    - gender: ENUM('female', 'male', 'other')

    - card_type: ENUM('produktif', 'afirmatif')

<!-- -->

2.  **Master Data Tables**

    - provinces: Data provinsi Indonesia (34 provinsi)

    - cities: Data kota/kabupaten Indonesia (514 kota)

<!-- -->

3.  **Authentication & Authorization Tables**

    - roles: Role definitions (superadmin, admin_screening,
      > admin_vendor, pelaku_usaha)

    - permissions: Permission definitions dengan hierarchy (parent_id)

    - role_permissions: Many-to-many mapping role ke permissions

    - users: User accounts dengan role assignment

<!-- -->

4.  **UMKM Data Table**

    - umkms: UMKM business information dengan encrypted sensitive data

> **Initial Data:**

- 34 provinsi

- 514 kota/kabupaten

- 4 roles dengan descriptions

- 20 permissions dengan hierarchy

- Role-permission mappings untuk setiap role

- 1 superadmin user (password: admin123)

> **20250906074403_create_applications_and_programs.sql**
>
> **Fungsi:** Create schema untuk programs dan benefits/requirements.
>
> **Objects Created:**

1.  **Custom Types**

    - program_type: ENUM('training', 'certification', 'funding')

    - training_type: ENUM('online', 'offline', 'hybrid')

<!-- -->

2.  **Tables**

    - programs: Program information dengan type-specific fields

    - program_benefits: Benefits associated dengan programs

    - program_requirements: Requirements associated dengan programs

> **Constraints:**

- Check constraint untuk training fields based on program type

> **20251108042550_create_applications.sql**
>
> **Fungsi:** Create schema untuk application workflow system.
>
> **Objects Created:**

1.  **Custom Types**

    - application_status: ENUM('screening', 'revised', 'final',
      > 'approved', 'rejected')

    - application_history_action: ENUM untuk tracking actions

    - application_type: ENUM('training', 'certification', 'funding')

    - document_type: ENUM untuk document classification

<!-- -->

2.  **Tables**

    - applications: Application submissions dengan status tracking

    - application_documents: Documents attached ke applications

    - application_histories: Audit trail untuk application decisions

> **20251113125118_create_sla.sql**
>
> **Fungsi:** Create SLA (Service Level Agreement) configuration table.
>
> **Objects Created:**

- slas: SLA configuration dengan max_days untuk setiap status

> **Initial Data:**

- Screening SLA: 7 hari

- Final SLA: 14 hari

> **20251114153138_create_otp.sql**
>
> **Fungsi:** Create OTP storage untuk mobile authentication.
>
> **Objects Created:**

1.  **Custom Types**

    - otp_status: ENUM('active', 'used')

<!-- -->

2.  **Tables**

    - otps: OTP codes dengan expiration dan temp token

> **Constraints:**

- Unique constraint pada temp_token

> **20251115080743_alter_nik_column.sql**
>
> **Fungsi:** Change NIK dan kartu_number column type untuk encrypted
> data.
>
> **Changes:**

- ALTER umkms.nik: VARCHAR(20) → TEXT

- ALTER umkms.kartu_number: VARCHAR(50) → TEXT

> **Reason:** Encrypted data memerlukan lebih banyak space
>
> **20251120013323_alter_umkms.sql**
>
> **Fungsi:** Add document fields ke UMKM table.
>
> **Changes:**

- ADD COLUMN revenue_record: TEXT

- ADD COLUMN business_permit: TEXT

> **20251120130113_alter_applications.sql**
>
> **Fungsi:** Create type-specific application tables.
>
> **Objects Created:**

- training_applications: Training-specific data

- certification_applications: Certification-specific data

- funding_applications: Funding-specific data (dengan financial data)

> **20251120135002_create_notification_vault_log.sql**
>
> **Fungsi:** Create notification system dan vault audit logging.
>
> **Objects Created:**

1.  **Custom Types**

    - decrypt_purpose: ENUM untuk categorize decrypt operations

    - notification_type: ENUM untuk notification categories

<!-- -->

2.  **Tables**

    - vault_decrypt_logs: Audit trail untuk decryption operations

    - notifications: Notification system untuk UMKM users

> **Indexes:**

- Multi-column indexes untuk fast queries

- Covering indexes untuk common query patterns

> **20251123073212_alter_umkms_add_pfp.sql**
>
> **Fungsi:** Add profile photo dan QR code fields.
>
> **Changes:**

- ADD COLUMN photo: TEXT

- ADD COLUMN qr_code: TEXT

> **20251123154625_alter_application_histories.sql**
>
> **Fungsi:** Make actioned_by nullable untuk system-generated history.
>
> **Changes:**

- ALTER application_histories.actioned_by: NOT NULL → NULLABLE

> **20251123161833_alter_application_documents.sql**
>
> **Fungsi:** Change document type ke text untuk flexibility.
>
> **Changes:**

- ALTER application_documents.type: ENUM → TEXT

> **20251128123439_create_news.sql**
>
> **Fungsi:** Create news/article management system.
>
> **Objects Created:**

1.  **Custom Types**

    - news_category: ENUM('announcement', 'success_story', 'event',
      > 'article')

<!-- -->

2.  **Tables**

    - news: News articles dengan SEO fields

    - news_tags: Tags untuk news categorization

> **Indexes:**

- slug, category, published status indexes

- Composite indexes untuk filtering

> **20251128143024_alter_umkms.sql**
>
> **Fungsi:** Change NIB dan NPWP ke TEXT type.
>
> **Changes:**

- ALTER umkms.nib: VARCHAR(50) → TEXT

- ALTER umkms.npwp: VARCHAR(50) → TEXT

> **20251129000600_new_role_permission.sql**
>
> **Fungsi:** Add news management permissions.
>
> **Changes:**

- INSERT 5 new permissions untuk news management

- Parent permission: MANAGE_NEWS

- Child permissions: CREATE_NEWS, EDIT_NEWS, DELETE_NEWS, VIEW_NEWS

#### **Database Seeders** {#database-seeders .unnumbered}

> **Location:** config/db/seeder/
>
> **Seeder Files:**
>
> **20251108085202_users.sql**
>
> **Fungsi:** Seed admin users untuk testing.
>
> **Data:**

- 2 Admin Screening users

- 2 Admin Vendor users

- Password: admin123 (hashed dengan bcrypt)

> **20251108085217_umkms.sql**
>
> **Fungsi:** Seed UMKM users dan business data.
>
> **Data:**

- 5 UMKM users dengan complete profile

- Diverse business types

- Sample NIK, phone, address data

- Mixed kartu types (produktif dan afirmatif)

> **20251108085230_programs.sql**
>
> **Fungsi:** Seed 30 sample programs (10 training, 10 certification, 10
> funding).
>
> **Data Structure:**

- Training programs: Dengan batch, schedule, location

- Certification programs: Dengan certification standards

- Funding programs: Dengan min/max amount, interest rate, tenure

> **Related Data:**

- Program benefits: 2-3 benefits per program

- Program requirements: 2-3 requirements per program

> **20251108085244_applications.sql**
>
> **Fungsi:** Seed 30 sample applications (10 per type).
>
> **Data:**

- Applications dari 5 UMKM users

- Different submission dates (varying age)

- Complete documents per application type

- Initial history entry (submit action)

### Logging Configuration

#### **config/log/logrus.go** {#configloglogrus.go .unnumbered}

> **Fungsi:** Setup structured logging menggunakan Logrus dengan custom
> formatter.
>
> **ApacheStyleFormatter**
>
> **Fungsi:** Custom log formatter dengan Apache/Nginx style dan color
> support.
>
> **Format Output:**
>
> \[02/Jan/2006:15:04:05 -0700\] LEVEL: message - field1=value1,
> field2=value2
>
> **Features:**

1.  **Timestamp**: Apache-style dengan timezone

2.  **Level**: Color-coded uppercase level string

3.  **Message**: Main log message

4.  **Fields**: Additional structured data

> **Color Codes by Level:**

- DEBUG: Cyan (1b\[36m)

- INFO: Green (1b\[32m)

- WARN: Yellow (1b\[33m)

- ERROR: Red (1b\[31m)

- FATAL/PANIC: Magenta (1b\[35m)

- TRACE: White (1b\[37m)

> **Color Support:**

- NoColors=true: Disable colors (untuk file output)

- NoColors=false: Enable colors (untuk console output)

> **SetupLogger Function**
>
> **Fungsi:** Initialize dan configure logger berdasarkan environment
> mode.
>
> **Configuration per Mode:**
>
> **Production Mode**
>
> **case** constant.PRODUCTION_MODE:  
> Log.SetLevel(logrus.InfoLevel)  
> Log.SetFormatter(&logrus.JSONFormatter{\...})  
> Log.SetOutput(file) *// .server.log*

- Level: INFO (minimal logging)

- Format: JSON untuk machine-readable logs

- Output: File .server.log

- Tujuan: Efficient logging untuk production environment

> **Staging Mode**
>
> **case** constant.STAGING_MODE:  
> Log.SetLevel(logrus.TraceLevel)  
> Log.SetFormatter(&ApacheStyleFormatter{NoColors: true})  
> Log.SetOutput(os.Stdout)

- Level: TRACE (verbose logging)

- Format: Apache-style tanpa colors

- Output: Console stdout

- Tujuan: Detailed logging untuk testing

> **Development Mode**
>
> **default**: *// Development  
> * Log.SetLevel(logrus.TraceLevel)  
> Log.SetFormatter(&ApacheStyleFormatter{NoColors: false})  
> Log.SetOutput(os.Stdout)

- Level: TRACE (maksimum detail)

- Format: Apache-style dengan colors

- Output: Console stdout

- Tujuan: Maximum visibility untuk development

> **Helper Functions**
>
> **Fungsi:** Wrapper functions untuk easier logging.
>
> **Info**
>
> **func** Info(msg string, fields \...**map**\[string\]**interface**{})

- Log informational messages

- Optional structured fields

> **Error**
>
> **func** Error(msg string, fields
> \...**map**\[string\]**interface**{})

- Log error messages

- Optional structured fields untuk error context

> **Warn**
>
> **func** Warn(msg string, fields \...**map**\[string\]**interface**{})

- Log warning messages

- Optional structured fields

> **Debug**
>
> **func** Debug(msg string, fields
> \...**map**\[string\]**interface**{})

- Log debug messages

- Only visible in TRACE level

> **Digunakan di:**

- Seluruh aplikasi untuk logging operations

- Error tracking

- Audit trail

- Performance monitoring

### Redis Configuration

#### **config/redis/config.go** {#configredisconfig.go .unnumbered}

> **Fungsi:** Setup Redis client connection.
>
> **Global Variables**
>
> **var** redis redisInstance  
>   
> **type** redisInstance **struct** {  
> Client \*redisPackage.Client  
> }

- Singleton pattern untuk Redis client

- Thread-safe access

> **SetupRedisDatabase Function**
>
> **Fungsi:** Initialize Redis client dengan configuration.
>
> **Process:**

1.  **Database Selection**

- **var** db int  
  > **if** env.Cfg.Server.Mode == constant.DEVELOPMENT_MODE {  
  > db = 1  
  > }

  - Development: Database 1

  - Production/Staging: Database 0 (default)

2.  **Create Client**

- rdb := redisPackage.NewClient(&redisPackage.Options{  
  > Addr: fmt.Sprintf(\"%s:%s\", cfg.RHost, cfg.RPort),  
  > DB: db,  
  > })

3.  **Test Connection**

- \_, err := rdb.Ping(rdb.Context()).Result()

  - Fatal error jika ping gagal

4.  **Set Global Instance**

- redis.Client = rdb

> **GetRedisRepository Function**
>
> **Fungsi:** Factory function untuk mendapatkan Redis repository
> instance.
>
> **Returns:** RedisRepository interface implementation
>
> **Validation:**

- Panic jika client belum initialized

- Tujuan: Enforce proper initialization order

#### **config/redis/redis.go** {#configredisredis.go .unnumbered}

> **Fungsi:** Redis repository implementation dengan common operations.
>
> **RedisRepository Interface**
>
> **Methods Available:**
>
> **Key-Value Operations**

- Set(ctx, key, value, exp): Set key dengan expiration

- Get(ctx, key): Get value by key

- Del(ctx, keys\...): Delete multiple keys

- Exists(ctx, keys\...): Check key existence

- Expire(ctx, key, exp): Set key expiration

> **Atomic Operations**

- SetNX(ctx, key, value, exp): Set if not exists

- Incr(ctx, key): Increment counter

> **Hash Operations**

- HSet(ctx, key, value, exp): Set hash field dengan expiration

- HGet(ctx, key, field): Get hash field value

> **Pub/Sub Operations**

- Publish(ctx, channel, message): Publish message ke channel

- Subscribe(ctx, channel): Subscribe ke channel

> **Bulk Operations**

- MGet(ctx, keys): Get multiple keys

- MSet(ctx, data): Set multiple keys

> **Scan Operations**

- Scan(ctx, match, count): Scan keys dengan pattern

- Keys(ctx, pattern): Get all keys matching pattern

> **Error Handling:**

- All methods return formatted errors dengan "redis:" prefix

- Consistent error wrapping untuk debugging

> **Digunakan di:**

- OTP storage dan validation

- Session management

- Caching operations

- Rate limiting

### MinIO Storage Configuration

#### **config/storage/constant.go** {#configstorageconstant.go .unnumbered}

> **Fungsi:** Define bucket names untuk different file types.
>
> **Bucket Constants:**
>
> **const** (  
> ProgramBucket = \"umkmgo-programs\"  
> UMKMBucket = \"umkmgo-umkms\"  
> ApplicationBucket = \"umkmgo-applications\"  
> NewsBucket = \"umkmgo-news\"  
> )
>
> **Usage:**

- Organize files by domain

- Separate concerns untuk different file types

- Easy bucket management

#### **config/storage/minio.go** {#configstorageminio.go .unnumbered}

> **Fungsi:** MinIO client manager dengan advanced features.
>
> **Data Structures**
>
> **MinIOConfig**
>
> **type** MinIOConfig **struct** {  
> Host string  
> AccessKey string  
> SecretKey string  
> UseSSL bool  
> MaxConnections int  
> ConnectTimeout time.Duration  
> RequestTimeout time.Duration  
> }
>
> **FileValidationConfig**
>
> **type** FileValidationConfig **struct** {  
> AllowedExtensions \[\]string  
> MaxFileSize int64 *// in bytes  
> * MinFileSize int64 *// in bytes  
> * }
>
> **UploadRequest**
>
> **type** UploadRequest **struct** {  
> Base64Data string  
> Prefix string  
> BucketName string  
> Validation \*FileValidationConfig  
> }
>
> **UploadResponse**
>
> **type** UploadResponse **struct** {  
> BucketName string  
> ObjectName string  
> Size int64  
> URL string  
> Ext string  
> ETag string  
> }
>
> **MinIOManager**
>
> **Fungsi:** Singleton manager untuk MinIO operations.
>
> **Features:**

1.  **Connection Pooling**

    - Reuse connections untuk efficiency

    - Thread-safe operations dengan mutex

<!-- -->

2.  **Bucket Caching**

- bucketCache **map**\[string\]bool

  - Cache bucket existence checks

  - Reduce redundant API calls

3.  **Health Checking**

- **func** (m \*MinIOManager) IsReady() bool

  - Check if client is ready

  - Used in health check endpoints

> **SetupMinio Function**
>
> **Fungsi:** Initialize global MinIO manager (singleton pattern).
>
> **Process:**

1.  Use sync.Once untuk ensure single initialization

2.  Create MinIOConfig dari environment variables

3.  Call newMinIOManager untuk create instance

4.  Panic jika initialization gagal

> **Core Methods**
>
> **UploadFile**
>
> **func** (m \*MinIOManager) UploadFile(ctx, request)
> (\*UploadResponse, error)
>
> **Process:**

1.  Validate base64 data

2.  Check MinIO client readiness

3.  Validate bucket existence

4.  Decode base64 dan extract content type

5.  Apply validation rules (size, extension)

6.  Generate unique object name dengan timestamp

7.  Upload ke MinIO

8.  Return response dengan URL

> **Features:**

- Automatic content type detection

- File size validation

- Extension validation

- Timestamp-based naming

- URL generation

> **GetFile**
>
> **func** (m \*MinIOManager) GetFile(ctx, bucketName, objectName)
> (\*minio.Object, error)

- Retrieve file object dari MinIO

- Return streaming object untuk download

> **DeleteFile**
>
> **func** (m \*MinIOManager) DeleteFile(ctx, bucketName, objectName)
> error

- Delete file dari MinIO

- Used saat update file (delete old, upload new)

> **GetPresignedURL**
>
> **func** (m \*MinIOManager) GetPresignedURL(ctx, bucketName,
> objectName, expires) (string, error)

- Generate temporary signed URL

- Used untuk secure direct access

- Configurable expiration time

> **ListObjects**
>
> **func** (m \*MinIOManager) ListObjects(ctx, bucketName, prefix)
> (\[\]minio.ObjectInfo, error)

- List objects dengan prefix filter

- Used untuk file management

> **Helper Functions**
>
> **getContentTypeFromData**
>
> **func** getContentTypeFromData(data \[\]byte) string

- Detect content type dari file signature

- Support common file types (JPEG, PNG, PDF, ZIP, dll)

> **getExtensionFromContentType**
>
> **func** getExtensionFromContentType(contentType string) string

- Map content type ke file extension

- Support 30+ content types

> **isBase64**
>
> **func** isBase64(s string) bool

- Validate base64 string

- Handle data URI prefix

- Regex validation

> **Validation Configs**
>
> **CreateDefaultValidationConfig**
>
> **func** CreateDefaultValidationConfig() \*FileValidationConfig

- Extensions: .jpg, .jpeg, .png, .gif, .pdf, .doc, .docx, .svg

- Max size: 10MB

- Min size: 1 byte

> **CreateImageValidationConfig**
>
> **func** CreateImageValidationConfig() \*FileValidationConfig

- Extensions: .jpg, .jpeg, .png, .gif, .webp, .pdf

- Max size: 5MB

- Min size: 1 byte

> **Digunakan di:**

- Program banner/logo upload

- UMKM document upload

- Application document upload

- News thumbnail upload

- Profile photo upload

### Vault Configuration

#### **config/vault/vault.go** {#configvaultvault.go .unnumbered}

> **Fungsi:** HashiCorp Vault integration untuk encryption/decryption
> operations.
>
> **Global Variables**
>
> **var** VaultClient \*vault.Client  
> **var** vaultConfig env.Vault
>
> **SetupVault Function**
>
> **Fungsi:** Initialize Vault client dengan AppRole authentication.
>
> **Process:**

1.  **Store Configuration**

- vaultConfig = vaultCfg

2.  **Create Client**

- cfg := vault.DefaultConfig()  
  > cfg.Address = vaultCfg.Addr  
  > client, err := vault.NewClient(cfg)

3.  **AppRole Login**

- loginAppRole(client)

  - Authenticate menggunakan role_id dan secret_id

  - Receive client token

4.  **Set Global Client**

- VaultClient = client

5.  **Start Auto Renewal**

- **go** autoRenewToken()

  - Background goroutine untuk token renewal

> **loginAppRole Function**
>
> **Fungsi:** Authenticate ke Vault menggunakan AppRole method.
>
> **Process:**

1.  Prepare login payload dengan role_id dan secret_id

2.  Write to auth/approle/login endpoint

3.  Extract client token dari response

4.  Set token ke client

5.  Log token TTL information

> **Error Handling:**

- Return error jika login gagal

- Return error jika no token in response

> **autoRenewToken Function**
>
> **Fungsi:** Background process untuk automatic token renewal.
>
> **Process:**

1.  Create ticker dengan interval 30 menit

2.  Loop setiap tick:

    - Attempt token renewal (extend TTL 3600 seconds)

    - Jika gagal: Re-authenticate menggunakan AppRole

    - Log success/failure

> **Tujuan:**

- Prevent token expiration

- Maintain continuous access ke Vault

- Automatic recovery dari token issues

> **Encryption/Decryption Methods**
>
> **EncryptTransit**
>
> **func** EncryptTransit(ctx, transitMount, transitKey, plaintext)
> (string, error)
>
> **Process:**

1.  Base64 encode plaintext (Transit requirement)

2.  Build path: {transitMount}/encrypt/{transitKey}

3.  Call Vault encrypt endpoint

4.  Return ciphertext string (format: vault:v1:\...)

> **Parameters:**

- transitMount: Transit engine mount path (e.g., "transit")

- transitKey: Encryption key name (e.g., "nik-key")

- plaintext: Data to encrypt (bytes)

> **Returns:**

- Ciphertext string safe untuk storage

> **DecryptTransit**
>
> **func** DecryptTransit(ctx, transitMount, transitKey, ciphertext)
> (\[\]byte, error)
>
> **Process:**

1.  Build path: {transitMount}/decrypt/{transitKey}

2.  Call Vault decrypt endpoint dengan ciphertext

3.  Extract base64 plaintext dari response

4.  Decode base64 ke bytes

5.  Return plaintext bytes

> **Parameters:**

- transitMount: Transit engine mount path

- transitKey: Encryption key name

- ciphertext: Encrypted data (vault:v1:... format)

> **Returns:**

- Decrypted plaintext bytes

> **Logging Integration**
>
> **DecryptParams**
>
> **type** DecryptParams **struct** {  
> UserID int  
> UMKMID \*int  
> FieldName string  
> TableName string  
> RecordID int  
> Purpose string  
> IPAddress string  
> UserAgent string  
> RequestID string  
> }
>
> **Fungsi:** Parameters untuk audit logging decryption operations.
>
> **DecryptWithLog**
>
> **func** DecryptWithLog(ctx, ciphertext, encryptionKey, params,
> vaultLogRepo) (string, error)
>
> **Process:**

1.  Perform decryption menggunakan DecryptTransit

2.  Create log entry dengan parameters

3.  Set success flag based on error

4.  Store error message if failed

5.  Save log entry to database (asynchronously)

6.  Return decrypted string atau error

> **Parameters:**

- ciphertext: Encrypted data (vault:v1:... format)

- encryptionKey: Transit key name untuk decryption

- params: DecryptParams struct dengan metadata

- vaultLogRepo: Repository untuk save log

> **Returns:**

- Decrypted string value

- Error jika decryption failed

> **Log Fields:**

- user_id: ID user yang melakukan decrypt

- umkm_id: ID UMKM yang data-nya di-decrypt (nullable)

- field_name: Nama field yang di-decrypt (e.g., "nik", "kartu_number")

- table_name: Nama table sumber data

- record_id: ID record yang di-decrypt

- purpose: Tujuan decryption (e.g., "view_detail", "export_data")

- ip_address: IP address user

- user_agent: Browser/client user agent

- request_id: Unique request identifier untuk tracing

- success: Boolean flag (true/false)

- error_message: Error detail jika gagal

- created_at: Timestamp log

> **Use Cases:**

- NIK decryption untuk detail view

- Kartu number decryption untuk admin review

- Audit trail untuk compliance

- Security monitoring

> **Digunakan di:**

- UMKM detail endpoints

- Application review endpoints

- Export operations dengan sensitive data

### HTTP Router Configuration

#### **interface/http/router/router.go** {#interfacehttprouterrouter.go .unnumbered}

> **Fungsi:** Setup HTTP server dan routing menggunakan Fiber framework.
>
> **Fiber Configuration**
>
> **Global Settings:**
>
> fiber.Config{  
> Prefork: env.Cfg.Server.Mode == constant.DEVELOPMENT_MODE,  
> ServerHeader: \"UMKMGo API\",  
> StrictRouting: false,  
> CaseSensitive: false,  
> BodyLimit: 50 \* 1024 \* 1024, *// 50MB  
> * ReadTimeout: time.Second \* 300,  
> WriteTimeout: time.Second \* 300,  
> IdleTimeout: time.Second \* 120,  
> DisableKeepalive: false,  
> }
>
> **Configuration Details:**
>
> **Prefork Mode**
>
> Prefork: env.Cfg.Server.Mode == constant.DEVELOPMENT_MODE

- **Development**: Prefork enabled untuk better debugging

- **Production/Staging**: Prefork disabled

- **Purpose**: Balance between development experience dan production
  > performance

> **Server Header**
>
> ServerHeader: \"UMKMGo API\"

- Custom server identification

- Shown in HTTP response headers

- Branding dan security (hide framework version)

> **Routing Configuration**
>
> StrictRouting: false  
> CaseSensitive: false

- **StrictRouting**: /path dan /path/ treated the same

- **CaseSensitive**: /Path dan /path treated the same

- **Purpose**: Flexible routing untuk user convenience

> **Body Limit**
>
> BodyLimit: 50 \* 1024 \* 1024 *// 50MB*

- Maximum request body size: 50MB

- **Purpose**: Allow file uploads (documents, images)

- **Validation**: Additional validation di handler level

> **Timeouts**
>
> ReadTimeout: time.Second \* 300 *// 5 minutes  
> * WriteTimeout: time.Second \* 300 *// 5 minutes  
> * IdleTimeout: time.Second \* 120 *// 2 minutes*

- **ReadTimeout**: Max time to read entire request

- **WriteTimeout**: Max time to write response

- **IdleTimeout**: Max time keep-alive connection can idle

- **Purpose**: Prevent hanging connections, handle long operations (file
  > uploads)

> **Keep-Alive**
>
> DisableKeepalive: false

- Keep-alive enabled untuk connection reuse

- **Performance**: Reduce connection overhead

- **Production**: Better throughput untuk multiple requests

> **SetupRouter Function**
>
> **Fungsi:** Initialize dan configure HTTP router dengan semua routes.
>
> **Process:**

1.  **Create Fiber App**

- router := fiber.New(fiber.Config{\...})

2.  **Apply Global Middleware**

- router.Use(middleware.CORS(), middleware.Logger())

  - CORS: Enable cross-origin requests

  - Logger: Log semua HTTP requests

3.  **Health Check Route**

- router.Get(\"test\", **func**(c \*fiber.Ctx) error {  
  > **return** c.JSON(fiber.Map{\"message\": \"Hello World!\"})  
  > })

  - Simple health check endpoint

  - Path: /test

  - No authentication required

4.  **API Version Group**

- version := router.Group(\"/v1\")

  - All API routes under /v1 prefix

  - **Purpose**: API versioning untuk backward compatibility

5.  **Register Route Modules**

- routes.UserRoutes(version, db.DB, redis.GetRedisRepository(),
  > storage.MinioClient)  
  > routes.ProgramRoutes(version, db.DB, redis.GetRedisRepository(),
  > storage.MinioClient)  
  > routes.ApplicationRoutes(version, db.DB,
  > redis.GetRedisRepository())  
  > routes.DashboardRoutes(version, db.DB)  
  > routes.SLARoutes(version, db.DB)  
  > routes.NewsRoutes(version, db.DB, storage.MinioClient)  
  > routes.MobileRoutes(version, db.DB, storage.MinioClient)

- **Dependencies Injection:**

  - db.DB: PostgreSQL database connection

  - redis.GetRedisRepository(): Redis client wrapper

  - storage.MinioClient: MinIO client manager

<!-- -->

- **Route Modules:**

  - **UserRoutes**: User management, authentication, UMKM profiles

  - **ProgramRoutes**: Program CRUD, benefits, requirements

  - **ApplicationRoutes**: Application submissions, workflow, documents

  - **DashboardRoutes**: Statistics, analytics, reports

  - **SLARoutes**: SLA configuration dan monitoring

  - **NewsRoutes**: News/article management

  - **MobileRoutes**: Mobile-specific endpoints (OTP, public data)

6.  **Route Logging**

- **for** \_, routes := **range** router.Stack() {  
  > **for** \_, r := **range** routes {  
  > log.Log.Infof(\"Registered route - METHOD: %s, PATH: %s\", r.Method,
  > r.Path)  
  > }  
  > }

  - Log semua registered routes saat startup

  - **Purpose**: Debugging, documentation, verification

7.  **Return Router**

- **return** router

  - Return configured Fiber app

  - Ready untuk .Listen(port)

> **Route Registration Pattern:**
>
> Setiap route module mengikuti pattern yang sama:
>
> **func** ModuleRoutes(  
> router fiber.Router,  
> db \*gorm.DB,  
> redisRepo redis.RedisRepository,  
> minioClient \*storage.MinIOManager,  
> ) {  
> *// 1. Initialize Repository  
> * repo := repository.NewModuleRepository(db)  
>   
> *// 2. Initialize Service  
> * service := service.NewModuleService(repo, redisRepo, minioClient)  
>   
> *// 3. Initialize Handler  
> * handler := handler.NewModuleHandler(service)  
>   
> *// 4. Register Routes  
> * module := router.Group(\"/module\")  
> {  
> module.Get(\"/\", middleware.Auth(roles\...), handler.GetAll)  
> module.Get(\"/:id\", middleware.Auth(roles\...), handler.GetByID)  
> module.Post(\"/\", middleware.Auth(roles\...), handler.Create)  
> module.Put(\"/:id\", middleware.Auth(roles\...), handler.Update)  
> module.Delete(\"/:id\", middleware.Auth(roles\...), handler.Delete)  
> }  
> }

## API Routes dan Handlers

### Web Authentication Routes

> **Base Path:** /v1/webauth
>
> Endpoints

- **POST** /login → User_handler.Login

  - Handler: Login untuk web dashboard

  - Dependencies: UsersService

- **POST** /register → User_handler.Register

  - Handler: Register user baru untuk web dashboard

  - Dependencies: UsersService

- **PUT** /profile → User_handler.UpdateProfile

  - Handler: Update profil user yang sedang login

  - Dependencies: UsersService

### Mobile Authentication Routes

> **Base Path:** /v1/mobileauth
>
> Endpoints

- **GET** /meta → User_handler.GetMeta

  - Handler: Mendapatkan data master (provinsi dan kota)

  - Dependencies: UsersService

- **POST** /login → User_handler.LoginMobile

  - Handler: Login untuk aplikasi mobile

  - Dependencies: UsersService

- **POST** /register → User_handler.RegisterMobile

  - Handler: Register user mobile (mengirim OTP)

  - Dependencies: UsersService

- **POST** /register/profile → User_handler.RegisterMobileProfile

  - Handler: Melengkapi profil setelah verifikasi OTP

  - Dependencies: UsersService

- **POST** /forgot-password → User_handler.ForgotPassword

  - Handler: Request reset password (mengirim OTP)

  - Dependencies: UsersService

- **POST** /reset-password → User_handler.ResetPassword

  - Handler: Reset password dengan temp token

  - Dependencies: UsersService

- **POST** /verify/otp → User_handler.VerifyOTP

  - Handler: Verifikasi kode OTP

  - Dependencies: UsersService

### User Management Routes

> **Base Path:** /v1/users **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** / → User_handler.GetAllUsers

  - Handler: Mendapatkan semua data user

  - Dependencies: UsersService

- **GET** /:id → User_handler.GetUserByID

  - Handler: Mendapatkan data user berdasarkan ID

  - Dependencies: UsersService

- **PUT** /:id → User_handler.UpdateUser

  - Handler: Update data user berdasarkan ID

  - Dependencies: UsersService

- **DELETE** /:id → User_handler.DeleteUser

  - Handler: Soft delete user berdasarkan ID

  - Dependencies: UsersService

### Permissions Routes

> **Base Path:** /v1 **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** /permissions → User_handler.GetListPermissions

  - Handler: Mendapatkan daftar semua permission

  - Dependencies: UsersService

- **GET** /role-permissions → User_handler.GetListRolePermissions

  - Handler: Mendapatkan mapping role dan permissions

  - Dependencies: UsersService

- **POST** /role-permissions → User_handler.UpdateRolePermissions

  - Handler: Update permissions untuk suatu role

  - Dependencies: UsersService

### Programs Routes

> **Base Path:** /v1/programs **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** / → programHandler.GetAllPrograms

  - Handler: Mendapatkan semua program (training, certification,
    > funding)

  - Dependencies: ProgramsService (ProgramsRepository, UsersRepository,
    > Redis, MinIO)

- **GET** /:id → programHandler.GetProgramByID

  - Handler: Mendapatkan detail program berdasarkan ID

  - Dependencies: ProgramsService

- **POST** / → programHandler.CreateProgram

  - Handler: Membuat program baru

  - Dependencies: ProgramsService

- **PUT** /:id → programHandler.UpdateProgram

  - Handler: Update data program

  - Dependencies: ProgramsService

- **PUT** /activate/:id → programHandler.ActivateProgram

  - Handler: Mengaktifkan program

  - Dependencies: ProgramsService

- **PUT** /deactivate/:id → programHandler.DeactivateProgram

  - Handler: Menonaktifkan program

  - Dependencies: ProgramsService

- **DELETE** /:id → programHandler.DeleteProgram

  - Handler: Soft delete program

  - Dependencies: ProgramsService

### Applications Routes

> **Base Path:** /v1/applications **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** / → applicationHandler.GetAllApplications

  - Handler: Mendapatkan semua aplikasi dengan filter optional

  - Dependencies: ApplicationsService (ApplicationsRepository,
    > UsersRepository, NotificationRepository, SLARepository,
    > VaultDecryptLogRepository)

- **GET** /:id → applicationHandler.GetApplicationByID

  - Handler: Mendapatkan detail aplikasi berdasarkan ID

  - Dependencies: ApplicationsService

> Screening Decision Endpoints

- **PUT** /screening-approve/:id → applicationHandler.ScreeningApprove

  - Handler: Approve aplikasi pada tahap screening

  - Dependencies: ApplicationsService

- **PUT** /screening-reject/:id → applicationHandler.ScreeningReject

  - Handler: Reject aplikasi pada tahap screening

  - Dependencies: ApplicationsService

- **PUT** /screening-revise/:id → applicationHandler.ScreeningRevise

  - Handler: Meminta revisi aplikasi pada tahap screening

  - Dependencies: ApplicationsService

> Final Decision Endpoints

- **PUT** /final-approve/:id → applicationHandler.FinalApprove

  - Handler: Approve aplikasi pada tahap final

  - Dependencies: ApplicationsService

- **PUT** /final-reject/:id → applicationHandler.FinalReject

  - Handler: Reject aplikasi pada tahap final

  - Dependencies: ApplicationsService

### Dashboard Routes

> **Base Path:** /v1/dashboard **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** /umkm-by-card-type → dashboardHandler.GetUMKMByCardType

  - Handler: Statistik UMKM berdasarkan tipe kartu

  - Dependencies: DashboardService (DashboardRepository)

- **GET** /application-status-summary →
  > dashboardHandler.GetApplicationStatusSummary

  - Handler: Ringkasan status aplikasi

  - Dependencies: DashboardService

- **GET** /application-status-detail →
  > dashboardHandler.GetApplicationStatusDetail

  - Handler: Detail status aplikasi

  - Dependencies: DashboardService

- **GET** /application-by-type → dashboardHandler.GetApplicationByType

  - Handler: Statistik aplikasi berdasarkan tipe (training,
    > certification, funding)

  - Dependencies: DashboardService

### SLA Routes

> **Base Path:** /v1/sla **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** /screening → slaHandler.GetSLAScreening

  - Handler: Mendapatkan konfigurasi SLA untuk screening

  - Dependencies: SLAService (SLARepository)

- **GET** /final → slaHandler.GetSLAFinal

  - Handler: Mendapatkan konfigurasi SLA untuk final decision

  - Dependencies: SLAService

- **PUT** /screening → slaHandler.UpdateSLAScreening

  - Handler: Update konfigurasi SLA screening

  - Dependencies: SLAService

- **PUT** /final → slaHandler.UpdateSLAFinal

  - Handler: Update konfigurasi SLA final

  - Dependencies: SLAService

- **POST** /export-applications → slaHandler.ExportApplications

  - Handler: Export data aplikasi ke PDF/Excel

  - Dependencies: SLAService

- **POST** /export-programs → slaHandler.ExportPrograms

  - Handler: Export data program ke PDF/Excel

  - Dependencies: SLAService

### News Routes

> **Base Path:** /v1/news **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** / → newsHandler.GetAllNews

  - Handler: Mendapatkan semua berita dengan filter

  - Dependencies: NewsService (NewsRepository, MinIO)

- **GET** /:id → newsHandler.GetNewsByID

  - Handler: Mendapatkan detail berita berdasarkan ID

  - Dependencies: NewsService

- **POST** / → newsHandler.CreateNews

  - Handler: Membuat berita baru

  - Dependencies: NewsService

- **PUT** /:id → newsHandler.UpdateNews

  - Handler: Update berita

  - Dependencies: NewsService

- **DELETE** /:id → newsHandler.DeleteNews

  - Handler: Soft delete berita

  - Dependencies: NewsService

- **PUT** /publish/:id → newsHandler.PublishNews

  - Handler: Publish berita

  - Dependencies: NewsService

- **PUT** /unpublish/:id → newsHandler.UnpublishNews

  - Handler: Unpublish berita

  - Dependencies: NewsService

### Mobile Routes

> **Base Path:** /v1/mobile **Middleware:** MobileAuthMiddleware

#### Dashboard {#dashboard .unnumbered}

- **GET** /dashboard → mobileHandler.GetDashboard

  - Handler: Mendapatkan dashboard data untuk mobile

  - Dependencies: MobileService (MobileRepository, ProgramsRepository,
    > NotificationRepository, VaultDecryptLogRepository,
    > ApplicationsRepository, SLARepository, MinIO)

#### Programs {#programs .unnumbered}

> **Base Path:** /v1/mobile/programs

- **GET** /training → mobileHandler.GetTrainingPrograms

  - Handler: Mendapatkan daftar program training

  - Dependencies: MobileService

- **GET** /certification → mobileHandler.GetCertificationPrograms

  - Handler: Mendapatkan daftar program certification

  - Dependencies: MobileService

- **GET** /funding → mobileHandler.GetFundingPrograms

  - Handler: Mendapatkan daftar program funding

  - Dependencies: MobileService

- **GET** /:id → mobileHandler.GetProgramDetail

  - Handler: Mendapatkan detail program

  - Dependencies: MobileService

#### Profile {#profile .unnumbered}

> **Base Path:** /v1/mobile/profile

- **GET** / → mobileHandler.GetUMKMProfile

  - Handler: Mendapatkan profil UMKM

  - Dependencies: MobileService

- **PUT** / → mobileHandler.UpdateUMKMProfile

  - Handler: Update profil UMKM

  - Dependencies: MobileService

#### Documents {#documents .unnumbered}

> **Base Path:** /v1/mobile/documents

- **GET** / → mobileHandler.GetUMKMDocuments

  - Handler: Mendapatkan daftar dokumen UMKM

  - Dependencies: MobileService

- **POST** /upload → mobileHandler.UploadDocument

  - Handler: Upload dokumen (NIB, NPWP, revenue record, business permit)

  - Dependencies: MobileService

#### Applications {#applications .unnumbered}

> **Base Path:** /v1/mobile/applications

- **POST** /training → mobileHandler.CreateTrainingApplication

  - Handler: Membuat aplikasi training

  - Dependencies: MobileService

- **POST** /certification → mobileHandler.CreateCertificationApplication

  - Handler: Membuat aplikasi certification

  - Dependencies: MobileService

- **POST** /funding → mobileHandler.CreateFundingApplication

  - Handler: Membuat aplikasi funding

  - Dependencies: MobileService

- **GET** / → mobileHandler.GetApplicationList

  - Handler: Mendapatkan daftar aplikasi user

  - Dependencies: MobileService

- **GET** /:id → mobileHandler.GetApplicationDetail

  - Handler: Mendapatkan detail aplikasi

  - Dependencies: MobileService

- **PUT** /:id → mobileHandler.ReviseApplication

  - Handler: Revisi aplikasi yang diminta perubahan

  - Dependencies: MobileService

#### Notifications {#notifications .unnumbered}

> **Base Path:** /v1/mobile/notifications

- **GET** / → mobileHandler.GetNotificationsByUMKMID

  - Handler: Mendapatkan daftar notifikasi

  - Dependencies: MobileService

- **GET** /unread-count → mobileHandler.GetUnreadCount

  - Handler: Mendapatkan jumlah notifikasi yang belum dibaca

  - Dependencies: MobileService

- **PUT** /mark-as-read/:id → mobileHandler.MarkNotificationsAsRead

  - Handler: Tandai notifikasi sebagai sudah dibaca

  - Dependencies: MobileService

- **PUT** /mark-all-as-read → mobileHandler.MarkAllNotificationsAsRead

  - Handler: Tandai semua notifikasi sebagai sudah dibaca

  - Dependencies: MobileService

#### News (Mobile) {#news-mobile .unnumbered}

> **Base Path:** /v1/mobile/news

- **GET** / → mobileHandler.GetPublishedNews

  - Handler: Mendapatkan berita yang sudah dipublish dengan filter

  - Dependencies: MobileService

- **GET** /:slug → mobileHandler.GetNewsDetailBySlug

  - Handler: Mendapatkan detail berita berdasarkan slug

  - Dependencies: MobileService

### Vault Decrypt Logs Routes

> **Base Path:** /v1/vault-decrypt-logs **Middleware:** AuthMiddleware
>
> Endpoints

- **GET** / → vaultDecryptLogHandler.GetLogs

  - Handler: Mendapatkan semua log dekripsi dengan pagination

  - Dependencies: VaultDecryptLogService (VaultDecryptLogRepository)

- **GET** /user → vaultDecryptLogHandler.GetLogsByUserID

  - Handler: Mendapatkan log dekripsi berdasarkan user ID

  - Dependencies: VaultDecryptLogService

- **GET** /umkm/:umkm_id → vaultDecryptLogHandler.GetLogsByUMKMID

  - Handler: Mendapatkan log dekripsi berdasarkan UMKM ID

  - Dependencies: VaultDecryptLogService

## Middleware

### AuthMiddleware  {#authmiddleware}

> Digunakan untuk endpoint web dashboard yang membutuhkan autentikasi
> admin.

### MobileAuthMiddleware  {#mobileauthmiddleware}

> Digunakan untuk endpoint mobile yang membutuhkan autentikasi pelaku
> usaha (UMKM).

### CORSMiddleware  {#corsmiddleware}

> Digunakan untuk pengaturan CORS dari aplikasi.

### LoggerMiddleware  {#loggermiddleware}

> Digunakan untuk memantau aktivitas dan performa API serta memudahkan
> debugging.

####  {#section .unnumbered}

## Service Layer

### Users Service

> Service untuk mengelola operasi terkait user, autentikasi, dan role
> permissions.
>
> **Dependencies:**

- UsersRepository

- OTPRepository

- RedisRepository

- MinIOManager

#### **Register** {#register .unnumbered}

> **Fungsi:** Registrasi user baru untuk web dashboard
>
> **Input:**

- ctx context.Context - Context untuk operasi database

- user dto.Users - Data user yang akan didaftarkan (name, email,
  > password, confirm_password, role_id)

> **Process:**

1.  Validasi input (name, email, password, confirm_password, role_id
    > tidak boleh kosong)

2.  Validasi format email menggunakan utils.EmailValidator

3.  Cek apakah email sudah terdaftar di database

4.  Validasi password (minimal 8 karakter, mengandung huruf dan angka)
    > menggunakan utils.PasswordValidator

5.  Validasi password dan confirm password harus sama

6.  Validasi role_id harus valid

7.  Hash password menggunakan utils.PasswordHashing (bcrypt)

8.  Simpan user baru ke database

> **Output:**

- dto.Users - Data user yang berhasil dibuat (id, name, email)

- error - Error jika terjadi kesalahan

> **Struct/Type yang Digunakan:**

- dto.Users

- model.User

#### **Login** {#login .unnumbered}

> **Fungsi:** Login user untuk web dashboard
>
> **Input:**

- ctx context.Context

- user dto.Users - Data login (email, password)

> **Process:**

1.  Validasi email dan password tidak boleh kosong

2.  Cek user berdasarkan email di database

3.  Verifikasi password menggunakan utils.ComparePass

4.  Validasi user harus aktif (is_active = true)

5.  Ambil role name dari database

6.  Update last_login_at ke waktu sekarang

7.  Ambil list permissions berdasarkan role_id

8.  Generate JWT token menggunakan utils.GenerateWebToken

> **Output:**

- \*string - JWT token

- error

> **Struct/Type yang Digunakan:**

- dto.Users

- model.User

- model.Role

#### **UpdateProfile** {#updateprofile .unnumbered}

> **Fungsi:** Update profil user yang sedang login
>
> **Input:**

- ctx context.Context

- id int - ID user yang akan diupdate

- userNew dto.Users - Data baru (name, email)

> **Process:**

1.  Ambil data user berdasarkan ID

2.  Validasi name dan email tidak boleh kosong

3.  Validasi format email

4.  Cek apakah email sudah digunakan user lain

5.  Update name dan email

6.  Simpan perubahan ke database

> **Output:**

- dto.Users - Data user yang berhasil diupdate

- error

#### **RegisterMobile** {#registermobile .unnumbered}

> **Fungsi:** Registrasi user mobile (mengirim OTP)
>
> **Input:**

- ctx context.Context

- email string

- phone string

> **Process:**

1.  Validasi email dan phone tidak boleh kosong

2.  Validasi format email

3.  Normalisasi nomor telepon menggunakan utils.NormalizePhone

4.  Cek apakah email sudah terdaftar

5.  Generate OTP code menggunakan utils.GenerateOTP

6.  Simpan OTP ke database dengan status 'active' dan expires_at 5 menit

7.  Kirim OTP ke WhatsApp menggunakan Fonnte API

> **Output:**

- error

> **Struct/Type yang Digunakan:**

- model.OTP

> **Utils yang Digunakan:**

- utils.EmailValidator

- utils.NormalizePhone

- utils.GenerateOTP

- otp.InitVendor

- otp.SendOTP

#### **VerifyOTP** {#verifyotp .unnumbered}

> **Fungsi:** Verifikasi kode OTP
>
> **Input:**

- ctx context.Context

- phone string

- code string - Kode OTP

> **Process:**

1.  Normalisasi nomor telepon

2.  Ambil data OTP berdasarkan phone number

3.  Validasi OTP tidak expired dan status masih 'active'

4.  Validasi kode OTP cocok

5.  Generate temp_token menggunakan utils.RandomString

6.  Update OTP dengan temp_token baru

> **Output:**

- \*string - Temp token untuk registrasi profil

- error

> **Utils yang Digunakan:**

- utils.NormalizePhone

- utils.RandomString

#### **RegisterMobileProfile** {#registermobileprofile .unnumbered}

> **Fungsi:** Melengkapi profil UMKM setelah verifikasi OTP
>
> **Input:**

- ctx context.Context

- user dto.UMKMMobile - Data lengkap profil UMKM

- tempToken string - Token dari verifikasi OTP

> **Process:**

1.  Validasi temp_token dan ambil data OTP

2.  Validasi semua input wajib (fullname, business_name, nik,
    > birth_date, gender, address, province_id, city_id, district,
    > postal_code, kartu_type, kartu_number, password)

3.  Validasi password (min 8 karakter, mengandung huruf dan angka)

4.  Normalisasi nomor telepon

5.  Parse birth_date ke format time.Time

6.  Hash password menggunakan bcrypt

7.  Ambil role 'pelaku_usaha' dari database

8.  Enkripsi NIK menggunakan Vault Transit dengan key khusus NIK

9.  Enkripsi Kartu Number menggunakan Vault Transit dengan key khusus
    > Kartu

10. Generate QR Code dari Kartu Number menggunakan utils.GenerateQRCode

11. Upload QR Code ke MinIO

12. Simpan data UMKM dan User ke database

13. Update status OTP menjadi 'used'

14. Generate JWT token untuk mobile menggunakan
    > utils.GenerateMobileToken

> **Output:**

- \*string - JWT token

- error

> **Struct/Type yang Digunakan:**

- dto.UMKMMobile

- model.UMKM

- model.User

- model.Role

> **Utils yang Digunakan:**

- utils.NormalizePhone

- utils.PasswordValidator

- utils.PasswordHashing

- utils.GenerateQRCode

- utils.GenerateFileName

- vault.EncryptTransit

- storage.MinIOManager.UploadFile

#### **LoginMobile** {#loginmobile .unnumbered}

> **Fungsi:** Login untuk aplikasi mobile
>
> **Input:**

- ctx context.Context

- user dto.UMKMMobile - Data login (phone, password)

> **Process:**

1.  Validasi phone dan password tidak boleh kosong

2.  Normalisasi nomor telepon

3.  Cek UMKM berdasarkan phone number

4.  Verifikasi password menggunakan utils.ComparePass

5.  Generate JWT token untuk mobile

> **Output:**

- \*string - JWT token

- error

> **Utils yang Digunakan:**

- utils.NormalizePhone

- utils.ComparePass

- utils.GenerateMobileToken

#### **ForgotPassword** {#forgotpassword .unnumbered}

> **Fungsi:** Request reset password (mengirim OTP)
>
> **Input:**

- ctx context.Context

- phone string

> **Process:**

1.  Validasi phone tidak boleh kosong

2.  Normalisasi nomor telepon

3.  Cek UMKM berdasarkan phone number

4.  Generate OTP code

5.  Simpan OTP ke database dengan status 'active' dan expires_at 5 menit

6.  Kirim OTP ke WhatsApp menggunakan Fonnte API

> **Output:**

- error

#### **ResetPassword** {#resetpassword .unnumbered}

> **Fungsi:** Reset password dengan temp token
>
> **Input:**

- ctx context.Context

- user dto.ResetPasswordMobile - Data password baru (password,
  > confirm_password)

- tempToken string - Token dari verifikasi OTP

> **Process:**

1.  Validasi password dan confirm_password tidak boleh kosong

2.  Ambil OTP berdasarkan temp_token

3.  Validasi OTP masih aktif

4.  Ambil user berdasarkan email dari OTP

5.  Validasi password (min 8 karakter, huruf dan angka)

6.  Validasi password dan confirm_password sama

7.  Hash password baru

8.  Update password user

9.  Update status OTP menjadi 'used'

> **Output:**

- error

> **Utils yang Digunakan:**

- utils.PasswordValidator

- utils.PasswordHashing

#### **GetAllUsers** {#getallusers .unnumbered}

> **Fungsi:** Mendapatkan semua data user
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil semua user dari database

2.  Map ke DTO dengan format yang sesuai

> **Output:**

- \[\]dto.Users

- error

#### **GetUserByID** {#getuserbyid .unnumbered}

> **Fungsi:** Mendapatkan data user berdasarkan ID
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil user berdasarkan ID dari database

2.  Map ke DTO

> **Output:**

- dto.Users

- error

#### **UpdateUser** {#updateuser .unnumbered}

> **Fungsi:** Update data user (admin only)
>
> **Input:**

- ctx context.Context

- id int

- userNew dto.Users - Data baru (name, email, role_id)

> **Process:**

1.  Ambil user berdasarkan ID

2.  Validasi name, email, role_id tidak boleh kosong

3.  Validasi format email

4.  Cek email tidak digunakan user lain

5.  Validasi role_id valid

6.  Update data user

7.  Simpan ke database

> **Output:**

- dto.Users

- error

#### **DeleteUser** {#deleteuser .unnumbered}

> **Fungsi:** Soft delete user
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil user berdasarkan ID

2.  Soft delete user

> **Output:**

- dto.Users

- error

#### **MetaCityAndProvince** {#metacityandprovince .unnumbered}

> **Fungsi:** Mendapatkan data master provinsi dan kota
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil semua provinsi dari database

2.  Ambil semua kota dari database

3.  Gabungkan dalam satu response

> **Output:**

- \[\]dto.MetaCityAndProvince

- error

#### **GetListPermissions** {#getlistpermissions .unnumbered}

> **Fungsi:** Mendapatkan daftar semua permission
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil semua permission dari database

2.  Map ke DTO

> **Output:**

- \[\]dto.Permissions

- error

#### **GetListRolePermissions** {#getlistrolepermissions .unnumbered}

> **Fungsi:** Mendapatkan mapping role dan permissions
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil role permissions dari database dengan format JSON

> **Output:**

- \[\]dto.RolePermissionsResponse

- error

#### **UpdateRolePermissions** {#updaterolepermissions .unnumbered}

> **Fungsi:** Update permissions untuk suatu role
>
> **Input:**

- ctx context.Context

- rolePermissions dto.RolePermissions - role_id dan array permission
  > codes

> **Process:**

1.  Validasi role_id ada

2.  Validasi semua permission codes valid

3.  Hapus semua permission lama untuk role tersebut

4.  Tambahkan permission baru

> **Output:**

- Error

### Programs Service

> Service untuk mengelola program (training, certification, funding).
>
> **Dependencies:**

- ProgramsRepository

- UsersRepository

- RedisRepository

- MinIOManager

#### **GetAllPrograms** {#getallprograms .unnumbered}

> **Fungsi:** Mendapatkan semua program dengan detail benefits dan
> requirements
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil semua program dari database

2.  Untuk setiap program:

    - Ambil benefits

    - Ambil requirements

    - Map ke DTO dengan format lengkap

> **Output:**

- \[\]dto.Programs

- error

> **Struct/Type yang Digunakan:**

- dto.Programs

- model.Program

- model.ProgramBenefit

- model.ProgramRequirement

#### **GetProgramByID** {#getprogrambyid .unnumbered}

> **Fungsi:** Mendapatkan detail program berdasarkan ID
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil program berdasarkan ID

2.  Ambil benefits dan requirements

3.  Map ke DTO

> **Output:**

- dto.Programs

- error

#### **CreateProgram** {#createprogram .unnumbered}

> **Fungsi:** Membuat program baru
>
> **Input:**

- ctx context.Context

- program dto.Programs - Data program lengkap

> **Process:**

1.  Validasi input wajib (title, type, application_deadline)

2.  Validasi type harus salah satu dari: training, certification,
    > funding

3.  Validasi training_type jika type training/certification (online,
    > offline, hybrid)

4.  Validasi creator user ada

5.  Upload banner ke MinIO jika ada (base64 → MinIO)

6.  Upload provider logo ke MinIO jika ada

7.  Simpan program ke database

8.  Simpan benefits ke database

9.  Simpan requirements ke database

> **Output:**

- dto.Programs

- error

> **Utils yang Digunakan:**

- utils.GenerateFileName

- storage.MinIOManager.UploadFile

#### **UpdateProgram** {#updateprogram .unnumbered}

> **Fungsi:** Update program existing
>
> **Input:**

- ctx context.Context

- id int

- program dto.Programs - Data program baru

> **Process:**

1.  Ambil program existing berdasarkan ID

2.  Validasi input sama seperti CreateProgram

3.  Upload banner baru ke MinIO jika ada dan hapus yang lama

4.  Upload provider logo baru ke MinIO jika ada dan hapus yang lama

5.  Update data program

6.  Hapus benefits lama dan simpan yang baru

7.  Hapus requirements lama dan simpan yang baru

> **Output:**

- dto.Programs

- error

> **Utils yang Digunakan:**

- utils.GenerateFileName

- storage.MinIOManager.UploadFile

- storage.MinIOManager.DeleteFile

- storage.ExtractObjectNameFromURL

#### **DeleteProgram** {#deleteprogram .unnumbered}

> **Fungsi:** Soft delete program
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil program berdasarkan ID

2.  Soft delete program (cascade ke benefits dan requirements)

> **Output:**

- dto.Programs

- error

#### **ActivateProgram** {#activateprogram .unnumbered}

> **Fungsi:** Mengaktifkan program
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil program berdasarkan ID

2.  Set is_active = true

3.  Update ke database

> **Output:**

- dto.Programs

- error

#### **DeactivateProgram** {#deactivateprogram .unnumbered}

> **Fungsi:** Menonaktifkan program
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil program berdasarkan ID

2.  Set is_active = false

3.  Update ke database

> **Output:**

- dto.Programs

- Error

### Applications Service

> Service untuk mengelola aplikasi/pengajuan UMKM ke program.
>
> **Dependencies:**

- ApplicationsRepository

- UsersRepository

- NotificationRepository

- SLARepository

- VaultDecryptLogRepository

#### **GetAllApplications** {#getallapplications .unnumbered}

> **Fungsi:** Mendapatkan semua aplikasi dengan filter optional
>
> **Input:**

- ctx context.Context

- userID int - ID user yang request

- filterType string - Filter berdasarkan type program
  > (training/certification/funding/all)

> **Process:**

1.  Ambil semua aplikasi dari database dengan filter

2.  Untuk setiap aplikasi:

    - Ambil documents

    - Ambil histories dengan nama user

    - Map ke DTO lengkap dengan relasi program dan UMKM

> **Output:**

- \[\]dto.Applications

- error

> **Struct/Type yang Digunakan:**

- dto.Applications

- dto.ApplicationDocuments

- dto.ApplicationHistories

- model.Application

#### **GetApplicationByID** {#getapplicationbyid .unnumbered}

> **Fungsi:** Mendapatkan detail aplikasi dengan dekripsi data sensitif
>
> **Input:**

- ctx context.Context

- userID int

- id int - ID aplikasi

> **Process:**

1.  Ambil aplikasi berdasarkan ID dengan eager loading relasi

2.  Dekripsi NIK menggunakan Vault dengan logging

3.  Dekripsi Kartu Number menggunakan Vault dengan logging

4.  Map documents, histories ke DTO

5.  Tambahkan data spesifik berdasarkan type
    > (training/certification/funding)

> **Output:**

- dto.Applications

- error

> **Struct/Type yang Digunakan:**

- dto.Applications

- dto.TrainingApplicationData

- dto.CertificationApplicationData

- dto.FundingApplicationData

> **Utils yang Digunakan:**

- vault.DecryptNIKWithLog

- vault.DecryptKartuNumberWithLog

- vault.GetContextInfo

#### **ScreeningApprove** {#screeningapprove .unnumbered}

> **Fungsi:** Approve aplikasi pada tahap screening
>
> **Input:**

- ctx context.Context

- userID int - ID admin yang approve

- applicationID int

> **Process:**

1.  Ambil aplikasi berdasarkan ID

2.  Validasi status harus 'screening'

3.  Ambil SLA untuk tahap 'final'

4.  Update status menjadi 'final'

5.  Update expired_at berdasarkan SLA final

6.  Create history dengan action 'approve_by_admin_screening'

7.  Create notification untuk UMKM

> **Output:**

- dto.Applications

- error

> **Utils yang Digunakan:**

- constant.NotificationApproved

- constant.NotificationTitleApproved

- constant.NotificationMessageApproved

#### **ScreeningReject** {#screeningreject .unnumbered}

> **Fungsi:** Reject aplikasi pada tahap screening
>
> **Input:**

- ctx context.Context

- userID int

- decision dto.ApplicationDecision - application_id dan notes (wajib)

> **Process:**

1.  Validasi notes tidak boleh kosong

2.  Ambil aplikasi berdasarkan ID

3.  Validasi status harus 'screening'

4.  Update status menjadi 'rejected'

5.  Create history dengan action 'reject_by_admin_screening'

6.  Create notification dengan notes rejection

> **Output:**

- dto.Applications

- error

#### **ScreeningRevise** {#screeningrevise .unnumbered}

> **Fungsi:** Meminta revisi aplikasi pada tahap screening
>
> **Input:**

- ctx context.Context

- userID int

- decision dto.ApplicationDecision - application_id dan notes (wajib)

> **Process:**

1.  Validasi notes tidak boleh kosong

2.  Ambil aplikasi berdasarkan ID

3.  Validasi status harus 'screening'

4.  Update status menjadi 'revised'

5.  Create history dengan action 'revise'

6.  Create notification dengan notes revisi

> **Output:**

- dto.Applications

- error

#### **FinalApprove** {#finalapprove .unnumbered}

> **Fungsi:** Approve aplikasi pada tahap final
>
> **Input:**

- ctx context.Context

- userID int

- applicationID int

> **Process:**

1.  Ambil aplikasi berdasarkan ID

2.  Validasi status harus 'final'

3.  Update status menjadi 'approved'

4.  Create history dengan action 'approve_by_admin_vendor'

5.  Create notification

> **Output:**

- dto.Applications

- error

#### **FinalReject** {#finalreject .unnumbered}

> **Fungsi:** Reject aplikasi pada tahap final
>
> **Input:**

- ctx context.Context

- userID int

- decision dto.ApplicationDecision

> **Process:**

1.  Validasi notes tidak boleh kosong

2.  Ambil aplikasi berdasarkan ID

3.  Validasi status harus 'final'

4.  Update status menjadi 'rejected'

5.  Create history dengan action 'reject_by_admin_vendor'

6.  Create notification dengan notes rejection

> **Output:**

- dto.Applications

- error

### Mobile Service

> Service untuk operasi mobile app (UMKM user).
>
> **Dependencies:**

- MobileRepository

- ProgramsRepository

- NotificationRepository

- VaultDecryptLogRepository

- ApplicationsRepository

- SLARepository

- MinIOManager

#### **GetDashboard** {#getdashboard .unnumbered}

> **Fungsi:** Mendapatkan data dashboard untuk mobile
>
> **Input:**

- ctx context.Context

- userID int

> **Process:**

1.  Ambil profil UMKM berdasarkan user_id

2.  Hitung unread notifications

3.  Dekripsi Kartu Number untuk ditampilkan

4.  Ambil semua aplikasi user

5.  Hitung total aplikasi dan yang approved

6.  Return data dashboard lengkap

> **Output:**

- dto.DashboardData

- error

> **Struct/Type yang Digunakan:**

- dto.DashboardData

> **Utils yang Digunakan:**

- vault.DecryptTransit

#### **GetTrainingPrograms** {#gettrainingprograms .unnumbered}

> **Fungsi:** Mendapatkan daftar program training yang aktif
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil program dengan type 'training'

2.  Map ke DTO mobile

> **Output:**

- \[\]dto.ProgramListMobile

- error

#### **GetCertificationPrograms** {#getcertificationprograms .unnumbered}

> **Fungsi:** Mendapatkan daftar program certification yang aktif
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil program dengan type 'certification'

2.  Map ke DTO mobile

> **Output:**

- \[\]dto.ProgramListMobile

- error

#### **GetFundingPrograms** {#getfundingprograms .unnumbered}

> **Fungsi:** Mendapatkan daftar program funding yang aktif
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil program dengan type 'funding'

2.  Map ke DTO mobile

> **Output:**

- \[\]dto.ProgramListMobile

- error

#### **GetProgramDetail** {#getprogramdetail .unnumbered}

> **Fungsi:** Mendapatkan detail program dengan benefits dan
> requirements
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil program detail berdasarkan ID

2.  Ambil benefits dan requirements

3.  Map ke DTO mobile detail

> **Output:**

- dto.ProgramDetailMobile

- error

#### **GetUMKMProfile** {#getumkmprofile .unnumbered}

> **Fungsi:** Mendapatkan profil UMKM dengan dekripsi data sensitif
>
> **Input:**

- ctx context.Context

- userID int

> **Process:**

1.  Ambil profil UMKM berdasarkan user_id

2.  Dekripsi NIK

3.  Dekripsi Kartu Number

4.  Map ke DTO lengkap

> **Output:**

- dto.UMKMProfile

- error

> **Utils yang Digunakan:**

- vault.DecryptTransit

#### **UpdateUMKMProfile** {#updateumkmprofile .unnumbered}

> **Fungsi:** Update profil UMKM
>
> **Input:**

- ctx context.Context

- userID int

- request dto.UpdateUMKMProfile

> **Process:**

1.  Ambil profil existing

2.  Parse birth_date ke time.Time

3.  Upload photo baru ke MinIO jika ada dan hapus yang lama

4.  Update fields profil

5.  Simpan ke database

6.  Return profil terbaru dengan dekripsi

> **Output:**

- dto.UMKMProfile

- error

> **Utils yang Digunakan:**

- utils.GenerateFileName

- storage.MinIOManager.UploadFile

- storage.MinIOManager.DeleteFile

#### **GetUMKMDocuments** {#getumkmdocuments .unnumbered}

> **Fungsi:** Mendapatkan daftar dokumen UMKM
>
> **Input:**

- ctx context.Context

- userID int

> **Process:**

1.  Ambil profil UMKM

2.  Build array dokumen yang ada (NIB, NPWP, Revenue Record, Business
    > Permit)

3.  Return list dokumen

> **Output:**

- \[\]dto.UMKMDocument

- error

#### **UploadDocument** {#uploaddocument .unnumbered}

> **Fungsi:** Upload atau update dokumen UMKM
>
> **Input:**

- ctx context.Context

- userID int

- doc dto.UploadDocumentRequest - type
  > (nib/npwp/revenue_record/business_permit) dan document (base64/url)

> **Process:**

1.  Ambil profil UMKM

2.  Validasi document type

3.  Upload document ke MinIO jika base64

4.  Hapus dokumen lama jika ada

5.  Update field dokumen di database

> **Output:**

- error

> **Utils yang Digunakan:**

- utils.GenerateFileName

- storage.MinIOManager.UploadFile

- storage.MinIOManager.DeleteFile

#### **CreateTrainingApplication** {#createtrainingapplication .unnumbered}

> **Fungsi:** Membuat aplikasi program training
>
> **Input:**

- ctx context.Context

- userID int

- request dto.CreateApplicationTraining

> **Process:**

1.  Validasi program ada dan type 'training'

2.  Cek user belum pernah apply program ini

3.  Ambil UMKM dengan dekripsi (untuk validasi profil lengkap)

4.  Ambil SLA screening untuk set expired_at

5.  Create base application dengan status 'screening'

6.  Create training application data (motivation, business_experience,
    > dll)

7.  Create history 'submit'

8.  Create notification

9.  Process dan upload documents ke MinIO (async/goroutine)

> **Output:**

- error

> **Struct/Type yang Digunakan:**

- dto.CreateApplicationTraining

- model.Application

- model.TrainingApplication

> **Utils yang Digunakan:**

- vault.DecryptNIKWithLog

- vault.DecryptKartuNumberWithLog

#### **CreateCertificationApplication** {#createcertificationapplication .unnumbered}

> **Fungsi:** Membuat aplikasi program certification
>
> **Input:**

- ctx context.Context

- userID int

- request dto.CreateApplicationCertification

> **Process:**

1.  Validasi program ada dan type 'certification'

2.  Cek user belum pernah apply program ini

3.  Ambil UMKM dengan dekripsi

4.  Ambil SLA screening

5.  Create base application

6.  Create certification application data (business_sector,
    > product_or_service, dll)

7.  Create history dan notification

8.  Process documents (async)

> **Output:**

- error

> **Struct/Type yang Digunakan:**

- dto.CreateApplicationCertification

- model.CertificationApplication

#### **CreateFundingApplication** {#createfundingapplication .unnumbered}

> **Fungsi:** Membuat aplikasi program funding
>
> **Input:**

- ctx context.Context

- userID int

- request dto.CreateApplicationFunding

> **Process:**

1.  Validasi program ada dan type 'funding'

2.  Validasi requested_amount dalam range min_amount dan max_amount
    > program

3.  Validasi requested_tenure_months tidak melebihi max_tenure_months
    > program

4.  Cek user belum pernah apply program ini

5.  Ambil UMKM dengan dekripsi

6.  Ambil SLA screening

7.  Create base application

8.  Create funding application data (requested_amount, fund_purpose,
    > collateral, dll)

9.  Create history dan notification

10. Process documents (async)

> **Output:**

- error

> **Struct/Type yang Digunakan:**

- dto.CreateApplicationFunding

- model.FundingApplication

#### **GetApplicationList** {#getapplicationlist .unnumbered}

> **Fungsi:** Mendapatkan daftar aplikasi user
>
> **Input:**

- ctx context.Context

- userID int

> **Process:**

1.  Ambil profil UMKM

2.  Ambil semua aplikasi UMKM tersebut

3.  Map ke DTO list mobile

> **Output:**

- \[\]dto.ApplicationListMobile

- error

#### **GetApplicationDetail** {#getapplicationdetail .unnumbered}

> **Fungsi:** Mendapatkan detail aplikasi dengan data spesifik per type
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil aplikasi detail dengan eager loading

2.  Ambil benefits dan requirements program

3.  Map documents dan histories

4.  Tambahkan data spesifik
    > (training_data/certification_data/funding_data) berdasarkan type

> **Output:**

- dto.ApplicationDetailMobile

- error

#### **ReviseApplication** {#reviseapplication .unnumbered}

> **Fungsi:** Revisi aplikasi yang diminta perubahan
>
> **Input:**

- ctx context.Context

- userID int

- applicationID int

- documents \[\]dto.UploadDocumentRequest

> **Process:**

1.  Ambil aplikasi detail

2.  Validasi status harus 'revised'

3.  Upload dokumen baru ke MinIO (async)

4.  Update status menjadi 'screening' dan submitted_at ke sekarang

5.  Hapus dokumen lama

6.  Create history 'submit' untuk resubmit

7.  Create notification

> **Output:**

- error

#### **GetNotificationsByUMKMID** {#getnotificationsbyumkmid .unnumbered}

> **Fungsi:** Mendapatkan daftar notifikasi UMKM
>
> **Input:**

- ctx context.Context

- umkmID int

> **Process:**

1.  Ambil notifikasi dengan limit 100

2.  Map ke DTO response

> **Output:**

- \[\]dto.NotificationResponse

- error

#### **GetUnreadCount** {#getunreadcount .unnumbered}

> **Fungsi:** Mendapatkan jumlah notifikasi belum dibaca
>
> **Input:**

- ctx context.Context

- umkmID int

> **Process:**

1.  Hitung notifikasi dengan is_read = false

> **Output:**

- int64

- error

#### **MarkNotificationsAsRead** {#marknotificationsasread .unnumbered}

> **Fungsi:** Tandai notifikasi sebagai sudah dibaca
>
> **Input:**

- ctx context.Context

- umkmID int

- notificationIDs int

> **Process:**

1.  Update is_read = true dan read_at = NOW untuk notifikasi tersebut

> **Output:**

- error

#### **MarkAllNotificationsAsRead** {#markallnotificationsasread .unnumbered}

> **Fungsi:** Tandai semua notifikasi sebagai sudah dibaca
>
> **Input:**

- ctx context.Context

- umkmID int

> **Process:**

1.  Update is_read = true dan read_at = NOW untuk semua notifikasi UMKM

> **Output:**

- error

#### **GetPublishedNews** {#getpublishednews .unnumbered}

> **Fungsi:** Mendapatkan berita yang sudah dipublish dengan filter
>
> **Input:**

- ctx context.Context

- params dto.NewsQueryParams - page, limit, category, search, tag

> **Process:**

1.  Ambil news yang is_published = true dengan filter

2.  Map ke DTO mobile

> **Output:**

- \[\]dto.NewsListMobile

- int64 - total records

- error

#### **GetNewsDetail** {#getnewsdetail .unnumbered}

> **Fungsi:** Mendapatkan detail berita berdasarkan slug
>
> **Input:**

- ctx context.Context

- slug string

> **Process:**

1.  Ambil news berdasarkan slug

2.  Increment views_count

3.  Map tags

4.  Return detail news

> **Output:**

- dto.NewsDetailMobile

- error

### Dashboard Service

> Service untuk statistik dan dashboard web.
>
> **Dependencies:**

- DashboardRepository

#### **GetUMKMByCardType** {#getumkmbycardtype .unnumbered}

> **Fungsi:** Statistik berdasarkan tipe kartu
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Query count UMKM group by kartu_type

2.  Map ke DTO

> **Output:**

- \[\]dto.UMKMByCardType

- Error

#### **GetApplicationStatusSummary** {#getapplicationstatussummary .unnumbered}

> **Fungsi:** Ringkasan status aplikasi
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Hitung total applications

2.  Hitung in_process (screening + revised + final)

3.  Hitung approved

4.  Hitung rejected

> **Output:**

- \[\]dto.ApplicationStatusSummary

- Error

#### **GetApplicationStatusDetail** {#getapplicationstatusdetail .unnumbered}

> **Fungsi:** Detail status aplikasi per status
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Hitung aplikasi untuk setiap status (screening, revised, final,
    > approved, rejected)

> **Output:**

- \[\]dto.ApplicationStatusDetail

- Error

#### **GetApplicationByType** {#getapplicationbytype .unnumbered}

> **Fungsi:** Statistik aplikasi berdasarkan type program
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Hitung aplikasi group by type (training, certification, funding)

> **Output:**

- \[\]dto.ApplicationByType

- Error

### SLA Service

> Service untuk konfigurasi SLA dan export report.
>
> **Dependencies:**

- SLARepository

#### **GetSLAScreening** {#getslascreening .unnumbered}

> **Fungsi:** Mendapatkan konfigurasi SLA untuk screening
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil SLA dengan status 'screening'

2.  Map ke DTO

> **Output:**

- dto.SLA

- Error

#### **GetSLAFinal** {#getslafinal .unnumbered}

> **Fungsi:** Mendapatkan konfigurasi SLA untuk final decision
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil SLA dengan status 'final'

2.  Map ke DTO

> **Output:**

- dto.SLA

- Error

#### **UpdateSLAScreening** {#updateslascreening .unnumbered}

> **Fungsi:** Update konfigurasi SLA screening
>
> **Input:**

- ctx context.Context

- slaDTO dto.SLA - max_days dan description

> **Process:**

1.  Validasi max_days \> 0

2.  Ambil SLA existing

3.  Update max_days dan description

4.  Simpan ke database

> **Output:**

- dto.SLA

- Error

#### **UpdateSLAFinal** {#updateslafinal .unnumbered}

> **Fungsi:** Update konfigurasi SLA final
>
> **Input:**

- ctx context.Context

- slaDTO dto.SLA

> **Process:**

1.  Validasi max_days \> 0

2.  Ambil SLA existing

3.  Update max_days dan description

4.  Simpan ke database

> **Output:**

- dto.SLA

- Error

#### **ExportApplications** {#exportapplications .unnumbered}

> **Fungsi:** Export data aplikasi ke PDF/Excel
>
> **Input:**

- ctx context.Context

- request dto.ExportRequest - file_type (pdf/excel) dan application_type
  > (all/training/certification/funding)

> **Process:**

1.  Ambil aplikasi dari database dengan filter type

2.  Generate file berdasarkan file_type:

    - PDF: text-based format dengan informasi aplikasi

    - Excel: CSV format dengan kolom terstruktur

<!-- -->

3.  Return file bytes dan filename

> **Output:**

- \[\]byte - File content

- string - Filename

- Error

#### **ExportPrograms** {#exportprograms .unnumbered}

> **Fungsi:** Export data program ke PDF/Excel
>
> **Input:**

- ctx context.Context

- request dto.ExportRequest

> **Process:**

1.  Ambil program dari database dengan filter type

2.  Generate file berdasarkan file_type

3.  Return file bytes dan filename

> **Output:**

- \[\]byte

- string

- Error

### News Service

> Service untuk mengelola berita/artikel.
>
> **Dependencies:**

- NewsRepository

- MinIOManager

#### **GetAllNews** {#getallnews .unnumbered}

> **Fungsi:** Mendapatkan semua berita dengan filter
>
> **Input:**

- ctx context.Context

- params dto.NewsQueryParams - page, limit, category, search, tag,
  > is_published

> **Process:**

1.  Query news dengan filter dan pagination

2.  Map ke DTO list

> **Output:**

- \[\]dto.NewsListResponse

- int64 - total records

- Error

#### **GetNewsByID** {#getnewsbyid .unnumbered}

> **Fungsi:** Mendapatkan detail berita berdasarkan ID
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil news berdasarkan ID dengan eager loading tags

2.  Map tags ke array string

3.  Map ke DTO response

> **Output:**

- dto.NewsResponse

- Error

#### **CreateNews** {#createnews .unnumbered}

> **Fungsi:** Membuat berita baru
>
> **Input:**

- ctx context.Context

- authorID int

- request dto.NewsRequest

> **Process:**

1.  Validasi title dan content tidak boleh kosong

2.  Generate slug dari title menggunakan helper generateSlug

3.  Cek slug unique, jika tidak unique tambahkan timestamp

4.  Upload thumbnail ke MinIO jika ada (base64 → MinIO)

5.  Simpan news ke database

6.  Set published_at jika is_published = true

7.  Create tags jika ada

> **Output:**

- dto.NewsResponse

- error

> **Utils yang Digunakan:**

- utils.GenerateFileName

- storage.MinIOManager.UploadFile

#### **UpdateNews** {#updatenews .unnumbered}

> **Fungsi:** Update berita existing
>
> **Input:**

- ctx context.Context

- id int

- request dto.NewsRequest

> **Process:**

1.  Ambil news existing

2.  Validasi title dan content

3.  Generate slug baru jika title berubah

4.  Upload thumbnail baru ke MinIO jika ada dan hapus yang lama

5.  Update fields news

6.  Handle publishing status (set/unset published_at)

7.  Hapus tags lama dan create yang baru

> **Output:**

- dto.NewsResponse

- Error

#### **DeleteNews** {#deletenews .unnumbered}

> **Fungsi:** Soft delete berita
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil news berdasarkan ID

2.  Hapus thumbnail dari MinIO

3.  Soft delete news (cascade ke tags)

> **Output:**

- Error

#### **PublishNews** {#publishnews .unnumbered}

> **Fungsi:** Publish berita
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil news berdasarkan ID

2.  Validasi belum published

3.  Set is_published = true dan published_at = NOW

4.  Update ke database

> **Output:**

- dto.NewsResponse

- Error

#### **UnpublishNews** {#unpublishnews .unnumbered}

> **Fungsi:** Unpublish berita
>
> **Input:**

- ctx context.Context

- id int

> **Process:**

1.  Ambil news berdasarkan ID

2.  Validasi sudah published

3.  Set is_published = false dan published_at = NULL

4.  Update ke database

> **Output:**

- dto.NewsResponse

- Error

### Vault Decrypt Log Service

> Service untuk audit log dekripsi data sensitif.
>
> **Dependencies:**

- VaultDecryptLogRepository

#### **GetLogs** {#getlogs .unnumbered}

> **Fungsi:** Mendapatkan semua log dekripsi dengan pagination
>
> **Input:**

- ctx context.Context

> **Process:**

1.  Ambil log dengan default limit 100 dan offset 0

> **Output:**

- \[\]model.VaultDecryptLog

- Error

#### **GetLogsByUserID** {#getlogsbyuserid .unnumbered}

> **Fungsi:** Mendapatkan log dekripsi berdasarkan user ID
>
> **Input:**

- ctx context.Context

- userID int

> **Process:**

1.  Ambil log untuk user tertentu dengan pagination

> **Output:**

- \[\]model.VaultDecryptLog

- Error

#### **GetLogsByUMKMID** {#getlogsbyumkmid .unnumbered}

> **Fungsi:** Mendapatkan log dekripsi berdasarkan UMKM ID
>
> **Input:**

- ctx context.Context

- umkmID int

> **Process:**

1.  Ambil log untuk UMKM tertentu dengan pagination

> **Output:**

- \[\]model.VaultDecryptLog

- error

## Repository Layer

### Applications Repository

> Repository yang menangani operasi database untuk entitas Applications
> (Pengajuan Program).

#### **GetAllApplications** {#getallapplications-1 .unnumbered}

> GetAllApplications(ctx context.Context, filterType string)
> (\[\]model.Application, error)
>
> **Deskripsi:** Mengambil semua data aplikasi/pengajuan dengan filter
> opsional berdasarkan tipe.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- filterType: String filter tipe aplikasi
  > (training/certification/funding), kosong untuk semua

> **Process:**

- Query database dengan preload relasi: Program, UMKM.User,
  > UMKM.City.Province, Documents, Histories.User

- Filter aplikasi yang tidak dihapus (deleted_at IS NULL)

- Jika filterType tidak kosong, tambahkan filter by type

> **Output:**

- Slice of model.Application dengan semua relasi

- Error jika query gagal

> **Dependencies:**

- gorm.DB

- model.Application

#### **GetApplicationByID** {#getapplicationbyid-1 .unnumbered}

> GetApplicationByID(ctx context.Context, id int) (model.Application,
> error)
>
> **Deskripsi:** Mengambil detail aplikasi berdasarkan ID dengan semua
> relasi dan data spesifik tipe.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID aplikasi

> **Process:**

- Query database dengan multiple preload

- Include TrainingApplication, CertificationApplication,
  > FundingApplication

- Filter by ID dan deleted_at IS NULL

> **Output:**

- model.Application lengkap dengan semua relasi

- Error "application not found" jika tidak ditemukan

> **Dependencies:**

- gorm.DB

- model.Application

#### **GetApplicationsByUMKMID** {#getapplicationsbyumkmid .unnumbered}

> GetApplicationsByUMKMID(ctx context.Context, umkmID int)
> (\[\]model.Application, error)
>
> **Deskripsi:** Mengambil semua aplikasi milik UMKM tertentu.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

> **Process:**

- Query dengan preload relasi lengkap

- Filter by umkm_id dan deleted_at IS NULL

> **Output:**

- Slice of model.Application

- Error jika query gagal

#### **CreateApplication** {#createapplication .unnumbered}

> CreateApplication(ctx context.Context, application model.Application)
> (model.Application, error)
>
> **Deskripsi:** Membuat aplikasi baru dengan mengisi timestamp
> otomatis.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- application: Struct model.Application yang akan dibuat

> **Process:**

- Set SubmittedAt ke waktu sekarang

- Set ExpiredAt ke 30 hari dari sekarang

- Insert ke database

> **Output:**

- model.Application yang telah dibuat

- Error "failed to create application" jika gagal

> **Dependencies:**

- time package untuk timestamp

#### **UpdateApplication** {#updateapplication .unnumbered}

> UpdateApplication(ctx context.Context, application model.Application)
> (model.Application, error)
>
> **Deskripsi:** Update data aplikasi yang sudah ada.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- application: Struct model.Application dengan data baru

> **Process:**

- Simpan perubahan ke database menggunakan GORM Save

> **Output:**

- model.Application yang telah diupdate

- Error "failed to update application" jika gagal

#### **DeleteApplication** {#deleteapplication .unnumbered}

> DeleteApplication(ctx context.Context, application model.Application)
> (model.Application, error)
>
> **Deskripsi:** Soft delete aplikasi (set deleted_at).
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- application: Struct model.Application yang akan dihapus

> **Process:**

- Soft delete menggunakan GORM Delete

> **Output:**

- model.Application yang telah dihapus

- Error "failed to delete application" jika gagal

#### **CreateApplicationDocuments** {#createapplicationdocuments .unnumbered}

> CreateApplicationDocuments(ctx context.Context, documents
> \[\]model.ApplicationDocument) error
>
> **Deskripsi:** Membuat multiple dokumen aplikasi sekaligus.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- documents: Slice of model.ApplicationDocument

> **Process:**

- Skip jika slice kosong

- Batch insert semua dokumen

> **Output:**

- Error "failed to create application documents" jika gagal

- nil jika sukses

#### **GetApplicationDocuments** {#getapplicationdocuments .unnumbered}

> GetApplicationDocuments(ctx context.Context, applicationID int)
> (\[\]model.ApplicationDocument, error)
>
> **Deskripsi:** Mengambil semua dokumen yang terkait dengan aplikasi.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- applicationID: Integer ID aplikasi

> **Process:**

- Query dokumen by application_id

- Filter deleted_at IS NULL

> **Output:**

- Slice of model.ApplicationDocument

- Error jika query gagal

#### **DeleteApplicationDocuments** {#deleteapplicationdocuments .unnumbered}

> DeleteApplicationDocuments(ctx context.Context, applicationID int)
> error
>
> **Deskripsi:** Hard delete semua dokumen aplikasi (unscoped).
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- applicationID: Integer ID aplikasi

> **Process:**

- Unscoped delete semua dokumen dengan application_id tertentu

> **Output:**

- Error "failed to delete application documents" jika gagal

- nil jika sukses

#### **CreateApplicationHistory** {#createapplicationhistory .unnumbered}

> CreateApplicationHistory(ctx context.Context, history
> model.ApplicationHistory) error
>
> **Deskripsi:** Mencatat riwayat aksi pada aplikasi.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- history: Struct model.ApplicationHistory

> **Process:**

- Set ActionedAt ke waktu sekarang

- Insert ke database

> **Output:**

- Error "failed to create application history" jika gagal

- nil jika sukses

#### **GetApplicationHistories** {#getapplicationhistories .unnumbered}

> GetApplicationHistories(ctx context.Context, applicationID int)
> (\[\]model.ApplicationHistory, error)
>
> **Deskripsi:** Mengambil riwayat aksi aplikasi terurut dari terbaru.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- applicationID: Integer ID aplikasi

> **Process:**

- Preload relasi User

- Filter by application_id dan deleted_at IS NULL

- Order by actioned_at DESC

> **Output:**

- Slice of model.ApplicationHistory

- Error jika query gagal

#### **GetProgramByID** {#getprogrambyid-1 .unnumbered}

> GetProgramByID(ctx context.Context, id int) (model.Program, error)
>
> **Deskripsi:** Validasi: mengambil data program untuk verifikasi.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID program

> **Process:**

- Query program by ID

- Filter deleted_at IS NULL

> **Output:**

- model.Program

- Error "program not found" jika tidak ada

#### **GetUMKMByUserID** {#getumkmbyuserid .unnumbered}

> GetUMKMByUserID(ctx context.Context, userID int) (model.UMKM, error)
>
> **Deskripsi:** Validasi: mengambil data UMKM berdasarkan user ID.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- userID: Integer ID user

> **Process:**

- Query UMKM by user_id

- Filter deleted_at IS NULL

> **Output:**

- model.UMKM

- Error "UMKM not found" jika tidak ada

#### **IsApplicationExists** {#isapplicationexists .unnumbered}

> IsApplicationExists(ctx context.Context, umkmID, programID int) bool
>
> **Deskripsi:** Cek apakah UMKM sudah punya aplikasi aktif untuk
> program tertentu.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

- programID: Integer ID program

> **Process:**

- Count aplikasi dengan umkm_id dan program_id tertentu

- Filter deleted_at IS NULL dan status NOT IN ('rejected')

> **Output:**

- Boolean true jika ada, false jika tidak

### Dashboard Repository

> Repository untuk agregasi data statistik dashboard.

#### **GetUMKMByCardType** {#getumkmbycardtype-1 .unnumbered}

> GetUMKMByCardType(ctx context.Context)
> (\[\]**map**\[string\]**interface**{}, error)
>
> **Deskripsi:** Mengambil statistik UMKM berdasarkan tipe kartu.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Raw SQL query dengan CASE untuk label

- GROUP BY kartu_type

- Filter deleted_at IS NULL dan kartu_type IS NOT NULL

> **Output:**

- Slice of map dengan key: name (string) dan count (int64)

- Error jika query gagal

> **Query:**
>
> **SELECT  
> ** **CASE  
> ** **WHEN** kartu_type = \'produktif\' **THEN** \'Kartu Produktif\'  
> **WHEN** kartu_type = \'afirmatif\' **THEN** \'Kartu Afirmatif\'  
> **ELSE** \'Unknown\'  
> **END** **as** name,  
> COUNT(\*) **as** count  
> **FROM** umkms  
> **WHERE** deleted_at **IS** **NULL** **AND** kartu_type **IS** **NOT**
> **NULL  
> ** **GROUP** **BY** kartu_type

#### **GetApplicationStatusSummary** {#getapplicationstatussummary-1 .unnumbered}

> GetApplicationStatusSummary(ctx context.Context)
> (**map**\[string\]int64, error)
>
> **Deskripsi:** Ringkasan status aplikasi untuk dashboard overview.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Multiple raw SQL COUNT query

- Hitung total_applications

- Hitung in_process (screening + revised + final)

- Hitung approved

- Hitung rejected

> **Output:**

- Map dengan key: total_applications, in_process, approved, rejected

- Error jika query gagal

#### **GetApplicationStatusDetail** {#getapplicationstatusdetail-1 .unnumbered}

> GetApplicationStatusDetail(ctx context.Context)
> (**map**\[string\]int64, error)
>
> **Deskripsi:** Detail status aplikasi per status.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Raw SQL query GROUP BY status

- Filter deleted_at IS NULL

> **Output:**

- Map dengan key status (screening, revised, final, approved, rejected)

- Error jika query gagal

#### **GetApplicationByType** {#getapplicationbytype-1 .unnumbered}

> GetApplicationByType(ctx context.Context) (**map**\[string\]int64,
> error)
>
> **Deskripsi:** Statistik aplikasi berdasarkan tipe program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Raw SQL query GROUP BY type

- Filter deleted_at IS NULL

> **Output:**

- Map dengan key type (training, certification, funding)

- Error jika query gagal

### Mobile Repository

> Repository untuk operasi mobile app (pelaku usaha).

#### **GetProgramsByType** {#getprogramsbytype .unnumbered}

> GetProgramsByType(ctx context.Context, programType string)
> (\[\]model.Program, error)
>
> **Deskripsi:** Mengambil daftar program berdasarkan tipe untuk mobile.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- programType: String tipe (training/certification/funding)

> **Process:**

- Query dengan filter type, is_active, deleted_at

- Order by created_at DESC

> **Output:**

- Slice of model.Program

- Error jika query gagal

#### **GetProgramDetailByID** {#getprogramdetailbyid .unnumbered}

> GetProgramDetailByID(ctx context.Context, id int) (model.Program,
> error)
>
> **Deskripsi:** Mengambil detail program untuk mobile dengan relasi
> User.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID program

> **Process:**

- Preload Users

- Filter is_active dan deleted_at

> **Output:**

- model.Program dengan relasi

- Error "program not found" jika tidak ada

#### **GetUMKMProfileByID** {#getumkmprofilebyid .unnumbered}

> GetUMKMProfileByID(ctx context.Context, userID int) (model.UMKM,
> error)
>
> **Deskripsi:** Mengambil profil UMKM lengkap dengan relasi.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- userID: Integer ID UMKM (bukan user_id)

> **Process:**

- Preload User, Province, City

- Filter by ID dan deleted_at

> **Output:**

- model.UMKM dengan relasi lengkap

- Error "UMKM profile not found" jika tidak ada

#### **UpdateUMKMProfile** {#updateumkmprofile-1 .unnumbered}

> UpdateUMKMProfile(ctx context.Context, umkm model.UMKM) (model.UMKM,
> error)
>
> **Deskripsi:** Update profil UMKM dalam transaction.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkm: Struct model.UMKM dengan data baru

> **Process:**

- Gunakan transaction

- Save UMKM data

- Jika user.name ada, update user.name terpisah

> **Output:**

- model.UMKM yang telah diupdate

- Error "failed to update UMKM profile" jika gagal

> **Dependencies:**

- GORM Transaction

#### **UpdateUMKMDocument** {#updateumkmdocument .unnumbered}

> UpdateUMKMDocument(ctx context.Context, umkmID int, field, value
> string) error
>
> **Deskripsi:** Update field dokumen UMKM secara spesifik.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

- field: String nama field (nib, npwp, revenue_record, business_permit)

- value: String URL dokumen baru

> **Process:**

- Update satu field saja menggunakan Model.Update

> **Output:**

- Error "failed to update document" jika gagal

- nil jika sukses

#### **CreateApplication (Mobile)** {#createapplication-mobile .unnumbered}

> CreateApplication(ctx context.Context, application model.Application)
> (model.Application, error)
>
> **Deskripsi:** Membuat aplikasi baru dari mobile app.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- application: Struct model.Application

> **Process:**

- Insert ke database

> **Output:**

- model.Application yang telah dibuat

- Error "failed to create application" jika gagal

#### **CreateApplicationDocuments (Mobile)** {#createapplicationdocuments-mobile .unnumbered}

> CreateApplicationDocuments(ctx context.Context, documents
> \[\]model.ApplicationDocument) error
>
> **Deskripsi:** Batch insert dokumen aplikasi mobile.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- documents: Slice of model.ApplicationDocument

> **Process:**

- Skip jika kosong

- Batch insert

> **Output:**

- Error "failed to create application documents" jika gagal

- nil jika sukses

#### **CreateApplicationHistory (Mobile)** {#createapplicationhistory-mobile .unnumbered}

> CreateApplicationHistory(ctx context.Context, history
> model.ApplicationHistory) error
>
> **Deskripsi:** Mencatat history submit aplikasi dari mobile.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- history: Struct model.ApplicationHistory

> **Process:**

- Insert history

> **Output:**

- Error "failed to create application history" jika gagal

- nil jika sukses

#### **CreateTrainingApplication** {#createtrainingapplication-1 .unnumbered}

> CreateTrainingApplication(ctx context.Context, training
> model.TrainingApplication) error
>
> **Deskripsi:** Menyimpan data spesifik training application.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- training: Struct model.TrainingApplication

> **Process:**

- Insert data training spesifik

> **Output:**

- Error "failed to create training application" jika gagal

- nil jika sukses

#### **CreateCertificationApplication** {#createcertificationapplication-1 .unnumbered}

> CreateCertificationApplication(ctx context.Context, certification
> model.CertificationApplication) error
>
> **Deskripsi:** Menyimpan data spesifik certification application.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- certification: Struct model.CertificationApplication

> **Process:**

- Insert data certification spesifik

> **Output:**

- Error "failed to create certification application" jika gagal

- nil jika sukses

#### **CreateFundingApplication** {#createfundingapplication-1 .unnumbered}

> CreateFundingApplication(ctx context.Context, funding
> model.FundingApplication) error
>
> **Deskripsi:** Menyimpan data spesifik funding application.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- funding: Struct model.FundingApplication

> **Process:**

- Insert data funding spesifik

> **Output:**

- Error "failed to create funding application" jika gagal

- nil jika sukses

#### **GetApplicationsByUMKMID (Mobile)** {#getapplicationsbyumkmid-mobile .unnumbered}

> GetApplicationsByUMKMID(ctx context.Context, umkmID int)
> (\[\]model.Application, error)
>
> **Deskripsi:** Mengambil list aplikasi milik UMKM untuk mobile.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

> **Process:**

- Preload Program dan data spesifik tipe

- Order by submitted_at DESC

> **Output:**

- Slice of model.Application

- Error jika query gagal

#### **GetApplicationDetailByID (Mobile)** {#getapplicationdetailbyid-mobile .unnumbered}

> GetApplicationDetailByID(ctx context.Context, id int)
> (model.Application, error)
>
> **Deskripsi:** Mengambil detail aplikasi untuk mobile app.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID aplikasi

> **Process:**

- Preload Program, Documents, Histories, dan data spesifik tipe

- Filter deleted_at

> **Output:**

- model.Application lengkap

- Error "application not found" jika tidak ada

#### **DeleteApplicationDocumentsByApplicationID** {#deleteapplicationdocumentsbyapplicationid .unnumbered}

> DeleteApplicationDocumentsByApplicationID(ctx context.Context,
> applicationID int) error
>
> **Deskripsi:** Menghapus dokumen aplikasi untuk proses revisi.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- applicationID: Integer ID aplikasi

> **Process:**

- Delete dokumen by application_id

> **Output:**

- Error "failed to delete application documents" jika gagal

- nil jika sukses

#### **GetProgramRequirements** {#getprogramrequirements .unnumbered}

> GetProgramRequirements(ctx context.Context, programID int)
> (\[\]model.ProgramRequirement, error)
>
> **Deskripsi:** Mengambil syarat-syarat program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- programID: Integer ID program

> **Process:**

- Query requirements by program_id

- Filter deleted_at

> **Output:**

- Slice of model.ProgramRequirement

- Error jika query gagal

#### **GetPublishedNews** {#getpublishednews-1 .unnumbered}

> GetPublishedNews(ctx context.Context, params dto.NewsQueryParams)
> (\[\]model.News, int64, error)
>
> **Deskripsi:** Mengambil berita yang sudah publish untuk mobile.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- params: Struct dto.NewsQueryParams (page, limit, category, search,
  > tag)

> **Process:**

- Filter is_published = true dan deleted_at IS NULL

- Optional filter by category

- Optional search by title (ILIKE)

- Optional filter by tag (join news_tags)

- Count total

- Pagination dengan limit dan offset

- Order by published_at DESC

> **Output:**

- Slice of model.News dengan relasi Author dan Tags

- Total count (int64)

- Error jika query gagal

> **Dependencies:**

- dto.NewsQueryParams

#### **GetPublishedNewsBySlug** {#getpublishednewsbyslug .unnumbered}

> GetPublishedNewsBySlug(ctx context.Context, slug string) (model.News,
> error)
>
> **Deskripsi:** Mengambil detail berita berdasarkan slug.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- slug: String slug berita

> **Process:**

- Query by slug

- Filter is_published dan deleted_at

- Preload Author dan Tags

> **Output:**

- model.News dengan relasi

- Error "news not found" jika tidak ada

#### **IncrementViews** {#incrementviews .unnumbered}

> IncrementViews(ctx context.Context, newsID int) error
>
> **Deskripsi:** Increment counter views berita.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- newsID: Integer ID berita

> **Process:**

- UpdateColumn views_count + 1 menggunakan Expr

> **Output:**

- Error jika update gagal

- nil jika sukses

### News Repository

> Repository untuk manajemen berita (web admin).

#### **GetAllNews** {#getallnews-1 .unnumbered}

> GetAllNews(ctx context.Context, params dto.NewsQueryParams)
> (\[\]model.News, int64, error)
>
> **Deskripsi:** Mengambil semua berita dengan filter untuk admin.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- params: Struct dto.NewsQueryParams

> **Process:**

- Preload Author dan Tags

- Optional filter by category

- Optional filter by is_published

- Optional search by title atau content (ILIKE)

- Optional filter by tag (join)

- Count total

- Pagination

- Order by created_at DESC

> **Output:**

- Slice of model.News

- Total count

- Error jika gagal

#### **GetNewsByID** {#getnewsbyid-1 .unnumbered}

> GetNewsByID(ctx context.Context, id int) (model.News, error)
>
> **Deskripsi:** Mengambil berita berdasarkan ID untuk admin.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID berita

> **Process:**

- Preload Author dan Tags

- Filter deleted_at

> **Output:**

- model.News

- Error "news not found" jika tidak ada

#### **GetNewsBySlug** {#getnewsbyslug .unnumbered}

> GetNewsBySlug(ctx context.Context, slug string) (model.News, error)
>
> **Deskripsi:** Mengambil berita berdasarkan slug.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- slug: String slug

> **Process:**

- Query by slug

- Preload relasi

- Filter deleted_at

> **Output:**

- model.News

- Error "news not found" jika tidak ada

#### **CreateNews** {#createnews-1 .unnumbered}

> CreateNews(ctx context.Context, news model.News) (model.News, error)
>
> **Deskripsi:** Membuat berita baru.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- news: Struct model.News

> **Process:**

- Insert ke database

> **Output:**

- model.News yang telah dibuat

- Error "failed to create news" jika gagal

#### **UpdateNews** {#updatenews-1 .unnumbered}

> UpdateNews(ctx context.Context, news model.News) (model.News, error)
>
> **Deskripsi:** Update data berita.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- news: Struct model.News dengan data baru

> **Process:**

- Save perubahan

> **Output:**

- model.News yang telah diupdate

- Error "failed to update news" jika gagal

#### **DeleteNews** {#deletenews-1 .unnumbered}

> DeleteNews(ctx context.Context, news model.News) error
>
> **Deskripsi:** Soft delete berita.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- news: Struct model.News

> **Process:**

- Soft delete

> **Output:**

- Error "failed to delete news" jika gagal

- nil jika sukses

#### **IsSlugExists** {#isslugexists .unnumbered}

> IsSlugExists(ctx context.Context, slug string, excludeID int) bool
>
> **Deskripsi:** Cek apakah slug sudah dipakai.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- slug: String slug yang akan dicek

- excludeID: Integer ID untuk dikecualikan (untuk update)

> **Process:**

- Count berita dengan slug tertentu

- Exclude ID jika \> 0

- Filter deleted_at

> **Output:**

- Boolean true jika ada, false jika tidak

#### **CreateNewsTags** {#createnewstags .unnumbered}

> CreateNewsTags(ctx context.Context, tags \[\]model.NewsTag) error
>
> **Deskripsi:** Batch insert tags berita.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- tags: Slice of model.NewsTag

> **Process:**

- Skip jika kosong

- Batch insert

> **Output:**

- Error "failed to create news tags" jika gagal

- nil jika sukses

#### **DeleteNewsTags** {#deletenewstags .unnumbered}

> DeleteNewsTags(ctx context.Context, newsID int) error
>
> **Deskripsi:** Hard delete semua tags berita.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- newsID: Integer ID berita

> **Process:**

- Unscoped delete by news_id

> **Output:**

- Error "failed to delete news tags" jika gagal

- nil jika sukses

#### **GetNewsTags** {#getnewstags .unnumbered}

> GetNewsTags(ctx context.Context, newsID int) (\[\]model.NewsTag,
> error)
>
> **Deskripsi:** Mengambil tags berita.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- newsID: Integer ID berita

> **Process:**

- Query tags by news_id

> **Output:**

- Slice of model.NewsTag

- Error jika gagal

### Notification Repository

> Repository untuk notifikasi mobile.

#### **CreateNotification** {#createnotification .unnumbered}

> CreateNotification(ctx context.Context, notification
> model.Notification) error
>
> **Deskripsi:** Membuat notifikasi baru.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- notification: Struct model.Notification

> **Process:**

- Insert ke database

> **Output:**

- Error jika gagal

- nil jika sukses

#### **GetNotificationsByUMKMID** {#getnotificationsbyumkmid-1 .unnumbered}

> GetNotificationsByUMKMID(ctx context.Context, umkmID int, limit,
> offset int) (\[\]model.Notification, error)
>
> **Deskripsi:** Mengambil notifikasi UMKM dengan pagination.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

- limit: Integer jumlah per halaman

- offset: Integer offset untuk pagination

> **Process:**

- Query by umkm_id

- Filter deleted_at

- Order by created_at DESC

- Limit dan offset

> **Output:**

- Slice of model.Notification

- Error jika gagal

#### **GetUnreadCount** {#getunreadcount-1 .unnumbered}

> GetUnreadCount(ctx context.Context, umkmID int) (int64, error)
>
> **Deskripsi:** Hitung jumlah notifikasi yang belum dibaca.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

> **Process:**

- Count notifikasi dengan is_read = false

- Filter by umkm_id dan deleted_at

> **Output:**

- Integer count

- Error jika gagal

#### **MarkAsRead** {#markasread .unnumbered}

> MarkAsRead(ctx context.Context, notificationIDs int, umkmID int) error
>
> **Deskripsi:** Tandai notifikasi sebagai sudah dibaca.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- notificationIDs: Integer ID notifikasi

- umkmID: Integer ID UMKM (untuk validasi kepemilikan)

> **Process:**

- Update is_read = true dan read_at = NOW()

- Filter by ID dan umkm_id

> **Output:**

- Error jika gagal

- nil jika sukses

#### **MarkAllAsRead** {#markallasread .unnumbered}

> MarkAllAsRead(ctx context.Context, umkmID int) error
>
> **Deskripsi:** Tandai semua notifikasi UMKM sebagai sudah dibaca.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

> **Process:**

- Update semua notifikasi dengan is_read = false

- Set is_read = true dan read_at = NOW()

> **Output:**

- Error jika gagal

- nil jika sukses

### OTP Repository

> Repository untuk manajemen OTP authentication.

#### **CreateOTP** {#createotp .unnumbered}

> CreateOTP(ctx context.Context, otp model.OTP) error
>
> **Deskripsi:** Menyimpan OTP baru.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- otp: Struct model.OTP

> **Process:**

- Insert ke database

> **Output:**

- Error jika gagal

- nil jika sukses

#### **GetOTPByPhone** {#getotpbyphone .unnumbered}

> GetOTPByPhone(ctx context.Context, phone string) (\*model.OTP, error)
>
> **Deskripsi:** Mengambil OTP terbaru berdasarkan nomor telepon.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- phone: String nomor telepon

> **Process:**

- Query by phone_number

- Order by created_at DESC

- Take 1

> **Output:**

- Pointer model.OTP atau nil jika tidak ada

- Error jika ada masalah query

> **Dependencies:**

- gorm.ErrRecordNotFound untuk deteksi tidak ada record

#### **GetOTPByTempToken** {#getotpbytemptoken .unnumbered}

> GetOTPByTempToken(ctx context.Context, tempToken string) (\*model.OTP,
> error)
>
> **Deskripsi:** Mengambil OTP berdasarkan temporary token.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- tempToken: String token temporary

> **Process:**

- Query by temp_token

- First record

> **Output:**

- Pointer model.OTP atau nil jika tidak ada

- Error jika ada masalah query

#### **UpdateOTP** {#updateotp .unnumbered}

> UpdateOTP(ctx context.Context, otp model.OTP) error
>
> **Deskripsi:** Update data OTP (biasanya status).
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- otp: Struct model.OTP dengan data baru

> **Process:**

- Update by phone_number dan otp_code

> **Output:**

- Error jika gagal

- nil jika sukses

### Programs Repository

> Repository untuk manajemen program (training/certification/funding).

#### **GetAllPrograms** {#getallprograms-1 .unnumbered}

> GetAllPrograms(ctx context.Context) (\[\]model.Program, error)
>
> **Deskripsi:** Mengambil semua program dengan relasi User (creator).
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Preload Users (creator)

- Filter deleted_at IS NULL

> **Output:**

- Slice of model.Program

- Error jika gagal

#### **GetProgramByID** {#getprogrambyid-2 .unnumbered}

> GetProgramByID(ctx context.Context, id int) (model.Program, error)
>
> **Deskripsi:** Mengambil detail program berdasarkan ID.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID program

> **Process:**

- Preload Users

- Filter deleted_at

> **Output:**

- model.Program

- Error "program not found" jika tidak ada

#### **CreateProgram** {#createprogram-1 .unnumbered}

> CreateProgram(ctx context.Context, program model.Program)
> (model.Program, error)
>
> **Deskripsi:** Membuat program baru.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- program: Struct model.Program

> **Process:**

- Insert ke database

> **Output:**

- model.Program yang telah dibuat

- Error "failed to create program" jika gagal

#### **UpdateProgram** {#updateprogram-1 .unnumbered}

> UpdateProgram(ctx context.Context, program model.Program)
> (model.Program, error)
>
> **Deskripsi:** Update data program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- program: Struct model.Program dengan data baru

> **Process:**

- Save perubahan

> **Output:**

- model.Program yang telah diupdate

- Error "failed to update program" jika gagal

#### **DeleteProgram** {#deleteprogram-1 .unnumbered}

> DeleteProgram(ctx context.Context, program model.Program)
> (model.Program, error)
>
> **Deskripsi:** Soft delete program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- program: Struct model.Program

> **Process:**

- Soft delete

> **Output:**

- model.Program yang telah dihapus

- Error "failed to delete program" jika gagal

#### **CreateProgramBenefits** {#createprogrambenefits .unnumbered}

> CreateProgramBenefits(ctx context.Context, benefits
> \[\]model.ProgramBenefit) error
>
> **Deskripsi:** Batch insert benefit program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- benefits: Slice of model.ProgramBenefit

> **Process:**

- Skip jika kosong

- Batch insert

> **Output:**

- Error "failed to create program benefits" jika gagal

- nil jika sukses

#### **CreateProgramRequirements** {#createprogramrequirements .unnumbered}

> CreateProgramRequirements(ctx context.Context, requirements
> \[\]model.ProgramRequirement) error
>
> **Deskripsi:** Batch insert syarat program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- requirements: Slice of model.ProgramRequirement

> **Process:**

- Skip jika kosong

- Batch insert

> **Output:**

- Error "failed to create program requirements" jika gagal

- nil jika sukses

#### **GetProgramBenefits** {#getprogrambenefits .unnumbered}

> GetProgramBenefits(ctx context.Context, programID int)
> (\[\]model.ProgramBenefit, error)
>
> **Deskripsi:** Mengambil benefit program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- programID: Integer ID program

> **Process:**

- Query by program_id

- Filter deleted_at

> **Output:**

- Slice of model.ProgramBenefit

- Error jika gagal

#### **GetProgramRequirements** {#getprogramrequirements-1 .unnumbered}

> GetProgramRequirements(ctx context.Context, programID int)
> (\[\]model.ProgramRequirement, error)
>
> **Deskripsi:** Mengambil syarat program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- programID: Integer ID program

> **Process:**

- Query by program_id

- Filter deleted_at

> **Output:**

- Slice of model.ProgramRequirement

- Error jika gagal

#### **DeleteProgramBenefits** {#deleteprogrambenefits .unnumbered}

> DeleteProgramBenefits(ctx context.Context, programID int) error
>
> **Deskripsi:** Hard delete semua benefit program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- programID: Integer ID program

> **Process:**

- Unscoped delete by program_id

> **Output:**

- Error "failed to delete program benefits" jika gagal

- nil jika sukses

#### **DeleteProgramRequirements** {#deleteprogramrequirements .unnumbered}

> DeleteProgramRequirements(ctx context.Context, programID int) error
>
> **Deskripsi:** Hard delete semua syarat program.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- programID: Integer ID program

> **Process:**

- Unscoped delete by program_id

> **Output:**

- Error "failed to delete program requirements" jika gagal

- nil jika sukses

### SLA Repository

> Repository untuk Service Level Agreement dan export data.

#### **GetSLAByStatus** {#getslabystatus .unnumbered}

> GetSLAByStatus(ctx context.Context, status string) (model.SLA, error)
>
> **Deskripsi:** Mengambil konfigurasi SLA berdasarkan status.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- status: String status (screening/final)

> **Process:**

- Query by status

- Filter deleted_at

> **Output:**

- model.SLA

- Error "SLA not found" jika tidak ada

#### **UpdateSLA** {#updatesla .unnumbered}

> UpdateSLA(ctx context.Context, sla model.SLA) (model.SLA, error)
>
> **Deskripsi:** Update konfigurasi SLA.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- sla: Struct model.SLA dengan data baru

> **Process:**

- Save perubahan

> **Output:**

- model.SLA yang telah diupdate

- Error "failed to update SLA" jika gagal

#### **GetApplicationsForExport** {#getapplicationsforexport .unnumbered}

> GetApplicationsForExport(ctx context.Context, applicationType string)
> (\[\]model.Application, error)
>
> **Deskripsi:** Mengambil data aplikasi untuk export.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- applicationType: String tipe atau "all"

> **Process:**

- Preload Program, UMKM.User, UMKM.City.Province

- Filter deleted_at

- Jika bukan "all", filter by type

> **Output:**

- Slice of model.Application dengan relasi lengkap

- Error jika gagal

#### **GetProgramsForExport** {#getprogramsforexport .unnumbered}

> GetProgramsForExport(ctx context.Context, applicationType string)
> (\[\]model.Program, error)
>
> **Deskripsi:** Mengambil data program untuk export.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- applicationType: String tipe atau "all"

> **Process:**

- Preload Users

- Filter deleted_at

- Jika bukan "all", filter by type

> **Output:**

- Slice of model.Program

- Error jika gagal

### Users Repository

> Repository untuk manajemen user, UMKM, dan role permissions.

#### **GetAllUsers** {#getallusers-1 .unnumbered}

> GetAllUsers(ctx context.Context) (\[\]model.User, error)
>
> **Deskripsi:** Mengambil semua user dengan relasi Role.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Preload Roles

- Find all users

> **Output:**

- Slice of model.User

- Error jika gagal

#### **GetUserByID** {#getuserbyid-1 .unnumbered}

> GetUserByID(ctx context.Context, id int) (model.User, error)
>
> **Deskripsi:** Mengambil user berdasarkan ID.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID user

> **Process:**

- Query by ID

> **Output:**

- model.User

- Error "user not found" jika tidak ada

#### **GetUserByEmail** {#getuserbyemail .unnumbered}

> GetUserByEmail(ctx context.Context, email string) (model.User, error)
>
> **Deskripsi:** Mengambil user berdasarkan email untuk login.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- email: String email

> **Process:**

- Query by email

> **Output:**

- model.User

- Error "user not found" jika tidak ada

#### **CreateUser** {#createuser .unnumbered}

> CreateUser(ctx context.Context, user model.User) (model.User, error)
>
> **Deskripsi:** Membuat user baru.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- user: Struct model.User

> **Process:**

- Insert ke database

> **Output:**

- model.User yang telah dibuat

- Error "failed to create user" jika gagal

#### **UpdateUser** {#updateuser-1 .unnumbered}

> UpdateUser(ctx context.Context, user model.User) (model.User, error)
>
> **Deskripsi:** Update data user.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- user: Struct model.User dengan data baru

> **Process:**

- Save perubahan

> **Output:**

- model.User yang telah diupdate

- Error "failed to update user" jika gagal

#### **DeleteUser** {#deleteuser-1 .unnumbered}

> DeleteUser(ctx context.Context, user model.User) (model.User, error)
>
> **Deskripsi:** Soft delete user.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- user: Struct model.User

> **Process:**

- Soft delete

> **Output:**

- model.User yang telah dihapus

- Error "failed to delete user" jika gagal

#### **CreateUMKM** {#createumkm .unnumbered}

> CreateUMKM(ctx context.Context, umkm model.UMKM, user model.User)
> (dto.UMKMMobile, error)
>
> **Deskripsi:** Membuat user dan UMKM sekaligus dalam transaction.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkm: Struct model.UMKM

- user: Struct model.User

> **Process:**

- Transaction

- Create user dulu

- Set umkm.UserID dari user.ID

- Create UMKM

- Map ke DTO response

> **Output:**

- dto.UMKMMobile dengan data gabungan

- Error jika gagal

> **Dependencies:**

- GORM Transaction

- dto.UMKMMobile

#### **GetUMKMByPhone** {#getumkmbyphone .unnumbered}

> GetUMKMByPhone(ctx context.Context, phone string) (model.UMKM, error)
>
> **Deskripsi:** Mengambil UMKM berdasarkan nomor telepon.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- phone: String nomor telepon

> **Process:**

- Preload User

- Query by phone

> **Output:**

- model.UMKM dengan relasi User

- Error "UMKM not found" jika tidak ada

#### **GetAllRoles** {#getallroles .unnumbered}

> GetAllRoles(ctx context.Context) (\[\]model.Role, error)
>
> **Deskripsi:** Mengambil semua role yang ada.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Find all roles

> **Output:**

- Slice of model.Role

- Error jika gagal

#### **GetRoleByID** {#getrolebyid .unnumbered}

> GetRoleByID(ctx context.Context, id int) (model.Role, error)
>
> **Deskripsi:** Mengambil role berdasarkan ID.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID role

> **Process:**

- Query by ID

> **Output:**

- model.Role

- Error "role not found" jika tidak ada

#### **GetRoleByName** {#getrolebyname .unnumbered}

> GetRoleByName(ctx context.Context, name string) (model.Role, error)
>
> **Deskripsi:** Mengambil role berdasarkan nama.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- name: String nama role

> **Process:**

- Query by name

> **Output:**

- model.Role

- Error "role not found" jika tidak ada

#### **IsRoleExist** {#isroleexist .unnumbered}

> IsRoleExist(ctx context.Context, id int) bool
>
> **Deskripsi:** Cek apakah role dengan ID tertentu ada.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- id: Integer ID role

> **Process:**

- Query by ID

> **Output:**

- Boolean true jika ada, false jika tidak

#### **IsPermissionExist** {#ispermissionexist .unnumbered}

> IsPermissionExist(ctx context.Context, ids \[\]string) (\[\]int, bool)
>
> **Deskripsi:** Validasi apakah semua permission code ada.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- ids: Slice string permission codes

> **Process:**

- Query permission IDs by codes (IN clause)

- Pluck IDs

- Compare jumlah hasil dengan jumlah input

> **Output:**

- Slice integer permission IDs

- Boolean true jika semua ada, false jika ada yang tidak ada

#### **GetListPermissions** {#getlistpermissions-1 .unnumbered}

> GetListPermissions(ctx context.Context) (\[\]model.Permission, error)
>
> **Deskripsi:** Mengambil semua permission yang punya parent (bukan
> root).
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Query dengan filter parent_id IS NOT NULL

> **Output:**

- Slice of model.Permission

- Error jika gagal

#### **GetListPermissionsByRoleID** {#getlistpermissionsbyroleid .unnumbered}

> GetListPermissionsByRoleID(ctx context.Context, roleID int)
> (\[\]string, error)
>
> **Deskripsi:** Mengambil permission codes dari suatu role.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- roleID: Integer ID role

> **Process:**

- Raw SQL dengan jsonb_agg untuk aggregate permission codes

- Join role_permissions, roles, permissions

- Unmarshal JSON result

> **Output:**

- Slice string permission codes

- Error jika gagal atau parsing gagal

> **Dependencies:**

- encoding/json untuk unmarshal

#### **GetListRolePermissions** {#getlistrolepermissions-1 .unnumbered}

> GetListRolePermissions(ctx context.Context)
> (\[\]dto.RolePermissionsResponse, error)
>
> **Deskripsi:** Mengambil mapping semua role dengan permissions-nya.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Raw SQL dengan jsonb_agg

- GROUP BY role

- Return role_id, role_name, permissions (jsonb)

> **Output:**

- Slice of dto.RolePermissionsResponse

- Error jika gagal

> **Dependencies:**

- dto.RolePermissionsResponse dengan field Permissions json.RawMessage

#### **DeletePermissionsByRoleID** {#deletepermissionsbyroleid .unnumbered}

> DeletePermissionsByRoleID(ctx context.Context, roleID int) error
>
> **Deskripsi:** Hard delete semua permissions dari role.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- roleID: Integer ID role

> **Process:**

- Unscoped delete by role_id

> **Output:**

- Error "failed to delete role permissions" jika gagal

- nil jika sukses

#### **AddRolePermissions** {#addrolepermissions .unnumbered}

> AddRolePermissions(ctx context.Context, roleID int, permissions
> \[\]int) error
>
> **Deskripsi:** Batch insert role permissions.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- roleID: Integer ID role

- permissions: Slice integer permission IDs

> **Process:**

- Loop create struct RolePermission

- Batch insert dengan Omit timestamps

> **Output:**

- Error "failed to add role permissions" jika gagal

- nil jika sukses

#### **GetProvinces** {#getprovinces .unnumbered}

> GetProvinces(ctx context.Context) (\[\]dto.Province, error)
>
> **Deskripsi:** Mengambil semua data provinsi.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Find all provinces

- Map ke DTO

> **Output:**

- Slice of dto.Province

- Error "failed to get provinces" jika gagal

> **Dependencies:**

- dto.Province

#### **GetCities** {#getcities .unnumbered}

> GetCities(ctx context.Context) (\[\]dto.City, error)
>
> **Deskripsi:** Mengambil semua data kota/kabupaten.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

> **Process:**

- Find all cities

- Map ke DTO

> **Output:**

- Slice of dto.City

- Error "failed to get cities" jika gagal

> **Dependencies:**

- dto.City

### Vault Decrypt Logs Repository

> Repository untuk logging aktivitas dekripsi data sensitif.

#### **LogDecrypt** {#logdecrypt .unnumbered}

> LogDecrypt(ctx context.Context, log model.VaultDecryptLog) error
>
> **Deskripsi:** Mencatat aktivitas dekripsi data.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- log: Struct model.VaultDecryptLog

> **Process:**

- Insert log ke database

> **Output:**

- Error jika gagal

- nil jika sukses

#### **GetLogs** {#getlogs-1 .unnumbered}

> GetLogs(ctx context.Context, limit, offset int)
> (\[\]model.VaultDecryptLog, error)
>
> **Deskripsi:** Mengambil semua log dekripsi dengan pagination.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- limit: Integer jumlah per halaman

- offset: Integer offset pagination

> **Process:**

- Query dengan order by decrypted_at DESC

- Limit dan offset

> **Output:**

- Slice of model.VaultDecryptLog

- Error jika gagal

#### **GetLogsByUserID** {#getlogsbyuserid-1 .unnumbered}

> GetLogsByUserID(ctx context.Context, userID int, limit, offset int)
> (\[\]model.VaultDecryptLog, error)
>
> **Deskripsi:** Mengambil log dekripsi berdasarkan user yang melakukan.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- userID: Integer ID user

- limit: Integer jumlah per halaman

- offset: Integer offset pagination

> **Process:**

- Filter by user_id

- Order by decrypted_at DESC

- Limit dan offset

> **Output:**

- Slice of model.VaultDecryptLog

- Error jika gagal

#### **GetLogsByUMKMID** {#getlogsbyumkmid-1 .unnumbered}

> GetLogsByUMKMID(ctx context.Context, umkmID int, limit, offset int)
> (\[\]model.VaultDecryptLog, error)
>
> **Deskripsi:** Mengambil log dekripsi data UMKM tertentu.
>
> **Input:**

- ctx: Context untuk kontrol eksekusi

- umkmID: Integer ID UMKM

- limit: Integer jumlah per halaman

- offset: Integer offset pagination

> **Process:**

- Filter by umkm_id

- Order by decrypted_at DESC

- Limit dan offset

> **Output:**

- Slice of model.VaultDecryptLog

- Error jika gagal

## Unit Test  {#unit-test}

### Users Service Tests

#### **TestNewUsersService** {#testnewusersservice .unnumbered}

> **Function:** TestNewUsersService
>
> **Purpose:** Memverifikasi bahwa constructor service users dapat
> membuat instance service dengan benar.
>
> **Input:**

- Mock UsersRepository

- Mock Redis client

- Logger instance

> **Process:**

- Memanggil NewUsersService() dengan dependencies

- Memverifikasi instance yang dibuat tidak nil

> **Output:**

- Instance UsersService yang valid

> **Dependencies:**

- repository.UsersRepository (mock)

- redis.Client (mock)

- logrus.Logger

> **Tests For:** Constructor NewUsersService

#### **TestLogin** {#testlogin .unnumbered}

> **Function:** TestLogin
>
> **Purpose:** Menguji proses login user untuk web dashboard.
>
> **Input:**

- dto.LoginRequest berisi email dan password

> **Process:**

- Mencari user berdasarkan email via repository

- Memverifikasi password menggunakan bcrypt

- Memeriksa apakah user aktif (deleted_at null)

- Generate JWT token jika login berhasil

> **Output:**

- dto.LoginResponse berisi user data dan token

- Error jika email tidak ditemukan, password salah, atau user sudah
  > dihapus

> **Test Cases:**

- Success: Login berhasil dengan kredensial valid

- Error: Email tidak ditemukan

- Error: Password salah

- Error: User sudah dihapus (soft deleted)

> **Utils & Types:**

- bcrypt untuk verifikasi password

- jwt untuk generate token

- entities.Users

- dto.LoginRequest

- dto.LoginResponse

> **Tests For:** Handler Login di users handler

#### **TestRegister** {#testregister .unnumbered}

> **Function:** TestRegister
>
> **Purpose:** Menguji proses registrasi user baru untuk web dashboard.
>
> **Input:**

- dto.RegisterRequest berisi name, email, password, role

> **Process:**

- Validasi email belum terdaftar

- Hash password menggunakan bcrypt

- Insert user baru ke database

- Generate JWT token untuk auto-login

> **Output:**

- dto.RegisterResponse berisi user data dan token

- Error jika email sudah terdaftar atau gagal insert

> **Test Cases:**

- Success: Registrasi berhasil dengan data valid

- Error: Email sudah terdaftar

- Error: Gagal hash password

- Error: Gagal insert ke database

> **Utils & Types:**

- bcrypt untuk hash password

- jwt untuk generate token

- entities.Users

- dto.RegisterRequest

- dto.RegisterResponse

> **Tests For:** Handler Register di users handler

#### **TestGetMeta** {#testgetmeta .unnumbered}

> **Function:** TestGetMeta
>
> **Purpose:** Menguji pengambilan data master provinsi dan kota.
>
> **Input:** Tidak ada parameter
>
> **Process:**

- Mengambil data provinsi dan kota dari repository

- Mengembalikan dalam format response

> **Output:**

- dto.MetaResponse berisi list provinsi dan kota

- Error jika gagal mengambil data

> **Test Cases:**

- Success: Berhasil mengambil data master

- Error: Gagal query database

> **Utils & Types:**

- entities.Province

- entities.City

- dto.MetaResponse

> **Tests For:** Handler GetMeta di users handler

#### **TestLoginMobile** {#testloginmobile .unnumbered}

> **Function:** TestLoginMobile
>
> **Purpose:** Menguji proses login untuk aplikasi mobile (UMKM).
>
> **Input:**

- dto.LoginMobileRequest berisi email dan password

> **Process:**

- Mencari user dengan role UMKM berdasarkan email

- Memverifikasi password

- Memeriksa status user (aktif/tidak)

- Generate JWT token khusus mobile

> **Output:**

- dto.LoginMobileResponse berisi UMKM data dan token

- Error jika kredensial salah atau user tidak aktif

> **Test Cases:**

- Success: Login mobile berhasil

- Error: User tidak ditemukan

- Error: Password salah

- Error: User bukan role UMKM

- Error: User sudah dihapus

> **Utils & Types:**

- bcrypt untuk verifikasi password

- jwt untuk generate token

- entities.Users

- dto.LoginMobileRequest

- dto.LoginMobileResponse

> **Tests For:** Handler LoginMobile di users handler

#### **TestRegisterMobile** {#testregistermobile .unnumbered}

> **Function:** TestRegisterMobile
>
> **Purpose:** Menguji proses registrasi mobile yang mengirim OTP.
>
> **Input:**

- dto.RegisterMobileRequest berisi email dan phone

> **Process:**

- Validasi email dan phone belum terdaftar

- Generate OTP 6 digit

- Simpan OTP ke Redis dengan TTL

- Kirim OTP via SMS/Email (mock)

- Insert user temporary dengan status pending

> **Output:**

- dto.RegisterMobileResponse berisi temp_token untuk verifikasi

- Error jika email/phone sudah terdaftar

> **Test Cases:**

- Success: OTP berhasil dikirim

- Error: Email sudah terdaftar

- Error: Phone sudah terdaftar

- Error: Gagal generate OTP

- Error: Gagal simpan ke Redis

> **Utils & Types:**

- redis untuk simpan OTP

- Random generator untuk OTP

- entities.Users

- dto.RegisterMobileRequest

- dto.RegisterMobileResponse

> **Tests For:** Handler RegisterMobile di users handler

#### **TestRegisterMobileProfile** {#testregistermobileprofile .unnumbered}

> **Function:** TestRegisterMobileProfile
>
> **Purpose:** Menguji proses melengkapi profil setelah verifikasi OTP.
>
> **Input:**

- dto.RegisterMobileProfileRequest berisi name, password, business_name,
  > dll

- temp_token dari proses RegisterMobile

> **Process:**

- Validasi temp_token

- Hash password

- Update user data dengan profil lengkap

- Generate JWT token final

- Hapus data temporary dari Redis

> **Output:**

- dto.LoginMobileResponse berisi user data dan token

- Error jika temp_token invalid

> **Test Cases:**

- Success: Profil berhasil dilengkapi

- Error: Temp token invalid/expired

- Error: Gagal update profil

- Error: Gagal generate token

> **Utils & Types:**

- bcrypt untuk hash password

- jwt untuk generate token

- redis untuk validasi temp_token

- entities.Users

- dto.RegisterMobileProfileRequest

> **Tests For:** Handler RegisterMobileProfile di users handler

#### **TestForgotPassword** {#testforgotpassword .unnumbered}

> **Function:** TestForgotPassword
>
> **Purpose:** Menguji proses request reset password yang mengirim OTP.
>
> **Input:**

- dto.ForgotPasswordRequest berisi email

> **Process:**

- Cari user berdasarkan email

- Generate OTP 6 digit

- Simpan OTP ke Redis dengan TTL

- Kirim OTP via email/SMS

- Generate temp_token untuk reset

> **Output:**

- dto.ForgotPasswordResponse berisi temp_token

- Error jika email tidak ditemukan

> **Test Cases:**

- Success: OTP berhasil dikirim

- Error: Email tidak ditemukan

- Error: User sudah dihapus

- Error: Gagal generate/kirim OTP

> **Utils & Types:**

- redis untuk simpan OTP

- Random generator untuk OTP

- entities.Users

- dto.ForgotPasswordRequest

- dto.ForgotPasswordResponse

> **Tests For:** Handler ForgotPassword di users handler

#### **TestResetPassword** {#testresetpassword .unnumbered}

> **Function:** TestResetPassword
>
> **Purpose:** Menguji proses reset password dengan temp token.
>
> **Input:**

- dto.ResetPasswordRequest berisi temp_token dan new_password

> **Process:**

- Validasi temp_token dari Redis

- Hash password baru

- Update password user di database

- Hapus temp_token dari Redis

> **Output:**

- Success response

- Error jika temp_token invalid atau expired

> **Test Cases:**

- Success: Password berhasil direset

- Error: Temp token invalid/expired

- Error: Gagal hash password

- Error: Gagal update database

> **Utils & Types:**

- bcrypt untuk hash password

- redis untuk validasi temp_token

- entities.Users

- dto.ResetPasswordRequest

> **Tests For:** Handler ResetPassword di users handler

#### **TestVerifyOTP** {#testverifyotp .unnumbered}

> **Function:** TestVerifyOTP
>
> **Purpose:** Menguji proses verifikasi kode OTP.
>
> **Input:**

- dto.VerifyOTPRequest berisi temp_token dan otp_code

> **Process:**

- Ambil OTP dari Redis berdasarkan temp_token

- Compare OTP yang diinput dengan yang tersimpan

- Jika cocok, tandai sebagai verified

> **Output:**

- Success response jika OTP valid

- Error jika OTP salah atau expired

> **Test Cases:**

- Success: OTP valid dan verified

- Error: OTP tidak ditemukan (expired)

- Error: OTP tidak cocok

- Error: Temp token invalid

> **Utils & Types:**

- redis untuk get/set OTP

- dto.VerifyOTPRequest

> **Tests For:** Handler VerifyOTP di users handler

#### **TestGetAllUsers** {#testgetallusers .unnumbered}

> **Function:** TestGetAllUsers
>
> **Purpose:** Menguji pengambilan daftar semua user dengan pagination
> dan filter.
>
> **Input:**

- Query params: page, limit, search, role, status

> **Process:**

- Build query dengan filter (search by name/email, filter by
  > role/status)

- Execute query dengan pagination

- Count total records

- Return paginated result

> **Output:**

- dto.PaginatedUsersResponse berisi list users dan metadata pagination

- Error jika query gagal

> **Test Cases:**

- Success: Get all users tanpa filter

- Success: Get users dengan search

- Success: Get users dengan filter role

- Success: Get users dengan pagination

- Error: Query database gagal

> **Utils & Types:**

- entities.Users

- dto.PaginatedUsersResponse

- Pagination helper

> **Tests For:** Handler GetAllUsers di users handler

#### **TestGetUserByID** {#testgetuserbyid .unnumbered}

> **Function:** TestGetUserByID
>
> **Purpose:** Menguji pengambilan data user berdasarkan ID.
>
> **Input:**

- User ID (integer)

> **Process:**

- Query user dari database berdasarkan ID

- Include relations (province, city, role)

- Return user data

> **Output:**

- entities.Users dengan relasi lengkap

- Error jika user tidak ditemukan

> **Test Cases:**

- Success: User ditemukan

- Error: User tidak ditemukan

- Error: Query gagal

> **Utils & Types:**

- entities.Users

- entities.Province

- entities.City

> **Tests For:** Handler GetUserByID di users handler

#### **TestUpdateUser** {#testupdateuser .unnumbered}

> **Function:** TestUpdateUser
>
> **Purpose:** Menguji proses update data user.
>
> **Input:**

- User ID

- dto.UpdateUserRequest berisi field yang akan diupdate

> **Process:**

- Validasi user exists

- Jika password diupdate, hash password baru

- Update data ke database

- Return updated user data

> **Output:**

- entities.Users yang sudah diupdate

- Error jika user tidak ditemukan atau gagal update

> **Test Cases:**

- Success: Update user tanpa password

- Success: Update user dengan password baru

- Error: User tidak ditemukan

- Error: Gagal hash password

- Error: Gagal update database

> **Utils & Types:**

- bcrypt untuk hash password

- entities.Users

- dto.UpdateUserRequest

> **Tests For:** Handler UpdateUser di users handler

#### **TestUpdateProfile** {#testupdateprofile .unnumbered}

> **Function:** TestUpdateProfile
>
> **Purpose:** Menguji update profil user yang sedang login.
>
> **Input:**

- User ID dari JWT token

- dto.UpdateProfileRequest berisi field profil

> **Process:**

- Get user dari context (authenticated user)

- Update field profil

- Save ke database

> **Output:**

- entities.Users yang sudah diupdate

- Error jika gagal update

> **Test Cases:**

- Success: Update profil berhasil

- Error: User tidak authenticated

- Error: Gagal update database

> **Utils & Types:**

- entities.Users

- dto.UpdateProfileRequest

> **Tests For:** Handler UpdateProfile di users handler

#### **TestDeleteUser** {#testdeleteuser .unnumbered}

> **Function:** TestDeleteUser
>
> **Purpose:** Menguji soft delete user.
>
> **Input:**

- User ID

> **Process:**

- Set deleted_at timestamp

- Update status menjadi inactive

- User masih ada di database tapi tidak bisa login

> **Output:**

- Success response

- Error jika user tidak ditemukan

> **Test Cases:**

- Success: Soft delete berhasil

- Error: User tidak ditemukan

- Error: User sudah dihapus sebelumnya

- Error: Gagal update database

> **Utils & Types:**

- entities.Users

- Soft delete pattern

> **Tests For:** Handler DeleteUser di users handler

#### **TestGetListPermissions** {#testgetlistpermissions .unnumbered}

> **Function:** TestGetListPermissions
>
> **Purpose:** Menguji pengambilan daftar semua permission yang
> tersedia.
>
> **Input:** Tidak ada parameter
>
> **Process:**

- Query semua permissions dari database

- Return list permissions

> **Output:**

- Array of entities.Permission

- Error jika query gagal

> **Test Cases:**

- Success: Get all permissions

- Error: Query database gagal

> **Utils & Types:**

- entities.Permission

> **Tests For:** Handler GetListPermissions di users handler

#### **TestGetListRolePermissions** {#testgetlistrolepermissions .unnumbered}

> **Function:** TestGetListRolePermissions
>
> **Purpose:** Menguji pengambilan mapping role dan permissions.
>
> **Input:** Optional role_id untuk filter
>
> **Process:**

- Query role_permissions table

- Include relasi permission

- Group by role jika diperlukan

> **Output:**

- Array of entities.RolePermission

- Error jika query gagal

> **Test Cases:**

- Success: Get all role permissions

- Success: Get role permissions by role_id

- Error: Query gagal

> **Utils & Types:**

- entities.RolePermission

- entities.Permission

> **Tests For:** Handler GetListRolePermissions di users handler

#### **TestUpdateRolePermissions** {#testupdaterolepermissions .unnumbered}

> **Function:** TestUpdateRolePermissions
>
> **Purpose:** Menguji update permissions untuk suatu role.
>
> **Input:**

- dto.UpdateRolePermissionsRequest berisi role_id dan array
  > permission_ids

> **Process:**

- Begin transaction

- Delete existing role permissions

- Insert new role permissions

- Commit transaction

> **Output:**

- Success response

- Error jika gagal (rollback transaction)

> **Test Cases:**

- Success: Update role permissions berhasil

- Error: Role tidak ditemukan

- Error: Permission tidak valid

- Error: Gagal update (rollback)

> **Utils & Types:**

- entities.RolePermission

- dto.UpdateRolePermissionsRequest

- Database transaction

> **Tests For:** Handler UpdateRolePermissions di users handler

### Programs Service Tests

#### **TestNewProgramsService** {#testnewprogramsservice .unnumbered}

> **Function:** TestNewProgramsService
>
> **Purpose:** Memverifikasi constructor service programs.
>
> **Input:**

- Mock ProgramsRepository

- Mock UsersRepository

- Mock Redis client

- Mock MinIO client

> **Process:**

- Memanggil NewProgramsService() dengan dependencies

- Memverifikasi instance tidak nil

> **Output:**

- Instance ProgramsService yang valid

> **Dependencies:**

- repository.ProgramsRepository (mock)

- repository.UsersRepository (mock)

- redis.Client (mock)

- minio.Client (mock)

> **Tests For:** Constructor NewProgramsService

#### **TestGetAllPrograms** {#testgetallprograms .unnumbered}

> **Function:** TestGetAllPrograms
>
> **Purpose:** Menguji pengambilan semua program dengan filter dan
> pagination.
>
> **Input:**

- Query params: page, limit, program_type, status, search

> **Process:**

- Build query dengan filter

- Apply pagination

- Count total records

- Return paginated result

> **Output:**

- dto.PaginatedProgramsResponse berisi list programs

- Error jika query gagal

> **Test Cases:**

- Success: Get all programs

- Success: Filter by program_type (training/certification/funding)

- Success: Filter by status (active/inactive)

- Success: Search by title

- Error: Query gagal

> **Utils & Types:**

- entities.Program

- dto.PaginatedProgramsResponse

- Pagination helper

> **Tests For:** Handler GetAllPrograms di programs handler

#### **TestGetProgramByID** {#testgetprogrambyid .unnumbered}

> **Function:** TestGetProgramByID
>
> **Purpose:** Menguji pengambilan detail program berdasarkan ID.
>
> **Input:**

- Program ID

> **Process:**

- Query program dari database

- Include image URL dari MinIO

- Return program detail

> **Output:**

- entities.Program dengan detail lengkap

- Error jika tidak ditemukan

> **Test Cases:**

- Success: Program ditemukan

- Error: Program tidak ditemukan

- Error: Query gagal

> **Utils & Types:**

- entities.Program

- MinIO untuk get image URL

> **Tests For:** Handler GetProgramByID di programs handler

#### **TestCreateProgram** {#testcreateprogram .unnumbered}

> **Function:** TestCreateProgram
>
> **Purpose:** Menguji pembuatan program baru.
>
> **Input:**

- dto.CreateProgramRequest berisi title, description, program_type, dll

- Image file (multipart)

> **Process:**

- Validate input

- Upload image ke MinIO

- Insert program ke database dengan image_url

- Return created program

> **Output:**

- entities.Program yang baru dibuat

- Error jika validasi gagal atau upload gagal

> **Test Cases:**

- Success: Create program dengan image

- Success: Create program tanpa image

- Error: Validasi gagal (required fields)

- Error: Upload image gagal

- Error: Insert database gagal

> **Utils & Types:**

- entities.Program

- dto.CreateProgramRequest

- MinIO untuk upload image

> **Tests For:** Handler CreateProgram di programs handler

#### **TestUpdateProgram** {#testupdateprogram .unnumbered}

> **Function:** TestUpdateProgram
>
> **Purpose:** Menguji update data program.
>
> **Input:**

- Program ID

- dto.UpdateProgramRequest berisi field yang akan diupdate

- Optional: image file baru

> **Process:**

- Validate program exists

- Jika ada image baru, hapus image lama dan upload baru

- Update data program

- Return updated program

> **Output:**

- entities.Program yang sudah diupdate

- Error jika program tidak ditemukan atau gagal update

> **Test Cases:**

- Success: Update program tanpa image

- Success: Update program dengan image baru

- Error: Program tidak ditemukan

- Error: Upload image gagal

- Error: Update database gagal

> **Utils & Types:**

- entities.Program

- dto.UpdateProgramRequest

- MinIO untuk upload/delete image

> **Tests For:** Handler UpdateProgram di programs handler

#### **TestActivateProgram** {#testactivateprogram .unnumbered}

> **Function:** TestActivateProgram
>
> **Purpose:** Menguji aktivasi program.
>
> **Input:**

- Program ID

> **Process:**

- Set status program menjadi active

- Set activated_at timestamp

- Update database

> **Output:**

- Success response

- Error jika program tidak ditemukan

> **Test Cases:**

- Success: Program berhasil diaktifkan

- Error: Program tidak ditemukan

- Error: Program sudah aktif

- Error: Update gagal

> **Utils & Types:**

- entities.Program

> **Tests For:** Handler ActivateProgram di programs handler

#### **TestDeactivateProgram** {#testdeactivateprogram .unnumbered}

> **Function:** TestDeactivateProgram
>
> **Purpose:** Menguji deaktivasi program.
>
> **Input:**

- Program ID

> **Process:**

- Set status program menjadi inactive

- Set deactivated_at timestamp

- Update database

> **Output:**

- Success response

- Error jika program tidak ditemukan

> **Test Cases:**

- Success: Program berhasil dinonaktifkan

- Error: Program tidak ditemukan

- Error: Program sudah inactive

- Error: Update gagal

> **Utils & Types:**

- entities.Program

> **Tests For:** Handler DeactivateProgram di programs handler

#### **TestDeleteProgram** {#testdeleteprogram .unnumbered}

> **Function:** TestDeleteProgram
>
> **Purpose:** Menguji soft delete program.
>
> **Input:**

- Program ID

> **Process:**

- Set deleted_at timestamp

- Hapus image dari MinIO

- Program tidak muncul di list tapi data masih ada

> **Output:**

- Success response

- Error jika program tidak ditemukan

> **Test Cases:**

- Success: Soft delete berhasil

- Error: Program tidak ditemukan

- Error: Program sudah dihapus

- Error: Gagal hapus image

- Error: Update database gagal

> **Utils & Types:**

- entities.Program

- MinIO untuk delete image

- Soft delete pattern

> **Tests For:** Handler DeleteProgram di programs handler

### Applications Service Tests

#### **TestNewApplicationsService** {#testnewapplicationsservice .unnumbered}

> **Function:** TestNewApplicationsService
>
> **Purpose:** Memverifikasi constructor service applications.
>
> **Input:**

- Mock ApplicationsRepository

- Mock UsersRepository

- Mock NotificationRepository

- Mock SLARepository

- Mock VaultDecryptLogRepository

> **Process:**

- Memanggil NewApplicationsService() dengan dependencies

- Memverifikasi instance tidak nil

> **Output:**

- Instance ApplicationsService yang valid

> **Dependencies:**

- repository.ApplicationsRepository (mock)

- repository.UsersRepository (mock)

- repository.NotificationRepository (mock)

- repository.SLARepository (mock)

- repository.VaultDecryptLogRepository (mock)

> **Tests For:** Constructor NewApplicationsService

#### **TestGetAllApplications** {#testgetallapplications .unnumbered}

> **Function:** TestGetAllApplications
>
> **Purpose:** Menguji pengambilan semua aplikasi dengan filter.
>
> **Input:**

- Query params: page, limit, status, program_type, umkm_id, date_range

> **Process:**

- Build complex query dengan multiple filters

- Apply pagination

- Include relasi (program, umkm)

- Calculate SLA status

- Return paginated result

> **Output:**

- dto.PaginatedApplicationsResponse dengan SLA info

- Error jika query gagal

> **Test Cases:**

- Success: Get all applications

- Success: Filter by status (pending/approved/rejected)

- Success: Filter by program_type

- Success: Filter by umkm_id

- Success: Filter by date range

- Error: Query gagal

> **Utils & Types:**

- entities.Application

- entities.Program

- entities.Users

- dto.PaginatedApplicationsResponse

- SLA calculation helper

> **Tests For:** Handler GetAllApplications di applications handler

#### **TestGetApplicationByID** {#testgetapplicationbyid .unnumbered}

> **Function:** TestGetApplicationByID
>
> **Purpose:** Menguji pengambilan detail aplikasi.
>
> **Input:**

- Application ID

> **Process:**

- Query application dengan semua relasi

- Decrypt sensitive data dari Vault

- Log decrypt activity

- Calculate SLA status

- Return detail lengkap

> **Output:**

- entities.Application dengan data lengkap dan decrypted

- Error jika tidak ditemukan

> **Test Cases:**

- Success: Application ditemukan dengan data decrypted

- Error: Application tidak ditemukan

- Error: Decrypt gagal

- Error: Query gagal

> **Utils & Types:**

- entities.Application

- Vault client untuk decrypt

- entities.VaultDecryptLog

- SLA calculation

> **Tests For:** Handler GetApplicationByID di applications handler

#### **TestScreeningApprove** {#testscreeningapprove .unnumbered}

> **Function:** TestScreeningApprove
>
> **Purpose:** Menguji approve aplikasi pada tahap screening.
>
> **Input:**

- Application ID

- User ID (admin yang approve)

- Optional: screening notes

> **Process:**

- Validate application status (harus pending_screening)

- Update status menjadi screening_approved

- Set screening_approved_by dan screening_approved_at

- Create notification untuk UMKM

- Log activity

> **Output:**

- Updated application

- Notification created

- Error jika status tidak valid

> **Test Cases:**

- Success: Screening approve berhasil

- Error: Application tidak ditemukan

- Error: Status tidak valid untuk approve

- Error: User tidak berhak approve

- Error: Update gagal

> **Utils & Types:**

- entities.Application

- entities.Notification

- Status validation helper

> **Tests For:** Handler ScreeningApprove di applications handler

#### **TestScreeningReject** {#testscreeningreject .unnumbered}

> **Function:** TestScreeningReject
>
> **Purpose:** Menguji reject aplikasi pada tahap screening.
>
> **Input:**

- Application ID

- User ID (admin yang reject)

- Rejection reason (required)

> **Process:**

- Validate application status

- Update status menjadi screening_rejected

- Set rejection_reason dan rejected_by

- Create notification untuk UMKM

- Log activity

> **Output:**

- Updated application

- Notification created

- Error jika reason kosong

> **Test Cases:**

- Success: Screening reject dengan reason

- Error: Rejection reason kosong

- Error: Application tidak ditemukan

- Error: Status tidak valid

- Error: Update gagal

> **Utils & Types:**

- entities.Application

- entities.Notification

- Rejection reason validation

> **Tests For:** Handler ScreeningReject di applications handler

#### **TestScreeningRevise** {#testscreeningrevise .unnumbered}

> **Function:** TestScreeningRevise
>
> **Purpose:** Menguji request revisi pada tahap screening.
>
> **Input:**

- Application ID

- Revision notes (required)

> **Process:**

- Update status menjadi revision_requested

- Set revision_notes

- Create notification untuk UMKM

- UMKM dapat update dan resubmit

> **Output:**

- Updated application

- Notification created

- Error jika notes kosong

> **Test Cases:**

- Success: Revisi berhasil direquest

- Error: Revision notes kosong

- Error: Application tidak ditemukan

- Error: Status tidak valid

- Error: Update gagal

> **Utils & Types:**

- entities.Application

- entities.Notification

- Revision notes validation

> **Tests For:** Handler ScreeningRevise di applications handler

#### **TestFinalApprove** {#testfinalapprove .unnumbered}

> **Function:** TestFinalApprove
>
> **Purpose:** Menguji final approval aplikasi.
>
> **Input:**

- Application ID

- User ID (admin yang approve)

- Optional: final notes

> **Process:**

- Validate status (harus screening_approved)

- Update status menjadi final_approved

- Set final_approved_by dan final_approved_at

- Create notification untuk UMKM

- Log completion

> **Output:**

- Updated application

- Notification created

- Error jika status tidak valid

> **Test Cases:**

- Success: Final approve berhasil

- Error: Application tidak dalam status screening_approved

- Error: Application tidak ditemukan

- Error: Update gagal

> **Utils & Types:**

- entities.Application

- entities.Notification

- Status workflow validation

> **Tests For:** Handler FinalApprove di applications handler

#### **TestFinalReject** {#testfinalreject .unnumbered}

> **Function:** TestFinalReject
>
> **Purpose:** Menguji final rejection aplikasi.
>
> **Input:**

- Application ID

- User ID (admin yang reject)

- Rejection reason (required)

> **Process:**

- Validate status

- Update status menjadi final_rejected

- Set rejection details

- Create notification

- Log final decision

> **Output:**

- Updated application

- Notification created

- Error jika reason kosong

> **Test Cases:**

- Success: Final reject dengan reason

- Error: Rejection reason kosong

- Error: Status tidak valid

- Error: Update gagal

> **Utils & Types:**

- entities.Application

- entities.Notification

- Final decision validation

> **Tests For:** Handler FinalReject di applications handler

### Dashboard Service Tests

#### **TestNewDashboardService** {#testnewdashboardservice .unnumbered}

> **Function:** TestNewDashboardService
>
> **Purpose:** Memverifikasi constructor service dashboard.
>
> **Input:**

- Mock DashboardRepository

> **Process:**

- Memanggil NewDashboardService() dengan repository

- Memverifikasi instance tidak nil

> **Output:**

- Instance DashboardService yang valid

> **Dependencies:**

- repository.DashboardRepository (mock)

> **Tests For:** Constructor NewDashboardService

#### **TestGetUMKMByCardType** {#testgetumkmbycardtype .unnumbered}

> **Function:** TestGetUMKMByCardType
>
> **Purpose:** Menguji statistik UMKM berdasarkan tipe kartu.
>
> **Input:**

- Optional date range filter

> **Process:**

- Query aggregate data dari users

- Group by card_type (Gold, Silver, Bronze)

- Count per kategori

- Return statistics

> **Output:**

- dto.UMKMByCardTypeResponse berisi count per card type

- Error jika query gagal

> **Test Cases:**

- Success: Get card type statistics

- Success: Filter by date range

- Error: Query gagal

> **Utils & Types:**

- dto.UMKMByCardTypeResponse

- Aggregate query helper

> **Tests For:** Handler GetUMKMByCardType di dashboard handler

#### **TestGetApplicationStatusSummary** {#testgetapplicationstatussummary .unnumbered}

> **Function:** TestGetApplicationStatusSummary
>
> **Purpose:** Menguji ringkasan status aplikasi.
>
> **Input:**

- Optional date range

- Optional program type filter

> **Process:**

- Count applications by status

- Calculate percentages

- Group by program type if filtered

- Return summary

> **Output:**

- dto.ApplicationStatusSummary berisi count per status

- Error jika query gagal

> **Test Cases:**

- Success: Get status summary

- Success: Filter by program type

- Success: Filter by date range

- Error: Query gagal

> **Utils & Types:**

- dto.ApplicationStatusSummary

- Status grouping helper

> **Tests For:** Handler GetApplicationStatusSummary di dashboard
> handler

#### **TestGetApplicationStatusDetail** {#testgetapplicationstatusdetail .unnumbered}

> **Function:** TestGetApplicationStatusDetail
>
> **Purpose:** Menguji detail status aplikasi dengan breakdown.
>
> **Input:**

- Status filter

- Optional date range

> **Process:**

- Get detailed application list by status

- Include program info

- Include UMKM info

- Calculate SLA for each

- Return detailed list

> **Output:**

- dto.ApplicationStatusDetail berisi list applications

- Error jika query gagal

> **Test Cases:**

- Success: Get detail by status

- Success: Filter by date range

- Success: Include SLA info

- Error: Query gagal

> **Utils & Types:**

- dto.ApplicationStatusDetail

- entities.Application

- SLA calculation

> **Tests For:** Handler GetApplicationStatusDetail di dashboard handler

#### **TestGetApplicationByType** {#testgetapplicationbytype .unnumbered}

> **Function:** TestGetApplicationByType
>
> **Purpose:** Menguji statistik aplikasi berdasarkan tipe program.
>
> **Input:**

- Optional date range

> **Process:**

- Count applications by program_type

- Group by training, certification, funding

- Calculate totals and percentages

- Return statistics

> **Output:**

- dto.ApplicationByTypeResponse berisi count per type

- Error jika query gagal

> **Test Cases:**

- Success: Get type statistics

- Success: Filter by date range

- Error: Query gagal

> **Utils & Types:**

- dto.ApplicationByTypeResponse

- Type grouping helper

> **Tests For:** Handler GetApplicationByType di dashboard handler

### SLA Service Tests

#### **TestNewSLAService** {#testnewslaservice .unnumbered}

> **Function:** TestNewSLAService
>
> **Purpose:** Memverifikasi constructor service SLA.
>
> **Input:**

- Mock SLARepository

> **Process:**

- Memanggil NewSLAService() dengan repository

- Memverifikasi instance tidak nil

> **Output:**

- Instance SLAService yang valid

> **Dependencies:**

- repository.SLARepository (mock)

> **Tests For:** Constructor NewSLAService

#### **TestGetSLAScreening** {#testgetslascreening .unnumbered}

> **Function:** TestGetSLAScreening
>
> **Purpose:** Menguji pengambilan konfigurasi SLA untuk screening.
>
> **Input:** Tidak ada parameter
>
> **Process:**

- Query SLA config untuk screening phase

- Get days threshold

- Return config

> **Output:**

- entities.SLAConfig untuk screening

- Error jika config tidak ditemukan

> **Test Cases:**

- Success: Get screening SLA config

- Error: Config tidak ada (default value)

- Error: Query gagal

> **Utils & Types:**

- entities.SLAConfig

> **Tests For:** Handler GetSLAScreening di SLA handler

#### **TestGetSLAFinal** {#testgetslafinal .unnumbered}

> **Function:** TestGetSLAFinal
>
> **Purpose:** Menguji pengambilan konfigurasi SLA untuk final decision.
>
> **Input:** Tidak ada parameter
>
> **Process:**

- Query SLA config untuk final phase

- Get days threshold

- Return config

> **Output:**

- entities.SLAConfig untuk final

- Error jika config tidak ditemukan

> **Test Cases:**

- Success: Get final SLA config

- Error: Config tidak ada (default value)

- Error: Query gagal

> **Utils & Types:**

- entities.SLAConfig

> **Tests For:** Handler GetSLAFinal di SLA handler

#### **TestUpdateSLAScreening** {#testupdateslascreening .unnumbered}

> **Function:** TestUpdateSLAScreening
>
> **Purpose:** Menguji update konfigurasi SLA screening.
>
> **Input:**

- dto.UpdateSLARequest berisi days threshold baru

> **Process:**

- Validate days (harus \> 0)

- Update SLA config untuk screening

- Log changes

- Return updated config

> **Output:**

- Updated entities.SLAConfig

- Error jika validasi gagal

> **Test Cases:**

- Success: Update screening SLA

- Error: Days invalid (≤ 0)

- Error: Update gagal

> **Utils & Types:**

- entities.SLAConfig

- dto.UpdateSLARequest

- Validation helper

> **Tests For:** Handler UpdateSLAScreening di SLA handler

#### **TestUpdateSLAFinal** {#testupdateslafinal .unnumbered}

> **Function:** TestUpdateSLAFinal
>
> **Purpose:** Menguji update konfigurasi SLA final.
>
> **Input:**

- dto.UpdateSLARequest berisi days threshold baru

> **Process:**

- Validate days

- Update SLA config untuk final

- Log changes

- Return updated config

> **Output:**

- Updated entities.SLAConfig

- Error jika validasi gagal

> **Test Cases:**

- Success: Update final SLA

- Error: Days invalid

- Error: Update gagal

> **Utils & Types:**

- entities.SLAConfig

- dto.UpdateSLARequest

> **Tests For:** Handler UpdateSLAFinal di SLA handler

#### **TestExportApplications** {#testexportapplications .unnumbered}

> **Function:** TestExportApplications
>
> **Purpose:** Menguji export data aplikasi ke PDF/Excel.
>
> **Input:**

- dto.ExportRequest berisi format (pdf/excel), filters

> **Process:**

- Query applications dengan filters

- Generate file (PDF atau Excel)

- Return file bytes atau URL

- Support berbagai format export

> **Output:**

- File bytes atau download URL

- Error jika export gagal

> **Test Cases:**

- Success: Export to PDF

- Success: Export to Excel

- Success: Export dengan filter

- Error: Format tidak didukung

- Error: Generate file gagal

> **Utils & Types:**

- dto.ExportRequest

- PDF generator library

- Excel generator library

- entities.Application

> **Tests For:** Handler ExportApplications di SLA handler

#### **TestExportPrograms** {#testexportprograms .unnumbered}

> **Function:** TestExportPrograms
>
> **Purpose:** Menguji export data program ke PDF/Excel.
>
> **Input:**

- dto.ExportRequest berisi format dan filters

> **Process:**

- Query programs dengan filters

- Generate file

- Return file untuk download

> **Output:**

- File bytes atau download URL

- Error jika export gagal

> **Test Cases:**

- Success: Export programs to PDF

- Success: Export programs to Excel

- Success: Export dengan filter

- Error: Generate gagal

> **Utils & Types:**

- dto.ExportRequest

- PDF/Excel generator

- entities.Program

> **Tests For:** Handler ExportPrograms di SLA handler

### News Service Tests

#### **TestNewNewsService** {#testnewnewsservice .unnumbered}

> **Function:** TestNewNewsService
>
> **Purpose:** Memverifikasi constructor service news.
>
> **Input:**

- Mock NewsRepository

- Mock MinIO client

> **Process:**

- Memanggil NewNewsService() dengan dependencies

- Memverifikasi instance tidak nil

> **Output:**

- Instance NewsService yang valid

> **Dependencies:**

- repository.NewsRepository (mock)

- minio.Client (mock)

> **Tests For:** Constructor NewNewsService

#### **TestGetAllNews** {#testgetallnews .unnumbered}

> **Function:** TestGetAllNews
>
> **Purpose:** Menguji pengambilan semua berita dengan filter.
>
> **Input:**

- Query params: page, limit, status, category, search

> **Process:**

- Build query dengan filters

- Apply pagination

- Include image URLs dari MinIO

- Return paginated result

> **Output:**

- dto.PaginatedNewsResponse

- Error jika query gagal

> **Test Cases:**

- Success: Get all news

- Success: Filter by status (published/draft)

- Success: Filter by category

- Success: Search by title/content

- Error: Query gagal

> **Utils & Types:**

- entities.News

- dto.PaginatedNewsResponse

- MinIO untuk image URLs

> **Tests For:** Handler GetAllNews di news handler

#### **TestGetNewsByID** {#testgetnewsbyid .unnumbered}

> **Function:** TestGetNewsByID
>
> **Purpose:** Menguji pengambilan detail berita berdasarkan ID.
>
> **Input:**

- News ID

> **Process:**

- Query news dari database

- Get image URL dari MinIO

- Increment view count

- Return news detail

> **Output:**

- entities.News dengan detail lengkap

- Error jika tidak ditemukan

> **Test Cases:**

- Success: News ditemukan

- Success: View count incremented

- Error: News tidak ditemukan

- Error: Query gagal

> **Utils & Types:**

- entities.News

- MinIO untuk image URL

> **Tests For:** Handler GetNewsByID di news handler

#### **TestCreateNews** {#testcreatenews .unnumbered}

> **Function:** TestCreateNews
>
> **Purpose:** Menguji pembuatan berita baru.
>
> **Input:**

- dto.CreateNewsRequest berisi title, content, category, dll

- Image file

> **Process:**

- Validate input

- Generate slug dari title

- Upload image ke MinIO

- Insert news ke database

- Return created news

> **Output:**

- entities.News yang baru dibuat

- Error jika validasi gagal

> **Test Cases:**

- Success: Create news dengan image

- Success: Generate unique slug

- Error: Title kosong

- Error: Upload image gagal

- Error: Insert database gagal

> **Utils & Types:**

- entities.News

- dto.CreateNewsRequest

- Slug generator helper

- MinIO untuk upload

> **Tests For:** Handler CreateNews di news handler

#### **TestUpdateNews** {#testupdatenews .unnumbered}

> **Function:** TestUpdateNews
>
> **Purpose:** Menguji update berita.
>
> **Input:**

- News ID

- dto.UpdateNewsRequest

- Optional: image baru

> **Process:**

- Validate news exists

- Update slug jika title berubah

- Upload image baru jika ada

- Update database

- Return updated news

> **Output:**

- Updated entities.News

- Error jika tidak ditemukan

> **Test Cases:**

- Success: Update news tanpa image

- Success: Update news dengan image baru

- Success: Update slug

- Error: News tidak ditemukan

- Error: Upload gagal

> **Utils & Types:**

- entities.News

- dto.UpdateNewsRequest

- MinIO untuk upload/delete

> **Tests For:** Handler UpdateNews di news handler

#### **TestDeleteNews** {#testdeletenews .unnumbered}

> **Function:** TestDeleteNews
>
> **Purpose:** Menguji soft delete berita.
>
> **Input:**

- News ID

> **Process:**

- Set deleted_at timestamp

- Hapus image dari MinIO

- News tidak muncul di list

> **Output:**

- Success response

- Error jika tidak ditemukan

> **Test Cases:**

- Success: Soft delete berhasil

- Error: News tidak ditemukan

- Error: Hapus image gagal

- Error: Update database gagal

> **Utils & Types:**

- entities.News

- MinIO untuk delete image

- Soft delete pattern

> **Tests For:** Handler DeleteNews di news handler

#### **TestPublishNews** {#testpublishnews .unnumbered}

> **Function:** TestPublishNews
>
> **Purpose:** Menguji publish berita.
>
> **Input:**

- News ID

> **Process:**

- Set status menjadi published

- Set published_at timestamp

- Update database

- News muncul di public API

> **Output:**

- Success response

- Error jika tidak ditemukan

> **Test Cases:**

- Success: News berhasil dipublish

- Error: News tidak ditemukan

- Error: News sudah published

- Error: Update gagal

> **Utils & Types:**

- entities.News

> **Tests For:** Handler PublishNews di news handler

#### **TestUnpublishNews** {#testunpublishnews .unnumbered}

> **Function:** TestUnpublishNews
>
> **Purpose:** Menguji unpublish berita.
>
> **Input:**

- News ID

> **Process:**

- Set status menjadi draft

- Clear published_at

- Update database

- News tidak muncul di public

> **Output:**

- Success response

- Error jika tidak ditemukan

> **Test Cases:**

- Success: News berhasil di-unpublish

- Error: News tidak ditemukan

- Error: News belum published

- Error: Update gagal

> **Utils & Types:**

- entities.News

> **Tests For:** Handler UnpublishNews di news handler

### Mobile Service Tests

#### **TestNewMobileService** {#testnewmobileservice .unnumbered}

> **Function:** TestNewMobileService
>
> **Purpose:** Memverifikasi constructor service mobile.
>
> **Input:**

- Mock MobileRepository

- Mock ProgramsRepository

- Mock NotificationRepository

- Mock VaultDecryptLogRepository

- Mock ApplicationsRepository

- Mock SLARepository

- Mock MinIO client

> **Process:**

- Memanggil NewMobileService() dengan dependencies

- Memverifikasi instance tidak nil

> **Output:**

- Instance MobileService yang valid

> **Dependencies:**

- Multiple repository mocks

- MinIO client mock

> **Tests For:** Constructor NewMobileService

#### **TestGetDashboard** {#testgetdashboard .unnumbered}

> **Function:** TestGetDashboard
>
> **Purpose:** Menguji pengambilan dashboard data untuk mobile.
>
> **Input:**

- User ID dari JWT token

> **Process:**

- Get UMKM profile summary

- Get active programs count

- Get user's applications summary

- Get unread notifications count

- Aggregate all data

> **Output:**

- dto.MobileDashboardResponse berisi summary data

- Error jika user tidak ditemukan

> **Test Cases:**

- Success: Get dashboard data

- Error: User tidak authenticated

- Error: Query gagal

> **Utils & Types:**

- dto.MobileDashboardResponse

- entities.Users

- Aggregate queries

> **Tests For:** Handler GetDashboard di mobile handler

#### **TestGetTrainingPrograms** {#testgettrainingprograms .unnumbered}

> **Function:** TestGetTrainingPrograms
>
> **Purpose:** Menguji pengambilan daftar program training untuk mobile.
>
> **Input:**

- Query params: page, limit, search

> **Process:**

- Query programs dengan type=training dan status=active

- Apply pagination

- Include image URLs

- Return list

> **Output:**

- dto.PaginatedProgramsResponse untuk training

- Error jika query gagal

> **Test Cases:**

- Success: Get training programs

- Success: Search training programs

- Error: Query gagal

> **Utils & Types:**

- entities.Program

- dto.PaginatedProgramsResponse

- Program type filter

> **Tests For:** Handler GetTrainingPrograms di mobile handler

#### **TestGetCertificationPrograms** {#testgetcertificationprograms .unnumbered}

> **Function:** TestGetCertificationPrograms
>
> **Purpose:** Menguji pengambilan daftar program certification.
>
> **Input:**

- Query params: page, limit, search

> **Process:**

- Query programs dengan type=certification dan status=active

- Apply pagination

- Return list

> **Output:**

- dto.PaginatedProgramsResponse untuk certification

- Error jika query gagal

> **Test Cases:**

- Success: Get certification programs

- Success: Search certification programs

- Error: Query gagal

> **Utils & Types:**

- entities.Program

- dto.PaginatedProgramsResponse

> **Tests For:** Handler GetCertificationPrograms di mobile handler

#### **TestGetFundingPrograms** {#testgetfundingprograms .unnumbered}

> **Function:** TestGetFundingPrograms
>
> **Purpose:** Menguji pengambilan daftar program funding.
>
> **Input:**

- Query params: page, limit, search

> **Process:**

- Query programs dengan type=funding dan status=active

- Apply pagination

- Return list

> **Output:**

- dto.PaginatedProgramsResponse untuk funding

- Error jika query gagal

> **Test Cases:**

- Success: Get funding programs

- Success: Search funding programs

- Error: Query gagal

> **Utils & Types:**

- entities.Program

- dto.PaginatedProgramsResponse

> **Tests For:** Handler GetFundingPrograms di mobile handler

#### **TestGetProgramDetail** {#testgetprogramdetail .unnumbered}

> **Function:** TestGetProgramDetail
>
> **Purpose:** Menguji pengambilan detail program untuk mobile.
>
> **Input:**

- Program ID

> **Process:**

- Query program detail

- Check if user already applied

- Include image URL

- Return detail

> **Output:**

- dto.ProgramDetailResponse dengan applied status

- Error jika tidak ditemukan

> **Test Cases:**

- Success: Get program detail

- Success: Include applied status

- Error: Program tidak ditemukan

- Error: Program tidak active

> **Utils & Types:**

- entities.Program

- dto.ProgramDetailResponse

> **Tests For:** Handler GetProgramDetail di mobile handler

#### **TestGetUMKMProfile** {#testgetumkmprofile .unnumbered}

> **Function:** TestGetUMKMProfile
>
> **Purpose:** Menguji pengambilan profil UMKM.
>
> **Input:**

- User ID dari JWT token

> **Process:**

- Get user data dengan relasi lengkap

- Include province dan city

- Decrypt sensitive data jika perlu

- Return profile

> **Output:**

- dto.UMKMProfileResponse

- Error jika user tidak ditemukan

> **Test Cases:**

- Success: Get UMKM profile

- Success: Include location data

- Error: User tidak authenticated

- Error: Query gagal

> **Utils & Types:**

- entities.Users

- dto.UMKMProfileResponse

- Vault decrypt

> **Tests For:** Handler GetUMKMProfile di mobile handler

#### **TestUpdateUMKMProfile** {#testupdateumkmprofile .unnumbered}

> **Function:** TestUpdateUMKMProfile
>
> **Purpose:** Menguji update profil UMKM.
>
> **Input:**

- User ID dari JWT token

- dto.UpdateUMKMProfileRequest

> **Process:**

- Validate input

- Update user data

- Encrypt sensitive data jika ada

- Return updated profile

> **Output:**

- Updated dto.UMKMProfileResponse

- Error jika validasi gagal

> **Test Cases:**

- Success: Update profile

- Error: Validasi gagal

- Error: User tidak authenticated

- Error: Update gagal

> **Utils & Types:**

- entities.Users

- dto.UpdateUMKMProfileRequest

- Vault encrypt untuk sensitive data

> **Tests For:** Handler UpdateUMKMProfile di mobile handler

#### **TestGetUMKMDocuments** {#testgetumkmdocuments .unnumbered}

> **Function:** TestGetUMKMDocuments
>
> **Purpose:** Menguji pengambilan daftar dokumen UMKM.
>
> **Input:**

- User ID dari JWT token

> **Process:**

- Query documents milik user

- Get document URLs dari MinIO

- Group by document type

- Return list

> **Output:**

- dto.UMKMDocumentsResponse berisi list documents

- Error jika user tidak authenticated

> **Test Cases:**

- Success: Get documents list

- Success: Group by type (NIB, NPWP, etc)

- Error: User tidak authenticated

- Error: Query gagal

> **Utils & Types:**

- entities.Document

- dto.UMKMDocumentsResponse

- MinIO untuk document URLs

> **Tests For:** Handler GetUMKMDocuments di mobile handler

#### **TestUploadDocument** {#testuploaddocument .unnumbered}

> **Function:** TestUploadDocument
>
> **Purpose:** Menguji upload dokumen UMKM.
>
> **Input:**

- User ID dari JWT token

- Document type (NIB, NPWP, revenue_record, business_permit)

- File upload

> **Process:**

- Validate document type

- Validate file format (PDF)

- Upload file ke MinIO

- Save document record ke database

- Return document info

> **Output:**

- entities.Document yang baru diupload

- Error jika validasi gagal

> **Test Cases:**

- Success: Upload NIB

- Success: Upload NPWP

- Success: Replace existing document

- Error: Invalid document type

- Error: Invalid file format

- Error: Upload gagal

> **Utils & Types:**

- entities.Document

- MinIO untuk upload

- File validation helper

> **Tests For:** Handler UploadDocument di mobile handler

#### **TestCreateTrainingApplication** {#testcreatetrainingapplication .unnumbered}

> **Function:** TestCreateTrainingApplication
>
> **Purpose:** Menguji pembuatan aplikasi training.
>
> **Input:**

- User ID dari JWT token

- dto.CreateTrainingApplicationRequest berisi program_id

> **Process:**

- Validate program exists dan active

- Check user belum apply ke program ini

- Validate dokumen lengkap

- Create application record

- Encrypt sensitive data

- Create notification

- Return application

> **Output:**

- entities.Application yang baru dibuat

- Error jika validasi gagal

> **Test Cases:**

- Success: Create training application

- Error: Program tidak ditemukan

- Error: Program tidak active

- Error: User sudah apply

- Error: Dokumen tidak lengkap

- Error: Create gagal

> **Utils & Types:**

- entities.Application

- dto.CreateTrainingApplicationRequest

- Vault encrypt

- entities.Notification

> **Tests For:** Handler CreateTrainingApplication di mobile handler

#### **TestCreateCertificationApplication** {#testcreatecertificationapplication .unnumbered}

> **Function:** TestCreateCertificationApplication
>
> **Purpose:** Menguji pembuatan aplikasi certification.
>
> **Input:**

- User ID dari JWT token

- dto.CreateCertificationApplicationRequest

> **Process:**

- Validate program dan prerequisites

- Check dokumen lengkap

- Create application

- Encrypt data

- Create notification

- Return application

> **Output:**

- entities.Application yang baru dibuat

- Error jika validasi gagal

> **Test Cases:**

- Success: Create certification application

- Error: Prerequisites tidak terpenuhi

- Error: Dokumen tidak lengkap

- Error: User sudah apply

- Error: Create gagal

> **Utils & Types:**

- entities.Application

- dto.CreateCertificationApplicationRequest

- Prerequisites validation

> **Tests For:** Handler CreateCertificationApplication di mobile
> handler

#### **TestCreateFundingApplication** {#testcreatefundingapplication .unnumbered}

> **Function:** TestCreateFundingApplication
>
> **Purpose:** Menguji pembuatan aplikasi funding.
>
> **Input:**

- User ID dari JWT token

- dto.CreateFundingApplicationRequest berisi amount requested

> **Process:**

- Validate program dan funding limit

- Validate card type eligibility

- Check dokumen keuangan lengkap

- Create application

- Encrypt financial data

- Create notification

- Return application

> **Output:**

- entities.Application yang baru dibuat

- Error jika validasi gagal

> **Test Cases:**

- Success: Create funding application

- Error: Amount melebihi limit

- Error: Card type tidak eligible

- Error: Dokumen keuangan tidak lengkap

- Error: User sudah apply

- Error: Create gagal

> **Utils & Types:**

- entities.Application

- dto.CreateFundingApplicationRequest

- Financial validation

- Card type validation

> **Tests For:** Handler CreateFundingApplication di mobile handler

#### **TestGetApplicationList** {#testgetapplicationlist .unnumbered}

> **Function:** TestGetApplicationList
>
> **Purpose:** Menguji pengambilan daftar aplikasi user.
>
> **Input:**

- User ID dari JWT token

- Query params: page, limit, status, program_type

> **Process:**

- Query applications milik user

- Apply filters

- Include program info

- Calculate SLA status

- Return paginated list

> **Output:**

- dto.PaginatedApplicationsResponse

- Error jika user tidak authenticated

> **Test Cases:**

- Success: Get all user's applications

- Success: Filter by status

- Success: Filter by program_type

- Error: User tidak authenticated

- Error: Query gagal

> **Utils & Types:**

- entities.Application

- dto.PaginatedApplicationsResponse

- SLA calculation

> **Tests For:** Handler GetApplicationList di mobile handler

#### **TestGetApplicationDetail** {#testgetapplicationdetail .unnumbered}

> **Function:** TestGetApplicationDetail
>
> **Purpose:** Menguji pengambilan detail aplikasi untuk mobile.
>
> **Input:**

- Application ID

- User ID dari JWT token

> **Process:**

- Validate application belongs to user

- Get application dengan relasi lengkap

- Decrypt sensitive data

- Log decrypt activity

- Calculate SLA

- Return detail

> **Output:**

- dto.ApplicationDetailResponse dengan data lengkap

- Error jika tidak authorized

> **Test Cases:**

- Success: Get application detail

- Error: Application bukan milik user

- Error: Application tidak ditemukan

- Error: Decrypt gagal

> **Utils & Types:**

- entities.Application

- dto.ApplicationDetailResponse

- Vault decrypt

- Authorization check

> **Tests For:** Handler GetApplicationDetail di mobile handler

#### **TestReviseApplication** {#testreviseapplication .unnumbered}

> **Function:** TestReviseApplication
>
> **Purpose:** Menguji revisi aplikasi yang diminta perubahan.
>
> **Input:**

- Application ID

- User ID dari JWT token

- dto.ReviseApplicationRequest berisi updated data

> **Process:**

- Validate application status (harus revision_requested)

- Validate application belongs to user

- Update application data

- Re-encrypt sensitive data

- Change status ke pending_screening

- Create notification untuk admin

- Return updated application

> **Output:**

- Updated entities.Application

- Error jika status tidak valid

> **Test Cases:**

- Success: Revise dan resubmit

- Error: Status bukan revision_requested

- Error: Application bukan milik user

- Error: Validasi data gagal

- Error: Update gagal

> **Utils & Types:**

- entities.Application

- dto.ReviseApplicationRequest

- Vault encrypt

- Status validation

> **Tests For:** Handler ReviseApplication di mobile handler

#### **TestGetNotificationsByUMKMID** {#testgetnotificationsbyumkmid .unnumbered}

> **Function:** TestGetNotificationsByUMKMID
>
> **Purpose:** Menguji pengambilan daftar notifikasi.
>
> **Input:**

- User ID dari JWT token

- Query params: page, limit, is_read filter

> **Process:**

- Query notifications milik user

- Apply pagination

- Filter by read status jika ada

- Order by created_at DESC

- Return list

> **Output:**

- dto.PaginatedNotificationsResponse

- Error jika user tidak authenticated

> **Test Cases:**

- Success: Get all notifications

- Success: Filter unread only

- Success: Filter read only

- Error: User tidak authenticated

- Error: Query gagal

> **Utils & Types:**

- entities.Notification

- dto.PaginatedNotificationsResponse

> **Tests For:** Handler GetNotificationsByUMKMID di mobile handler

#### **TestGetUnreadCount** {#testgetunreadcount .unnumbered}

> **Function:** TestGetUnreadCount
>
> **Purpose:** Menguji pengambilan jumlah notifikasi yang belum dibaca.
>
> **Input:**

- User ID dari JWT token

> **Process:**

- Count notifications dengan is_read=false

- Return count

> **Output:**

- dto.UnreadCountResponse berisi count

- Error jika user tidak authenticated

> **Test Cases:**

- Success: Get unread count

- Error: User tidak authenticated

- Error: Query gagal

> **Utils & Types:**

- dto.UnreadCountResponse

- Count query

> **Tests For:** Handler GetUnreadCount di mobile handler

#### **TestMarkNotificationsAsRead** {#testmarknotificationsasread .unnumbered}

> **Function:** TestMarkNotificationsAsRead
>
> **Purpose:** Menguji menandai notifikasi sebagai sudah dibaca.
>
> **Input:**

- Notification ID

- User ID dari JWT token

> **Process:**

- Validate notification belongs to user

- Update is_read menjadi true

- Set read_at timestamp

- Return success

> **Output:**

- Success response

- Error jika tidak authorized

> **Test Cases:**

- Success: Mark as read

- Error: Notification bukan milik user

- Error: Notification tidak ditemukan

- Error: Update gagal

> **Utils & Types:**

- entities.Notification

- Authorization check

> **Tests For:** Handler MarkNotificationsAsRead di mobile handler

#### **TestMarkAllNotificationsAsRead** {#testmarkallnotificationsasread .unnumbered}

> **Function:** TestMarkAllNotificationsAsRead
>
> **Purpose:** Menguji menandai semua notifikasi sebagai sudah dibaca.
>
> **Input:**

- User ID dari JWT token

> **Process:**

- Update semua notifications milik user

- Set is_read=true untuk semua

- Set read_at timestamp

- Return count updated

> **Output:**

- Success response dengan count

- Error jika user tidak authenticated

> **Test Cases:**

- Success: Mark all as read

- Error: User tidak authenticated

- Error: Update gagal

> **Utils & Types:**

- entities.Notification

- Batch update

> **Tests For:** Handler MarkAllNotificationsAsRead di mobile handler

#### **TestGetPublishedNews** {#testgetpublishednews .unnumbered}

> **Function:** TestGetPublishedNews
>
> **Purpose:** Menguji pengambilan berita yang sudah dipublish untuk
> mobile.
>
> **Input:**

- Query params: page, limit, category, search

> **Process:**

- Query news dengan status=published

- Apply filters

- Apply pagination

- Include image URLs

- Order by published_at DESC

- Return list

> **Output:**

- dto.PaginatedNewsResponse untuk published news

- Error jika query gagal

> **Test Cases:**

- Success: Get published news

- Success: Filter by category

- Success: Search news

- Error: Query gagal

> **Utils & Types:**

- entities.News

- dto.PaginatedNewsResponse

- MinIO untuk images

> **Tests For:** Handler GetPublishedNews di mobile handler

#### **TestGetNewsDetailBySlug** {#testgetnewsdetailbyslug .unnumbered}

> **Function:** TestGetNewsDetailBySlug
>
> **Purpose:** Menguji pengambilan detail berita berdasarkan slug.
>
> **Input:**

- News slug (string)

> **Process:**

- Query news berdasarkan slug

- Validate status=published

- Get image URL

- Increment view count

- Return detail

> **Output:**

- entities.News dengan detail lengkap

- Error jika tidak ditemukan atau unpublished

> **Test Cases:**

- Success: Get news by slug

- Success: View count incremented

- Error: News tidak ditemukan

- Error: News belum published

> **Utils & Types:**

- entities.News

- Slug lookup

- View counter

> **Tests For:** Handler GetNewsDetailBySlug di mobile handler

### Vault Decrypt Log Service Tests

#### **TestNewVaultDecryptLogService** {#testnewvaultdecryptlogservice .unnumbered}

> **Function:** TestNewVaultDecryptLogService
>
> **Purpose:** Memverifikasi constructor service vault decrypt log.
>
> **Input:**

- Mock VaultDecryptLogRepository

> **Process:**

- Memanggil NewVaultDecryptLogService() dengan repository

- Memverifikasi instance tidak nil

> **Output:**

- Instance VaultDecryptLogService yang valid

> **Dependencies:**

- repository.VaultDecryptLogRepository (mock)

> **Tests For:** Constructor NewVaultDecryptLogService

#### **TestGetLogs** {#testgetlogs .unnumbered}

> **Function:** TestGetLogs
>
> **Purpose:** Menguji pengambilan semua log dekripsi dengan pagination.
>
> **Input:**

- Query params: page, limit, date_range

> **Process:**

- Query vault decrypt logs

- Apply pagination

- Include user dan UMKM info

- Order by created_at DESC

- Return paginated list

> **Output:**

- dto.PaginatedVaultDecryptLogsResponse

- Error jika query gagal

> **Test Cases:**

- Success: Get all logs

- Success: Filter by date range

- Success: Pagination

- Error: Query gagal

> **Utils & Types:**

- entities.VaultDecryptLog

- dto.PaginatedVaultDecryptLogsResponse

- Date range filter

> **Tests For:** Handler GetLogs di vault decrypt log handler

#### **TestGetLogsByUserID** {#testgetlogsbyuserid .unnumbered}

> **Function:** TestGetLogsByUserID
>
> **Purpose:** Menguji pengambilan log dekripsi berdasarkan user ID.
>
> **Input:**

- User ID

- Query params: page, limit

> **Process:**

- Query logs where decrypted_by=user_id

- Apply pagination

- Include UMKM info

- Return list

> **Output:**

- dto.PaginatedVaultDecryptLogsResponse untuk user

- Error jika user tidak ditemukan

> **Test Cases:**

- Success: Get logs by user

- Success: Pagination

- Error: User tidak ditemukan

- Error: Query gagal

> **Utils & Types:**

- entities.VaultDecryptLog

- dto.PaginatedVaultDecryptLogsResponse

> **Tests For:** Handler GetLogsByUserID di vault decrypt log handler

#### **TestGetLogsByUMKMID** {#testgetlogsbyumkmid .unnumbered}

> **Function:** TestGetLogsByUMKMID
>
> **Purpose:** Menguji pengambilan log dekripsi berdasarkan UMKM ID.
>
> **Input:**

- UMKM ID

- Query params: page, limit

> **Process:**

- Query logs where umkm_id=umkm_id

- Apply pagination

- Include user info (who decrypted)

- Return list

> **Output:**

- dto.PaginatedVaultDecryptLogsResponse untuk UMKM

- Error jika UMKM tidak ditemukan

> **Test Cases:**

- Success: Get logs by UMKM

- Success: Show who accessed data

- Success: Pagination

- Error: UMKM tidak ditemukan

- Error: Query gagal

> **Utils & Types:**

- entities.VaultDecryptLog

- dto.PaginatedVaultDecryptLogsResponse

- Audit trail data

> **Tests For:** Handler GetLogsByUMKMID di vault decrypt log handler

## Model

Model digunakan untuk representasi struktur data dalam database
menggunakan GORM ORM. Setiap model merepresentasikan satu tabel dalam
database.

### Base Model

> Model dasar yang digunakan oleh semua model lainnya untuk menyimpan
> metadata waktu.
>
> **Fields:**

- CreatedAt (time.Time) - Timestamp saat record dibuat

- UpdatedAt (time.Time) - Timestamp saat record terakhir diupdate

- DeletedAt (gorm.DeletedAt) - Timestamp soft delete dengan index

### User Model

> Model untuk menyimpan data pengguna sistem (admin, pelaku usaha, dll).
>
> **Fields:**

- ID (int) - Primary key

- Name (string) - Nama lengkap user

- Email (string) - Email user (unique)

- Password (string) - Password terenkripsi

- RoleID (int) - Foreign key ke tabel roles

- IsActive (bool) - Status aktif user (default: true)

- LastLoginAt (time.Time) - Timestamp login terakhir

- Base - Embedded base fields

> **Relations:**

- Roles (Role) - Relasi ke model Role (belongsTo)

### Role Model

> Model untuk menyimpan peran/role dalam sistem.
>
> **Fields:**

- ID (int) - Primary key

- Name (string) - Nama role

- Description (string) - Deskripsi role

- Base - Embedded base fields

### Permission Model

> Model untuk menyimpan izin akses dalam sistem.
>
> **Fields:**

- ID (int) - Primary key

- ParentID (\*int) - Foreign key ke parent permission (nullable)

- Name (string) - Nama permission

- Code (string) - Kode permission

- Description (string) - Deskripsi permission

- Base - Embedded base fields

### RolePermission Model

> Model pivot untuk relasi many-to-many antara Role dan Permission.
>
> **Fields:**

- RoleID (int) - Primary key, foreign key ke roles

- PermissionID (int) - Primary key, foreign key ke permissions

- Base - Embedded base fields

### Province Model

> Model untuk menyimpan data provinsi di Indonesia.
>
> **Fields:**

- ID (int) - Primary key

- Name (string) - Nama provinsi

- Base - Embedded base fields

> **Relations:**

- UMKMs (\[\]UMKM) - Relasi ke model UMKM (hasMany)

- Cities (\[\]City) - Relasi ke model City (hasMany)

### City Model

> Model untuk menyimpan data kota/kabupaten di Indonesia.
>
> **Fields:**

- ID (int) - Primary key

- Name (string) - Nama kota/kabupaten

- ProvinceID (int) - Foreign key ke provinces

- Base - Embedded base fields

> **Relations:**

- UMKMs (\[\]UMKM) - Relasi ke model UMKM (hasMany)

- Province (Province) - Relasi ke model Province (belongsTo)

### UMKM Model

> Model untuk menyimpan data profil UMKM (Usaha Mikro Kecil Menengah).
>
> **Fields:**

- ID (int) - Primary key

- UserID (int) - Foreign key ke users

- BusinessName (string) - Nama usaha

- NIK (string) - NIK terenkripsi (text)

- Gender (string) - Jenis kelamin (enum: male, female, other)

- BirthDate (time.Time) - Tanggal lahir

- Phone (string) - Nomor telepon

- Address (string) - Alamat lengkap

- ProvinceID (int) - Foreign key ke provinces

- CityID (int) - Foreign key ke cities

- District (string) - Kecamatan

- Subdistrict (string) - Kelurahan

- PostalCode (string) - Kode pos

- NIB (string) - Nomor Induk Berusaha

- NPWP (string) - Nomor Pokok Wajib Pajak

- RevenueRecord (string) - URL dokumen catatan omzet

- BusinessPermit (string) - URL dokumen izin usaha

- KartuType (string) - Tipe kartu (enum: produktif, afirmatif)

- KartuNumber (string) - Nomor kartu terenkripsi

- Photo (string) - URL foto profil

- QRCode (string) - URL QR code

- Base - Embedded base fields

> **Relations:**

- User (User) - Relasi ke model User (belongsTo)

- Province (Province) - Relasi ke model Province (belongsTo)

- City (City) - Relasi ke model City (belongsTo)

- Applications (\[\]Application) - Relasi ke model Application (hasMany)

### Program Model

> Model untuk menyimpan data program bantuan (training, certification,
> funding).
>
> **Fields:**

- ID (int) - Primary key

- Title (string) - Judul program

- Description (string) - Deskripsi program

- Banner (string) - URL banner program

- Provider (string) - Penyedia program

- ProviderLogo (string) - URL logo penyedia

- Type (string) - Tipe program (enum: training, certification, funding)

- TrainingType (\*string) - Tipe pelatihan (enum: online, offline,
  > hybrid) - nullable

- Batch (\*int) - Nomor batch - nullable

- BatchStartDate (\*string) - Tanggal mulai batch - nullable

- BatchEndDate (\*string) - Tanggal selesai batch - nullable

- Location (\*string) - Lokasi program - nullable

- MinAmount (\*float64) - Jumlah minimal pendanaan - nullable

- MaxAmount (\*float64) - Jumlah maksimal pendanaan - nullable

- InterestRate (\*float64) - Persentase bunga - nullable

- MaxTenureMonths (\*int) - Maksimal tenor dalam bulan - nullable

- ApplicationDeadline (string) - Batas waktu pendaftaran

- IsActive (bool) - Status aktif program (default: true)

- CreatedBy (int) - Foreign key ke users (pembuat program)

- Base - Embedded base fields

> **Relations:**

- Users (User) - Relasi ke model User (belongsTo)

### ProgramBenefit Model

> Model untuk menyimpan manfaat/benefit dari program.
>
> **Fields:**

- ID (int) - Primary key

- ProgramID (int) - Foreign key ke programs

- Name (string) - Nama benefit

- Base - Embedded base fields

### ProgramRequirement Model

> Model untuk menyimpan persyaratan program.
>
> **Fields:**

- ID (int) - Primary key

- ProgramID (int) - Foreign key ke programs

- Name (string) - Nama persyaratan

- Base - Embedded base fields

### Application Model {#application-model .unnumbered}

> Model untuk menyimpan data pengajuan aplikasi program.
>
> **Fields:**

- ID (int) - Primary key

- UMKMID (int) - Foreign key ke umkms

- ProgramID (int) - Foreign key ke programs

- Type (string) - Tipe aplikasi (enum: training, certification, funding)

- Status (string) - Status aplikasi (enum: screening, revised, final,
  > approved, rejected)

- SubmittedAt (time.Time) - Waktu pengajuan (default: NOW())

- ExpiredAt (time.Time) - Waktu kadaluarsa

- Base - Embedded base fields

> **Relations:**

- Documents (\[\]ApplicationDocument) - Relasi ke model
  > ApplicationDocument (hasMany)

- Histories (\[\]ApplicationHistory) - Relasi ke model
  > ApplicationHistory (hasMany)

- Program (Program) - Relasi ke model Program (belongsTo)

- UMKM (UMKM) - Relasi ke model UMKM (belongsTo)

- TrainingApplication (\*TrainingApplication) - Relasi ke model
  > TrainingApplication (hasOne) - nullable

- CertificationApplication (\*CertificationApplication) - Relasi ke
  > model CertificationApplication (hasOne) - nullable

- FundingApplication (\*FundingApplication) - Relasi ke model
  > FundingApplication (hasOne) - nullable

### ApplicationDocument Model

> Model untuk menyimpan dokumen yang dilampirkan pada aplikasi.
>
> **Fields:**

- ID (int) - Primary key

- ApplicationID (int) - Foreign key ke applications

- Type (string) - Tipe dokumen (enum: ktp, nib, npwp, proposal,
  > portfolio, rekening, other)

- File (string) - URL file dokumen

- Base - Embedded base fields

### ApplicationHistory Model {#applicationhistory-model .unnumbered}

> Model untuk menyimpan riwayat perubahan status aplikasi.
>
> **Fields:**

- ID (int) - Primary key

- ApplicationID (int) - Foreign key ke applications

- Status (string) - Action yang dilakukan (enum: submit, revise,
  > approve_by_admin_screening, reject_by_admin_screening,
  > approve_by_admin_vendor, reject_by_admin_vendor)

- Notes (string) - Catatan terkait action

- ActionedAt (time.Time) - Waktu action dilakukan (default: NOW())

- ActionedBy (\*int) - Foreign key ke users (yang melakukan action) -
  > nullable

- Base - Embedded base fields

> **Relations:**

- User (User) - Relasi ke model User (belongsTo)

### TrainingApplication Model

> Model untuk menyimpan data spesifik aplikasi training.
>
> **Fields:**

- ID (int) - Primary key

- ApplicationID (int) - Foreign key ke applications

- Motivation (string) - Motivasi mengikuti training

- BusinessExperience (string) - Pengalaman bisnis

- LearningObjectives (string) - Tujuan pembelajaran

- AvailabilityNotes (string) - Catatan ketersediaan waktu

- Base - Embedded base fields

> **Relations:**

- Application (Application) - Relasi ke model Application (belongsTo)

### CertificationApplication Model

> Model untuk menyimpan data spesifik aplikasi certification.
>
> **Fields:**

- ID (int) - Primary key

- ApplicationID (int) - Foreign key ke applications

- BusinessSector (string) - Sektor bisnis

- ProductOrService (string) - Produk atau layanan yang akan
  > disertifikasi

- BusinessDescription (string) - Deskripsi usaha

- YearsOperating (\*int) - Lama usaha beroperasi (tahun) - nullable

- CurrentStandards (string) - Standar yang sudah diterapkan

- CertificationGoals (string) - Tujuan sertifikasi

- Base - Embedded base fields

> **Relations:**

- Application (Application) - Relasi ke model Application (belongsTo)

### FundingApplication Model

> Model untuk menyimpan data spesifik aplikasi funding.
>
> **Fields:**

- ID (int) - Primary key

- ApplicationID (int) - Foreign key ke applications

- BusinessSector (string) - Sektor bisnis

- BusinessDescription (string) - Deskripsi usaha

- YearsOperating (\*int) - Lama usaha beroperasi (tahun) - nullable

- RequestedAmount (float64) - Jumlah dana yang diminta

- FundPurpose (string) - Tujuan penggunaan dana

- BusinessPlan (string) - Rencana bisnis

- RevenueProjection (\*float64) - Proyeksi pendapatan - nullable

- MonthlyRevenue (\*float64) - Pendapatan bulanan - nullable

- RequestedTenureMonths (int) - Tenor yang diminta (bulan)

- CollateralDescription (string) - Deskripsi jaminan

- Base - Embedded base fields

> **Relations:**

- Application (Application) - Relasi ke model Application (belongsTo)

### SLA Model

> Model untuk menyimpan konfigurasi Service Level Agreement.
>
> **Fields:**

- ID (int) - Primary key

- Status (string) - Status SLA (unique, values: screening, final)

- MaxDays (int) - Maksimal hari untuk proses

- Description (string) - Deskripsi SLA

- Base - Embedded base fields

### Notification Model

> Model untuk menyimpan notifikasi kepada UMKM.
>
> **Fields:**

- ID (int) - Primary key

- UMKMID (int) - Foreign key ke umkms

- ApplicationID (\*int) - Foreign key ke applications - nullable

- Type (string) - Tipe notifikasi (enum: application_submitted,
  > screening_approved, screening_rejected, screening_revised,
  > final_approved, final_rejected, program_reminder, document_required,
  > general_info)

- Title (string) - Judul notifikasi

- Message (string) - Isi pesan notifikasi

- IsRead (bool) - Status sudah dibaca (default: false)

- ReadAt (\*time.Time) - Waktu dibaca - nullable

- Metadata (string) - Data tambahan dalam format JSON

- Base - Embedded base fields

> **Relations:**

- UMKM (UMKM) - Relasi ke model UMKM (belongsTo)

- Application (\*Application) - Relasi ke model Application
  > (belongsTo) - nullable

### News Model

> Model untuk menyimpan artikel berita.
>
> **Fields:**

- ID (int) - Primary key

- Title (string) - Judul berita

- Slug (string) - URL-friendly identifier (unique)

- Excerpt (string) - Ringkasan berita

- Content (string) - Isi lengkap berita

- Thumbnail (string) - URL thumbnail

- Category (string) - Kategori berita (enum: announcement,
  > success_story, event, article)

- AuthorID (int) - Foreign key ke users

- IsPublished (bool) - Status publish (default: false)

- PublishedAt (\*time.Time) - Waktu publish - nullable

- ViewsCount (int) - Jumlah views (default: 0)

- Base - Embedded base fields

> **Relations:**

- Author (User) - Relasi ke model User (belongsTo)

- Tags (\[\]NewsTag) - Relasi ke model NewsTag (hasMany)

### NewsTag Model

> Model untuk menyimpan tag berita.
>
> **Fields:**

- ID (int) - Primary key

- NewsID (int) - Foreign key ke news

- TagName (string) - Nama tag

- CreatedAt (time.Time) - Timestamp dibuat (default: NOW())

### OTP Model

> Model untuk menyimpan data One-Time Password untuk verifikasi.
>
> **Fields:**

- PhoneNumber (string) - Nomor telepon

- Email (string) - Email

- OTPCode (string) - Kode OTP

- TempToken (\*string) - Token temporary untuk verifikasi - nullable

- Status (string) - Status OTP (enum: active, used)

- ExpiresAt (time.Time) - Waktu kadaluarsa

### VaultDecryptLog Model

> Model untuk menyimpan log aktivitas dekripsi data sensitif.
>
> **Fields:**

- ID (int64) - Primary key

- UserID (int) - Foreign key ke users (yang melakukan dekripsi)

- UMKMID (\*int) - Foreign key ke umkms - nullable

- FieldName (string) - Nama field yang didekripsi

- TableName (string) - Nama tabel sumber data

- RecordID (int) - ID record yang didekripsi

- Purpose (string) - Tujuan dekripsi (enum: profile_view,
  > application_review, application_creation, profile_update,
  > admin_verification, report_generation, compliance_audit,
  > system_process)

- IPAddress (string) - IP address yang melakukan request

- UserAgent (string) - User agent browser/app

- RequestID (string) - ID request untuk tracking

- Success (bool) - Status keberhasilan (default: true)

- ErrorMessage (string) - Pesan error jika gagal

- DecryptedAt (time.Time) - Waktu dekripsi (default: NOW())

## DTO

DTO digunakan untuk transfer data antara layer aplikasi, terutama untuk
request dan response API.

### Users DTO

> DTO untuk operasi terkait user.
>
> **Fields:**

- ID (int) - ID user

- Name (string) - Nama lengkap

- Email (string) - Email (dengan validasi email)

- Password (string) - Password (minimal 8 karakter)

- ConfirmPassword (string) - Konfirmasi password (minimal 8 karakter)

- RoleID (\*int) - ID role - nullable

- RoleName (string) - Nama role

- LastLoginAt (string) - Timestamp login terakhir

- Permissions (\[\]string) - Daftar permission codes

- CreatedAt (string) - Timestamp dibuat

- UpdatedAt (string) - Timestamp diupdate

- IsActive (bool) - Status aktif

### OTP DTO

> DTO untuk verifikasi OTP.
>
> **Fields:**

- Email (string) - Email user

- OTP (string) - Kode OTP

### Permissions DTO

> DTO untuk data permission.
>
> **Fields:**

- ID (int) - ID permission

- Name (string) - Nama permission

- Code (string) - Kode permission

- Description (string) - Deskripsi permission

### RolePermissions DTO

> DTO untuk update permissions pada role.
>
> **Fields:**

- RoleID (int) - ID role (required)

- Permissions (\[\]string) - Array kode permissions (required, setiap
  > element \> 0)

### RolePermissionsResponse DTO

> DTO untuk response daftar role dan permissions-nya.
>
> **Fields:**

- RoleID (int) - ID role

- RoleName (string) - Nama role

- Permissions (json.RawMessage) - Data permissions dalam format JSON

###  {#section-1 .unnumbered}

### Province DTO

> DTO untuk data provinsi.
>
> **Fields:**

- ID (int) - ID provinsi

- Name (string) - Nama provinsi

###  {#section-2 .unnumbered}

### City DTO

> DTO untuk data kota/kabupaten.
>
> **Fields:**

- ID (int) - ID kota

- Name (string) - Nama kota

- ProvinceID (int) - ID provinsi

###  {#section-3 .unnumbered}

### User DTO

> DTO untuk data user sederhana.
>
> **Fields:**

- ID (int) - ID user

- Name (string) - Nama user

- Email (string) - Email user

- Address (string) - Alamat user

###  {#section-4 .unnumbered}

### RegisterMobile DTO

> DTO untuk registrasi mobile.
>
> **Fields:**

- Email (string) - Email

- Phone (string) - Nomor telepon

- OTPCode (string) - Kode OTP

###  {#section-5 .unnumbered}

### ResetPasswordMobile DTO

> DTO untuk reset password mobile.
>
> **Fields:**

- Password (string) - Password baru

- ConfirmPassword (string) - Konfirmasi password baru

### UMKMMobile DTO

> DTO lengkap untuk data UMKM mobile.
>
> **Fields:**

- ID (int) - ID UMKM

- UserID (int) - ID user

- Fullname (string) - Nama lengkap

- BusinessName (string) - Nama usaha

- NIK (string) - NIK

- Gender (string) - Jenis kelamin

- BirthDate (string) - Tanggal lahir (format: YYYY-MM-DD)

- Password (string) - Password

- Phone (string) - Nomor telepon

- Email (string) - Email

- Address (string) - Alamat

- ProvinceID (int) - ID provinsi

- CityID (int) - ID kota

- District (string) - Kecamatan

- Subdistrict (string) - Kelurahan

- PostalCode (string) - Kode pos

- NIB (string) - Nomor Induk Berusaha

- NPWP (string) - NPWP

- KartuType (string) - Tipe kartu

- KartuNumber (string) - Nomor kartu

### UMKMWeb DTO

> DTO untuk data UMKM di web dashboard.
>
> **Fields:**

- ID (int) - ID UMKM

- UserID (int) - ID user

- BusinessName (string) - Nama usaha

- NIK (string) - NIK

- Gender (string) - Jenis kelamin

- BirthDate (string) - Tanggal lahir

- Phone (string) - Nomor telepon

- Address (string) - Alamat

- ProvinceID (int) - ID provinsi

- CityID (int) - ID kota

- District (string) - Kecamatan

- Subdistrict (string) - Kelurahan

- PostalCode (string) - Kode pos

- NIB (string) - NIB

- NPWP (string) - NPWP

- KartuType (string) - Tipe kartu

- KartuNumber (string) - Nomor kartu

- User (User) - Data user

- Province (Province) - Data provinsi

- City (City) - Data kota

###  {#section-6 .unnumbered}

### MetaCityAndProvince DTO

> DTO untuk data master provinsi dan kota.
>
> **Fields:**

- Provinces (\[\]Province) - Daftar provinsi

- Cities (\[\]City) - Daftar kota

###  {#section-7 .unnumbered}

### Programs DTO

> DTO untuk data program.
>
> **Fields:**

- ID (int) - ID program

- Title (string) - Judul program (required)

- Description (string) - Deskripsi program

- Banner (string) - URL banner

- Provider (string) - Penyedia program

- ProviderLogo (string) - URL logo penyedia

- Type (string) - Tipe program (required, values: training,
  > certification, funding)

- TrainingType (\*string) - Tipe training (values: online, offline,
  > hybrid) - nullable

- Batch (\*int) - Nomor batch - nullable

- BatchStartDate (\*string) - Tanggal mulai batch - nullable

- BatchEndDate (\*string) - Tanggal selesai batch - nullable

- Location (\*string) - Lokasi program - nullable

- MinAmount (\*float64) - Jumlah minimal pendanaan - nullable

- MaxAmount (\*float64) - Jumlah maksimal pendanaan - nullable

- InterestRate (\*float64) - Persentase bunga - nullable

- MaxTenureMonths (\*int) - Maksimal tenor (bulan) - nullable

- ApplicationDeadline (string) - Deadline pendaftaran (required)

- IsActive (bool) - Status aktif

- CreatedBy (int) - ID pembuat

- CreatedByName (string) - Nama pembuat

- CreatedAt (string) - Timestamp dibuat

- UpdatedAt (string) - Timestamp diupdate

- Benefits (\[\]string) - Daftar benefit

- Requirements (\[\]string) - Daftar persyaratan

### Applications DTO

> DTO untuk data aplikasi/pengajuan program.
>
> **Fields:**

- ID (int) - ID aplikasi

- UMKMID (int) - ID UMKM (required)

- ProgramID (int) - ID program (required)

- Type (string) - Tipe aplikasi (required, values: training,
  > certification, funding)

- Status (string) - Status aplikasi

- SubmittedAt (string) - Waktu submit

- ExpiredAt (string) - Waktu kadaluarsa

- CreatedAt (string) - Timestamp dibuat

- UpdatedAt (string) - Timestamp diupdate

- Documents (\[\]ApplicationDocuments) - Daftar dokumen

- Histories (\[\]ApplicationHistories) - Riwayat perubahan

- Program (\*Programs) - Data program

- UMKM (\*UMKMWeb) - Data UMKM

- TrainingData (\*TrainingApplicationData) - Data spesifik training -
  > nullable

- CertificationData (\*CertificationApplicationData) - Data spesifik
  > certification - nullable

- FundingData (\*FundingApplicationData) - Data spesifik funding -
  > nullable

###  {#section-8 .unnumbered}

### ApplicationDocuments DTO

> DTO untuk dokumen aplikasi.
>
> **Fields:**

- ID (int) - ID dokumen

- ApplicationID (int) - ID aplikasi

- Type (string) - Tipe dokumen (required, values: ktp, nib, npwp,
  > proposal, portfolio, rekening, other)

- File (string) - URL file (required)

- CreatedAt (string) - Timestamp dibuat

- UpdatedAt (string) - Timestamp diupdate

###  {#section-9 .unnumbered}

### ApplicationHistories DTO

> DTO untuk riwayat aplikasi.
>
> **Fields:**

- ID (int) - ID history

- ApplicationID (int) - ID aplikasi

- Status (string) - Status action (required)

- Notes (string) - Catatan

- ActionedAt (string) - Waktu action

- ActionedBy (\*int) - ID yang melakukan action - nullable

- ActionedByName (string) - Nama yang melakukan action

- CreatedAt (string) - Timestamp dibuat

- UpdatedAt (string) - Timestamp diupdate

###  {#section-10 .unnumbered}

### ApplicationDecision DTO

> DTO untuk keputusan terhadap aplikasi.
>
> **Fields:**

- ApplicationID (int) - ID aplikasi (required)

- Action (string) - Action yang dilakukan (required, values: approve,
  > reject, revise)

- Notes (string) - Catatan keputusan

###  {#section-11 .unnumbered}

### TrainingApplicationData DTO

> DTO untuk data spesifik aplikasi training.
>
> **Fields:**

- Motivation (string) - Motivasi mengikuti training

- BusinessExperience (string) - Pengalaman bisnis

- LearningObjectives (string) - Tujuan pembelajaran

- AvailabilityNotes (string) - Catatan ketersediaan

###  {#section-12 .unnumbered}

### CertificationApplicationData DTO

> DTO untuk data spesifik aplikasi certification.
>
> **Fields:**

- BusinessSector (string) - Sektor bisnis

- ProductOrService (string) - Produk/layanan yang disertifikasi

- BusinessDescription (string) - Deskripsi usaha

- YearsOperating (\*int) - Lama operasi (tahun) - nullable

- CurrentStandards (string) - Standar yang diterapkan

- CertificationGoals (string) - Tujuan sertifikasi

###  {#section-13 .unnumbered}

### FundingApplicationData DTO

> DTO untuk data spesifik aplikasi funding.
>
> **Fields:**

- BusinessSector (string) - Sektor bisnis

- BusinessDescription (string) - Deskripsi usaha

- YearsOperating (\*int) - Lama operasi (tahun) - nullable

- RequestedAmount (float64) - Jumlah dana diminta

- FundPurpose (string) - Tujuan dana

- BusinessPlan (string) - Rencana bisnis

- RevenueProjection (\*float64) - Proyeksi pendapatan - nullable

- MonthlyRevenue (\*float64) - Pendapatan bulanan - nullable

- RequestedTenureMonths (int) - Tenor diminta (bulan)

- CollateralDescription (string) - Deskripsi jaminan

###  {#section-14 .unnumbered}

### SLA DTO

> DTO untuk konfigurasi SLA.
>
> **Fields:**

- ID (int) - ID SLA

- Status (string) - Status SLA (required, values: screening, final)

- MaxDays (int) - Maksimal hari (required, minimal 1)

- Description (string) - Deskripsi

- CreatedAt (string) - Timestamp dibuat

- UpdatedAt (string) - Timestamp diupdate

###  {#section-15 .unnumbered}

### ExportRequest DTO

> DTO untuk request export data.
>
> **Fields:**

- FileType (string) - Tipe file export (required, values: pdf, excel)

- ApplicationType (string) - Tipe aplikasi yang diexport (required,
  > values: all, funding, training, certification)

###  {#section-16 .unnumbered}

### UMKMByCardType DTO

> DTO untuk statistik UMKM berdasarkan tipe kartu.
>
> **Fields:**

- Name (string) - Nama tipe kartu

- Count (int64) - Jumlah UMKM

###  {#section-17 .unnumbered}

### ApplicationStatusSummary DTO

> DTO untuk ringkasan status aplikasi.
>
> **Fields:**

- TotalApplications (int64) - Total aplikasi

- InProcess (int64) - Aplikasi dalam proses

- Approved (int64) - Aplikasi disetujui

- Rejected (int64) - Aplikasi ditolak

###  {#section-18 .unnumbered}

### ApplicationStatusDetail DTO

> DTO untuk detail status aplikasi.
>
> **Fields:**

- Screening (int64) - Jumlah di screening

- Revised (int64) - Jumlah revisi

- Final (int64) - Jumlah di final

- Approved (int64) - Jumlah disetujui

- Rejected (int64) - Jumlah ditolak

### ApplicationByType DTO

> DTO untuk statistik aplikasi berdasarkan tipe.
>
> **Fields:**

- Funding (int64) - Jumlah aplikasi funding

- Certification (int64) - Jumlah aplikasi certification

- Training (int64) - Jumlah aplikasi training

###  {#section-19 .unnumbered}

### UserData DTO

> DTO untuk data user dari JWT token.
>
> **Fields:**

- ID (float64) - ID user

- Name (string) - Nama user

- Email (string) - Email user

- Role (float64) - ID role

- RoleName (string) - Nama role

- BusinessName (string) - Nama usaha (untuk UMKM)

- KartuType (string) - Tipe kartu (untuk UMKM)

- Phone (string) - Nomor telepon (untuk UMKM)

###  {#section-20 .unnumbered}

### ProgramListMobile DTO

> DTO untuk list program di mobile.
>
> **Fields:**

- ID (int) - ID program

- Title (string) - Judul program

- Description (string) - Deskripsi

- Banner (string) - URL banner

- Provider (string) - Penyedia

- ProviderLogo(string) - URL logo penyedia

- Type (string) - Tipe program

- TrainingType (\*string) - Tipe training - nullable

- Batch (\*int) - Nomor batch - nullable

- BatchStartDate (\*string) - Tanggal mulai - nullable

- BatchEndDate (\*string) - Tanggal selesai - nullable

- Location (\*string) - Lokasi - nullable

- MinAmount (\*float64) - Minimal pendanaan - nullable

- MaxAmount (\*float64) - Maksimal pendanaan - nullable

- InterestRate (\*float64) - Bunga - nullable

- MaxTenureMonths (\*int) - Maksimal tenor - nullable

- ApplicationDeadline (string) - Deadline

- IsActive (bool) - Status aktif

###  {#section-21 .unnumbered}

### ProgramDetailMobile DTO

> DTO untuk detail program di mobile (extends ProgramListMobile).
>
> **Additional Fields:**

- Benefits (\[\]string) - Daftar benefit

- Requirements (\[\]string) - Daftar persyaratan

###  {#section-22 .unnumbered}

### UMKMProfile DTO

> DTO untuk profil UMKM.
>
> **Fields:**

- ID (int) - ID UMKM

- UserID (int) - ID user

- BusinessName (string) - Nama usaha

- NIK (string) - NIK

- Gender (string) - Jenis kelamin

- BirthDate (string) - Tanggal lahir

- Phone (string) - Nomor telepon

- Address (string) - Alamat

- ProvinceID (int) - ID provinsi

- CityID (int) - ID kota

- District (string) - Kecamatan

- Subdistrict (string) - Kelurahan

- PostalCode (string) - Kode pos

- NIB (string) - NIB

- NPWP (string) - NPWP

- RevenueRecord (string) - URL dokumen omzet

- BusinessPermit (string) - URL izin usaha

- KartuType (string) - Tipe kartu

- KartuNumber (string) - Nomor kartu

- Photo (string) - URL foto

- Province (Province) - Data provinsi

- City (City) - Data kota

- User (User) - Data user

###  {#section-23 .unnumbered}

### UpdateUMKMProfile DTO

> DTO untuk update profil UMKM.
>
> **Fields:**

- BusinessName (string) - Nama usaha (required)

- Gender (string) - Jenis kelamin (required, values: male, female,
  > other)

- BirthDate (string) - Tanggal lahir (required)

- Address (string) - Alamat (required)

- ProvinceID (int) - ID provinsi (required)

- CityID (int) - ID kota (required)

- District (string) - Kecamatan (required)

- PostalCode (string) - Kode pos (required)

- Name (string) - Nama lengkap (required)

- Photo (string) - Foto profil base64

###  {#section-24 .unnumbered}

### UploadDocumentRequest DTO

> DTO untuk upload dokumen UMKM.
>
> **Fields:**

- Type (string) - Tipe dokumen (required, values: nib, npwp,
  > revenue_record, business_permit)

- Document (string) - Data dokumen base64 atau URL (required)

###  {#section-25 .unnumbered}

### CreateApplicationTraining DTO

> DTO untuk membuat aplikasi training.
>
> **Fields:**

- ProgramID (int) - ID program (required)

- Motivation (string) - Motivasi (required)

- BusinessExperience (string) - Pengalaman bisnis

- LearningObjectives (string) - Tujuan pembelajaran

- AvailabilityNotes (string) - Catatan ketersediaan

- Documents (map\[string\]string) - Map tipe dokumen ke data base64/URL
  > (required)

###  {#section-26 .unnumbered}

### CreateApplicationCertification DTO

> DTO untuk membuat aplikasi certification.
>
> **Fields:**

- ProgramID (int) - ID program (required)

- BusinessSector (string) - Sektor bisnis (required)

- ProductOrService (string) - Produk/layanan (required)

- BusinessDescription (string) - Deskripsi usaha (required)

- YearsOperating (\*int) - Lama operasi - nullable

- CurrentStandards (string) - Standar diterapkan

- CertificationGoals (string) - Tujuan sertifikasi (required)

- Documents (map\[string\]string) - Map tipe dokumen ke data base64/URL
  > (required)

###  {#section-27 .unnumbered}

### CreateApplicationFunding DTO

> DTO untuk membuat aplikasi funding.
>
> **Fields:**

- ProgramID (int) - ID program (required)

- BusinessSector (string) - Sektor bisnis (required)

- BusinessDescription (string) - Deskripsi usaha (required)

- YearsOperating (\*int) - Lama operasi - nullable

- RequestedAmount (float64) - Jumlah diminta (required)

- FundPurpose (string) - Tujuan dana (required)

- BusinessPlan (string) - Rencana bisnis

- RevenueProjection (\*float64) - Proyeksi pendapatan - nullable

- MonthlyRevenue (\*float64) - Pendapatan bulanan - nullable

- RequestedTenureMonths (int) - Tenor diminta (required)

- CollateralDescription (string) - Deskripsi jaminan

- Documents (map\[string\]string) - Map tipe dokumen ke data base64/URL
  > (required)

### ApplicationListMobile DTO

> DTO untuk list aplikasi di mobile.
>
> **Fields:**

- ID (int) - ID aplikasi

- ProgramID (int) - ID program

- ProgramName (string) - Nama program

- Type (string) - Tipe aplikasi

- Status (string) - Status aplikasi

- SubmittedAt (string) - Waktu submit

- ExpiredAt (string) - Waktu kadaluarsa

###  {#section-28 .unnumbered}

### ApplicationDetailMobile DTO

> DTO untuk detail aplikasi di mobile.
>
> **Fields:**

- ID (int) - ID aplikasi

- UMKMID (int) - ID UMKM

- ProgramID (int) - ID program

- Type (string) - Tipe aplikasi

- Status (string) - Status aplikasi

- SubmittedAt (string) - Waktu submit

- ExpiredAt (string) - Waktu kadaluarsa

- Documents (\[\]ApplicationDocuments) - Daftar dokumen

- Histories (\[\]ApplicationHistories) - Riwayat perubahan

- Program (ProgramDetailMobile) - Data program

- TrainingData (\*TrainingApplicationData) - Data training - nullable

- CertificationData (\*CertificationApplicationData) - Data
  > certification - nullable

- FundingData (\*FundingApplicationData) - Data funding - nullable

###  {#section-29 .unnumbered}

### DashboardData DTO

> DTO untuk data dashboard mobile.
>
> **Fields:**

- Name (string) - Nama user

- KartuType (string) - Tipe kartu

- KartuNumber (string) - Nomor kartu

- QRCode (string) - URL QR code

- NotificationsCount (int) - Jumlah notifikasi belum dibaca

- TotalApplications (int) - Total aplikasi

- ApprovedApplications (int) - Aplikasi disetujui

###  {#section-30 .unnumbered}

### UMKMDocument DTO

> DTO untuk dokumen UMKM.
>
> **Fields:**

- DocumentType (string) - Tipe dokumen

- DocumentURL (string) - URL dokumen

### NotificationResponse DTO

> DTO untuk response notifikasi.
>
> **Fields:**

- ID (int) - ID notifikasi

- Type (string) - Tipe notifikasi

- Title (string) - Judul

- Message (string) - Pesan

- IsRead (bool) - Status baca

- ReadAt (\*string) - Waktu dibaca - nullable

- Metadata (map\[string\]interface{}) - Data tambahan

- CreatedAt (string) - Waktu dibuat

- ApplicationID (\*int) - ID aplikasi terkait - nullable

###  {#section-31 .unnumbered}

### NewsRequest DTO

> DTO untuk request create/update berita.
>
> **Fields:**

- Title (string) - Judul berita (required)

- Excerpt (string) - Ringkasan

- Content (string) - Isi berita (required)

- Thumbnail (string) - Thumbnail base64 atau URL

- Category (string) - Kategori (required, values: announcement, event,
  > program_update, success_story, tips, regulation, general)

- IsPublished (bool) - Status publish

- Tags (\[\]string) - Daftar tag

###  {#section-32 .unnumbered}

### NewsResponse DTO

> DTO untuk response detail berita.
>
> **Fields:**

- ID (int) - ID berita

- Title (string) - Judul

- Slug (string) - Slug URL

- Excerpt (string) - Ringkasan

- Content (string) - Isi

- Thumbnail (string) - URL thumbnail

- Category (string) - Kategori

- AuthorID (int) - ID penulis

- AuthorName (string) - Nama penulis

- IsPublished (bool) - Status publish

- PublishedAt (\*string) - Waktu publish - nullable

- ViewsCount (int) - Jumlah views

- CreatedAt (string) - Waktu dibuat

- UpdatedAt (string) - Waktu diupdate

- Tags (\[\]string) - Daftar tag

###  {#section-33 .unnumbered}

### NewsListResponse DTO

> DTO untuk list berita di web.
>
> **Fields:**

- ID (int) - ID berita

- Title (string) - Judul

- Slug (string) - Slug URL

- Excerpt (string) - Ringkasan

- Thumbnail (string) - URL thumbnail

- Category (string) - Kategori

- AuthorName (string) - Nama penulis

- IsPublished (bool) - Status publish

- PublishedAt (\*string) - Waktu publish - nullable

- ViewsCount (int) - Jumlah views

- CreatedAt (string) - Waktu dibuat

###  {#section-34 .unnumbered}

### NewsListMobile DTO

> DTO untuk list berita di mobile.
>
> **Fields:**

- ID (int) - ID berita

- Title (string) - Judul

- Slug (string) - Slug URL

- Excerpt (string) - Ringkasan

- Thumbnail (string) - URL thumbnail

- Category (string) - Kategori

- AuthorName (string) - Nama penulis

- ViewsCount (int) - Jumlah views

- CreatedAt (string) - Waktu dibuat

###  {#section-35 .unnumbered}

### NewsDetailMobile DTO

> DTO untuk detail berita di mobile.
>
> **Fields:**

- ID (int) - ID berita

- Title (string) - Judul

- Slug (string) - Slug URL

- Content (string) - Isi lengkap

- Thumbnail (string) - URL thumbnail

- Category (string) - Kategori

- AuthorName (string) - Nama penulis

- ViewsCount (int) - Jumlah views

- CreatedAt (string) - Waktu dibuat

- Tags (\[\]string) - Daftar tag

###  {#section-36 .unnumbered}

### NewsQueryParams DTO

> DTO untuk query parameter berita.
>
> **Fields:**

- Page (int) - Nomor halaman

- Limit (int) - Jumlah per halaman

- Category (string) - Filter kategori

- Search (string) - Kata kunci pencarian

- Tag (string) - Filter tag

- IsPublished (\*bool) - Filter status publish - nullable

## Helper  {#helper}

### Password Management

#### **PasswordValidator** {#passwordvalidator .unnumbered}

> **func** PasswordValidator(str string) (bool, bool, bool)
>
> **Fungsi:** Memvalidasi kekuatan password dengan memeriksa kriteria
> keamanan.
>
> **Return Value:**

- hasMinLen (bool): Password minimal 8 karakter

- hasLetter (bool): Password mengandung minimal 1 huruf

- hasDigit (bool): Password mengandung minimal 1 angka

> **Digunakan di:**

- internal/service/users.go → Register (web)

- internal/service/users.go → RegisterMobileProfile

- internal/service/users.go → ResetPassword

> **Tujuan:** Memastikan password memenuhi standar keamanan minimum
> sebelum disimpan ke database.

#### **PasswordHashing** {#passwordhashing .unnumbered}

> **func** PasswordHashing(str string) (string, error)
>
> **Fungsi:** Mengenkripsi password menggunakan algoritma bcrypt dengan
> cost minimum.
>
> **Parameter:**

- str: Password plaintext yang akan di-hash

> **Return:** Password yang sudah di-hash atau error
>
> **Digunakan di:**

- internal/service/users.go → Register (web)

- internal/service/users.go → RegisterMobileProfile

- internal/service/users.go → ResetPassword

> **Tujuan:** Mengamankan password user sebelum disimpan ke database
> dengan one-way encryption.

#### **ComparePass** {#comparepass .unnumbered}

> **func** ComparePass(hashPassword, reqPassword string) bool
>
> **Fungsi:** Membandingkan password hash dengan password plaintext
> untuk validasi login.
>
> **Parameter:**

- hashPassword: Password yang sudah di-hash dari database

- reqPassword: Password plaintext dari request user

> **Return:** true jika password cocok, false jika tidak
>
> **Digunakan di:**

- internal/service/users.go → Login (web)

- internal/service/users.go → LoginMobile

> **Tujuan:** Memverifikasi kredensial user saat login tanpa mendekripsi
> password asli.

### JWT Token Management

#### **GenerateWebToken** {#generatewebtoken .unnumbered}

> **func** GenerateWebToken(user dto.Users) (string, error)
>
> **Fungsi:** Membuat JWT token untuk user web dashboard (admin) dengan
> claims yang berisi informasi user dan permissions.
>
> **Claims yang disimpan:**

- id: User ID

- name: Nama user

- email: Email user

- role: Role ID

- role_name: Nama role

- permissions: List permissions

- iat: Issued at timestamp

- exp: Expiration time (3 hari)

- is_admin: Boolean flag true

> **Digunakan di:**

- internal/service/users.go → Login

> **Tujuan:** Menghasilkan token autentikasi untuk sesi web dashboard
> dengan permission-based access control.

#### **GenerateMobileToken** {#generatemobiletoken .unnumbered}

> **func** GenerateMobileToken(user dto.UMKMMobile) (string, error)
>
> **Fungsi:** Membuat JWT token untuk user mobile (pelaku usaha/UMKM).
>
> **Claims yang disimpan:**

- id: UMKM ID

- name: Nama lengkap

- business_name: Nama usaha

- email: Email

- phone: Nomor telepon

- kartu_type: Tipe kartu UMKM

- iat: Issued at timestamp

- exp: Expiration time (3 hari)

- is_admin: Boolean flag false

> **Digunakan di:**

- internal/service/users.go → RegisterMobileProfile

- internal/service/users.go → LoginMobile

> **Tujuan:** Menghasilkan token autentikasi untuk sesi mobile app
> dengan data UMKM.

#### **VerifyToken** {#verifytoken .unnumbered}

> **func** VerifyToken(jwtToken string) (dto.UserData, error)
>
> **Fungsi:** Memverifikasi dan mem-parse JWT token untuk mengekstrak
> informasi user.
>
> **Parameter:**

- jwtToken: Token JWT string

> **Return:** Data user yang ter-extract dari token atau error
>
> **Digunakan di:**

- Middleware autentikasi (AuthMiddleware dan MobileAuthMiddleware)

> **Tujuan:** Validasi token dan ekstraksi data user untuk authorization
> di setiap request yang memerlukan autentikasi.

### Email Validation {#email-validation .unnumbered}

#### **EmailValidator** {#emailvalidator .unnumbered}

> **func** EmailValidator(str string) bool
>
> **Fungsi:** Memvalidasi format email menggunakan regular expression.
>
> **Pattern:** \^\[a-zA-Z0-9.\_%+-\]+@\[a-zA-Z0-9.-\]+\\\[a-zA-Z\]{2,}\$
>
> **Digunakan di:**

- internal/service/users.go → Register

- internal/service/users.go → UpdateProfile

- internal/service/users.go → UpdateUser

- internal/service/users.go → SetOTP

- internal/service/users.go → ValidateOTP

- internal/service/users.go → RegisterMobile

> **Tujuan:** Memastikan format email valid sebelum disimpan atau
> diproses.

### Phone Number Management

#### **NormalizePhone** {#normalizephone .unnumbered}

> **func** NormalizePhone(phone string) (string, error)
>
> **Fungsi:** Menormalisasi nomor telepon Indonesia ke format standar
> (dimulai dari '8').
>
> **Konversi yang dilakukan:**

- +62812345678 → 812345678

- 62812345678 → 812345678

- 0812345678 → 812345678

- 812345678 → 812345678

> **Validasi:**

- Minimal 9 digit setelah normalisasi

- Hanya menerima format Indonesia (0, 62, +62, atau 8)

> **Digunakan di:**

- internal/service/users.go → RegisterMobile

- internal/service/users.go → VerifyOTP

- internal/service/users.go → LoginMobile

- internal/service/users.go → ForgotPassword

> **Tujuan:** Menyeragamkan format nomor telepon untuk konsistensi
> penyimpanan dan pengiriman OTP.

#### **DenormalizePhone** {#denormalizephone .unnumbered}

> **func** DenormalizePhone(phone string) string
>
> **Fungsi:** Mengkonversi nomor telepon dari format normalized ke
> format internasional (+62).
>
> **Konversi yang dilakukan:**

- 812345678 → +62812345678

- 0812345678 → +62812345678

- 62812345678 → +62812345678

> **Digunakan di:**

- internal/service/applications.go → GetApplicationByID (untuk
  > menampilkan nomor)

> **Tujuan:** Menampilkan nomor telepon dalam format yang user-friendly
> dan standar internasional.

### OTP Management

#### **GenerateOTP** {#generateotp .unnumbered}

> **func** GenerateOTP() string
>
> **Fungsi:** Menghasilkan kode OTP 6 digit random.
>
> **Format:** String 6 digit (000000 - 999999)
>
> **Digunakan di:**

- internal/service/users.go → RegisterMobile

- internal/service/users.go → ForgotPassword

> **Tujuan:** Membuat kode verifikasi untuk proses registrasi dan reset
> password.

### NIK Validation

#### **NIKValidator** {#nikvalidator .unnumbered}

> **func** NIKValidator(nik string) error
>
> **Fungsi:** Memvalidasi format NIK (Nomor Induk Kependudukan)
> Indonesia.
>
> **Validasi yang dilakukan:**

1.  Panjang harus 16 digit

2.  Hanya berisi angka

3.  Validasi tanggal lahir (digit 7-8):

    - Jika \> 40: perempuan (dikurangi 40)

    - Range 1-31 untuk tanggal

<!-- -->

4.  Validasi bulan (digit 9-10): Range 1-12

5.  Validasi tahun (digit 11-12): Tidak boleh di masa depan

> **Digunakan di:**

- Saat ini di-comment di internal/service/users.go →
  > RegisterMobileProfile

> **Tujuan:** Memastikan NIK yang diinput valid sesuai format standar
> Indonesia.

### File and String Utilities

#### **GenerateRequestID** {#generaterequestid .unnumbered}

> **func** GenerateRequestID() string
>
> **Fungsi:** Membuat request ID unik berupa string alphanumeric random
> sepanjang 10 karakter.
>
> **Charset:** a-z, A-Z, 0-9
>
> **Digunakan di:**

- Context tracking untuk logging dan audit trail

> **Tujuan:** Memberikan identifier unik untuk setiap request untuk
> keperluan tracing dan debugging.

#### **GenerateFileName** {#generatefilename .unnumbered}

> **func** GenerateFileName(originalName, prefix string) string
>
> **Fungsi:** Membuat nama file yang standar untuk upload ke MinIO
> storage.
>
> **Format output:** original_name_lowercase/prefix\_
>
> **Contoh:**

- Input: \"Program Pelatihan\", \"banner\_\"

- Output: \"program_pelatihan/banner\_\"

> **Digunakan di:**

- internal/service/programs.go → Upload banner dan provider logo

- internal/service/news.go → Upload thumbnail berita

- internal/service/mobile.go → Upload dokumen aplikasi

- internal/service/users.go → Upload QR code

> **Tujuan:** Menyediakan naming convention yang konsisten dan
> terorganisir untuk file storage.

#### **RandomString** {#randomstring .unnumbered}

> **func** RandomString(size int) string
>
> **Fungsi:** Menghasilkan string hexadecimal random dengan panjang yang
> ditentukan.
>
> **Parameter:**

- size: Panjang string yang diinginkan (0 untuk full 64 chars)

> **Digunakan di:**

- internal/service/users.go → Generate temp token untuk OTP verification

> **Tujuan:** Membuat token temporary yang secure untuk proses
> verifikasi multi-step.

#### **MaskMiddle** {#maskmiddle .unnumbered}

> **func** MaskMiddle(s string) string
>
> **Fungsi:** Menyembunyikan bagian tengah string dengan "XXXXXXXX"
> untuk keamanan data sensitif.
>
> **Logika:**

- Tampilkan 1/3 awal

- Mask dengan "XXXXXXXX"

- Tampilkan 1/3 akhir

> **Contoh:**

- Input: \"3201234567890123\"

- Output: \"32012XXXXXXXX0123\"

> **Digunakan di:**

- Saat ini belum digunakan dalam codebase, siap untuk diimplementasikan

> **Tujuan:** Menyembunyikan data sensitif seperti NIK atau nomor kartu
> saat ditampilkan di UI.

### QR Code Generation

#### **GenerateQRCode** {#generateqrcode .unnumbered}

> **func** GenerateQRCode(data string, size int) (string, error)
>
> **Fungsi:** Menghasilkan QR code dalam format base64 dari data string.
>
> **Parameter:**

- data: String yang akan di-encode ke QR code

- size: Ukuran QR code dalam pixel (default 256 jika \<= 0)

> **Error Correction Level:** Medium
>
> **Digunakan di:**

- internal/service/users.go → RegisterMobileProfile (untuk kartu UMKM)

> **Tujuan:** Membuat QR code dari nomor kartu UMKM untuk kemudahan
> scanning dan verifikasi.

### Email Sending Utilities

#### **SMTPInterface** {#smtpinterface .unnumbered}

> **type** SMTPInterface **interface** {  
> GetAuth() smtp.Auth  
> GetAddress() string  
> GetUser() string  
> }
>
> **Fungsi:** Interface untuk konfigurasi SMTP server.
>
> **Implementation:** zohoSMTP struct untuk Zoho Mail
>
> **Digunakan di:**

- Email service untuk mengirim notifikasi dan OTP via email

> **Tujuan:** Abstraksi untuk berbagai SMTP provider (saat ini
> menggunakan Zoho).

#### **NewSMTPClient** {#newsmtpclient .unnumbered}

> **func** NewSMTPClient(smtpInterface SMTPInterface)
> SMTPClientInterface
>
> **Fungsi:** Membuat client SMTP untuk pengiriman email.
>
> **Method yang tersedia:**

- SendSingleEmail(to, subject, htmlFile, data): Kirim email dengan
  > template HTML

> **Template Functions:**

- formatDateMY: Format tanggal ke "January 2006"

- formatDateMDY: Format tanggal ke "January 02, 2006"

- formatDateMDYT: Format tanggal dengan waktu

- convertBToMB: Konversi bytes ke megabytes

> **Digunakan di:**

- OTP verification email

- Password reset email

- Notifikasi email lainnya

> **Tujuan:** Mengirim email transaksional dengan template HTML yang
> dinamis.

## Constants  {#constants}

### Environment Modes

> DEVELOPMENT_MODE, STAGING_MODE, PRODUCTION_MODE
>
> **const** (  
> DEVELOPMENT_MODE = \"development\"  
> STAGING_MODE = \"staging\"  
> PRODUCTION_MODE = \"production\"  
> )
>
> **Fungsi:** Menentukan environment aplikasi berjalan.
>
> **Digunakan di:**

- Configuration loading

- Database connection

- Logging level

- Feature flags

> **Tujuan:** Memisahkan konfigurasi dan behavior berdasarkan
> environment.

###  Connection Settings {#connection-settings}

> DefaultConnectionTimeout
>
> **const** DefaultConnectionTimeout = 30 \* time.Second
>
> **Fungsi:** Timeout default untuk koneksi database dan external
> service.
>
> **Digunakan di:**

- Database connection setup

- HTTP client requests

- External API calls

> **Tujuan:** Mencegah hanging request dan memberikan consistent timeout
> behavior.

### User Roles

> Role Constants
>
> **const** (  
> RoleSuperAdmin = \"superadmin\"  
> RoleAdminScreening = \"admin_screening\"  
> RoleAdminVendor = \"admin_vendor\"  
> RoleUMKM = \"pelaku_usaha\"  
> )
>
> **Fungsi:** Definisi role user dalam sistem.
>
> **Hierarchy:**

- RoleSuperAdmin: Full access ke semua fitur

- RoleAdminScreening: Review aplikasi tahap screening

- RoleAdminVendor: Review aplikasi tahap final

- RoleUMKM: Pelaku usaha (mobile user)

> **Digunakan di:**

- internal/service/users.go → RegisterMobileProfile (mencari role UMKM)

- Authorization middleware

- Permission checking

> **Tujuan:** Implementasi role-based access control (RBAC) untuk
> authorization.

### OTP Status

> OTP Status Constants
>
> **const** (  
> OTPStatusActive = \"active\"  
> OTPStatusUsed = \"used\"  
> )
>
> **Fungsi:** Status lifecycle OTP.
>
> **States:**

- OTPStatusActive: OTP valid dan belum digunakan

- OTPStatusUsed: OTP sudah digunakan (tidak bisa dipakai lagi)

> **Digunakan di:**

- internal/service/users.go → RegisterMobile (create OTP)

- internal/service/users.go → VerifyOTP (validasi status)

- internal/service/users.go → RegisterMobileProfile (update ke used)

- internal/service/users.go → ForgotPassword (create OTP)

- internal/service/users.go → ResetPassword (update ke used)

> **Tujuan:** Memastikan OTP hanya bisa digunakan sekali dan tracking
> lifecycle-nya.

### Application Status

> Application Status Constants
>
> **const** (  
> ApplicationStatusScreening = \"screening\"  
> ApplicationStatusRevised = \"revised\"  
> ApplicationStatusFinal = \"final\"  
> ApplicationStatusApproved = \"approved\"  
> ApplicationStatusRejected = \"rejected\"  
> )
>
> **Fungsi:** Status workflow aplikasi program UMKM.
>
> **Flow:**

1.  screening: Aplikasi baru masuk, dalam review tahap screening

2.  revised: Diminta revisi oleh admin screening

3.  final: Lolos screening, masuk tahap final review

4.  approved: Disetujui pada tahap final

5.  rejected: Ditolak (bisa dari screening atau final)

> **Digunakan di:**

- internal/service/applications.go → Semua decision methods

- internal/service/mobile.go → CreateApplication, ReviseApplication

- internal/service/mobile.go → GetDashboard (count approved)

> **Tujuan:** Tracking workflow approval dengan multi-stage process.

### Notification Types

> Notification Type Constants
>
> **const** (  
> NotificationSubmitted = \"application_submitted\"  
> NotificationApproved = \"screening_approved\"  
> NotificationRejected = \"screening_rejected\"  
> NotificationRevised = \"screening_revised\"  
> NotificationFinalApproved = \"final_approved\"  
> NotificationFinalRejected = \"final_rejected\"  
> NotificationProgramReminder = \"program_reminder\"  
> NotificationDocumentRequired = \"document_required\"  
> NotificationGeneralInfo = \"general_info\"  
> )
>
> **Fungsi:** Tipe notifikasi untuk kategorisasi dan filtering.
>
> **Categories:**

- Application workflow: submitted, approved, rejected, revised

- Final decisions: final_approved, final_rejected

- Reminders: program_reminder, document_required

- General: general_info

> **Digunakan di:**

- internal/service/applications.go → Create notification saat decision

- internal/service/mobile.go → Create notification saat submit/revise

> **Tujuan:** Kategorisasi notifikasi untuk UI filtering dan
> notification handling.

### Notification Titles

> Notification Title Constants
>
> **const** (  
> NotificationTitleSubmitted = \"Pengajuan Dikirim\"  
> NotificationTitleResubmitted = \"Pengajuan Dikirim Ulang\"  
> NotificationTitleApproved = \"Pengajuan Disetujui pada Tahap
> Screening\"  
> NotificationTitleRejected = \"Pengajuan Ditolak pada Tahap
> Screening\"  
> NotificationTitleRevised = \"Pengajuan Direvisi pada Tahap
> Screening\"  
> NotificationTitleFinalApproved = \"Pengajuan Disetujui pada Tahap
> Final\"  
> NotificationTitleFinalRejected = \"Pengajuan Ditolak pada Tahap
> Final\"  
> NotificationTitleProgramReminder = \"Pengingat Program\"  
> NotificationTitleDocumentRequired = \"Dokumen Diperlukan\"  
> NotificationTitleGeneralInfo = \"Informasi Umum\"  
> )
>
> **Fungsi:** Template judul notifikasi dalam Bahasa Indonesia.
>
> **Digunakan di:**

- internal/service/applications.go → ScreeningApprove, ScreeningReject,
  > ScreeningRevise, FinalApprove, FinalReject

- internal/service/mobile.go → CreateTrainingApplication,
  > CreateCertificationApplication, CreateFundingApplication,
  > ReviseApplication

> **Tujuan:** Standarisasi judul notifikasi untuk konsistensi UI dan
> user experience.

### Notification Messages

> Notification Message Constants
>
> **const** (  
> NotificationMessageSubmitted = \"Pengajuan Anda telah berhasil
> dikirim. Silakan tunggu proses screening.\"  
> NotificationMessageResubmitted = \"Pengajuan ulang Anda telah berhasil
> dikirim. Silakan tunggu proses screening.\"  
> NotificationMessageApproved = \"Pengajuan Anda telah disetujui pada
> tahap screening. Silakan menunggu lanjut ke tahap final.\"  
> NotificationMessageRejected = \"Pengajuan Anda telah ditolak pada
> tahap screening. Karena %s. Silakan periksa kembali data yang Anda
> kirim.\"  
> NotificationMessageRevised = \"Pengajuan Anda perlu direvisi pada
> tahap screening. Karena %s. Silakan periksa kembali data yang Anda
> kirim.\"  
> NotificationMessageFinalApproved = \"Pengajuan Anda telah disetujui
> pada tahap final. Selamat!\"  
> NotificationMessageFinalRejected = \"Pengajuan Anda telah ditolak pada
> tahap final. Karena %s. Silakan periksa kembali data yang Anda
> kirim.\"  
> NotificationMessageProgramReminder = \"Ingatkan program yang akan
> datang.\"  
> NotificationMessageDocumentRequired = \"Dokumen tambahan diperlukan
> untuk melanjutkan proses pengajuan.\"  
> NotificationMessageGeneralInfo = \"Informasi umum terkait program atau
> aplikasi.\"  
> )
>
> **Fungsi:** Template pesan notifikasi dengan support untuk dynamic
> content menggunakan fmt.Sprintf.
>
> **Dynamic Messages:**

- NotificationMessageRejected: %s diganti dengan notes rejection

- NotificationMessageRevised: %s diganti dengan notes revision

- NotificationMessageFinalRejected: %s diganti dengan notes rejection

> **Digunakan di:**

- internal/service/applications.go → ScreeningReject, ScreeningRevise,
  > FinalReject (dengan fmt.Sprintf untuk notes)

- internal/service/applications.go → ScreeningApprove, FinalApprove
  > (static message)

- internal/service/mobile.go → Create dan revise application (static
  > message)

> **Tujuan:** Standarisasi pesan notifikasi dengan kemampuan
> customization untuk konteks spesifik.

### Document Types

> Document Type Constants
>
> **const** (  
> DocumentTypeNib = \"nib\"  
> DocumentTypeNPWP = \"npwp\"  
> DocumentTypeRevenueRecord = \"revenue_record\"  
> DocumentTypeBusinessPermit = \"business_permit\"  
> )
>
> **Fungsi:** Tipe dokumen UMKM yang diperlukan.
>
> **Documents:**

- nib: Nomor Induk Berusaha

- npwp: Nomor Pokok Wajib Pajak

- revenue_record: Catatan omzet/pendapatan

- business_permit: Izin usaha

> **Digunakan di:**

- internal/service/mobile.go → GetUMKMDocuments (mapping field ke
  > document type)

- internal/service/mobile.go → UploadDocument (validasi document type)

> **Tujuan:** Standarisasi tipe dokumen untuk upload, validasi, dan
> display consistency.

## Deployments  {#deployments}

### GitLab CI/CD Pipeline

#### **.gitlab-ci.yml Stages** {#gitlab-ci.yml-stages .unnumbered}

> stages**:  
> ** **-** test  
> **-** build  
> **-** release  
> **-** deploy  
> **-** migrate
>
> **Stage Functions:**
>
> **test**

- **Job:** unit_test

- **Image:** golang:1.24.4

- **Trigger:** Tag dengan format v\[year\].\[version\] (e.g., v2024.1)

- **Actions:**

  - Install go-junit-report untuk test reporting

  - Run unit tests dengan coverage

  - Generate JUnit report XML

  - Generate coverage report

  - Extract coverage percentage dengan regex

<!-- -->

- **Artifacts:** report.xml, coverage.out

- **Digunakan untuk:** Quality assurance sebelum build

> **build**

- **Jobs:** build_backend, build_migrate

- **Image:** docker:latest dengan docker:dind service

- **Trigger:** Tag dengan format v\[year\].\[version\]

- **Actions:**

  - Build Docker image dari Dockerfile.api dan Dockerfile.migrate

  - Tag dengan latest dan version tag

  - Push ke GitLab Container Registry

  - Save job ID ke artifact

<!-- -->

- **Artifacts:** Build job ID files

- **Digunakan untuk:** Container packaging

> **release**

- **Jobs:** prepare_release, create_release

- **Actions:**

  - Upload coverage report sebagai generic package

  - Export job IDs untuk artifact download links

  - Create GitLab release dengan links ke:

    - Backend image artifacts

    - Migration image artifacts

    - Coverage report

<!-- -->

- **Digunakan untuk:** Release documentation dan artifact management

> **deploy**

- **Jobs:** deploy-staging, deploy-production

- **Image:** alpine:latest

- **Trigger:** Tag dengan format v\[year\].\[version\]

- **Mode:** staging auto, production manual

- **Actions:**

  - SSH ke deployment server

  - Set environment variables

  - Docker login ke registry

  - Pull latest images

  - Stop old containers

  - Start new containers dengan docker-compose

<!-- -->

- **Environments:**

  - Staging: umkmgo-staging-\* containers

  - Production: umkmgo-production-\* containers

<!-- -->

- **Digunakan untuk:** Automated deployment ke server

> **migrate**

- **Jobs:** migrate-staging, migrate-production

- **Image:** alpine:latest

- **Trigger:** Manual execution after deploy

- **Actions:**

  - SSH ke deployment server

  - Pull migration image

  - Run database migration dengan goose

<!-- -->

- **Dependencies:** Needs corresponding deploy job

- **Digunakan untuk:** Database schema updates

> **Environment Variables Required:**

- CI_REGISTRY_USER: GitLab registry username

- CI_REGISTRY_PASSWORD: GitLab registry password

- SSH_PRIVATE_KEY: SSH key untuk access deployment server

- SSH_USER: SSH username

- SSH_HOST: Deployment server hostname

- SSH_PORT: SSH port

### Docker Configurations

#### **Dockerfile.api** {#dockerfile.api .unnumbered}

> **Purpose:** Build backend API application
>
> **Build Stages:**

1.  **Builder Stage:**

    - Base: golang:1.24.4-alpine

    - Copy go.mod dan go.sum

    - Download dependencies

    - Build binary dengan CGO disabled untuk alpine compatibility

    - Output: /app/main

<!-- -->

2.  **Runtime Stage:**

    - Base: alpine:latest (minimal image)

    - Copy binary from builder

    - Expose port 8080

    - Run binary

> **Optimization:** Multi-stage build untuk image size yang minimal

#### **Dockerfile.migrate** {#dockerfile.migrate .unnumbered}

> **Purpose:** Build database migration container
>
> **Build Stages:**

1.  **Builder Stage:**

    - Base: golang:1.24.4-alpine

    - Install git dan ca-certificates

    - Clone goose migration tool

    - Build goose dengan tag excludes untuk unused databases

    - Excluded: clickhouse, libsql, mssql, mysql, sqlite3, vertica, ydb

<!-- -->

2.  **Runtime Stage:**

    - Base: alpine:latest

    - Install postgresql-client untuk pg_isready check

    - Copy goose binary

    - Copy migration files dan seeder

    - Copy dan set executable migrate.sh script

> **Purpose:** Dedicated container untuk database migrations dengan
> goose

### Migration Script

#### **migrate.sh** {#migrate.sh .unnumbered}

> **Purpose:** Shell script untuk menjalankan database migrations
>
> **Environment Variables:**

- MODE: Environment mode (staging/production), default: staging

- GOOSE_CMD: Goose command (up/down/status), default: up

> **Configuration per Environment:**

- **Staging:**

  - Host: umkmgo-staging-postgres

  - Port: 5432

  - Database: umkmgo

<!-- -->

- **Production:**

  - Host: umkmgo-production-postgres

  - Port: 5432

  - Database: umkmgo

> **Execution Flow:**

1.  Set goose environment variables

2.  Wait for database to be ready (pg_isready check)

3.  Run goose command with provided arguments

4.  Log success message

> **Usage Examples:**

- ./migrate.sh up - Run all pending migrations

- ./migrate.sh down - Rollback last migration

- ./migrate.sh status - Check migration status

> **Digunakan di:** GitLab CI/CD migrate jobs
