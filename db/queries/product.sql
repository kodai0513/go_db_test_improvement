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

-- 集計クエリ
-- name: ListProductCountAndAvgPriceByCategory :many
-- カテゴリごとの商品数と平均価格を集計する。
SELECT
    c.id AS category_id,
    c.name AS category_name,
    COUNT(DISTINCT pc.product_id) AS product_count,
    COALESCE(AVG(p.price_yen), 0)::float8 AS avg_price_yen
FROM categories c
JOIN product_categories pc ON pc.category_id = c.id
JOIN products p ON p.id = pc.product_id
GROUP BY c.id, c.name
ORDER BY product_count DESC, c.name;

-- name: ListVariantPriceRangeByProduct :many
-- 商品ごとのバリエーション数と価格レンジ(最安値・最高値)を集計する。
SELECT
    p.id AS product_id,
    p.name AS product_name,
    COUNT(pv.id) AS variant_count,
    MIN(pv.price_yen) AS min_price_yen,
    MAX(pv.price_yen) AS max_price_yen
FROM products p
JOIN product_variants pv ON pv.product_id = p.id
GROUP BY p.id, p.name
ORDER BY variant_count DESC, p.name;

-- name: ListProductCountByTag :many
-- 商品タグごとの商品数を集計する。
SELECT
    pt.id AS product_tag_id,
    pt.name AS tag_name,
    COUNT(pta.product_id) AS product_count
FROM product_tags pt
JOIN product_tag_assignments pta ON pta.product_tag_id = pt.id
GROUP BY pt.id, pt.name
ORDER BY product_count DESC, pt.name;
