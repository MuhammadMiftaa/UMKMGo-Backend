-- +goose Up

TRUNCATE TABLE programs RESTART IDENTITY CASCADE;

-- +goose StatementBegin
-- Seed Programs - Training (10 programs)
INSERT INTO programs (title, description, banner, provider, provider_logo, type, training_type, batch, batch_start_date, batch_end_date, location, application_deadline, is_active, created_by, created_at, updated_at) VALUES
('Pelatihan Digital Marketing untuk UMKM', 'Pelatihan komprehensif tentang digital marketing untuk meningkatkan penjualan UMKM melalui platform online', 'https://example.com/banner/digital-marketing.jpg', 'Kementerian Koperasi dan UKM', 'https://example.com/logo/kemenkop.png', 'training', 'hybrid', 1, '2025-12-01', '2025-12-15', 'Jakarta Convention Center & Online', '2025-11-20', true, 1, NOW(), NOW()),
('Pelatihan Akuntansi Dasar UMKM', 'Pelatihan pembukuan dan akuntansi sederhana untuk UMKM', 'https://example.com/banner/accounting.jpg', 'Kementerian Koperasi dan UKM', 'https://example.com/logo/kemenkop.png', 'training', 'online', 2, '2025-11-15', '2025-11-30', 'Full Online', '2025-11-05', true, 1, NOW(), NOW()),
('Pelatihan Manajemen Usaha', 'Pelatihan manajemen usaha untuk meningkatkan efisiensi bisnis', 'https://example.com/banner/management.jpg', 'Kementerian Koperasi dan UKM', 'https://example.com/logo/kemenkop.png', 'training', 'offline', 1, '2025-12-10', '2025-12-20', 'Bandung', '2025-11-25', true, 1, NOW(), NOW()),
('Pelatihan Export Import', 'Pelatihan ekspor impor untuk UMKM yang ingin go international', 'https://example.com/banner/export.jpg', 'Kementerian Perdagangan', 'https://example.com/logo/kemendag.png', 'training', 'hybrid', 1, '2025-12-05', '2025-12-25', 'Jakarta & Online', '2025-11-15', true, 1, NOW(), NOW()),
('Pelatihan Fotografi Produk', 'Pelatihan fotografi produk untuk meningkatkan visual branding', 'https://example.com/banner/photography.jpg', 'Bekraf', 'https://example.com/logo/bekraf.png', 'training', 'offline', 2, '2025-11-20', '2025-12-05', 'Surabaya', '2025-11-10', true, 1, NOW(), NOW()),
('Pelatihan E-Commerce', 'Pelatihan menggunakan platform e-commerce untuk berjualan online', 'https://example.com/banner/ecommerce.jpg', 'Tokopedia Academy', 'https://example.com/logo/tokopedia.png', 'training', 'online', 3, '2025-12-01', '2025-12-15', 'Full Online', '2025-11-20', true, 1, NOW(), NOW()),
('Pelatihan Social Media Marketing', 'Pelatihan pemasaran melalui media sosial', 'https://example.com/banner/socialmedia.jpg', 'Kementerian Koperasi dan UKM', 'https://example.com/logo/kemenkop.png', 'training', 'hybrid', 1, '2025-11-25', '2025-12-10', 'Yogyakarta & Online', '2025-11-12', true, 1, NOW(), NOW()),
('Pelatihan Packaging Design', 'Pelatihan desain kemasan produk yang menarik', 'https://example.com/banner/packaging.jpg', 'Bekraf', 'https://example.com/logo/bekraf.png', 'training', 'offline', 1, '2025-12-15', '2025-12-20', 'Bandung', '2025-11-30', true, 1, NOW(), NOW()),
('Pelatihan Branding UMKM', 'Pelatihan membangun brand untuk UMKM', 'https://example.com/banner/branding.jpg', 'Kementerian Koperasi dan UKM', 'https://example.com/logo/kemenkop.png', 'training', 'online', 2, '2025-12-05', '2025-12-20', 'Full Online', '2025-11-25', true, 1, NOW(), NOW()),
('Pelatihan Content Marketing', 'Pelatihan membuat konten marketing yang efektif', 'https://example.com/banner/content.jpg', 'Google Digital Garage', 'https://example.com/logo/google.png', 'training', 'hybrid', 1, '2025-11-28', '2025-12-18', 'Jakarta & Online', '2025-11-18', true, 1, NOW(), NOW());
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Programs - Certification (10 programs)
INSERT INTO programs (title, description, banner, provider, provider_logo, type, training_type, batch, batch_start_date, batch_end_date, location, application_deadline, is_active, created_by, created_at, updated_at) VALUES
('Sertifikasi Halal untuk Produk UMKM', 'Program sertifikasi halal untuk produk makanan dan minuman UMKM sesuai standar LPPOM MUI', 'https://example.com/banner/halal-cert.jpg', 'LPPOM MUI', 'https://example.com/logo/mui.png', 'certification', 'offline', 2, '2025-11-15', '2025-11-30', 'Jakarta, Surabaya, Medan', '2025-11-01', true, 1, NOW(), NOW()),
('Sertifikasi SNI Produk', 'Sertifikasi Standar Nasional Indonesia untuk produk UMKM', 'https://example.com/banner/sni.jpg', 'BSN', 'https://example.com/logo/bsn.png', 'certification', 'offline', 1, '2025-12-01', '2025-12-20', 'Jakarta', '2025-11-15', true, 1, NOW(), NOW()),
('Sertifikasi ISO 9001', 'Sertifikasi ISO 9001 untuk sistem manajemen mutu', 'https://example.com/banner/iso.jpg', 'BSN', 'https://example.com/logo/bsn.png', 'certification', 'hybrid', 1, '2025-12-10', '2026-01-10', 'Bandung & Online', '2025-11-25', true, 1, NOW(), NOW()),
('Sertifikasi HACCP', 'Sertifikasi HACCP untuk keamanan pangan', 'https://example.com/banner/haccp.jpg', 'BPOM', 'https://example.com/logo/bpom.png', 'certification', 'offline', 2, '2025-11-20', '2025-12-15', 'Jakarta', '2025-11-10', true, 1, NOW(), NOW()),
('Sertifikasi Organic', 'Sertifikasi produk organik untuk produk pertanian', 'https://example.com/banner/organic.jpg', 'Kementan', 'https://example.com/logo/kementan.png', 'certification', 'hybrid', 1, '2025-12-05', '2025-12-25', 'Yogyakarta & Online', '2025-11-20', true, 1, NOW(), NOW()),
('Sertifikasi PIRT', 'Sertifikasi Produksi Pangan Industri Rumah Tangga', 'https://example.com/banner/pirt.jpg', 'Dinkes', 'https://example.com/logo/dinkes.png', 'certification', 'offline', 3, '2025-11-25', '2025-12-10', 'Bandung', '2025-11-15', true, 1, NOW(), NOW()),
('Sertifikasi GMP', 'Sertifikasi Good Manufacturing Practice', 'https://example.com/banner/gmp.jpg', 'BPOM', 'https://example.com/logo/bpom.png', 'certification', 'offline', 1, '2025-12-15', '2026-01-05', 'Jakarta', '2025-11-30', true, 1, NOW(), NOW()),
('Sertifikasi Fairtrade', 'Sertifikasi perdagangan yang adil', 'https://example.com/banner/fairtrade.jpg', 'Fairtrade Indonesia', 'https://example.com/logo/fairtrade.png', 'certification', 'hybrid', 1, '2025-12-01', '2025-12-20', 'Surabaya & Online', '2025-11-18', true, 1, NOW(), NOW()),
('Sertifikasi Ecolabel', 'Sertifikasi produk ramah lingkungan', 'https://example.com/banner/ecolabel.jpg', 'Kementerian Lingkungan Hidup', 'https://example.com/logo/klh.png', 'certification', 'online', 2, '2025-11-28', '2025-12-18', 'Full Online', '2025-11-20', true, 1, NOW(), NOW()),
('Sertifikasi Keamanan Pangan', 'Sertifikasi keamanan pangan untuk produk makanan', 'https://example.com/banner/foodsafety.jpg', 'BPOM', 'https://example.com/logo/bpom.png', 'certification', 'offline', 1, '2025-12-08', '2025-12-28', 'Jakarta', '2025-11-25', true, 1, NOW(), NOW());
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Programs - Funding (10 programs)
INSERT INTO programs (title, description, banner, provider, provider_logo, type, min_amount, max_amount, interest_rate, max_tenure_months, application_deadline, is_active, created_by, created_at, updated_at) VALUES
('KUR Mikro - Kredit Usaha Rakyat', 'Program pinjaman modal usaha dengan bunga rendah untuk UMKM yang ingin mengembangkan bisnis', 'https://example.com/banner/kur-mikro.jpg', 'Bank Indonesia - KUR', 'https://example.com/logo/bi.png', 'funding', 1000000, 50000000, 6.0, 36, '2025-12-31', true, 1, NOW(), NOW()),
('KUR Kecil', 'Kredit Usaha Rakyat untuk usaha kecil dengan plafon lebih besar', 'https://example.com/banner/kur-kecil.jpg', 'Bank Indonesia - KUR', 'https://example.com/logo/bi.png', 'funding', 50000000, 500000000, 6.0, 48, '2025-12-31', true, 1, NOW(), NOW()),
('Pembiayaan Ultra Mikro', 'Pembiayaan untuk usaha ultra mikro dengan proses mudah', 'https://example.com/banner/ultra-mikro.jpg', 'Bank BRI', 'https://example.com/logo/bri.png', 'funding', 500000, 10000000, 7.0, 24, '2025-11-30', true, 1, NOW(), NOW()),
('Modal Ventura UMKM', 'Pembiayaan modal ventura untuk UMKM dengan potensi tinggi', 'https://example.com/banner/modal-ventura.jpg', 'PT PNM Ventura', 'https://example.com/logo/pnm.png', 'funding', 100000000, 1000000000, 0.0, 60, '2025-12-15', true, 1, NOW(), NOW()),
('Pinjaman Syariah Mikro', 'Pembiayaan syariah untuk UMKM tanpa riba', 'https://example.com/banner/syariah.jpg', 'Bank Syariah Indonesia', 'https://example.com/logo/bsi.png', 'funding', 1000000, 100000000, 8.0, 36, '2025-12-20', true, 1, NOW(), NOW()),
('Dana Bergulir UMKM', 'Program dana bergulir dari pemerintah untuk UMKM', 'https://example.com/banner/dana-bergulir.jpg', 'Kementerian Koperasi dan UKM', 'https://example.com/logo/kemenkop.png', 'funding', 5000000, 200000000, 3.0, 48, '2025-12-31', true, 1, NOW(), NOW()),
('Kredit Modal Kerja', 'Kredit untuk modal kerja UMKM', 'https://example.com/banner/modal-kerja.jpg', 'Bank Mandiri', 'https://example.com/logo/mandiri.png', 'funding', 10000000, 500000000, 9.0, 36, '2025-12-25', true, 1, NOW(), NOW()),
('Pinjaman Investasi', 'Pinjaman untuk investasi aset produktif', 'https://example.com/banner/investasi.jpg', 'Bank BNI', 'https://example.com/logo/bni.png', 'funding', 50000000, 1000000000, 10.0, 60, '2025-12-30', true, 1, NOW(), NOW()),
('Pembiayaan P2P Lending', 'Pembiayaan melalui platform peer-to-peer lending', 'https://example.com/banner/p2p.jpg', 'Modalku', 'https://example.com/logo/modalku.png', 'funding', 5000000, 200000000, 12.0, 12, '2025-12-15', true, 1, NOW(), NOW()),
('Kredit Usaha Produktif', 'Kredit untuk usaha produktif dengan bunga kompetitif', 'https://example.com/banner/produktif.jpg', 'Bank BTN', 'https://example.com/logo/btn.png', 'funding', 20000000, 500000000, 8.5, 48, '2025-12-28', true, 1, NOW(), NOW());
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Program Benefits for all programs
INSERT INTO program_benefits (program_id, name, created_at, updated_at) VALUES
-- Training Programs Benefits
(1, 'Sertifikat resmi dari Kemenkop', NOW(), NOW()),
(1, 'Modul pembelajaran lengkap', NOW(), NOW()),
(1, 'Mentoring selama 3 bulan', NOW(), NOW()),
(2, 'Sertifikat pelatihan', NOW(), NOW()),
(2, 'Software akuntansi gratis', NOW(), NOW()),
(3, 'Sertifikat manajemen', NOW(), NOW()),
(3, 'Konsultasi bisnis gratis', NOW(), NOW()),
(4, 'Sertifikat ekspor', NOW(), NOW()),
(4, 'Networking internasional', NOW(), NOW()),
(5, 'Sertifikat fotografi', NOW(), NOW()),
(5, 'Portfolio produk', NOW(), NOW()),
(6, 'Sertifikat e-commerce', NOW(), NOW()),
(6, 'Voucher Tokopedia', NOW(), NOW()),
(7, 'Sertifikat social media', NOW(), NOW()),
(7, 'Template konten', NOW(), NOW()),
(8, 'Sertifikat packaging design', NOW(), NOW()),
(8, 'Konsultasi desain', NOW(), NOW()),
(9, 'Sertifikat branding', NOW(), NOW()),
(9, 'Brand guideline', NOW(), NOW()),
(10, 'Sertifikat content marketing', NOW(), NOW()),
(10, 'Content calendar template', NOW(), NOW()),
-- Certification Programs Benefits
(11, 'Sertifikat halal resmi MUI', NOW(), NOW()),
(11, 'Logo halal untuk kemasan', NOW(), NOW()),
(12, 'Sertifikat SNI', NOW(), NOW()),
(12, 'Label SNI', NOW(), NOW()),
(13, 'Sertifikat ISO 9001', NOW(), NOW()),
(13, 'Audit assistance', NOW(), NOW()),
(14, 'Sertifikat HACCP', NOW(), NOW()),
(14, 'Panduan keamanan pangan', NOW(), NOW()),
(15, 'Sertifikat organic', NOW(), NOW()),
(15, 'Label organic', NOW(), NOW()),
(16, 'Sertifikat PIRT', NOW(), NOW()),
(16, 'Nomor PIRT', NOW(), NOW()),
(17, 'Sertifikat GMP', NOW(), NOW()),
(17, 'SOP produksi', NOW(), NOW()),
(18, 'Sertifikat Fairtrade', NOW(), NOW()),
(18, 'Logo Fairtrade', NOW(), NOW()),
(19, 'Sertifikat Ecolabel', NOW(), NOW()),
(19, 'Logo ramah lingkungan', NOW(), NOW()),
(20, 'Sertifikat keamanan pangan', NOW(), NOW()),
(20, 'Panduan BPOM', NOW(), NOW()),
-- Funding Programs Benefits
(21, 'Bunga rendah 6% per tahun', NOW(), NOW()),
(21, 'Proses pencairan cepat', NOW(), NOW()),
(21, 'Tanpa agunan untuk pinjaman < 10 juta', NOW(), NOW()),
(22, 'Bunga rendah 6% per tahun', NOW(), NOW()),
(22, 'Plafon hingga 500 juta', NOW(), NOW()),
(23, 'Proses mudah dan cepat', NOW(), NOW()),
(23, 'Pencairan 3 hari kerja', NOW(), NOW()),
(24, 'Tanpa bunga', NOW(), NOW()),
(24, 'Pembagian profit sharing', NOW(), NOW()),
(25, 'Pembiayaan syariah', NOW(), NOW()),
(25, 'Tanpa riba', NOW(), NOW()),
(26, 'Bunga subsidi 3%', NOW(), NOW()),
(26, 'Dana dari pemerintah', NOW(), NOW()),
(27, 'Plafon besar', NOW(), NOW()),
(27, 'Tenor fleksibel', NOW(), NOW()),
(28, 'Untuk investasi aset', NOW(), NOW()),
(28, 'Tenor hingga 5 tahun', NOW(), NOW()),
(29, 'Online approval', NOW(), NOW()),
(29, 'Pencairan cepat', NOW(), NOW()),
(30, 'Bunga kompetitif', NOW(), NOW()),
(30, 'Tenor panjang', NOW(), NOW());
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Program Requirements for all programs
INSERT INTO program_requirements (program_id, name, created_at, updated_at) VALUES
-- Training Programs Requirements
(1, 'Memiliki usaha yang sudah berjalan minimal 1 tahun', NOW(), NOW()),
(1, 'Memiliki NIB dan NPWP', NOW(), NOW()),
(2, 'Memiliki usaha aktif', NOW(), NOW()),
(2, 'Berkomitmen mengikuti pelatihan penuh', NOW(), NOW()),
(3, 'Pemilik atau pengelola UMKM', NOW(), NOW()),
(3, 'Berusia minimal 18 tahun', NOW(), NOW()),
(4, 'Memiliki produk untuk ekspor', NOW(), NOW()),
(4, 'Memiliki legalitas usaha', NOW(), NOW()),
(5, 'Memiliki produk UMKM', NOW(), NOW()),
(5, 'Membawa kamera sendiri', NOW(), NOW()),
(6, 'Memiliki smartphone atau laptop', NOW(), NOW()),
(6, 'Koneksi internet stabil', NOW(), NOW()),
(7, 'Memiliki akun media sosial', NOW(), NOW()),
(7, 'Bersedia berbagi konten', NOW(), NOW()),
(8, 'Memiliki produk UMKM', NOW(), NOW()),
(8, 'Membawa sample produk', NOW(), NOW()),
(9, 'Memiliki usaha minimal 6 bulan', NOW(), NOW()),
(9, 'Berkomitmen membangun brand', NOW(), NOW()),
(10, 'Memiliki platform digital', NOW(), NOW()),
(10, 'Bersedia membuat konten rutin', NOW(), NOW()),
-- Certification Programs Requirements
(11, 'Memiliki NIB dan izin usaha', NOW(), NOW()),
(11, 'Produk sudah diproduksi konsisten', NOW(), NOW()),
(12, 'Memiliki legalitas usaha', NOW(), NOW()),
(12, 'Produk sesuai standar SNI', NOW(), NOW()),
(13, 'Memiliki sistem manajemen', NOW(), NOW()),
(13, 'Bersedia audit berkala', NOW(), NOW()),
(14, 'Produksi makanan atau minuman', NOW(), NOW()),
(14, 'Fasilitas produksi memadai', NOW(), NOW()),
(15, 'Produksi pertanian organik', NOW(), NOW()),
(15, 'Tidak menggunakan pestisida kimia', NOW(), NOW()),
(16, 'Usaha pangan rumah tangga', NOW(), NOW()),
(16, 'Memiliki tempat produksi', NOW(), NOW()),
(17, 'Fasilitas produksi memenuhi standar', NOW(), NOW()),
(17, 'SOP produksi tersedia', NOW(), NOW()),
(18, 'Komitmen perdagangan adil', NOW(), NOW()),
(18, 'Produk berkualitas', NOW(), NOW()),
(19, 'Produk ramah lingkungan', NOW(), NOW()),
(19, 'Proses produksi sustainable', NOW(), NOW()),
(20, 'Produksi makanan', NOW(), NOW()),
(20, 'Lulus uji laboratorium', NOW(), NOW()),
-- Funding Programs Requirements
(21, 'Memiliki usaha produktif minimal 6 bulan', NOW(), NOW()),
(21, 'NIB dan NPWP aktif', NOW(), NOW()),
(22, 'Usaha berjalan minimal 1 tahun', NOW(), NOW()),
(22, 'Omset sesuai persyaratan', NOW(), NOW()),
(23, 'Usaha mikro aktif', NOW(), NOW()),
(23, 'Tidak sedang menerima kredit lain', NOW(), NOW()),
(24, 'Usaha dengan potensi tinggi', NOW(), NOW()),
(24, 'Business plan lengkap', NOW(), NOW()),
(25, 'Memahami prinsip syariah', NOW(), NOW()),
(25, 'Usaha halal', NOW(), NOW()),
(26, 'Terdaftar di Kemenkop', NOW(), NOW()),
(26, 'Usaha produktif', NOW(), NOW()),
(27, 'Usaha stabil minimal 2 tahun', NOW(), NOW()),
(27, 'Laporan keuangan lengkap', NOW(), NOW()),
(28, 'Rencana investasi jelas', NOW(), NOW()),
(28, 'Agunan tersedia', NOW(), NOW()),
(29, 'Terdaftar di platform', NOW(), NOW()),
(29, 'Rekening bank aktif', NOW(), NOW()),
(30, 'Usaha produktif', NOW(), NOW()),
(30, 'Legalitas lengkap', NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM program_benefits WHERE program_id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30);
DELETE FROM program_requirements WHERE program_id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30);
DELETE FROM programs WHERE id IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30);
-- +goose StatementEnd
