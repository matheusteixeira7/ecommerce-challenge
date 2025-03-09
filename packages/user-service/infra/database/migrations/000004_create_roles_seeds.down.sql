DELETE
FROM roles
WHERE name IN (
               'customer', 'seller', 'support',
               'inventory_manager', 'admin', 'finance', 'delivery'
    );
