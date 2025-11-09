-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, email, password, role_id, is_active, created_at, updated_at) VALUES
-- Pelaku Usaha
(101 ,'UMKM User 1', 'umkm1@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 4, true, NOW(), NOW()),
(102 ,'UMKM User 2', 'umkm2@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 4, true, NOW(), NOW()),
(103 ,'UMKM User 3', 'umkm3@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 4, true, NOW(), NOW()),
(104 ,'UMKM User 4', 'umkm4@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 4, true, NOW(), NOW()),
(105 ,'UMKM User 5', 'umkm5@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 4, true, NOW(), NOW());
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO umkms (user_id, business_name, nik, gender, birth_date, phone, address, province_id, city_id, district, subdistrict, postal_code, nib, npwp, kartu_type, kartu_number, created_at, updated_at) VALUES
(101, 'Warung Makan Berkah', '3201234567890001', 'male', '1985-05-15', '081234567890', 'Jl. Sudirman No. 123', 32, 3273, 'Karang Pilang', 'Kedurus', '40123', '1234567890123', '123456789012345', 'produktif', 'KUR001234567', NOW(), NOW()),
(102, 'Toko Baju Fashion', '3201234567890002', 'female', '1990-08-20', '081234567891', 'Jl. Asia Afrika No. 456', 32, 3273, 'Karang Pilang', 'Kedurus', '40124', '1234567890124', '123456789012346', 'produktif', 'KUR001234568', NOW(), NOW()),
(103, 'Kafe Kopi Nusantara', '3201234567890003', 'male', '1988-03-10', '081234567892', 'Jl. Braga No. 789', 32, 3273, 'Karang Pilang', 'Kedurus', '40125', '1234567890125', '123456789012347', 'produktif', 'KUR001234569', NOW(), NOW()),
(104, 'Toko Oleh-Oleh Bandung', '3201234567890004', 'female', '1992-11-25', '081234567893', 'Jl. Cihampelas No. 321', 32, 3273, 'Karang Pilang', 'Kedurus', '40126', '1234567890126', '123456789012348', 'afirmatif', 'KUR001234570', NOW(), NOW()),
(105, 'Bengkel Motor Jaya', '3201234567890005', 'male', '1987-07-18', '081234567894', 'Jl. Dago No. 654', 32, 3273, 'Karang Pilang', 'Kedurus', '40127', '1234567890127', '123456789012349', 'produktif', 'KUR001234571', NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM umkms WHERE user_id IN (101, 102, 103, 104, 105);
DELETE FROM users WHERE id IN (101, 102, 103, 104, 105);
-- +goose StatementEnd
