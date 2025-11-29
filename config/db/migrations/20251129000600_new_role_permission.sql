-- +goose Up
-- +goose StatementBegin
INSERT INTO permissions (id, parent_id, name, code, description) VALUES
(21, NULL, 'News', 'MANAGE_NEWS', 'Izin untuk mengelola artikel berita dan konten terkait.'),
(22, 21, 'Create News', 'CREATE_NEWS', 'Izin untuk membuat artikel berita.'),
(23, 21, 'Edit News', 'EDIT_NEWS', 'Izin untuk mengedit artikel berita yang sudah ada.'),
(24, 21, 'Delete News', 'DELETE_NEWS', 'Izin untuk menghapus artikel berita.'),
(25, 21, 'View News', 'VIEW_NEWS', 'Izin untuk melihat artikel berita.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM permissions WHERE id IN (21, 22, 23, 24, 25);
-- +goose StatementEnd
