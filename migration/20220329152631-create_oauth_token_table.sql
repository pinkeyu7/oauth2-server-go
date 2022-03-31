-- +migrate Up
CREATE TABLE `oauth_token` (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    code varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    access varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    refresh varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    expired_at int(11) NOT NULL,
    data varchar(2048) COLLATE utf8mb4_unicode_ci NOT NULL,
    PRIMARY KEY (id),
    KEY idx_oauth_token_code (code),
    KEY idx_oauth_token_expired_at (expired_at),
    KEY idx_oauth_token_access (access),
    KEY idx_oauth_token_refresh (refresh)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down
DROP TABLE `oauth_token`;
