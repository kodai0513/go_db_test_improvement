-- name: CreateProduct :one
INSERT INTO products (id, name, price_yen, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- カテゴリ
-- name: CreateCategory :one
INSERT INTO categories (id, name, parent_category_id, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1;

-- name: ListCategories :many
SELECT * FROM categories ORDER BY created_at DESC;

-- name: ListCategoriesByParentCategoryID :many
SELECT * FROM categories WHERE parent_category_id = $1 ORDER BY created_at DESC;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- 商品とカテゴリの中間テーブル
-- name: CreateProductCategory :one
INSERT INTO product_categories (id, product_id, category_id, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProductCategory :one
SELECT * FROM product_categories WHERE id = $1;

-- name: ListProductCategories :many
SELECT * FROM product_categories ORDER BY created_at DESC;

-- name: ListProductCategoriesByProductID :many
SELECT * FROM product_categories WHERE product_id = $1 ORDER BY created_at DESC;

-- name: ListProductCategoriesByCategoryID :many
SELECT * FROM product_categories WHERE category_id = $1 ORDER BY created_at DESC;

-- name: DeleteProductCategory :exec
DELETE FROM product_categories WHERE id = $1;

-- 商品属性マスタ
-- name: CreateProductAttribute :one
INSERT INTO product_attributes (id, name, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProductAttribute :one
SELECT * FROM product_attributes WHERE id = $1;

-- name: ListProductAttributes :many
SELECT * FROM product_attributes ORDER BY created_at DESC;

-- name: DeleteProductAttribute :exec
DELETE FROM product_attributes WHERE id = $1;

-- 商品属性値
-- name: CreateProductAttributeValue :one
INSERT INTO product_attribute_values (id, product_id, product_attribute_id, value)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProductAttributeValue :one
SELECT * FROM product_attribute_values WHERE id = $1;

-- name: ListProductAttributeValues :many
SELECT * FROM product_attribute_values ORDER BY id DESC;

-- name: ListProductAttributeValuesByProductID :many
SELECT * FROM product_attribute_values WHERE product_id = $1 ORDER BY id DESC;

-- name: ListProductAttributeValuesByProductAttributeID :many
SELECT * FROM product_attribute_values WHERE product_attribute_id = $1 ORDER BY id DESC;

-- name: DeleteProductAttributeValue :exec
DELETE FROM product_attribute_values WHERE id = $1;

-- 商品バリエーション(SKU)
-- name: CreateProductVariant :one
INSERT INTO product_variants (id, product_id, sku, price_yen, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProductVariant :one
SELECT * FROM product_variants WHERE id = $1;

-- name: ListProductVariants :many
SELECT * FROM product_variants ORDER BY created_at DESC;

-- name: ListProductVariantsByProductID :many
SELECT * FROM product_variants WHERE product_id = $1 ORDER BY created_at DESC;

-- name: DeleteProductVariant :exec
DELETE FROM product_variants WHERE id = $1;

-- 商品画像
-- name: CreateProductImage :one
INSERT INTO product_images (id, product_id, url, sort_order)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetProductImage :one
SELECT * FROM product_images WHERE id = $1;

-- name: ListProductImages :many
SELECT * FROM product_images ORDER BY id DESC;

-- name: ListProductImagesByProductID :many
SELECT * FROM product_images WHERE product_id = $1 ORDER BY id DESC;

-- name: DeleteProductImage :exec
DELETE FROM product_images WHERE id = $1;

-- ブランド
-- name: CreateBrand :one
INSERT INTO brands (id, name, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBrand :one
SELECT * FROM brands WHERE id = $1;

-- name: ListBrands :many
SELECT * FROM brands ORDER BY created_at DESC;

-- name: DeleteBrand :exec
DELETE FROM brands WHERE id = $1;

-- 商品タグマスタ
-- name: CreateProductTag :one
INSERT INTO product_tags (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetProductTag :one
SELECT * FROM product_tags WHERE id = $1;

-- name: ListProductTags :many
SELECT * FROM product_tags ORDER BY id DESC;

-- name: DeleteProductTag :exec
DELETE FROM product_tags WHERE id = $1;

-- 商品とタグの割り当て
-- name: CreateProductTagAssignment :one
INSERT INTO product_tag_assignments (id, product_id, product_tag_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProductTagAssignment :one
SELECT * FROM product_tag_assignments WHERE id = $1;

-- name: ListProductTagAssignments :many
SELECT * FROM product_tag_assignments ORDER BY id DESC;

-- name: ListProductTagAssignmentsByProductID :many
SELECT * FROM product_tag_assignments WHERE product_id = $1 ORDER BY id DESC;

-- name: ListProductTagAssignmentsByProductTagID :many
SELECT * FROM product_tag_assignments WHERE product_tag_id = $1 ORDER BY id DESC;

-- name: DeleteProductTagAssignment :exec
DELETE FROM product_tag_assignments WHERE id = $1;
