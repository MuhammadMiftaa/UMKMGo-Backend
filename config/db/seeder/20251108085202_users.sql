-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, email, password, role_id, is_active, created_at, updated_at) VALUES
-- Admin Screening
('Admin Screening 1', 'adminscreening1@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 2, true, NOW(), NOW()),
('Admin Screening 2', 'adminscreening2@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 2, true, NOW(), NOW()),
-- Admin Vendor
('Admin Vendor 1', 'adminvendor1@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 3, true, NOW(), NOW()),
('Admin Vendor 2', 'adminvendor2@umkm.go.id', '$2a$04$qTicwGjrvEBZ1Cd6QwYuS.ENR2PRzu01/TSwzIeFsJKLu5P8.q.SS', 3, true, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email IN (
    'adminscreening1@umkm.go.id',
    'adminscreening2@umkm.go.id',
    'adminvendor1@umkm.go.id',
    'adminvendor2@umkm.go.id'
);
-- +goose StatementEnd
