-- +goose Up
-- +goose StatementBegin

TRUNCATE TABLE applications CASCADE;

-- Seed Applications for Training (10 applications)
INSERT INTO applications (umkm_id, program_id, type, status, submitted_at, expired_at, created_at, updated_at) VALUES
(1, 1, 'training', 'screening', NOW() - INTERVAL '5 days', NOW() + INTERVAL '25 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
(2, 2, 'training', 'screening', NOW() - INTERVAL '4 days', NOW() + INTERVAL '26 days', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(3, 3, 'training', 'screening', NOW() - INTERVAL '3 days', NOW() + INTERVAL '27 days', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(4, 4, 'training', 'screening', NOW() - INTERVAL '2 days', NOW() + INTERVAL '28 days', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
(5, 5, 'training', 'screening', NOW() - INTERVAL '1 day', NOW() + INTERVAL '29 days', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),
(1, 6, 'training', 'screening', NOW() - INTERVAL '6 days', NOW() + INTERVAL '24 days', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),
(2, 7, 'training', 'screening', NOW() - INTERVAL '7 days', NOW() + INTERVAL '23 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
(3, 8, 'training', 'screening', NOW() - INTERVAL '8 days', NOW() + INTERVAL '22 days', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),
(4, 9, 'training', 'screening', NOW() - INTERVAL '9 days', NOW() + INTERVAL '21 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),
(5, 10, 'training', 'screening', NOW() - INTERVAL '10 days', NOW() + INTERVAL '20 days', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days');
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Applications for Certification (10 applications)
INSERT INTO applications (umkm_id, program_id, type, status, submitted_at, expired_at, created_at, updated_at) VALUES
(1, 11, 'certification', 'screening', NOW() - INTERVAL '5 days', NOW() + INTERVAL '25 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
(2, 12, 'certification', 'screening', NOW() - INTERVAL '4 days', NOW() + INTERVAL '26 days', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(3, 13, 'certification', 'screening', NOW() - INTERVAL '3 days', NOW() + INTERVAL '27 days', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(4, 14, 'certification', 'screening', NOW() - INTERVAL '2 days', NOW() + INTERVAL '28 days', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
(5, 15, 'certification', 'screening', NOW() - INTERVAL '1 day', NOW() + INTERVAL '29 days', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),
(1, 16, 'certification', 'screening', NOW() - INTERVAL '6 days', NOW() + INTERVAL '24 days', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),
(2, 17, 'certification', 'screening', NOW() - INTERVAL '7 days', NOW() + INTERVAL '23 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
(3, 18, 'certification', 'screening', NOW() - INTERVAL '8 days', NOW() + INTERVAL '22 days', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),
(4, 19, 'certification', 'screening', NOW() - INTERVAL '9 days', NOW() + INTERVAL '21 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),
(5, 20, 'certification', 'screening', NOW() - INTERVAL '10 days', NOW() + INTERVAL '20 days', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days');
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Applications for Funding (10 applications)
INSERT INTO applications (umkm_id, program_id, type, status, submitted_at, expired_at, created_at, updated_at) VALUES
(1, 21, 'funding', 'screening', NOW() - INTERVAL '5 days', NOW() + INTERVAL '25 days', NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
(2, 22, 'funding', 'screening', NOW() - INTERVAL '4 days', NOW() + INTERVAL '26 days', NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(3, 23, 'funding', 'screening', NOW() - INTERVAL '3 days', NOW() + INTERVAL '27 days', NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(4, 24, 'funding', 'screening', NOW() - INTERVAL '2 days', NOW() + INTERVAL '28 days', NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
(5, 25, 'funding', 'screening', NOW() - INTERVAL '1 day', NOW() + INTERVAL '29 days', NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),
(1, 26, 'funding', 'screening', NOW() - INTERVAL '6 days', NOW() + INTERVAL '24 days', NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),
(2, 27, 'funding', 'screening', NOW() - INTERVAL '7 days', NOW() + INTERVAL '23 days', NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
(3, 28, 'funding', 'screening', NOW() - INTERVAL '8 days', NOW() + INTERVAL '22 days', NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),
(4, 29, 'funding', 'screening', NOW() - INTERVAL '9 days', NOW() + INTERVAL '21 days', NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),
(5, 30, 'funding', 'screening', NOW() - INTERVAL '10 days', NOW() + INTERVAL '20 days', NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days');
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Application Documents for all applications
INSERT INTO application_documents (application_id, type, file, created_at, updated_at) VALUES
-- Training Applications Documents (1-10)
(1, 'ktp', 'https://storage.example.com/documents/ktp_umkm1_training1.pdf', NOW(), NOW()),
(1, 'nib', 'https://storage.example.com/documents/nib_umkm1_training1.pdf', NOW(), NOW()),
(1, 'npwp', 'https://storage.example.com/documents/npwp_umkm1_training1.pdf', NOW(), NOW()),
(1, 'proposal', 'https://storage.example.com/documents/proposal_umkm1_training1.pdf', NOW(), NOW()),
(2, 'ktp', 'https://storage.example.com/documents/ktp_umkm2_training2.pdf', NOW(), NOW()),
(2, 'nib', 'https://storage.example.com/documents/nib_umkm2_training2.pdf', NOW(), NOW()),
(2, 'npwp', 'https://storage.example.com/documents/npwp_umkm2_training2.pdf', NOW(), NOW()),
(3, 'ktp', 'https://storage.example.com/documents/ktp_umkm3_training3.pdf', NOW(), NOW()),
(3, 'nib', 'https://storage.example.com/documents/nib_umkm3_training3.pdf', NOW(), NOW()),
(3, 'proposal', 'https://storage.example.com/documents/proposal_umkm3_training3.pdf', NOW(), NOW()),
(4, 'ktp', 'https://storage.example.com/documents/ktp_umkm4_training4.pdf', NOW(), NOW()),
(4, 'npwp', 'https://storage.example.com/documents/npwp_umkm4_training4.pdf', NOW(), NOW()),
(4, 'portfolio', 'https://storage.example.com/documents/portfolio_umkm4_training4.pdf', NOW(), NOW()),
(5, 'ktp', 'https://storage.example.com/documents/ktp_umkm5_training5.pdf', NOW(), NOW()),
(5, 'nib', 'https://storage.example.com/documents/nib_umkm5_training5.pdf', NOW(), NOW()),
(6, 'ktp', 'https://storage.example.com/documents/ktp_umkm1_training6.pdf', NOW(), NOW()),
(6, 'npwp', 'https://storage.example.com/documents/npwp_umkm1_training6.pdf', NOW(), NOW()),
(7, 'ktp', 'https://storage.example.com/documents/ktp_umkm2_training7.pdf', NOW(), NOW()),
(7, 'nib', 'https://storage.example.com/documents/nib_umkm2_training7.pdf', NOW(), NOW()),
(8, 'ktp', 'https://storage.example.com/documents/ktp_umkm3_training8.pdf', NOW(), NOW()),
(8, 'proposal', 'https://storage.example.com/documents/proposal_umkm3_training8.pdf', NOW(), NOW()),
(9, 'ktp', 'https://storage.example.com/documents/ktp_umkm4_training9.pdf', NOW(), NOW()),
(9, 'nib', 'https://storage.example.com/documents/nib_umkm4_training9.pdf', NOW(), NOW()),
(10, 'ktp', 'https://storage.example.com/documents/ktp_umkm5_training10.pdf', NOW(), NOW()),
(10, 'npwp', 'https://storage.example.com/documents/npwp_umkm5_training10.pdf', NOW(), NOW()),
-- Certification Applications Documents (11-20)
(11, 'ktp', 'https://storage.example.com/documents/ktp_umkm1_cert1.pdf', NOW(), NOW()),
(11, 'nib', 'https://storage.example.com/documents/nib_umkm1_cert1.pdf', NOW(), NOW()),
(11, 'npwp', 'https://storage.example.com/documents/npwp_umkm1_cert1.pdf', NOW(), NOW()),
(12, 'ktp', 'https://storage.example.com/documents/ktp_umkm2_cert2.pdf', NOW(), NOW()),
(12, 'nib', 'https://storage.example.com/documents/nib_umkm2_cert2.pdf', NOW(), NOW()),
(13, 'ktp', 'https://storage.example.com/documents/ktp_umkm3_cert3.pdf', NOW(), NOW()),
(13, 'npwp', 'https://storage.example.com/documents/npwp_umkm3_cert3.pdf', NOW(), NOW()),
(14, 'ktp', 'https://storage.example.com/documents/ktp_umkm4_cert4.pdf', NOW(), NOW()),
(14, 'nib', 'https://storage.example.com/documents/nib_umkm4_cert4.pdf', NOW(), NOW()),
(15, 'ktp', 'https://storage.example.com/documents/ktp_umkm5_cert5.pdf', NOW(), NOW()),
(15, 'portfolio', 'https://storage.example.com/documents/portfolio_umkm5_cert5.pdf', NOW(), NOW()),
(16, 'ktp', 'https://storage.example.com/documents/ktp_umkm1_cert6.pdf', NOW(), NOW()),
(16, 'nib', 'https://storage.example.com/documents/nib_umkm1_cert6.pdf', NOW(), NOW()),
(17, 'ktp', 'https://storage.example.com/documents/ktp_umkm2_cert7.pdf', NOW(), NOW()),
(17, 'npwp', 'https://storage.example.com/documents/npwp_umkm2_cert7.pdf', NOW(), NOW()),
(18, 'ktp', 'https://storage.example.com/documents/ktp_umkm3_cert8.pdf', NOW(), NOW()),
(18, 'nib', 'https://storage.example.com/documents/nib_umkm3_cert8.pdf', NOW(), NOW()),
(19, 'ktp', 'https://storage.example.com/documents/ktp_umkm4_cert9.pdf', NOW(), NOW()),
(19, 'proposal', 'https://storage.example.com/documents/proposal_umkm4_cert9.pdf', NOW(), NOW()),
(20, 'ktp', 'https://storage.example.com/documents/ktp_umkm5_cert10.pdf', NOW(), NOW()),
(20, 'nib', 'https://storage.example.com/documents/nib_umkm5_cert10.pdf', NOW(), NOW()),
-- Funding Applications Documents (21-30)
(21, 'ktp', 'https://storage.example.com/documents/ktp_umkm1_fund1.pdf', NOW(), NOW()),
(21, 'nib', 'https://storage.example.com/documents/nib_umkm1_fund1.pdf', NOW(), NOW()),
(21, 'npwp', 'https://storage.example.com/documents/npwp_umkm1_fund1.pdf', NOW(), NOW()),
(21, 'proposal', 'https://storage.example.com/documents/proposal_umkm1_fund1.pdf', NOW(), NOW()),
(21, 'rekening', 'https://storage.example.com/documents/rekening_umkm1_fund1.pdf', NOW(), NOW()),
(22, 'ktp', 'https://storage.example.com/documents/ktp_umkm2_fund2.pdf', NOW(), NOW()),
(22, 'nib', 'https://storage.example.com/documents/nib_umkm2_fund2.pdf', NOW(), NOW()),
(22, 'npwp', 'https://storage.example.com/documents/npwp_umkm2_fund2.pdf', NOW(), NOW()),
(22, 'rekening', 'https://storage.example.com/documents/rekening_umkm2_fund2.pdf', NOW(), NOW()),
(23, 'ktp', 'https://storage.example.com/documents/ktp_umkm3_fund3.pdf', NOW(), NOW()),
(23, 'nib', 'https://storage.example.com/documents/nib_umkm3_fund3.pdf', NOW(), NOW()),
(23, 'proposal', 'https://storage.example.com/documents/proposal_umkm3_fund3.pdf', NOW(), NOW()),
(23, 'rekening', 'https://storage.example.com/documents/rekening_umkm3_fund3.pdf', NOW(), NOW()),
(24, 'ktp', 'https://storage.example.com/documents/ktp_umkm4_fund4.pdf', NOW(), NOW()),
(24, 'npwp', 'https://storage.example.com/documents/npwp_umkm4_fund4.pdf', NOW(), NOW()),
(24, 'proposal', 'https://storage.example.com/documents/proposal_umkm4_fund4.pdf', NOW(), NOW()),
(24, 'rekening', 'https://storage.example.com/documents/rekening_umkm4_fund4.pdf', NOW(), NOW()),
(25, 'ktp', 'https://storage.example.com/documents/ktp_umkm5_fund5.pdf', NOW(), NOW()),
(25, 'nib', 'https://storage.example.com/documents/nib_umkm5_fund5.pdf', NOW(), NOW()),
(25, 'npwp', 'https://storage.example.com/documents/npwp_umkm5_fund5.pdf', NOW(), NOW()),
(25, 'rekening', 'https://storage.example.com/documents/rekening_umkm5_fund5.pdf', NOW(), NOW()),
(26, 'ktp', 'https://storage.example.com/documents/ktp_umkm1_fund6.pdf', NOW(), NOW()),
(26, 'nib', 'https://storage.example.com/documents/nib_umkm1_fund6.pdf', NOW(), NOW()),
(26, 'proposal', 'https://storage.example.com/documents/proposal_umkm1_fund6.pdf', NOW(), NOW()),
(26, 'rekening', 'https://storage.example.com/documents/rekening_umkm1_fund6.pdf', NOW(), NOW()),
(27, 'ktp', 'https://storage.example.com/documents/ktp_umkm2_fund7.pdf', NOW(), NOW()),
(27, 'npwp', 'https://storage.example.com/documents/npwp_umkm2_fund7.pdf', NOW(), NOW()),
(27, 'rekening', 'https://storage.example.com/documents/rekening_umkm2_fund7.pdf', NOW(), NOW()),
(28, 'ktp', 'https://storage.example.com/documents/ktp_umkm3_fund8.pdf', NOW(), NOW()),
(28, 'nib', 'https://storage.example.com/documents/nib_umkm3_fund8.pdf', NOW(), NOW()),
(28, 'proposal', 'https://storage.example.com/documents/proposal_umkm3_fund8.pdf', NOW(), NOW()),
(28, 'rekening', 'https://storage.example.com/documents/rekening_umkm3_fund8.pdf', NOW(), NOW()),
(29, 'ktp', 'https://storage.example.com/documents/ktp_umkm4_fund9.pdf', NOW(), NOW()),
(29, 'nib', 'https://storage.example.com/documents/nib_umkm4_fund9.pdf', NOW(), NOW()),
(29, 'rekening', 'https://storage.example.com/documents/rekening_umkm4_fund9.pdf', NOW(), NOW()),
(30, 'ktp', 'https://storage.example.com/documents/ktp_umkm5_fund10.pdf', NOW(), NOW()),
(30, 'npwp', 'https://storage.example.com/documents/npwp_umkm5_fund10.pdf', NOW(), NOW()),
(30, 'proposal', 'https://storage.example.com/documents/proposal_umkm5_fund10.pdf', NOW(), NOW()),
(30, 'rekening', 'https://storage.example.com/documents/rekening_umkm5_fund10.pdf', NOW(), NOW());
-- +goose StatementEnd

-- +goose StatementBegin
-- Seed Application Histories (only submit action for all applications)
INSERT INTO application_histories (application_id, status, notes, actioned_at, actioned_by, created_at, updated_at) VALUES
-- Training Applications Histories (1-10)
(1, 'submit', 'Application submitted for Digital Marketing Training', NOW() - INTERVAL '5 days', 6, NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
(2, 'submit', 'Application submitted for Basic Accounting Training', NOW() - INTERVAL '4 days', 7, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(3, 'submit', 'Application submitted for Business Management Training', NOW() - INTERVAL '3 days', 8, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(4, 'submit', 'Application submitted for Export Import Training', NOW() - INTERVAL '2 days', 9, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
(5, 'submit', 'Application submitted for Product Photography Training', NOW() - INTERVAL '1 day', 10, NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),
(6, 'submit', 'Application submitted for E-Commerce Training', NOW() - INTERVAL '6 days', 6, NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),
(7, 'submit', 'Application submitted for Social Media Marketing Training', NOW() - INTERVAL '7 days', 7, NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
(8, 'submit', 'Application submitted for Packaging Design Training', NOW() - INTERVAL '8 days', 8, NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),
(9, 'submit', 'Application submitted for UMKM Branding Training', NOW() - INTERVAL '9 days', 9, NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),
(10, 'submit', 'Application submitted for Content Marketing Training', NOW() - INTERVAL '10 days', 10, NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),
-- Certification Applications Histories (11-20)
(11, 'submit', 'Application submitted for Halal Certification', NOW() - INTERVAL '5 days', 6, NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
(12, 'submit', 'Application submitted for SNI Product Certification', NOW() - INTERVAL '4 days', 7, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(13, 'submit', 'Application submitted for ISO 9001 Certification', NOW() - INTERVAL '3 days', 8, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(14, 'submit', 'Application submitted for HACCP Certification', NOW() - INTERVAL '2 days', 9, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
(15, 'submit', 'Application submitted for Organic Certification', NOW() - INTERVAL '1 day', 10, NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),
(16, 'submit', 'Application submitted for PIRT Certification', NOW() - INTERVAL '6 days', 6, NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),
(17, 'submit', 'Application submitted for GMP Certification', NOW() - INTERVAL '7 days', 7, NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
(18, 'submit', 'Application submitted for Fairtrade Certification', NOW() - INTERVAL '8 days', 8, NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),
(19, 'submit', 'Application submitted for Ecolabel Certification', NOW() - INTERVAL '9 days', 9, NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),
(20, 'submit', 'Application submitted for Food Safety Certification', NOW() - INTERVAL '10 days', 10, NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days'),
-- Funding Applications Histories (21-30)
(21, 'submit', 'Application submitted for KUR Mikro', NOW() - INTERVAL '5 days', 6, NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
(22, 'submit', 'Application submitted for KUR Kecil', NOW() - INTERVAL '4 days', 7, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(23, 'submit', 'Application submitted for Ultra Mikro Financing', NOW() - INTERVAL '3 days', 8, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(24, 'submit', 'Application submitted for UMKM Venture Capital', NOW() - INTERVAL '2 days', 9, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
(25, 'submit', 'Application submitted for Syariah Mikro Loan', NOW() - INTERVAL '1 day', 10, NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day'),
(26, 'submit', 'Application submitted for Revolving Fund UMKM', NOW() - INTERVAL '6 days', 6, NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days'),
(27, 'submit', 'Application submitted for Working Capital Credit', NOW() - INTERVAL '7 days', 7, NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days'),
(28, 'submit', 'Application submitted for Investment Loan', NOW() - INTERVAL '8 days', 8, NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days'),
(29, 'submit', 'Application submitted for P2P Lending', NOW() - INTERVAL '9 days', 9, NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days'),
(30, 'submit', 'Application submitted for Productive Business Credit', NOW() - INTERVAL '10 days', 10, NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE application_histories CASCADE;
TRUNCATE TABLE application_documents CASCADE;
TRUNCATE TABLE applications CASCADE;
-- +goose StatementEnd
