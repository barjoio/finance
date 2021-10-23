SELECT
    id,
    name,
    projected,
    Actual,
    (CASE WHEN Difference is NULL THEN projected ELSE Difference END) As Difference
FROM (
    SELECT
        categories.id,          -- Category ID
        categories.name,        -- Category Name
        categories.projected,   -- Projected
        SUM(payments.value) AS Actual,    -- Actual
        (categories.projected - SUM(payments.value)) AS Difference -- Difference
    FROM categories
    LEFT JOIN (SELECT value, category_id FROM payments WHERE year = ? AND month = ?) AS payments
        ON categories.id = payments.category_id
    GROUP BY categories.id
)