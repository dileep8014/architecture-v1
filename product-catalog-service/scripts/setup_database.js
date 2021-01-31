db.auth("admin", "admin")

db.createUser(
    {
        user: "pc-admin",
        pwd: "pc-pass",
        roles: [
            {
                role: "readWrite",
                db : "product-catalog-database"
            }
        ]
    }
)

db.createCollection("categories", { autoIndexID: true })

db.createCollection("products", { autoIndexID: true })