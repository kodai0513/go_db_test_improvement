-- name: CreateBanner :one
INSERT INTO banners (id, title, image_url, starts_at, ends_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetBanner :one
SELECT * FROM banners WHERE id = $1;

-- name: ListBanners :many
SELECT * FROM banners ORDER BY id DESC;

-- name: DeleteBanner :exec
DELETE FROM banners WHERE id = $1;

-- name: CreatePage :one
INSERT INTO pages (id, slug, title, body)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPage :one
SELECT * FROM pages WHERE id = $1;

-- name: ListPages :many
SELECT * FROM pages ORDER BY id DESC;

-- name: DeletePage :exec
DELETE FROM pages WHERE id = $1;

-- name: CreateBlogPost :one
INSERT INTO blog_posts (id, title, body, published_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetBlogPost :one
SELECT * FROM blog_posts WHERE id = $1;

-- name: ListBlogPosts :many
SELECT * FROM blog_posts ORDER BY id DESC;

-- name: DeleteBlogPost :exec
DELETE FROM blog_posts WHERE id = $1;

-- name: CreateBlogCategory :one
INSERT INTO blog_categories (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetBlogCategory :one
SELECT * FROM blog_categories WHERE id = $1;

-- name: ListBlogCategories :many
SELECT * FROM blog_categories ORDER BY id DESC;

-- name: DeleteBlogCategory :exec
DELETE FROM blog_categories WHERE id = $1;

-- name: CreateBlogPostTag :one
INSERT INTO blog_post_tags (id, blog_post_id, tag)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBlogPostTag :one
SELECT * FROM blog_post_tags WHERE id = $1;

-- name: ListBlogPostTags :many
SELECT * FROM blog_post_tags ORDER BY id DESC;

-- name: ListBlogPostTagsByBlogPostID :many
SELECT * FROM blog_post_tags WHERE blog_post_id = $1 ORDER BY id DESC;

-- name: DeleteBlogPostTag :exec
DELETE FROM blog_post_tags WHERE id = $1;

-- name: CreateMediaAsset :one
INSERT INTO media_assets (id, url, media_type, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMediaAsset :one
SELECT * FROM media_assets WHERE id = $1;

-- name: ListMediaAssets :many
SELECT * FROM media_assets ORDER BY created_at DESC;

-- name: DeleteMediaAsset :exec
DELETE FROM media_assets WHERE id = $1;

-- name: CreateSiteSetting :one
INSERT INTO site_settings (id, setting_key, setting_value)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSiteSetting :one
SELECT * FROM site_settings WHERE id = $1;

-- name: ListSiteSettings :many
SELECT * FROM site_settings ORDER BY id DESC;

-- name: DeleteSiteSetting :exec
DELETE FROM site_settings WHERE id = $1;

-- name: CreateMenu :one
INSERT INTO menus (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetMenu :one
SELECT * FROM menus WHERE id = $1;

-- name: ListMenus :many
SELECT * FROM menus ORDER BY id DESC;

-- name: DeleteMenu :exec
DELETE FROM menus WHERE id = $1;

-- name: CreateMenuItem :one
INSERT INTO menu_items (id, menu_id, label, url, sort_order)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMenuItem :one
SELECT * FROM menu_items WHERE id = $1;

-- name: ListMenuItems :many
SELECT * FROM menu_items ORDER BY id DESC;

-- name: ListMenuItemsByMenuID :many
SELECT * FROM menu_items WHERE menu_id = $1 ORDER BY id DESC;

-- name: DeleteMenuItem :exec
DELETE FROM menu_items WHERE id = $1;

-- name: CreateContentTranslation :one
INSERT INTO content_translations (id, page_id, locale, translated_body)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetContentTranslation :one
SELECT * FROM content_translations WHERE id = $1;

-- name: ListContentTranslations :many
SELECT * FROM content_translations ORDER BY id DESC;

-- name: ListContentTranslationsByPageID :many
SELECT * FROM content_translations WHERE page_id = $1 ORDER BY id DESC;

-- name: DeleteContentTranslation :exec
DELETE FROM content_translations WHERE id = $1;
